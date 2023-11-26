module github.com/go-orb/plugins/registry/nats

go 1.21.4

require (
	github.com/go-orb/go-orb v0.0.0-20231126172125-9f19ded9b372
	github.com/go-orb/plugins/log/slog v0.0.0-20230726151712-9ff06382f83b
	github.com/go-orb/plugins/registry/tests v0.0.0-20230713091520-67e7b5a34489
	github.com/google/uuid v1.4.0
	github.com/nats-io/nats.go v1.31.0
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.8.4
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/klauspost/compress v1.17.3 // indirect
	github.com/nats-io/nkeys v0.4.6 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	golang.org/x/crypto v0.15.0 // indirect
	golang.org/x/exp v0.0.0-20231110203233-9a3e6036ecaa // indirect
	golang.org/x/sys v0.14.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/go-orb/plugins/registry/tests => ../tests

replace github.com/go-orb/plugins/log/slog => ../../log/slog
