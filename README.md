# dialog [![Build Status](https://travis-ci.com/xyproto/dialog.svg?branch=master)](https://travis-ci.com/xyproto/dialog) [![Go Report Card](https://goreportcard.com/badge/github.com/xyproto/dialog)](https://goreportcard.com/report/github.com/xyproto/dialog) [![License](https://img.shields.io/badge/License-MIT-brightgreen)](https://raw.githubusercontent.com/xyproto/dialog/master/LICENSE)

Simple wrapper for the `dialog` executable.

## Online documentation

* [godoc.org](https://godoc.org/github.com/xyproto/dialog)
* [pkg.go.dev](https://pkg.go.dev/github.com/xyproto/dialog?tab=doc)

## Screenshots

![screenshot 1](img/dialog_screenshot1.png)
![screenshot 2](img/dialog_screenshot2.png)

## Features and limitations

Supports only these types:

* Message box
* Yes/No box
* Menu box

## Example use

```go
package main

import (
	"fmt"
	"github.com/xyproto/dialog"
)

func main() {
	d := dialog.New(80, 24)
	answeredYes, err := d.YesNo("Do you want cake?")
	if err != nil {
		panic(err)
	}
	if answeredYes {
		fmt.Println("You answered: yes")
		fmt.Println("The cake is a lie. Haha!")
	} else {
		fmt.Println("You answered: no")
		fmt.Println("Fine.")
	}
}
```

## Requirements

* Go >= 1.10

## General info

* Version: 1.1.0
* License: MIT
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
