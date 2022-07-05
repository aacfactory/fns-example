package modules

import (
	_ "github.com/aacfactory/fns-contrib/authorizations/encoding/jwt"
	"github.com/aacfactory/fns-contrib/databases/sql"
	"github.com/aacfactory/fns/service"
	"github.com/aacfactory/fns/service/builtin/authorizations"
	_ "github.com/lib/pq"
)

func dependencies() (services []service.Service) {
	services = append(
		services,
		sql.Service(),
		authorizations.Service(),
	)
	return
}
