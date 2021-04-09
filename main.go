package main

import (
	"github.com/secure-dns/lists/ads"
	"github.com/secure-dns/lists/childprotection"
	"github.com/secure-dns/lists/cryptojacking"
	"github.com/secure-dns/lists/security"
	"github.com/secure-dns/lists/tracking"
)

func main() {
	security.Build()
	cryptojacking.Build()
	childprotection.Build()
	ads.Build()
	tracking.Build()
}
