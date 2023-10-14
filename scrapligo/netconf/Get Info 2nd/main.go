package main

import (
	"fmt"
	"os"

	"github.com/scrapli/scrapligo/driver/netconf"
	"github.com/scrapli/scrapligo/driver/options"
	"gopkg.in/yaml.v3"
)

type Host struct {
	Name string   `yaml:"name"`
	IP   string   `yaml:"ip"`
	Cmd  []string `yaml:"cmds"`
}

// read the file passed from main and return a slice of the Host Struct and an error value
func readHosts(file string) ([]Host, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Failed to read hosts from file; error: %v\n", err)
	}

	var hostList []Host
	err = yaml.Unmarshal(data, &hostList)
	if err != nil {
		fmt.Printf("Failed to parse hosts from YAML; error: %v\n", err)
	}

	//fmt.Printf("%v\n", hostList)

	return hostList, nil

}

func createDriver(hostList []Host) (*netconf.Driver, error) {

	driver, err := netconf.NewDriver(
		host.IP,
		options.WithAuthNoStrictKey(),
		options.WithAuthUsername("lab"),
		options.WithAuthPassword("lab123"),
		options.WithPort(830),
		options.WithTransportType("standard"),
		options.WithNetconfForceSelfClosingTags(),
	)
	if err != nil {
		fmt.Printf("Failed to create driver for %s; error %+v\n", host, err)
	}
	err = driver.Open()
	if err != nil {
		fmt.Printf("Failed to open driver for %s; error: %+v\n", host, err)
	}
	return driver, nil

}

func executeRPC(hostList []Host) {

	for _, host := range hostList {
	}
}

func main() {
	filepath := "hosts.yaml"
	hosts, err := readHosts(filepath)
	if err != nil {
		fmt.Printf("Got passed an error from readHosts func %v\n", err)
	} else {
		createDriver(hosts)
	}

}
