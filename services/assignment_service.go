package services

import (
	"Go/packages/mongodb"
	"Go/repositories"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

func InsertListAssignments() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := mongodb.NewMongoDBConnect(ctx)

	file, err := os.Open("./inputs/numbers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ch := make(chan string)
	var wg sync.WaitGroup
	scanner := bufio.NewScanner(file)

	// start the workers
	for t := 0; t < 10; t++ {
		wg.Add(1)
		go func(ch chan string, wg *sync.WaitGroup) {
			defer wg.Done()
			for line := range ch {
				repositories.InsertAssignment(line, mongoClient)
				fmt.Println(line)
			}
		}(ch, &wg)
	}

	for scanner.Scan() {
		ch <- scanner.Text()
	}
	close(ch)
	wg.Wait()
}
