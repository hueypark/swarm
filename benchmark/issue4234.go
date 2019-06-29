package benchmark

import (
	"database/sql"
	"log"
	"sync"
)

// Issue4234 is benchmark for issue 4234
// https://github.com/cockroachdb/docs/issues/4234
func issue4234(db *sql.DB, insertRowCount int, goroutineCount int) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS accounts (balance INT)")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = db.Exec("TRUNCATE TABLE accounts")
	if err != nil {
		log.Fatalln(err)
	}

	var wg sync.WaitGroup

	for i := 0; i < goroutineCount; i++ {
		wg.Add(1)

		go func() {
			for i := 0; i < insertRowCount; i++ {
				_, err = db.Exec("INSERT INTO accounts (balance) VALUES (1000)")
				if err != nil {
					log.Fatalln(err)
				}
			}

			wg.Done()
		}()
	}

	wg.Wait()
}
