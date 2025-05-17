package main

import (
	"fmt"
	"time"

	"github.com/okarpova/my-app/pkg/grpcclient"
)

func main() {
	// Create client
	now := time.Now()
	profile, err := grpcclient.New("localhost:50051")
	if err != nil {
		panic(err)
	}
	fmt.Println("grpcclient.New:", time.Since(now))
	defer profile.Close()

	// First request
	now = time.Now()
	id, err := profile.Create("John", 25, "john@gmail.com", "+73003002020")
	if err != nil {
		panic(err)
	}
	fmt.Println("Create:", time.Since(now))

	// Get requests
	start := time.Now()

	for range 5 {
		now = time.Now()
		_, err = profile.Get(id.String())
		if err != nil {
			panic(err)
		}
		fmt.Println("Get:", time.Since(now))
	}

	fmt.Println("Total Get time:", time.Since(start))
}
