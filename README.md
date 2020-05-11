# hasher

Library for easy hashing and verifying passwords.

Hashing algorithm is Argon2id.

The parameters are:

* Memory: **65536** B
* Iterations: **1**
* Parallelism: **4**
* Salt length: **16**


## Usage

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
match, err := hasher.Verify(password, hash, salt)
if match {
    // The password is correct
}
```

## License

This library is licensed under the [MIT](./LICENSE) license

