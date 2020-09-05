# go-hms-push

[![GoDoc](https://godoc.org/github.com/icecream78/go-hms-push?status.svg)](https://godoc.org/github.com/icecream78/go-hms-push)

This project was reworked version of [Huawei demo pack](https://github.com/HMS-Core/hms-push-serverdemo-go).

Golang client library for APIs of the HUAWEI Push Kit server. Implemented only [HTTP client](https://developer.huawei.com/consumer/en/doc/development/HMS-References/push-sendapi#h1-1576155232538).

More information on [Huawei HMS Core](https://developer.huawei.com/consumer/en/doc/overview/HMS-4-0)

## Getting Started

To install go-hms-push, use `go get`:

```bash
go get github.com/icecream78/go-hms-push
```

## Sample Usage

Here is a simple example illustrating how to use HMS push library:

```go
package main

import (
	"context"
	"log"

	hms "github.com/icecream78/go-hms-push"
)

const (
	appId     string = "xxxxxx"
	appSecret string = "xxxxxx"
	clientToken  string = "xxxxxx"
)

func main() {
	client, err := hms.NewHuaweiClient(appId, appSecret)
	if err != nil {
		log.Fatal(err)
	}

	msg := hms.GetDefaultAndroidNotificationMessage([]string{clientToken})

	resp, err := client.SendMessage(context.Background(), msg)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v\n", resp)
}

```
