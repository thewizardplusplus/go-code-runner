# go-code-runner

[![GoDoc](https://godoc.org/github.com/thewizardplusplus/go-code-runner?status.svg)](https://godoc.org/github.com/thewizardplusplus/go-code-runner)
[![Go Report Card](https://goreportcard.com/badge/github.com/thewizardplusplus/go-code-runner)](https://goreportcard.com/report/github.com/thewizardplusplus/go-code-runner)
[![Build Status](https://travis-ci.org/thewizardplusplus/go-code-runner.svg?branch=master)](https://travis-ci.org/thewizardplusplus/go-code-runner)
[![codecov](https://codecov.io/gh/thewizardplusplus/go-code-runner/branch/master/graph/badge.svg)](https://codecov.io/gh/thewizardplusplus/go-code-runner)

## Installation

Prepare the directory:

```
$ mkdir --parents "$(go env GOPATH)/src/github.com/thewizardplusplus/"
$ cd "$(go env GOPATH)/src/github.com/thewizardplusplus/"
```

Clone this repository:

```
$ git clone https://github.com/thewizardplusplus/go-code-runner.git
$ cd go-code-runner
```

Install dependencies with the [dep](https://golang.github.io/dep/) tool:

```
$ dep ensure -vendor-only
```

## License

The MIT License (MIT)

Copyright &copy; 2021 thewizardplusplus
