package services

import (
	"Go/repositories"
	"bufio"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"mime/multipart"
	"sync"
)

func InsertListAssignments(ctx context.Context, file multipart.File, numbersCollection *mongo.Collection) error {
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
				if result, err := repositories.InsertAssignment(ctx, line, numbersCollection); !result {
					log.Fatal(err)
				}
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
