package modules

import (
	"github.com/aacfactory/fns-contrib/databases/sql"
	"github.com/aacfactory/fns/service"
	_ "github.com/lib/pq"
)

func dependencies() (services []service.Service) {
	services = append(
		services,
		sql.Service(),
	)
	return
}
