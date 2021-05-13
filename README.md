# <img align="right" src="https://avatars.githubusercontent.com/u/56905970?s=60&v=4" alt="alhamsya-gdifferento" title="alhamsya-different" /> Different

[![Sourcegraph](https://sourcegraph.com/github.com/alhamsya/different/-/badge.svg)](https://sourcegraph.com/github.com/alhamsya/different?badge)
[![Documentation](https://godoc.org/github.com/alhamsya/different?status.svg)](https://godoc.org/github.com/alhamsya/different)
[![codecov](https://codecov.io/gh/alhamsya/different/branch/master/graph/badge.svg?token=P7LSAI3V6S)](https://codecov.io/gh/alhamsya/different)
[![Awesome](https://cdn.rawgit.com/alhamsya/different/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/alhamsya/different)
[![rcard](https://goreportcard.com/badge/github.com/alhamsya/different)](https://goreportcard.com/report/github.com/alhamsya/different)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/alhamsya/different/master/LICENSE)

Library different value of struct.

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
