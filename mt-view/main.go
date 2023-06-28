package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

func randMax(a int, b int) int {
	min := a
	max := b
	return rand.Intn(max-min) + min
}

var regions = [4]string{"north", "south", "east", "west"}

func newRow() (int, int, string, int, string) {
	return randMax(1, 10), randMax(1, 10), regions[randMax(0, 3)], randMax(1000, 2000), "indian post"
}

var PerQ = 800

func createQuery() string {
	s := "INSERT INTO orders (jam_id, quantity, region, user_id, shipper) VALUES "

	for i := 0; i < PerQ; i++ {
		a, b, c, d, e := newRow()
		s = s + fmt.Sprintf("(%d, %d, '%s', %d, '%s')", a, b, c, d, e) + ","
	}

	s = s[:len(s)-1]
	s = s + ";"
	return s

}

func main() {
	// PostgreSQL connection parameters

	// Construct the connection string
	connStr := os.Getenv("postgres_url")
	log.Println(connStr)
	// Connect to the PostgreSQL database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
	}

	fmt.Println("Connected to the PostgreSQL database!")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {

		for i := 0; i < 1000; i++ {
			_, err := db.Exec(createQuery())
			if err != nil {
				log.Println(err)
				return
			}

			log.Println("inserted 1000 rows")
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {

		for i := 0; i < 1000; i++ {
			_, err := db.Exec(createQuery())
			if err != nil {
				log.Println(err)
				return
			}

			log.Println("inserted 1000 rows")
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {

		for i := 0; i < 1000; i++ {
			_, err := db.Exec(createQuery())
			if err != nil {
				log.Println(err)
				return
			}

			log.Println("inserted 1000 rows")
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {

		for i := 0; i < 1000; i++ {
			_, err := db.Exec(createQuery())
			if err != nil {
				log.Println(err)
				return
			}

			log.Println("inserted 1000 rows")
		}
		wg.Done()
	}()
	wg.Wait()
	// Perform database operations here...

}
