# logger
[![Go Reference](https://pkg.go.dev/badge/github.com/pravinba9495/logger.svg)](https://pkg.go.dev/github.com/pravinba9495/logger) ![Go Report Card](https://goreportcard.com/badge/github.com/pravinba9495/logger) ![Issues](https://img.shields.io/github/issues-raw/pravinba9495/logger) ![License](https://img.shields.io/github/license/pravinba9495/logger) ![Release](https://img.shields.io/github/v/release/pravinba9495/logger?include_prereleases)

A minimalisitic logger module for basic logging, written in Go.

## Table of Contents
- [Example](#example)
- [Documentation](#documentation)
- [Development](#development)
- [Maintainers](#maintainers)
- [License](#license)

## Example
### Client
```go
package main

import (
    "github.com/pravinba9495/logger"
)

func main() {
    // Initialize logger
    opts := &logger.LoggerOptions{
	LogLevel:    logger.LevelTrace,
	LogFilePath: "logfile.txt",
    }
    close, err := logger.Init(opts)
    if err != nil {
        panic(err)
    }
    defer close()

    logger.Trace("This is just a trace");            // Prints on stdout with a cyan colored text
    logger.Print("This is just a simple log");       // Prints on stdout with a white colored text
    logger.Warn("This is just a warning log");       // Prints on stdout with a yellow colored text
    logger.Success("This is just a success log");    // Prints on stdout with a green colored text
    logger.Error("This is just an error log");       // Prints on stdout with a red colored text
}
```

## Documentation
Logger documentation is hosted at [Read the docs](https://pkg.go.dev/github.com/pravinba9495/logger).

## Development
Logger is still under development. Contributions are always welcome!

## Maintainers
* [@pravinba9495](https://github.com/pravinba9495)
## License
MIT
