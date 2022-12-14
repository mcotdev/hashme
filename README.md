# hashme

An API that returns a sha256 hash from a given string.

# Usage
Run it behind a reverse proxy, and enter string after the endpoint `/hash/`

# Example

```
http://localhost:8080/hash/hello

Returns: 81b637d8fcd4c6da6359e6963113a1170de722e4b725b84d1e0b4cfd9ec58cf10
```

Todo:
- Add a salt for better security
- Modify so that it can be used to generate a password from a master password. Perhaps using a salt as a master password.
- Store generated hash etc to a database
- Refactor to Rust or C to improve speed.
