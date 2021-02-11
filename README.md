# hasher

Library for easy hashing and verifying passwords.

The hashing algorithm is Argon2id.

The parameters are:

* Memory: **64** MiB
* Iterations: **1**
* Parallelism: **4**
* Salt length: **32**

The result is:

* Hash: base64 string with length of **43**
* Salt: base64 string with length of **43**

## Usage

### Import

```go
import (
    "github.com/salif/hasher"
)
```

### Hash

```go
password := "password"
hash, salt, err := hasher.Hash(password)
// save hash and salt to database if err is nil
```

### Verify

```go
password := "password"
// get hash and salt from database
match, err := hasher.Verify(password, hash, salt)
if err == nil && match {
    // The password is correct
}
```

## CLI tool

### Install

```fish
go get github.com/salif/hasher/cmd/hasher
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
