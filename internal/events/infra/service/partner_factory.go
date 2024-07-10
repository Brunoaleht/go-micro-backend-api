package service

import "fmt"

type PartnerFactory interface {
	GetPartner(partnerID int) (Partner, error)
}

type DefaultPartnerFactory struct {
	partnerBaseURls map[int]string
}

func NewPartnerFactory(partnerBaseURls map[int]string) PartnerFactory {
	return &DefaultPartnerFactory{
		partnerBaseURls: partnerBaseURls,
	}
}

func (f *DefaultPartnerFactory) GetPartner(partnerID int) (Partner, error) {
	baseURL, ok := f.partnerBaseURls[partnerID]
	if !ok {
		return nil, fmt.Errorf("partner with id %d not found", partnerID)
	}

	switch partnerID {
	case 1:
		return &Partner1{BaseURL: baseURL}, nil
	case 2:
		return &Partner2{BaseURL: baseURL}, nil
	default:
		return nil, fmt.Errorf("partner with id %d not found", partnerID)
	}

}
