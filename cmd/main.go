package main

import (
	eventstreamserver "github.com/jamieyoung5/kwikmedical-eventstream/cmd/eventstream"
	"os"
	"time"
)

func main() {

	err := eventstreamserver.Start()
	for err != nil {
		err = eventstreamserver.Start()
		if err != nil {
			os.Exit(1)
		}

		time.Sleep(5 * time.Second)
	}

}
