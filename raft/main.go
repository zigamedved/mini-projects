package main

import (
	"flag"
	"fmt"
)

func main() {
	addr := flag.String("raft", "", "raft server address")
	join := flag.String("join", "", "join cluster address")
	api := flag.String("api", "", "api server address")
	state := flag.String("state_dir", "", "raft state directory (WAL, Snapshots)")
	flag.Parse()

	fmt.Printf("parsed flags: %s %s %s %s", *addr, *join, *api, *state)
}
