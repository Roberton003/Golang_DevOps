package main

import (
	
	"log"
	"sync"
	"time"
	
	)

	


func worker(id int){
	log.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	log.Printf("Worker %d\n done\n", id)
}

func main() {
	var wg sync.WaitGroup
	time.Sleep(time.Second)
	for i := 1; i <= 5; i++ {
		go func (id int){
		    defer wg.Done()
			worker(id)	
		} (i)
	}
	wg.Wait()
}