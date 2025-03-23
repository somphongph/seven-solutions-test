package services

type Servicer interface {
	CountBeefSummary(str string) BeefCount
}

type service struct {
}

func NewServices() *service {
	return &service{}
}
