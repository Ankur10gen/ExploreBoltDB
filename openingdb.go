package main

// Open a database in BoltDB.
// Try to open the same database in a goroutine. It is locked.
// Wait for 10 seconds and close the first DB.
// The goroutine completes.
// Also tested some unrelated bit shift operators
// and some os functions while reading the database Open code

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math"
	"os"
	"runtime"
	//"syscall"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	wg.Add(1)
	go func() {
		defer wg.Done()
		db1, err := bolt.Open("my.db", 0600, nil)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Can I reach here?")
		db1.Close()
	}()

	func() {
		time.Sleep(10 * time.Second)
		db.Close()
	}()
	fmt.Println(runtime.GOOS, runtime.NumCPU())
	fmt.Println((1 << 30) / 1024 / 1024 / 1024)
	fmt.Println(1 << 30)
	fmt.Println(math.Pow(2, 30))
	fmt.Println(1073741824 >> 30)
	fmt.Println(os.Getpagesize())
	//syscall.Truncate()
	wg.Wait()
}
