package main

import (
	"flag"
	"fmt"
	"regexp"
	"time"

	"github.com/CaliDog/certstream-go"
	"github.com/fatih/color"
)

var (
	green = color.New(color.FgGreen).SprintFunc()
	blue  = color.New(color.FgBlue).SprintFunc()
	cyan  = color.New(color.FgCyan).SprintFunc()
)

func main() {
	pattern := flag.String("E", "", "Regex pattern to filter domains")
	silent := flag.Bool("s", false, "Silent mode - print only matching domains")
	flag.Parse()

	var regex *regexp.Regexp
	var err error
	if *pattern != "" {
		regex, err = regexp.Compile(*pattern)
		if err != nil {
			fmt.Printf("Invalid regex pattern: %v\n", err)
			return
		}
	}

	if !*silent {
		fmt.Printf("%s Certificate Transparency Monitor %s\n\n",
			cyan("🔍"),
			blue("v1.0"))

		if regex != nil {
			fmt.Printf("Filtering domains with pattern: %s\n\n", *pattern)
		}
	}

	stream, errStream := certstream.CertStreamEventStream(false)
	for {
		select {
		case jq := <-stream:
			messageType, err := jq.String("message_type")
			if err != nil {
				continue
			}

			if messageType == "certificate_update" {
				// Extract domain
				domains, err := jq.ArrayOfStrings("data", "leaf_cert", "all_domains")
				if err != nil || len(domains) == 0 {
					continue
				}

				// Skip if doesn't match regex pattern
				if regex != nil {
					matches := false
					for _, domain := range domains {
						if regex.MatchString(domain) {
							matches = true
							break
						}
					}
					if !matches {
						continue
					}
				}

				// Get timestamp and format it
				seen, _ := jq.Float("data", "seen")
				timestamp := time.Unix(int64(seen), 0).Format("15:04:05")

				// Get the certificate source
				source, _ := jq.String("data", "source", "name")

				if *silent {
					fmt.Println(domains[0])
				} else {
					// Print in a nice format
					fmt.Printf("[%s] %s %s (%s)\n",
						blue(timestamp),
						green("✓"),
						domains[0],
						cyan(source))
				}
			}
		case err := <-errStream:
			if !*silent {
				fmt.Printf("Error: %v\n", err)
			}
		}
	}
}
