package main

import (
	"fmt"

	"github.com/qdrow-tch/errors_workout/logging"
)

func main() {
	err := logging.Push_to_logfile("ffffffff")
	if err != nil {
		fmt.Println("erroorrrrr")
	}
}
