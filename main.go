package main

import (
	"flag"
	"fmt"

	. "github.com/hto/redis-memory-analysis"
)

var (
	host       = flag.String("host", "127.0.0.1", "The host of redis")
	port       = flag.String("port", "6379", "The port of redis")
	password   = flag.String("password", "", "The password of redis")
	reportPath = flag.String("reportPath", "./reports", "The csv file path of analysis result")
)

func main() {

	flag.Parse()

	// ElastiCache Endpoint
	Connection(*host, *port, *password)

	// sorted by key priority
	Start([]string{":", "_"})

	// //Find the csv file in default target folder: ./reports
	// //CSV file name format: redis-analysis-{host:port}-{db}.csv
	// //The keys order by count desc
	err := SaveReports(*reportPath)
	if err == nil {
		fmt.Println("\ndone")
	} else {
		fmt.Println("error:", err)
	}

	Close()
}
