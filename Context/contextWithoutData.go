package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// func main() {
// 	var wg sync.WaitGroup
// 	ctx, cancel := context.WithCancel(context.Background())
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		if err := printGreeting(ctx); err != nil {
// 			fmt.Println(err)
// 			cancel()
// 		}
// 	}()

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		if err := printFarwell(ctx); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	wg.Wait()
// }

func printGreeting(ctx context.Context) error {
	greeting, err := genGreeting(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world", greeting)
	return nil
}

func printFarwell(ctx context.Context) error {
	farwell, err := genFarwell(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%s world", farwell)
	return nil
}

func genGreeting(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "hello", nil
	}
	return "", fmt.Errorf("Unsupported locale")
}

func genFarwell(ctx context.Context) (string, error) {
	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "Goodbye", nil
	}
	return "", fmt.Errorf("Unsupported locale")
}

func locale(ctx context.Context) (string, error) {
	// if deadline, ok := ctx.Deadline(); ok {
	// 	if deadline.Sub(time.Now().Add(5*time.Second)) <= 0 {
	// 		return "", context.DeadlineExceeded
	// 	}
	// }
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(1 * time.Second):
	}
	return "EN/US", nil
}
