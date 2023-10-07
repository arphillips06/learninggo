package main

import (
	"fmt"

	"github.com/scrapli/scrapligo/driver/generic"
	"github.com/scrapli/scrapligo/driver/options"
)

func main() {
	d, err := generic.NewDriver(
		"192.168.69.230",
		options.WithAuthNoStrictKey(),
		options.WithAuthUsername("lab"),
		options.WithAuthPassword("lab123"),
		options.WithPort(22),
		options.WithTransportType("standard"),
	)

	if err != nil {
		fmt.Printf("failed to create driver; error %+v\n", err)
		return
	}

	err = d.Open()
	if err != nil {
		fmt.Printf("failed to open driver; error: %+v\n", err)
		return
	}
	defer d.Close()

	prompt, err := d.GetPrompt()
	if err != nil {
		fmt.Printf("failed to get prompt; error: %+v\n", err)
		return
	}

	fmt.Printf("Found prompt: %s\n\n\n", prompt)

	output, err := d.Channel.SendInput("show configuration")
	if err != nil {
		fmt.Printf("failed to send input to device; error: %+v\n", err)
		return
	}

	fmt.Printf("output received (SendInput): \n %s\n\n\n", output)

}
