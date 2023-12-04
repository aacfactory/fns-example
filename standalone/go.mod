module github.com/aacfactory/fns-example/standalone

go 1.21.0

replace (
	github.com/aacfactory/fns v1.1.3 => ../../fns
	github.com/aacfactory/fns-contrib/databases/postgres v1.0.0 => ../../fns-contrib/databases/postgres
	github.com/aacfactory/fns-contrib/databases/sql v1.0.30 => ../../fns-contrib/databases/sql
)

require (
	github.com/aacfactory/errors v1.13.5
	github.com/aacfactory/fns v1.1.3
	github.com/aacfactory/fns-contrib/databases/postgres v1.0.0
	github.com/aacfactory/fns-contrib/databases/sql v1.0.30
	github.com/aacfactory/gcg v1.0.5
	github.com/lib/pq v1.10.9
)

require (
	github.com/aacfactory/afssl v1.10.0 // indirect
	github.com/aacfactory/cases v1.1.0 // indirect
	github.com/aacfactory/configures v1.12.1 // indirect
	github.com/aacfactory/json v1.16.5 // indirect
	github.com/aacfactory/logs v1.13.0 // indirect
	github.com/aacfactory/workers v1.8.4 // indirect
	github.com/andybalholm/brotli v1.0.6 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.3 // indirect
	github.com/dgrr/http2 v0.3.6-0.20231023141632-12370d352f5f // indirect
	github.com/fatih/color v1.16.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.16.0 // indirect
	github.com/goccy/go-yaml v1.11.2 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.3 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/rs/xid v1.5.0 // indirect
	github.com/rs/zerolog v1.31.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/tidwall/gjson v1.17.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tidwall/sjson v1.2.5 // indirect
	github.com/urfave/cli/v2 v2.25.7 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.51.0 // indirect
	github.com/valyala/fastrand v1.1.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	go.uber.org/automaxprocs v1.5.3 // indirect
	golang.org/x/crypto v0.16.0 // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sync v0.5.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/tools v0.16.0 // indirect
	golang.org/x/xerrors v0.0.0-20231012003039-104605ab7028 // indirect
)
