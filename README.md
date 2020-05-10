# hasher

Library for easy hashing and verifying passwords.

Hashing algorithm is Argon2id.

## Usage

### Install

```fish
go get -u github.com/salifm/hasher
```

### Import

```go
import (
    "github.com/salifm/hasher"
)
```

### Hash

```go
var password string = "password"
hash, salt, err := hasher.Hash(password)
// save hash and salt to database
```

### Verify

```go
var password string = "password"
// get hash and salt from database
result, err := hasher.Verify(password, hash, salt)
if result {
    // The password is correct
}
```

## License

This library is licensed under the [MIT](./LICENSE) license