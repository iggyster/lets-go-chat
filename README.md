# Let's Go Chat

[![Go Reference](https://pkg.go.dev/badge/github.com/iggyster/lets-go-chat/pkg/hasher.svg)](https://pkg.go.dev/github.com/iggyster/lets-go-chat/pkg/hasher) [![Go Report Card](https://goreportcard.com/badge/github.com/iggyster/lets-go-chat)](https://goreportcard.com/report/github.com/iggyster/lets-go-chat)

This an educational project to learn Go lang. Not for production usage.

### Import

To import Hasher package: `go get github.com/iggyster/lets-go-chat/pkg/hasher`
To import Tokengen package: `go get github.com/iggyster/lets-go-chat/pkg/tokengen`

## Testing

Run all tests: `go test ./... -coverprofile c.out`
Output coverage in HTML: `go tool cover -html=c.out -o ./web/c.html`
Check overall coverage percentage: `go tool cover -func=c.out`
