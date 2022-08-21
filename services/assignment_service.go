package services

import (
	"Go/packages/mongodb"
	"Go/repositories"
	"bufio"
	"context"
	"mime/multipart"
	"sync"
	"time"
)

func InsertListAssignments(file multipart.File) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Open mongodb connection
	mongoClient, err := mongodb.NewMongoDBConnect(ctx)
	if err != nil {
		return err
	}

	//Create channel
	ch := make(chan string)
	//Create waitgroup
	var wg sync.WaitGroup
	scanner := bufio.NewScanner(file)

	//Start 10 workers
	for t := 0; t < 10; t++ {
		//Add new worker
		wg.Add(1)
		//User goroutines for insert line in channel into mongo
		go func(ch chan string, wg *sync.WaitGroup) {
			//Mark waitgroup done
			defer wg.Done()
			for line := range ch {
				repositories.InsertAssignment(line, mongoClient)
			}
		}(ch, &wg)
	}

	//Read line and send to channel
	for scanner.Scan() {
		ch <- scanner.Text()
	}
	close(ch)
	wg.Wait()

	return nil
}
