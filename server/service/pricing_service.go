package service

type PricingService struct {
}

func (s *PricingService) GetSimplePrice(ticker string) float64 {
	return 150.0
}
