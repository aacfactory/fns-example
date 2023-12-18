package modules

import (
	"github.com/aacfactory/fns/services"
)

func Services() (v []services.Service) {
	v = append(
		dependencies(),
		endpoints()...,
	)
	return
}

func dependencies() (v []services.Service) {
	v = []services.Service{
		// add dependencies here
	}
	return
}
