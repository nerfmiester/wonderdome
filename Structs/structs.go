package Structs

import "encoding/xml"

type OTAVehAvailRateRS struct {
	XMLName        xml.Name       `xml:"OTA_VehAvailRateRS"`
	SequenceNmbr   int            `xml:"SequenceNmbr,attr"`
	Success        string         `xml:"Success"`
	VehAvailRSCore VehAvailRSCore `xml:"VehAvailRSCore"`
}
type VehAvailRSCore struct {
	XMLName         xml.Name        `xml:"VehAvailRSCore"`
	VehRentalCore   VehRentalCore   `xml:"VehRentalCore"`
	VehVendorAvails VehVendorAvails `xml:"VehVendorAvails"`
}
type VehVendorAvails struct {
	XMLName        xml.Name       `xml:"VehVendorAvails"`
	VehVendorAvail VehVendorAvail `xml:"VehVendorAvail"`
}
type VehVendorAvail struct {
	XMLName   xml.Name `xml:"VehVendorAvail"`
	Vendor    Vendor
	VehAvails VehAvails `xml:"VehAvails"`
}
type VehAvails struct {
	XMLName  xml.Name   `xml:"VehAvails"`
	VehAvail []VehAvail `xml:"VehAvail"`
}

type VehAvail struct {
	XMLName      xml.Name     `xml:"VehAvail"`
	VehAvailCore VehAvailCore `xml:"VehAvailCore"`
	VehAvailInfo VehAvailInfo
}
type VehRentalCore struct {
	XMLName        xml.Name `xml:"VehRentalCore"`
	PickUpDateTime string   `xml:"PickUpDateTime,attr"`
	ReturnDateTime string   `xml:"ReturnDateTime,attr"`
	PickUpLocation PickUpLocation
	ReturnLocation ReturnLocation
}
type PickUpLocation struct {
	XMLName      xml.Name `xml:"PickUpLocation"`
	LocationCode string   `xml:"LocationCode,attr"`
}
type ReturnLocation struct {
	XMLName      xml.Name `xml:"ReturnLocation"`
	LocationCode string   `xml:"LocationCode,attr"`
}

type Vendor struct {
	Code             string `xml:"Code,attr"`
	CompanyShortName string `xml:"CompanyShortName,attr"`
}
type VehAvailCore struct {
	XMLName xml.Name `xml:"VehAvailCore"`
	Status  string   `xml:"Status,attr"`
	Vehicle Vehicle
}
type VehAvailInfo struct {
	XMLName xml.Name `xml:"VehAvailInfo"`
}
type Vehicle struct {
	TransmissionType  string `xml:"TransmissionType,attr"`
	AirConditionInd   bool   `xml:"AirConditionInd,attr"`
	BaggageQuantity   int    `xml:"BaggageQuantity,attr"`
	PassengerQuantity string `xml:"PassengerQuantity,attr"`
	VehType           VehType
	VehClass          VehClass
	VehMakeModel      VehMakeModel
	PictureURL        string `xml:"PictureURL"`
}
type VehType struct {
	VehicleCategory string `xml:"VehicleCategory,attr"`
}
type VehClass struct {
	Size string `xml:"Size,attr"`
}
type VehMakeModel struct {
	Code string `xml:"Code,attr"`
	Name string `xml:"Name,attr"`
}
