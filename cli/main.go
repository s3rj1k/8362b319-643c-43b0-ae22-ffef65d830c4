package main

import (
	"flag"
	"fmt"
	"slices"
	"strings"

	log "github.com/sirupsen/logrus"

	"code.local/homework/api"
	"code.local/homework/config"
	"code.local/homework/logcfg"
	"code.local/homework/result"
	"code.local/homework/utils"
)

func main() {
	configPath := flag.String("config", "config.ini", "Path to the configuration file")
	logLevel := flag.String("log_level", "debug", "Set the log verbosity level (debug, info, etc.)")
	logPath := flag.String("log", "", "Path to the log file (logs to stdout if empty)")

	flag.Parse()

	logcfg.Setup(*logLevel, *logPath)

	log.Debug("Program started")
	log.Debugf("Using config: %s", *configPath)

	config, err := config.Parse("config.ini")
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	log.Debugf("Parsed config: %s", *configPath)
	log.Infof("API URL: %s", config.GetAPIURL())
	log.Infof("Loaded subnet: %s", config.GetSubnet())
	log.Infof("New subnet prefix: %d", config.GetNewPrefix())

	subnets, err := utils.DivideSubnet4(config.GetSubnet(), config.GetNewPrefix())
	if err != nil {
		log.Fatalf("Failed to divide IPv4 subnet: %v", err)
	}

	if len(subnets) == 0 {
		log.Fatalf("Failed to divide IPv4 subnet: empty result")
	}

	log.Infof("Generated subnets: %v", subnets)

	log.Debugf("Pulling info from: %s", config.GetAPIURL())

	var servers []api.Response
	if err := utils.FetchAndUnmarshal(config.GetAPIURL(), &servers); err != nil {
		log.Fatalf("Failed to fetch data from API endpoint: %v", err)
	}

	log.Debugf("Loaded %d servers from API", len(servers))

	fservers := slices.DeleteFunc(servers, func(server api.Response) bool {
		return !strings.EqualFold(server.Status, "online")
	})

	log.Debugf("Filtered %d online servers from API", len(fservers))

	result := result.New(subnets)

	for _, srvr := range fservers {
		addrs4 := srvr.GetIPv4List()
		matchCount := 0

		for _, addr := range addrs4 {
			for _, subnet := range subnets {
				match := subnet.Contains(addr)
				log.Debugf("Host: %s IPv4: %s Subnet: %s Match: %t",
					srvr.Hostname, addr, subnet, match,
				)

				if match {
					matchCount++

					result.AddMatched(subnet, srvr.Hostname)
				}
			}
		}

		if matchCount == 0 {
			result.AddUnmatched(srvr.Hostname)
		}
	}

	matched := result.GetMatched()
	unmatched := result.GetUnmatched()

	for k, v := range matched {
		log.Infof("Subnet: %s Hosts [%d]: %v", k, len(v), v)
	}

	log.Infof("Hosts with no subnet match [%d]: %v", len(unmatched), unmatched)

	if *logPath != "" { // when logging into file also print results into stdout
		for k, v := range matched {
			fmt.Printf("Subnet: %s Hosts [%d]: %v\n", k, len(v), v)
		}

		fmt.Printf("Hosts with no subnet match [%d]: %v\n", len(unmatched), unmatched)
	}
}
