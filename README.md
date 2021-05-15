# <img align="right" src="https://avatars.githubusercontent.com/u/56905970?s=60&v=4" alt="alhamsya-gdifferento" title="alhamsya-different" /> Different

[![Sourcegraph](https://sourcegraph.com/github.com/alhamsya/different/-/badge.svg)](https://sourcegraph.com/github.com/alhamsya/different?badge)
[![Documentation](https://godoc.org/github.com/alhamsya/different?status.svg)](https://godoc.org/github.com/alhamsya/different)
[![codecov](https://codecov.io/gh/alhamsya/different/branch/master/graph/badge.svg?token=P7LSAI3V6S)](https://codecov.io/gh/alhamsya/different)
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

## Benchmark

| Function     | CPU | Number of Operation | Time Required | Allocation Bytes | Allocation Times |
| ------------ | --- | ------------------- | ------------- | ---------------- | ---------------- |
| GenerateDiff | 1   | 419802              | 2392 ns/op    | 1440 B/op        | 28 allocs/op     |
| GenerateDiff | 4   | 429812              | 2392 ns/op    | 1440 B/op        | 28 allocs/op     |
| GenerateDiff | 8   | 427735              | 2451 ns/op    | 1440 B/op        | 28 allocs/op     |

```go
Showing nodes accounting for 1290ms, 56.09% of 2300ms total
Dropped 64 nodes (cum <= 11.50ms)
Showing top 20 nodes out of 142
      flat  flat%   sum%        cum   cum%
     200ms  8.70%  8.70%      740ms 32.17%  runtime.mallocgc
     140ms  6.09% 14.78%      140ms  6.09%  runtime.nextFreeFast (inline)
      90ms  3.91% 18.70%       90ms  3.91%  github.com/json-iterator/go.(*Stream).WriteStringWithHTMLEscaped
      90ms  3.91% 22.61%      170ms  7.39%  runtime.heapBitsSetType
      90ms  3.91% 26.52%      100ms  4.35%  runtime.scanobject
      70ms  3.04% 29.57%      670ms 29.13%  github.com/r3labs/diff/v2.(*Differ).diff
      70ms  3.04% 32.61%       90ms  3.91%  runtime.mapiternext
      60ms  2.61% 35.22%       60ms  2.61%  runtime.memclrNoHeapPointers
      60ms  2.61% 37.83%      500ms 21.74%  runtime.newobject
      50ms  2.17% 40.00%       50ms  2.17%  aeshashbody
      50ms  2.17% 42.17%       50ms  2.17%  indexbytebody
      50ms  2.17% 44.35%       50ms  2.17%  runtime.arenaIndex (inline)
      50ms  2.17% 46.52%      180ms  7.83%  runtime.growslice
      40ms  1.74% 48.26%       40ms  1.74%  runtime.futex
      30ms  1.30% 49.57%      550ms 23.91%  github.com/r3labs/diff/v2.(*Differ).diffStruct
      30ms  1.30% 50.87%       30ms  1.30%  github.com/r3labs/diff/v2.are
      30ms  1.30% 52.17%       30ms  1.30%  runtime.(*pallocBits).summarize
      30ms  1.30% 53.48%       40ms  1.74%  runtime.ifaceeq
      30ms  1.30% 54.78%      100ms  4.35%  runtime.mapaccess2
      30ms  1.30% 56.09%       50ms  2.17%  runtime.nilinterhash

```

## Supported Types

A diffable value can be/contain any of the following types:

| Type         | Supported |
| ------------ | --------- |
| struct       | ✔         |
| slice        | ✔         |
| string       | ✔         |
| int          | ✔         |
| bool         | ✔         |
| map          | ✔         |
| pointer      | ✔         |
| custom types | ✔         |

### Tags

In order for struct fields to be compared, they must be tagged with a given name. All tag values are prefixed with `diff`. i.e. `diff:"items"`.

| Tag           | Usage                                                                                                                                                                                                                                                                                           |
| ------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `-`           | Excludes a value from being diffed                                                                                                                                                                                                                                                              |
| `identifier`  | If you need to compare arrays by a matching identifier and not based on order, you can specify the `identifier` tag. If an identifiable element is found in both the from and to structures, they will be directly compared. i.e. `diff:"name, identifier"`                                     |
| `immutable`   | Will omit this struct field from diffing. When using `diff.StructValues()` these values will be added to the returned changelog. It's use case is for when we have nothing to compare a struct to and want to show all of its relevant values.                                                  |
| `nocreate`    | The default patch action is to allocate instances in the target strut, map or slice should they not exist. Adding this flag will tell patch to skip elements that it would otherwise need to allocate. This is separate from immutable, which is also honored while patching.                   |
| `omitunequal` | Patching is a 'best effort' operation, and will by default attempt to update the 'correct' member of the target even if the underlying value has already changed to something other than the value in the change log 'from'. This tag will selectively ignore values that are not a 100% match. |

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
