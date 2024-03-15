package controller

type Controller struct {
	HeroPoolController
	HeroBaseController
}

func New() *Controller {
	Controller := &Controller{}
	return Controller
}
