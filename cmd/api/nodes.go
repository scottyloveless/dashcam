package main

import (
	"fmt"
	"net/http"
	"net/netip"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type Location struct {
	latitude  float32
	longitude float32
}

type Node struct {
	ipaddress   netip.Addr
	location    Location
	class       string
	snmp        bool
	snmpversion int
}

func (app *application) createNodeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new node")
}

func (app *application) showNodeHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of node %d\n", id)
}
