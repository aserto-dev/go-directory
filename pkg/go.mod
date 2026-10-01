module github.com/aserto-dev/go-directory/pkg

go 1.27.1

replace github.com/aserto-dev/go-directory/aserto => ../aserto

require (
	github.com/aserto-dev/errors v0.34.1
	github.com/aserto-dev/go-directory/aserto v0.0.0-00010101000000-000000000000
	github.com/go-http-utils/headers v0.0.0-20181008091004-fed159eddc2a
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.31.0
	github.com/pkg/errors v0.9.1
	github.com/samber/lo v1.53.0
	github.com/stretchr/testify v1.12.1
	google.golang.org/grpc v1.84.0
	google.golang.org/protobuf v1.36.12
)

require (
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10 // indirect
	github.com/rs/zerolog v1.35.1 // indirect
	go.yaml.in/yaml/v3 v3.0.5 // indirect
	golang.org/x/net v0.59.0 // indirect
	golang.org/x/sys v0.48.0 // indirect
	golang.org/x/text v0.42.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20260928230214-8a89bd6388cc // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260921155816-b14227669459 // indirect
)
