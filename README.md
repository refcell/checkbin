## checkbin

[![Run Tests][tests]][tests-link]
[![Go Reference][ref]][ref-link]
[![Go Report Card][go-card]][go-card]

[tests]: https://github.com/refcell/checkbin/actions/workflows/ci.yml/badge.svg
[tests-link]: https://github.com/refcell/checkbin/actions/workflows/cli.yml
[ref]: https://pkg.go.dev/badge/github.com/refcell/checkbin/v0.svg
[ref-link]: https://pkg.go.dev/github.com/refcell/checkbin/v0
[go-card]: https://goreportcard.com/badge/github.com/refcell/checkbin/v1


Minimal binary to calculate the checksum for a binary in Golang.

### Installation

```sh
go install github.com/refcell/checkbin
```

### Usage

Get the checksum of a local binary file, for example `binary1`.

```sh
checkbin binary1
```

outputs

```
59bdccc87b9167828dbeb0e108a8150efa7a94d423f33fb1733eaee68900fe16
```

### Contributions

All contributions are welcome.

### License

[MIT License](./LICENSE)
