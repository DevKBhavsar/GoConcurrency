package main

import (
	"context"
	"log"
	"os"
	"sync"
)

func main() {
	defer log.Println("DONe")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	apiConnection := Open()
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ReadFile(context.Background())
			if err != nil {
				log.Printf("Cannot ReadFile %v", err)
			}
			log.Println("readfile")
		}()
	}

	
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ResolveAddress(context.Background())
			if err != nil {
				log.Printf("Cannot Resolve File %v", err)
			}
			log.Println("RESOLVEDAddres")
		}()
	}

	wg.Wait()
}
