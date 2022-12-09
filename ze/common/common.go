package common

import (
	"fmt"
	"os"
	"strings"
)

// TODO: Implement
func ValidateAuthToken(auth string) {
	if len(auth) == 0 {
		fmt.Println("Auth token must be set using the -auth flag")
		os.Exit(1)
	}
}

// TODO: Implement
func ValidateZapiUrl(url string) {
	if len(url) == 0 {
		fmt.Println("URL must be set using the -url flag")
		os.Exit(1)
	}
}

func ValidateUpMetadata(filename string, logype string, logstash bool, batchId string, cfgs string) {
	//Make sure log type is specified
	if len(filename) == 0 {
		if logype == "" && !logstash {
			fmt.Println("Error: logtype must be specified for streaming with --log")
			os.Exit(1)
		}
	}

	if strings.Contains(cfgs, "ze_batch_id") && batchId != "" {
		fmt.Println("ze_batch_id is defined in cfgs put also specified with --batch.  Please correct conflict.")
		os.Exit(1)
	}

}
