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

	p.SetName(v.VehAvailRSCore.VehVendorAvails.VehVendorAvail.Vendor.CompanyShortName)
	for _, veh := range v.VehAvailRSCore.VehVendorAvails.VehVendorAvail.VehAvails.VehAvail {

		mk.SetName(veh.VehAvailCore.Vehicle.VehMakeModel.Name)
		mk.SetCode(veh.VehAvailCore.Vehicle.VehMakeModel.Code)

		tv.SetAirConditionInd(veh.VehAvailCore.Vehicle.AirConditionInd)
		tv.SetBaggageQuantity(veh.VehAvailCore.Vehicle.BaggageQuantity)
		tv.SetVehMkeModel(mk)

		s1 = append(s1, tv)

	}
	p.SetTypVehicle(s1)

	for _, sx := range s1 {
		fmt.Println("AirCond = ", sx.AirConditionInd)
		fmt.Println("Baggage qty = ", sx.BaggageQuantity)
		fmt.Println("Vehicle name = ", sx.VehMakeModel.Name)
		fmt.Println("Vehicle Code = ", sx.VehMakeModel.Code)
	}
	s = append(s, p)
	p.SetName("billy")
	s = append(s, p)
	p.SetName("frankie")
	s = append(s, p)
	p.SetName("doozie")
	s = append(s, p)
	t := s
	j := t

	for _, s2 := range s {
		s2.SetTypVehicle(s1)
	}

	for _, sd := range s {
		fmt.Println("Provider -> ", sd.GetProvider())

	}

	arr := j

	b, err := json.Marshal(arr)
	if err != nil {
		fmt.Println("Error Marshalling struct j:", err)
		return
	}
	w.Write([]byte(b))
}
