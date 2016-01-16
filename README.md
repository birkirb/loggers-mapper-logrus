# loggers-mapper-logrus
Golang [Loggers](https://gopkg.in/birkirb/loggers.v1) mapper for [Logrus](https://github.com/Sirupsen/logrus).

[![GoDoc](https://godoc.org/github.com/birkirb/loggers-mapper-logrus?status.svg)](https://godoc.org/github.com/birkirb/loggers-mapper-logrus)
[![Build Status](https://travis-ci.org/birkirb/loggers-mapper-logrus.svg?branch=master)](http://travis-ci.org/birkirb/loggers-mapper-logrus)

## Pre-recquisite

See https://gopkg.in/birkirb/loggers.v1

## Installation

    go get github.com/birkirb/loggers-mapper-logrus

## Usage

Assuming you are using loggers in your code, and you want to use logrus as your logger implementation. Start by configuring your logrus, and then pass it to the mapper and assign it to your loggers interface (embedded use) or the log.Logger (direct package).

### Example

```Go
package main

import (
	"os"

	"github.com/Sirupsen/logrus"
    "gopkg.in/birkirb/loggers.v1"
    mapper "github.com/birkirb/loggers-mapper-logrus/"
)

// Log is my default logger.
var Log loggers.Contextual

func main() {
	l := logrus.New()
	l.Out = os.Stdout
	l.Level = logrus.InfoLevel

	m := mapper.NewLogger(l)
	Log = &m

	Log.Info("My program has started")
}
```
