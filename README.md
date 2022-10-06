# LBA

[![Run Tests](https://github.com/ductnn/lba/actions/workflows/test.yml/badge.svg)](https://github.com/ductnn/lba/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ductnn/lba)](https://goreportcard.com/report/github.com/ductnn/lba)

**lba** (LoadBalancer Algorithm) is a [go](https://go.dev/) implementation of the balancing algorithm.
- Round Robin
- IP Hash
- Least Connections
- **...updating...**

## Installation

First, install [Go](https://go.dev/doc/install), and install **lba** package:

```sh
go get -u github.com/ductnn/lba
```

Then, import package in your code:

```go
import "github.com/ductnn/lba"
```

## Example

- **Round Robin**:

```go
package main

import (
	"net/url"

	roundrobin "github.com/stsmdt/round-robin"
)

func main() {
	rr, _ := roundrobin.New(
		[]url.URL{
			{Host: "192.168.1.1"},
			{Host: "192.168.1.2"},
			{Host: "192.168.1.3"},
			{Host: "192.168.1.4"},
			{Host: "192.168.1.5"},
		},
	)

	rr.Next() // {Host: "192.168.1.1"}
	rr.Next() // {Host: "192.168.1.2"}
	rr.Next() // {Host: "192.168.1.3"}
	rr.Next() // {Host: "192.168.1.4"}
	rr.Next() // {Host: "192.168.1.5"}
	rr.Next() // {Host: "192.168.1.1"}
}
```

- **Least Connections**:

```go
lc, err := New([]*url.URL{
    {Host: "192.168.1.10"},
    {Host: "192.168.1.11"},
    {Host: "192.168.1.12"},
})

src1, done1 := lc.Next() // {Host: "192.168.1.10"}

src2, done2 := lc.Next() // {Host: "192.168.1.11"}

done1() // Reduce connection of src1

src3, done3 := lc.Next() // {Host: "192.168.1.10"}
```

- **IP Hash**:

```go
ip, _ := iphash.New([]*url.URL{
    {Host: "192.168.1.10"},
    {Host: "192.168.1.11"},
    {Host: "192.168.1.12"},
})

ip.Next(&url.URL{Host: "192.168.1.10"})  // {Host: "192.168.1.10"}
ip.Next(&url.URL{Host: "192.168.1.10"})  // {Host: "192.168.1.10"}
ip.Next(&url.URL{Host: "192.168.1.44"})  // {Host: "192.168.1.11"}
ip.Next(&url.URL{Host: "192.168.1.44"})  // {Host: "192.168.1.11"}
```
