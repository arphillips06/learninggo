package main

import (
	"fmt"

	"github.com/scrapli/scrapligo/driver/netconf"
	"github.com/scrapli/scrapligo/driver/options"
)

func main() {

	GET_CONFIG := "candidate"

	driver, err := netconf.NewDriver(
		"192.168.69.230",
		options.WithAuthNoStrictKey(),
		options.WithAuthUsername("lab"),
		options.WithAuthPassword("lab123"),
		options.WithPort(830),
		options.WithTransportType("standard"),
		options.WithNetconfForceSelfClosingTags(),
	)

	if err != nil {
		fmt.Printf("failed to create driver; error: %+v\n", err)
		return
	}
	err = driver.Open()
	if err != nil {
		fmt.Printf("failed to open driver; error: %+v\n", err)
		return
	}
	defer driver.Close()

	running, err := driver.GetConfig(GET_CONFIG)
	if err != nil {
		fmt.Printf("failed executing Get SW info; error: %+v\n", err)
		return
	}
	if running.Failed != nil {
		fmt.Printf("response object indicates failure: %+v\n", running.Failed)
		return
	}

	fmt.Printf("Result: %s", running.Result)
}
