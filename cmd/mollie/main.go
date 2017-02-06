package main

import "localhost/he/go-mollie-api/mollie"

func main() {
	m := mollie.Get("test_pQ2c9R3DDj2WbQdcaqFNxcjQQ6qSaU")
	m.Issuers.List()
	m.Issuers.Get("ideal_TESTNL99")
}
