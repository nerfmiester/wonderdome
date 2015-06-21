package Structs

type ProviderResp struct {
	Providers []Provider `json:"Provider"`
}

type Provider struct {
	Name string `json:"name"`
}

func (p *Provider) SetName(s string) {
	p.Name = s
}
