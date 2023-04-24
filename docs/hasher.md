# Hasher package

A decorator for the `golang.org/x/crypto/bcrypt` package which is providing easy access to functions.

## Functions

### func [HashPassword](https://github.com/iggyster/lets-go-chat/blob/main/pkg/hasher/hasher.go#L7)

Returns a hashed string encrypted with the bcrypt algorithm.

```go
func HashPassword(password string) (string, error)
```

### func [CheckPasswordHash](https://github.com/iggyster/lets-go-chat/blob/main/pkg/hasher/hasher.go#L15)

Returns `true` if the password match the hash or false if it's not.

```go
func CheckPasswordHash(password, hash string) bool
```