package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tecowl/slack-knocker"
)

var configPath = flag.String("c", "slack-knocker.json", "Path to configuration JSON file")
var dryrun = flag.Bool("dryrun", false, "Will not send message but show it")

func main() {
	flag.Parse()

	config, err := slackknocker.LoadConfigFile(*configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load config file %s because of %v\n", *configPath, err)
		os.Exit(1)
	}

	knocker := slackknocker.NewKnocker(config)
	payload := knocker.BuildPayload(flag.Args())
	if *dryrun {
		fmt.Fprintf(os.Stdout, "payload: %v\n", payload)
		os.Exit(0)
	}

	if err := knocker.Post(payload); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to post %v to %s\n", payload, knocker.Config.WebhookURL)
		os.Exit(1)
	}
}
