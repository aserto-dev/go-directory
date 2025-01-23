module github.com/aserto-dev/go-directory

go 1.22.11

toolchain go1.23.5

replace github.com/bufbuild/protovalidate-go => github.com/bufbuild/protovalidate-go v0.7.3

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.36.3-20241127180247-a33202765966.1
	github.com/aserto-dev/errors v0.0.13
	github.com/bufbuild/protovalidate-go v0.7.3
	github.com/go-http-utils/headers v0.0.0-20181008091004-fed159eddc2a
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.26.0
	github.com/pkg/errors v0.9.1
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10
	github.com/samber/lo v1.47.0
	github.com/stretchr/testify v1.10.0
	google.golang.org/genproto/googleapis/api v0.0.0-20250122153221-138b5a5a4fd4
	google.golang.org/grpc v1.70.0
	google.golang.org/protobuf v1.36.3
)

require (
	cel.dev/expr v0.19.0 // indirect
	github.com/antlr4-go/antlr/v4 v4.13.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/cel-go v0.22.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rs/zerolog v1.33.0 // indirect
	github.com/stoewer/go-strcase v1.3.0 // indirect
	golang.org/x/exp v0.0.0-20241210194714-1829a127f884 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250122153221-138b5a5a4fd4 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
