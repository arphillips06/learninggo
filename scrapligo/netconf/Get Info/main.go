package main

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/scrapli/scrapligo/driver/netconf"
	"github.com/scrapli/scrapligo/driver/opoptions"
	"github.com/scrapli/scrapligo/driver/options"
)

type RPCResult struct {
	Result string
}

type ShowArp struct {
	ArpTableInformation struct {
		ArpTableEntry []struct {
			MacAddress         string `xml:"mac-address"`
			IpAddress          string `xml:"ip-address"`
			Hostname           string `xml:"hostname"`
			InterfaceName      string `xml:"interface-name"`
			ArpTableEntryFlags struct {
				Text string `xml:",chardata"`
				None string `xml:"none"`
			} `xml:"arp-table-entry-flags"`
		} `xml:"arp-table-entry"`
		ArpEntryCount string `xml:"arp-entry-count"`
	} `xml:"arp-table-information"`
}

// Used to open connection to devices
func createOpenDriver() (*netconf.Driver, error) {

	host := "192.168.69.230"

	d, err := netconf.NewDriver(
		host,
		options.WithAuthNoStrictKey(),
		options.WithAuthUsername("lab"),
		options.WithAuthPassword("lab123"),
		options.WithPort(830),
		options.WithTransportType("standard"),
		options.WithNetconfForceSelfClosingTags(),
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to create driver for %s; error %+v\n", host, err)
	}
	err = d.Open()
	if err != nil {
		return nil, fmt.Errorf("Failed to open driver forr %s; error: %+v\n", host, err)

	}
	return d, nil
}

func main() {

	filterString := `<get-arp-table-information/>`

	driver, err := createOpenDriver()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer driver.Close()

	rpc := filterString

	request, err := driver.RPC(opoptions.WithFilter(rpc))

	if err != nil {
		fmt.Println(err)
	}

	rc1 := regexp.MustCompile(">( 2).*")
	request.Result = rc1.ReplaceAllString(request.Result, ">")

	rc2 := regexp.MustCompile("( ).*<")
	request.Result = rc2.ReplaceAllString(request.Result, "<")
	fmt.Println(reflect.TypeOf(request))

}
