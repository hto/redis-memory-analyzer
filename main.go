package main

import (
	"flag"
	"fmt"

	. "github.com/hto/redis-memory-analysis"
)

var (
	ip         = flag.String("ip", "127.0.0.1", "The host of redis")
	port       = flag.Uint("port", 6379, "The port of redis")
	password   = flag.String("password", "", "The password of redis")
	reportPath = flag.String("reportPath", "./reports", "The csv file path of analysis result")
)

func main() {

	flag.Parse()

	analysis, err := NewAnalysisConnection(*ip, uint16(*port), *password)
	if err != nil {
		fmt.Println("something wrong:", err)
		return
	}
	defer analysis.Close()

	analysis.Start([]string{":"})

	//Find the csv file in default target folder: ./reports
	//CSV file name format: redis-analysis-{host:port}-{db}.csv
	//The keys order by count desc
	err = analysis.SaveReports(*reportPath)
	if err == nil {
		fmt.Println("\ndone")
	} else {
		fmt.Println("error:", err)
	}
}
