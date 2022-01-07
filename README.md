# logger
[![Go Reference](https://pkg.go.dev/badge/github.com/pravinba9495/logger.svg)](https://pkg.go.dev/github.com/pravinba9495/logger) ![Go Report Card](https://goreportcard.com/badge/github.com/pravinba9495/logger)

A minimalisitic logger module for basic logging, written in Go.

## Usage
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

## License
MIT
