module github.com/aserto-dev/go-directory

go 1.26

toolchain go1.26.7

// replace github.com/aserto-dev/errors => ../errors

require (
	github.com/aserto-dev/errors v0.34.1-0.20260824131513-21f43c970ba0
	github.com/go-http-utils/headers v0.0.0-20181008091004-fed159eddc2a
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.30.0
	github.com/pkg/errors v0.9.1
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10
	github.com/samber/lo v1.53.0
	github.com/stretchr/testify v1.12.1
	google.golang.org/genproto/googleapis/api v0.0.0-20260819154853-08b0e4226688
	google.golang.org/grpc v1.83.1
	google.golang.org/protobuf v1.36.12
)

require (
	github.com/mattn/go-colorable v0.1.15 // indirect
	github.com/mattn/go-isatty v0.0.24 // indirect
	github.com/rs/zerolog v1.35.1 // indirect
	go.yaml.in/yaml/v3 v3.0.5 // indirect
	golang.org/x/net v0.58.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.41.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260819154853-08b0e4226688 // indirect
)
