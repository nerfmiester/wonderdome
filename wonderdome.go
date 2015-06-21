package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/nerfmiester/wonderdome/Structs"
)

var v Structs.OTAVehAvailRateRS

func main() {

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

	}
}
