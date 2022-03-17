package cmd

import "flag"

var (
	//System port
	Port string
	//Whether open criu-ns
	FlagNs bool
)

func init() {
	flag.StringVar(&Port, "port", "8080", "System port")
	flag.BoolVar(&FlagNs, "flag_ns", true, "Whether open criu-ns")
}
