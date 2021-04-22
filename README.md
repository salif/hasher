# hasher

Library and cli tool for password hashing

The hashing algorithm is Argon2id.

The result is:

* Hash: hexadecimal string with length of **64**
* Salt: hexadecimal string with length of **32**
* Version: int

## Usage

### Get

```sh
go get -u salif.eu/go/hasher
```

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

```sh
go get -u salif.eu/go/hasher/cmd/hasher
```

### Hash

```sh
echo -n "password" | hasher
```

#### or

```sh
hasher "filename"
```

### Verify

```sh
echo -n "password" | hasher "hash" "salt" 2
```

#### or

```sh
hasher "filename" "hash" "salt" 2
```

## License

This library is licensed under the [MIT](./LICENSE) license
