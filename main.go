package main

import (
	"log"
	"time"

	"encoding/json"

	procinfo "github.com/c9s/goprocinfo/linux"
)

func showStats() {
	stat, err := procinfo.ReadStat("/proc/stat")
	if err != nil {
		log.Fatal("stat read failed")
	}

	for _, s := range stat.CPUStats {
		v, _ := json.Marshal(s)
		log.Printf("%s", v)
	}
}

func main() {
	for {
		showStats()
		time.Sleep(10 * time.Second)
	}
}
