package controller

type Instance struct {
}

type Controller interface {
	SetInstance(instance Instance) error
	GetInstance() *Instance
}
