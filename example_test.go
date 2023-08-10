package maxminddb_test

import (
	"fmt"
	"log"
	"net"
	"testing"

	"github.com/ahfuzhang/maxminddb-golang"
)

// This example shows how to decode to a struct.
func Test_ExampleReader_Lookup_struct(t *testing.T) {
	db, err := maxminddb.Open("test-data/test-data")
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()
	var record struct {
		Country struct {
			ISOCode string `maxminddb:"iso_code"`
		} `maxminddb:"country"`
	} // Or any appropriate struct
	// 45.232.32.87  CL
	arr := []string{"10.140.70.221", "246.80.248.86", "45.232.32.87", "42.84.232.144", "39.144.219.158", "81.2.69.142"}
	for _, item := range arr {
		ip := net.ParseIP(item)
		err = db.Lookup(ip, &record)
		if err != nil {
			t.Error(err)
			return
		}
		t.Logf("ip=%s, addr=[%s]", ip.String(), record.Country.ISOCode)
	}
	//fmt.Print(record.Country.ISOCode)
	// Output:
	// GB
}

// This example demonstrates how to decode to an interface{}.
func ExampleReader_Lookup_interface() {
	db, err := maxminddb.Open("test-data/test-data/GeoIP2-City-Test.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ip := net.ParseIP("81.2.69.142")

	var record interface{}
	err = db.Lookup(ip, &record)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%v", record)
}

// This example demonstrates how to iterate over all networks in the
// database.
func ExampleReader_Networks() {
	db, err := maxminddb.Open("test-data/test-data/GeoIP2-Connection-Type-Test.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	record := struct {
		Domain string `maxminddb:"connection_type"`
	}{}

	networks := db.Networks(maxminddb.SkipAliasedNetworks)
	for networks.Next() {
		subnet, err := networks.Network(&record)
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("%s: %s\n", subnet.String(), record.Domain)
	}
	if networks.Err() != nil {
		log.Panic(networks.Err())
	}
	// Output:
	// 1.0.0.0/24: Cable/DSL
	// 1.0.1.0/24: Cellular
	// 1.0.2.0/23: Cable/DSL
	// 1.0.4.0/22: Cable/DSL
	// 1.0.8.0/21: Cable/DSL
	// 1.0.16.0/20: Cable/DSL
	// 1.0.32.0/19: Cable/DSL
	// 1.0.64.0/18: Cable/DSL
	// 1.0.128.0/17: Cable/DSL
	// 2.125.160.216/29: Cable/DSL
	// 67.43.156.0/24: Cellular
	// 80.214.0.0/20: Cellular
	// 96.1.0.0/16: Cable/DSL
	// 96.10.0.0/15: Cable/DSL
	// 96.69.0.0/16: Cable/DSL
	// 96.94.0.0/15: Cable/DSL
	// 108.96.0.0/11: Cellular
	// 149.101.100.0/28: Cellular
	// 175.16.199.0/24: Cable/DSL
	// 187.156.138.0/24: Cable/DSL
	// 201.243.200.0/24: Corporate
	// 207.179.48.0/20: Cellular
	// 216.160.83.56/29: Corporate
	// 2003::/24: Cable/DSL
}

// This example demonstrates how to iterate over all networks in the
// database which are contained within an arbitrary network.
func ExampleReader_NetworksWithin() {
	db, err := maxminddb.Open("test-data/test-data/GeoIP2-Connection-Type-Test.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	record := struct {
		Domain string `maxminddb:"connection_type"`
	}{}

	_, network, err := net.ParseCIDR("1.0.0.0/8")
	if err != nil {
		log.Panic(err)
	}

	networks := db.NetworksWithin(network, maxminddb.SkipAliasedNetworks)
	for networks.Next() {
		subnet, err := networks.Network(&record)
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("%s: %s\n", subnet.String(), record.Domain)
	}
	if networks.Err() != nil {
		log.Panic(networks.Err())
	}

	// Output:
	// 1.0.0.0/24: Cable/DSL
	// 1.0.1.0/24: Cellular
	// 1.0.2.0/23: Cable/DSL
	// 1.0.4.0/22: Cable/DSL
	// 1.0.8.0/21: Cable/DSL
	// 1.0.16.0/20: Cable/DSL
	// 1.0.32.0/19: Cable/DSL
	// 1.0.64.0/18: Cable/DSL
	// 1.0.128.0/17: Cable/DSL
}
