# jsst

[![Build Status][travis-image]][travis-url]
[![Doc][godoc-image]][godoc-url]
[![License][license-image]][license-url]

A CLI Generating Go struct file from JSON Hyper Schema (via [prmd](https://github.com/interagent/prmd)).

***WIP***

## Demo

```
# from file
$ jsst schema.json > struct.go
# from stdin
$ cat schema.json | jsst > struct.go
# change package
$ jsst schema.json -p schema > struct.go
```

## Installation

```
$ go get github.com/moqada/jsst
```

## Usage

```
usage: jsst [<flags>] [<file>]

Flags:
  -h, --help            Show context-sensitive help (also try --help-long and --help-man).
  -p, --package="main"  Package name for Go struct file
      --version         Show application version.

Args:
  [<file>]  Path of JSON Schema
```

Output Example: [./convertor_test.go](./convertor_test.go)

## Todo

- [ ] Add tests
- [ ] Support `anyOf`, `allOf`, `oneOf`
- [ ] Support Multiple types


[godoc-url]: https://godoc.org/github.com/moqada/jsst
[godoc-image]: https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square
[travis-url]: https://travis-ci.org/moqada/jsst
[travis-image]: https://img.shields.io/travis/moqada/jsst.svg?style=flat-square
[license-url]: http://opensource.org/licenses/MIT
[license-image]: https://img.shields.io/github/license/moqada/jsst.svg?style=flat-square
