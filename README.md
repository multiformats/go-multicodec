# go-multicodec

![Multicodecs update](https://github.com/multiformats/go-multicodec/workflows/Multicodecs%20update/badge.svg) [![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg)](https://github.com/RichardLitt/standard-readme) [![Go Report Card](https://goreportcard.com/badge/multiformats/go-multicodec)](https://goreportcard.com/report/multiformats/go-multicodec)

> Periodically self-generated Go constants of [multicodecs](https://github.com/multiformats/multicodec) used by the [multiformats](https://github.com/multiformats/multiformats) projects.

## Table of Contents

- [Workflow](#workflow)
- [Install](#install)
- [Usage](#usage)
- [Generator](#generator)
- [Maintainers](#maintainers)
- [Contribute](#contribute)
- [License](#license)

## Workflow

Every night a few minutes after midnight (UTC) a GitHub-Action fetches the latest [multicodecs table](https://raw.githubusercontent.com/multiformats/multicodec/master/table.csv), generates the constants, commits possible changes and creates a pull request. It will update the same pull request if subsequent runs find different changes and it won't create a pull request if no changes were detected.

## Install

`go-multicodec` is a standard Go module which can be installed with:

```sh
go get github.com/multiformats/go-multicodec
```

## Usage

```go
package main

import "github.com/multiformats/go-multicodec"

func main() {
    _ = multicodec.Sha2_256
}
```

The corresponding `name` value for each codec from the [multicodecs table](https://raw.githubusercontent.com/multiformats/multicodec/master/table.csv) can be accessed via its `String()` method. E.g. `multicodec.Sha2_256.String()` will return `sha2-256`.

## Generator

To generate the constants yourself checkout the repository like

```shell
git clone https://github.com/multiformats/go-multicodec.git
```

Then run `go generate` in the root of the repository. Currently this will invoke (see [`init.go`](./init.go)):

```shell
go run ./gen/gen.go
gofmt -w codec.go
stringer -type=Codec -linecomment
```

Note: You may need to install `stringer` via `go install golang.org/x/tools/cmd/stringer`.

## Maintainers

[@mvdan](https://github.com/mvdan).

## Contribute

Contributions welcome. Please check out [the issues](https://github.com/multiformats/go-multicodec/issues).

Check out our [contributing document](https://github.com/multiformats/multiformats/blob/master/contributing.md) for more information on how we work, and about contributing in general. Please be aware that all interactions related to multiformats are subject to the IPFS [Code of Conduct](https://github.com/ipfs/community/blob/master/code-of-conduct.md).

Small note: If editing the README, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

[MIT](LICENSE) © 2014 Juan Batiz-Benet