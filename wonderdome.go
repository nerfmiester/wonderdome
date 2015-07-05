package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"

	"github.com/gorilla/mux"
	"github.com/nerfmiester/wonderdome/Structs"
)

var v Structs.OTAVehAvailRateRS

var j Structs.ProviderResp

var p Structs.Provider

var t []Structs.Provider

var mk Structs.VehMkeModel

var tv Structs.TypVehicle

var atv []Structs.TypVehicle

var rates Structs.Rates

var chrg Structs.Charge

var qual Structs.Qualifier

var loc Structs.LocDetails

var vRLI Structs.VehRentLocInf

var infoM Structs.InfoM

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/provider/{providername}", ProviderHandler)
	fmt.Println("Allright geezer")
	//v := OTAVehAvailRateRS{}
	xmlFile, err := os.Open("./resp.xml")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	XMLdata, _ := ioutil.ReadAll(xmlFile)

	err = xml.Unmarshal([]byte(XMLdata), &v)

	fmt.Printf("Here we are, you know %s with sequence number %d and a pickup location of %s\n", v.XMLName, v.SequenceNmbr, v.VehAvailRSCore.VehRentalCore.PickUpLocation.LocationCode)
	fmt.Printf("Now we are here, you know %s \n", v.VehAvailRSCore.VehVendorAvails)
	fmt.Println("type of atv", reflect.TypeOf(atv))

	for _, vehs := range v.VehAvailRSCore.VehVendorAvails.VehVendorAvail.VehAvails.VehAvail {
		fmt.Printf("We have a vehicle of type %s \n", vehs.VehAvailCore.Status)

		http.ListenAndServe(":8080", r)

	}
}

// ProviderHandler = the function will handle provider requests
func ProviderHandler(w http.ResponseWriter, r *http.Request) {

	provider := mux.Vars(r)["providername"]
	fmt.Println("Provider = ", provider)

	w.Write([]byte(fmt.Sprintf("The provider is %s ", provider)))

	s := make([]Structs.Provider, 0)
	s1 := make([]Structs.TypVehicle, 0)
	sCharge := make([]Structs.Charge, 0)
	sLocD := make([]Structs.VehRentLocInf, 0)

	p.SetName(v.VehAvailRSCore.VehVendorAvails.VehVendorAvail.Vendor.CompanyShortName)
	for _, veh := range v.VehAvailRSCore.VehVendorAvails.VehVendorAvail.VehAvails.VehAvail {

		mk.SetName(veh.VehAvailCore.Vehicle.VehMakeModel.Name)
		mk.SetCode(veh.VehAvailCore.Vehicle.VehMakeModel.Code)

		tv.SetAirConditionInd(veh.VehAvailCore.Vehicle.AirConditionInd)
		tv.SetBaggageQuantity(veh.VehAvailCore.Vehicle.BaggageQuantity)
		tv.SetPassengerQuantity(veh.VehAvailCore.Vehicle.PassengerQuantity)
		tv.SetTransmissionType(veh.VehAvailCore.Vehicle.TransmissionType)
		tv.SetPictureURL(veh.VehAvailCore.Vehicle.PictureURL)
		tv.SetVehicleCategory(veh.VehAvailCore.Vehicle.VehType.VehicleCategory)
		tv.SetSize(veh.VehAvailCore.Vehicle.VehClass.Size)
		tv.SetVehMkeModel(mk)

		// Rates

		rates.SetUnitDistance(veh.VehAvailCore.RentalRate.RateDistance.DistUnitName)
		rates.SetUnLimited(veh.VehAvailCore.RentalRate.RateDistance.Unlimited)
		rates.SetVehiclePeriodUnitName(veh.VehAvailCore.RentalRate.RateDistance.VehiclePeriodUnitName)

		qual.SetCorpDiscountNmbr(veh.VehAvailCore.RentalRate.RateQualifier.CorpDiscountNmbr)
		qual.SetRateQualifier(veh.VehAvailCore.RentalRate.RateQualifier.RateQualifier)
		qual.SetRatePeriod(veh.VehAvailCore.RentalRate.RateQualifier.RatePeriod)

		rates.SetQualifier(qual)

		// Charges

		for _, chg := range veh.VehAvailCore.RentalRate.VehicleCharges.VehicleCharge {

			chrg.Amount = chg.Amount
			chrg.CurrencyCode = chg.CurrencyCode
			chrg.Description = chg.Description
			chrg.Purpose = chg.Purpose
			chrg.TaxInclusive = chg.TaxInclusive

			sCharge = append(sCharge, chrg)

		}

		rates.SetCharges(sCharge)

		sCharge = nil

		tv.SetRates(rates)

		fmt.Println("rates -> ", rates)

		s1 = append(s1, tv)

	}
	p.SetTypVehicle(s1)
	sLocD = nil
	for _, ix := range v.VehAvailRSCore.VehVendorAvails.VehVendorAvail.Info.LocationDetails.AdditionalInfo.VehRentLocInfos.VehRentLocInfo {
		vRLI.SetTitle(ix.Title)
		vRLI.SetText(ix.SubSection.Paragraph.Text)
		vRLI.SetType(ix.Type)
		sLocD = append(sLocD, vRLI)
	}
	loc.SetVehRentLocInfo(sLocD)
	infoM.SetLocDetails(loc)
	p.SetInfoM(infoM)
	s = append(s, p)

	t := s
	j := t

	arr := j

	b, err := json.Marshal(arr)
	if err != nil {
		fmt.Println("Error Marshalling struct j:", err)
		return
	}
	w.Write([]byte(b))
}
