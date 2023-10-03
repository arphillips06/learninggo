package main

import "fmt"

type Router struct {
	hostname  string
	platform  string
	username  string
	password  string
	strictkey bool
}

type Inventory struct {
	Routers []Router
}

func main() {
	var r1 Router
	r1.hostname = "router1.example.com"
	r2 := new(Router)
	r2.hostname = "router2.example.com"
	r3 := Router{
		hostname:  "router3.example.com",
		platform:  "junos",
		username:  "user",
		password:  "secret",
		strictkey: false,
	}
	inv := Inventory{
		Routers: []Router{r1, *r2, r3},
	}
	fmt.Printf("Inventory: %v\n", inv)
}
