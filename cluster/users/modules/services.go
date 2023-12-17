package modules

import (
	"github.com/aacfactory/fns-contrib/databases/postgres"
	"github.com/aacfactory/fns/services"
	_ "github.com/lib/pq"
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
		postgres.New(),
	}
	return
}
