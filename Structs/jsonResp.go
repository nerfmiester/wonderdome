package Structs

import "fmt"

type ProviderResp struct {
	Providers []Provider `json:"Provider"`
}

type Provider struct {
	Name        string       `json:"name"`
	TypVehicles []TypVehicle `json:"TypVehicle"`
}

type TypVehicle struct {
	AirConditionInd   bool        `json:"AirConditionInd"`
	BaggageQuantity   int         `json:"BaggageQuantity"`
	PassengerQuantity int         `json:"PassengerQuantity"`
	TransmissionType  string      `json:"TransmissionType"`
	VehicleCategory   int         `json:"VehicleCategory"`
	Size              int         `json:"Size"`
	VehMakeModel      VehMkeModel `json:"VehMakeModel"`
	PictureURL        string      `json:"PictureURL"`
}

func (vm *TypVehicle) SetAirConditionInd(b bool) {
	vm.AirConditionInd = b
}
func (vm *TypVehicle) SetBaggageQuantity(i int) {
	vm.BaggageQuantity = i
}
func (vm *TypVehicle) SetVehMkeModel(vmx VehMkeModel) {
	vm.VehMakeModel = vmx
}

type VehMkeModel struct {
	Code string `json:"Code"`
	Name string `json:"Name"`
}

func (vm *VehMkeModel) SetCode(s string) {
	vm.Code = s
}
func (vm *VehMkeModel) SetName(s string) {
	vm.Name = s
}

func (p *Provider) SetName(s string) {
	p.Name = s
}
func (p *Provider) SetTypVehicle(s TypVehicle, i int) {
	fmt.Println("whoops : ", s)
	p.TypVehicles[i] = s
}
