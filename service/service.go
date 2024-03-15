package service

type Service struct {
	HeroPoolService
	HeroBaseService
}

func New() *Service {
	service := &Service{}
	return service
}
