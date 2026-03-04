package main

import (
	"fmt"
	"net/http"
)

// type Location struct {
// 	latitude  float32
// 	longitude float32
// }

// type Node struct {
// 	ipaddress   netip.Addr
// 	location    Location
// 	class       string
// 	snmp        bool
// 	snmpversion int
// }

func (app *application) createNodeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new node")
}

func (app *application) showNodeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of node %d\n", id)
}
