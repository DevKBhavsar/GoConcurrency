package main

import (
	"fmt"
	"log"
	"os"
)

func handleError(key int, err error, message string) {
	log.SetPrefix(fmt.Sprintf("[logid %v]", key))
	log.Printf("%#v", err)
	fmt.Printf("[%v] %v", key, message)
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime| log.LUTC)

	err := runjob("1")
	if err != nil {
		msg := "UNexpected issue"
		if _, ok := err.(IntermediateErr); ok {
			msg = err.Error()
		}
		handleError(1,err,msg)
	}
}
