package main

import (
	"os"

	_ "github.com/chennqqi/filebeat.nsq.output/nsqout"
	"github.com/elastic/beats/v7/filebeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
