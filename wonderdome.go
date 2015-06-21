package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nerfmiester/wonderdome/Structs"
)

var v Structs.OTAVehAvailRateRS

var j Structs.ProviderResp

var p Structs.Provider

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
	fmt.Println("work")

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

	p.SetName(v.VehAvailRSCore.VehVendorAvails.VehVendorAvail.Vendor.CompanyShortName)

	arr := Structs.ProviderResp{
		[]Structs.Provider{
			p,
		},
	}

	b, err := json.Marshal(arr)
	if err != nil {
		fmt.Println("Error Marshalling struct j:", err)
		return
	}
	w.Write([]byte(b))
}
