module github.com/aacfactory/fns-example/standalone

go 1.17

replace (
	github.com/aacfactory/fns v0.3.0 => ../../fns
	github.com/aacfactory/fns-contrib/authorizations/jwt v0.3.0 => ../../fns-contrib/authorizations/jwt
	github.com/aacfactory/fns-contrib/databases/sql v0.3.0 => ../../fns-contrib/databases/sql
)

require (
	github.com/aacfactory/configuares v1.2.2
	github.com/aacfactory/errors v1.5.0
	github.com/aacfactory/fns v0.2.0
	github.com/aacfactory/fns-contrib/authorizations/jwt v0.0.0-20210904003123-1e4b4c8b7f02
	github.com/aacfactory/fns-contrib/databases/sql v0.0.0-20210904003123-1e4b4c8b7f02
	github.com/aacfactory/json v1.4.2
	github.com/lib/pq v1.10.3
)

require (
	github.com/aacfactory/logs v1.1.3 // indirect
	github.com/aacfactory/workers v1.2.0 // indirect
	github.com/andybalholm/brotli v1.0.3 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgraph-io/ristretto v0.1.0 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/fatih/color v1.12.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/goccy/go-yaml v1.9.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.0.0 // indirect
	github.com/golang/glog v1.0.0 // indirect
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/klauspost/compress v1.13.5 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rs/xid v1.3.0 // indirect
	github.com/rs/zerolog v1.24.0 // indirect
	github.com/tidwall/gjson v1.9.0 // indirect
	github.com/tidwall/match v1.0.3 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/tidwall/sjson v1.2.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.29.0 // indirect
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
	golang.org/x/sys v0.0.0-20210806184541-e5e7981a1069 // indirect
	golang.org/x/text v0.3.6 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
)
