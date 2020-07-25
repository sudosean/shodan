package main

import (
	"fmt"
	"github.com/sudosean/shodan/pkg/shodan"
	"log"
	"os"
)

func main() {
	fmt.Println("Shodan Client")

	if len(os.Args) != 2 {
		log.Fatalln("Useage: shodan searchterm")
	}
	apiKey :=  "INSERT_API_KEY_HERE" //os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf(
		"Query Credits: %d\nScan Credits %d \n\n",
		info.QueryCredits,
		info.ScanCredits,
	)
	hostSearch, err := s.HostSearch(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}
	for _, host := range hostSearch.Matches {
		fmt.Printf("%18s%8d\n", host.IPString, host.Port)
	}
	for _, host := range hostSearch.Matches {
		for _, domain := range host.Domains {
			fmt.Printf("%18s\n", domain)
		}
	}
}
