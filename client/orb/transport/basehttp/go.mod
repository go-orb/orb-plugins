module github.com/go-orb/plugins/client/orb/transport/basehttp

go 1.20

require (
	github.com/go-orb/go-orb v0.0.0-20231119181816-8fb44c1953fd
	github.com/go-orb/plugins/client/orb v0.0.0-20231124134616-38fd310569fc
)

require golang.org/x/exp v0.0.0-20231110203233-9a3e6036ecaa // indirect

replace github.com/go-orb/plugins/client/orb => ../..
