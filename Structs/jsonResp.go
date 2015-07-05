package Structs

type ProviderResp struct {
	Providers []Provider `json:"Provider"`
}

type Provider struct {
	Name        string       `json:"name"`
	TypVehicles []TypVehicle `json:"TypVehicle"`
	InfoM       InfoM        `json:"Info"`
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
	Rates             Rates       `json:"Rates"`
}

type Rates struct {
	UnitDistance          string    `json:"UnitDistance"`
	UnLimited             bool      `json:"Unlimited"`
	VehiclePeriodUnitName string    `json:"VehiclePeriodUnitName"`
	Charges               []Charge  `json:"Charges"`
	Qualifier             Qualifier `json:"Qualifier"`
}
type Qualifier struct {
	CorpDiscountNmbr string `json:"CorpDiscountNmbr"`
	RateQualifier    string `json:"RateQualifier"`
	RatePeriod       string `json:"RatePeriod"`
}

func (r *Rates) SetUnitDistance(s string) {
	r.UnitDistance = s
}
func (r *Rates) SetUnLimited(b bool) {
	r.UnLimited = b
}
func (r *Rates) SetVehiclePeriodUnitName(s string) {
	r.VehiclePeriodUnitName = s
}
func (r *Rates) SetQualifier(s Qualifier) {
	r.Qualifier = s
}
func (r *Rates) SetCharges(s []Charge) {
	r.Charges = s
}
func (q *Qualifier) SetCorpDiscountNmbr(s string) {
	q.CorpDiscountNmbr = s
}
func (q *Qualifier) SetRateQualifier(s string) {
	q.RateQualifier = s
}
func (q *Qualifier) SetRatePeriod(s string) {
	q.RatePeriod = s
}

type Charge struct {
	Amount       float64 `json:"Amount"`
	CurrencyCode string  `json:"CurrencyCode"`
	TaxInclusive bool    `json:"TaxInclusive"`
	Purpose      int     `json:"Purpose"`
	Description  string  `json:"Description"`
}

func (vm *TypVehicle) SetAirConditionInd(b bool) {
	vm.AirConditionInd = b
}
func (vm *TypVehicle) SetBaggageQuantity(i int) {
	vm.BaggageQuantity = i
}
func (vm *TypVehicle) SetPassengerQuantity(i int) {
	vm.PassengerQuantity = i
}
func (vm *TypVehicle) SetTransmissionType(s string) {
	vm.TransmissionType = s
}
func (vm *TypVehicle) SetPictureURL(s string) {
	vm.PictureURL = s
}
func (vm *TypVehicle) SetVehicleCategory(i int) {
	vm.VehicleCategory = i
}
func (vm *TypVehicle) SetSize(i int) {
	vm.Size = i
}
func (vm *TypVehicle) SetRates(s Rates) {
	vm.Rates = s
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
func (p *Provider) GetProvider() (prov *Provider) {
	return p
}
func (p *Provider) SetTypVehicle(s []TypVehicle) {
	p.TypVehicles = s
}
func (p *Provider) SetInfoM(in InfoM) {
	p.InfoM = in
}

type InfoM struct {
	LocationDetails LocDetails
	TPA_Extensions  TPA_ExtensionsM
}

func (i *InfoM) SetLocDetails(l LocDetails) {
	i.LocationDetails = l
}

type TPA_ExtensionsM struct {
	RentalDuration int    `json:"RentalDuration"`
	Type           string `json:"Type"`
	CustDropOff    string `json:"CustDropOff"`
	CustPickUp     string `json:"CustPickUp"`
}
type LocDetails struct {
	VehRentLocInfo []VehRentLocInf
}

func (l *LocDetails) SetVehRentLocInfo(v []VehRentLocInf) {
	l.VehRentLocInfo = v
}

type VehRentLocInf struct {
	Title string `json:"Title"`
	Type  int    `json:"Type"`
	Text  string `json:"Text"`
}

func (v *VehRentLocInf) SetTitle(s string) {
	v.Title = s
}
func (v *VehRentLocInf) SetType(s int) {
	v.Type = s
}
func (v *VehRentLocInf) SetText(s string) {
	v.Text = s
}
