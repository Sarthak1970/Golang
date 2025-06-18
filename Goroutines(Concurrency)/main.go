package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)
var m=sync.Mutex{} // Mutex to protect shared data
var wg=sync.WaitGroup{}
var dbData=[]string{"id1","id2","id3","id4","id5","id6","id7","id8","id9","id10"}

func dbCall(id int) {
	// Simulating a database call with a random sleep time
	fmt.Printf("Starting database call for %s\n", dbData[id])
	sleepTime := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(sleepTime)
	m.Lock() // Lock the mutex to protect shared data
	fmt.Printf("Database call for %s completed after %v\n", dbData[id], sleepTime)
	m.Unlock() // Unlock the mutex after accessing shared data
	wg.Done() // Decrement the WaitGroup counter
}

func main(){
	t0:=time.Now()
	for i:=0;i<len(dbData);i++{
		wg.Add(1) // Increment the WaitGroup counter
		go dbCall(i)
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Printf("Time taken for goroutines: %v\n", time.Since(t0))

}