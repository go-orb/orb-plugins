module github.com/go-orb/plugins/client/tests

go 1.21.4

require (
	github.com/go-orb/go-orb v0.0.0-20231126172125-9f19ded9b372
	github.com/go-orb/plugins/codecs/jsonpb v0.0.0-20231126111023-0a8c6d6cb2ee
	github.com/go-orb/plugins/codecs/proto v0.0.0-20231126111023-0a8c6d6cb2ee
	github.com/go-orb/plugins/codecs/yaml v0.0.0-20231126111023-0a8c6d6cb2ee
	github.com/go-orb/plugins/config/source/cli/urfave v0.0.0-20231126111023-0a8c6d6cb2ee
	github.com/go-orb/plugins/config/source/file v0.0.0-20231126111023-0a8c6d6cb2ee
	github.com/go-orb/plugins/log/slog v0.0.0-20231126111023-0a8c6d6cb2ee
	github.com/go-orb/plugins/registry/mdns v0.0.0-20231126111023-0a8c6d6cb2ee
	github.com/go-orb/plugins/server/http v0.0.0-20231126111023-0a8c6d6cb2ee
	github.com/google/wire v0.5.0
	github.com/hashicorp/consul/sdk v0.15.0
	github.com/stretchr/testify v1.8.4
	golang.org/x/exp v0.0.0-20231110203233-9a3e6036ecaa
	google.golang.org/genproto/googleapis/api v0.0.0-20231120223509-83a465c0220f
	google.golang.org/grpc v1.59.0
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.3 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/go-chi/chi v1.5.5 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/subcommands v1.2.0 // indirect
	github.com/google/uuid v1.4.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/miekg/dns v1.1.57 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/urfave/cli/v2 v2.25.7 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/net v0.18.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/tools v0.15.0 // indirect
	google.golang.org/genproto v0.0.0-20231120223509-83a465c0220f // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231120223509-83a465c0220f // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/go-orb/plugins/codecs/jsonpb => ../../codecs/jsonpb

replace github.com/go-orb/plugins/codecs/proto => ../../codecs/proto

replace github.com/go-orb/plugins/codecs/yaml => ../../codecs/yaml

replace github.com/go-orb/plugins/config/source/file => ../../config/source/file

replace github.com/go-orb/plugins/server/http => ../../server/http

replace github.com/go-orb/plugins/registry/mdns => ../../registry/mdns

replace github.com/go-orb/plugins/log/slog => ../../log/slog

replace github.com/go-orb/plugins/config/source/cli/urfave => ../../config/source/cli/urfave

replace github.com/go-orb/plugins/registry/regutil => ../../registry/regutil
