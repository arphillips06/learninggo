package main

import (
	"fmt"
	"os"

	"github.com/scrapli/scrapligo/driver/netconf"
	"github.com/scrapli/scrapligo/driver/opoptions"
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

// uses Scrapligo to create driver, is passed parameters from func executeRPC
// returns driver and error vars
// func is for Juniper, some option flags may be superfluous
// The parameter "(host Host)" is looking for a single host to be passed, not the whole struct
func createDriver(host Host) (*netconf.Driver, error) {

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
		fmt.Printf("Failed to create driver for %s; error %+v\n", host.IP, err)
	}
	err = driver.Open()
	if err != nil {
		fmt.Printf("Failed to open driver for %s; error: %+v\n", host.IP, err)
	}
	//fmt.Println(host.IP)
	//fmt.Println(driver)

	return driver, nil

}

// Func to run the commands found in hosts.yaml
// The parameter "(hostList []Host)" is expecting the slice to be passed to it
// The "for range" loop will then access the slice via index's
// The "_" means it does not care about the index # nor how many there are
func executeRPC(hostList []Host) {

	for _, host := range hostList {
		driver, err := createDriver(host)
		if err != nil {
			fmt.Printf("failed to create driver; error: %+v\n", err)
		}
		defer driver.Close()
		for _, cmd := range host.Cmd {
			rpc := cmd
			rpcCall, err := driver.RPC(opoptions.WithFilter(rpc))
			if err != nil {
				fmt.Printf("Failed executing RPC call for host %s; error: %v\n", host.Name, err)
				continue
			}
			println(rpcCall.Result)
		}
	}

}

// main func to call other funcs
func main() {
	filepath := "hosts.yaml"

	hosts, err := readHosts(filepath)
	if err != nil {
		fmt.Printf("Got passed an error from readHosts func %v\n", err)
	} else {
		executeRPC(hosts)
	}

}
