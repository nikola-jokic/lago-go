# Lago Go Client

This is a Go client for Lago API

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://spdx.org/licenses/MIT.html)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/nikola-jokic/lago-go)](https://pkg.go.dev/github.com/nikola-jokic/lago-go)

## Installation

```shell
go get github.com/nikola-jokic/lago-go
```

<!-- ## Usage -->
<!-- TODO: add usage example once the API is stable -->

## Documentation

The Lago documentation is available at [doc.getlago.com](https://doc.getlago.com/docs/api/intro).

## License

Lago Go client is distributed under [MIT license](LICENSE).

## Disclaimer

This repository is heavily under development and is a modified version of [lago-go-client](https://github.com/getlago/lago-go-client)

It is heavily under development and is not recommended for production use. Once the code changes are stabilized, the stable release will be issued.

The motivation behind this repository is to:
- Avoid using resty as a client. Users should be able to pick the client they want.
- Faster bug fixes. This client is used for [BountyHub](https://bountyhub.org) and bug fixes are needed to be done faster. Company maintainers are most likely overwhelmed with other tasks and this repository is a way to speed up the process.
