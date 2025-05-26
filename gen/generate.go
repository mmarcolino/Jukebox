package gen

//go:generate go run github.com/ogen-go/ogen/cmd/ogen --package openapi --target ./openapi --clean ../api/openapi.yaml
//go:generate go run github.com/sqlc-dev/sqlc/cmd/sqlc generate --no-remote -f ../sqlc.yaml
