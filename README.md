# <img align="right" src="https://avatars.githubusercontent.com/u/56905970?s=60&v=4" alt="alhamsya-gdifferento" title="alhamsya-different" /> Different

[![Sourcegraph](https://sourcegraph.com/github.com/alhamsya/different/-/badge.svg)](https://sourcegraph.com/github.com/alhamsya/different?badge)
[![Documentation](https://godoc.org/github.com/alhamsya/different?status.svg)](https://godoc.org/github.com/alhamsya/different)
[![Awesome](https://cdn.rawgit.com/alhamsya/different/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/alhamsya/different)
[![rcard](https://goreportcard.com/badge/github.com/alhamsya/different)](https://goreportcard.com/report/github.com/alhamsya/different)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/alhamsya/different/master/LICENSE)

Go library that provides fuzzy string matching optimized for filenames and code symbols in the style of Sublime Text,
VSCode, IntelliJ IDEA et al. This library is external dependency-free. It only depends on the Go standard library.

## Installation

```bash
go get github.com/alhamsya/different
```

## Library Dependencies

| Type          | Supported     |
| ------------- | ------------- |
| diff          | r3labs        |
| json-iterator | json-iterator |

## Usage

The following example comparison value of two struct:

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/alhamsya/different"
)

type User struct {
	Name    string
	Age     int `diff:"-"`
	Address Address
	Telephone
}

type Address struct {
	City   string
	Street string
}

type Telephone struct {
	Type        string
	PhoneNumber string `diff:"phone_number"`
}

func main() {
	origin := User{
		Name: "Alhamsya",
		Age:  23,
		Address: Address{
			City:   "Kediri",
			Street: "New street",
		},
		Telephone: Telephone{
			Type:        "Rumah",
			PhoneNumber: "081",
		},
	}

	newData := User{
		Name: "Bintang",
		Age:  25,
		Address: Address{
			City:   "Kediri",
			Street: "Old Street",
		},
		Telephone: Telephone{
			Type:        "Rumah",
			PhoneNumber: "085",
		},
	}

	diff, err := different.GenerateDiff(origin, newData)
	if err != nil {
		panic(err)
	}

	fmt.Print(string(diff))
}
```

The result of comparison struct would be type data `[]byte` and could uses json unmarshal.

```json
[
  {
    "Name": {
      "after": "Bintang",
      "before": "Alhamsya"
    }
  },
  {
    "Address": {
      "Street": {
        "before": "New street",
        "after": "Old Street"
      }
    }
  },
  {
    "Telephone": {
      "phone_number": {
        "before": "081",
        "after": "085"
      }
    }
  }
]
```

## Contributing

Everyone is welcome to contribute. Please send me a pull request or file an issue. I promise
to respond promptly.

## License

The MIT License (MIT)

Copyright (c) 2021 Alhamsya Bintang Dyasta

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
