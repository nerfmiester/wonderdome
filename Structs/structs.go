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
	Info      Info      `xml:"Info"`
}
type Info struct {
	LocationDetails LocationDetails
	TPA_Extensions  TPA_Extensions
}

type LocationDetails struct {
	AdditionalInfo AdditionalInfo
}

type AdditionalInfo struct {
	VehRentLocInfos VehRentLocInfos
}
type VehRentLocInfos struct {
	VehRentLocInfo []VehRentLocInfo
}
type VehRentLocInfo struct {
	Title      string `xml:"Title,attr"`
	Type       int    `xml:"Type,attr"`
	SubSection SubSection
}

type SubSection struct {
	Paragraph Paragraph
}

type Paragraph struct {
	Text string `xml:"Text"`
}

type TPA_Extensions struct {
	TPA_Extensions_Inf  TPA_Extensions_Inf
	TPA_Extension_Flags TPA_Extension_Flags
}

type TPA_Extensions_Inf struct {
	RentalDuration int `xml:"RentalDuration,attr"`
}

type TPA_Extension_Flags struct {
	Type        string `xml:"Type,attr"`
	CustDropOff string `xml:"CustDropOff,attr"`
	CustPickUp  string `xml:"CustPickUp,attr"`
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
	XMLName    xml.Name `xml:"VehAvailCore"`
	Status     string   `xml:"Status,attr"`
	Vehicle    Vehicle
	RentalRate RentalRate
}
type VehAvailInfo struct {
	XMLName xml.Name `xml:"VehAvailInfo"`
}
type Vehicle struct {
	TransmissionType  string `xml:"TransmissionType,attr"`
	AirConditionInd   bool   `xml:"AirConditionInd,attr"`
	BaggageQuantity   int    `xml:"BaggageQuantity,attr"`
	PassengerQuantity int    `xml:"PassengerQuantity,attr"`
	VehType           VehType
	VehClass          VehClass
	VehMakeModel      VehMakeModel
	PictureURL        string `xml:"PictureURL"`
}
type VehType struct {
	VehicleCategory int `xml:"VehicleCategory,attr"`
}
type VehClass struct {
	Size int `xml:"Size,attr"`
}
type VehMakeModel struct {
	Code string `xml:"Code,attr"`
	Name string `xml:"Name,attr"`
}
type RentalRate struct {
	RateDistance   RateDistance
	VehicleCharges VehicleCharges
	RateQualifier  RateQualifier
}
type RateDistance struct {
	XMLName               xml.Name `xml:"RateDistance"`
	DistUnitName          string   `xml:"DistUnitName,attr"`
	VehiclePeriodUnitName string   `xml:"VehiclePeriodUnitName,attr"`
	Unlimited             bool     `xml:"Unlimited,attr"`
}
type VehicleCharges struct {
	XMLName       xml.Name        `xml:"VehicleCharges"`
	VehicleCharge []VehicleCharge `xml:"VehicleCharge"`
}
type VehicleCharge struct {
	Amount       float64 `xml:"Amount,attr"`
	CurrencyCode string  `xml:"CurrencyCode,attr"`
	TaxInclusive bool    `xml:"TaxInclusive,attr"`
	Purpose      int     `xml:"Purpose,attr"`
	Description  string  `xml:"Description,attr"`
}
type RateQualifier struct {
	CorpDiscountNmbr string `xml:"CorpDiscountNmbr,attr"`
	RateQualifier    string `xml:"RateQualifier,attr"`
	RatePeriod       string `xml:"RatePeriod,attr"`
}
