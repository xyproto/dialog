# dialog

Simple wrapper for the `dialog` executable.

[![GoDoc](https://godoc.org/github.com/xyproto/dialog?status.svg)](http://godoc.org/github.com/xyproto/dialog) [![License](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/xyproto/dialog/master/LICENSE) [![Report Card](https://img.shields.io/badge/go_report-A+-brightgreen.svg?style=flat)](http://goreportcard.com/report/xyproto/dialog)

![screenshot 1](img/dialog_screenshot1.png)
![screenshot 2](img/dialog_screenshot2.png)

## Example use

```go
package main

import (
	"fmt"
	"github.com/xyproto/dialog"
)

func main() {
	d := dialog.New(80, 20)
	d.MsgBox("hi")
	answer := d.Menu("hi", 6, map[string]string{"a": "A", "b": "B"})
	fmt.Printf("\n\n\n\n%s\n", answer)
}
```

The path to `dialog` is `/usr/bin/dialog` by default, but can be changed with the `SetPath` function.

## General info

* Version: 1.0
* License: MIT
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
