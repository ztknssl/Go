package main

import (
	"log"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		log.Fatalf(err.Error())
	}

	os.Stdout.WriteString(time.String())
}
