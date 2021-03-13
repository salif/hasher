# hasher

Library and cli tool for password hashing

The hashing algorithm is Argon2id.

The result is:

* Hash: hexadecimal string with length of **64**
* Salt: hexadecimal string with length of **32**
* Version: int

## Usage

### Import

```go
import (
    "salif.eu/go/hasher"
)
```

### Hash

```go
var password = "password"
var hash, salt, version = hasher.Hash(password)
// save hash, salt and version to database
```

### Verify

```go
var password = "password"
// get hash, salt and version from database
var ok = hasher.Verify(password, hash, salt, version)
if ok {
    // The password is correct
}
```

## CLI tool

### Install

```fish
go get -u salif.eu/go/hasher/cmd/hasher
```

### Hash

```fish
hasher "password"
```

#### or

```fish
echo -n "password" | hasher
```

### Verify

```fish
hasher "password" "hash" "salt"
```

#### or

```fish
echo -n "password" | hasher "hash" "salt"
```

## License

This library is licensed under the [MIT](./LICENSE) license
