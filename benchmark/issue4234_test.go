package benchmark

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

func BenchmarkIssue4234(b *testing.B) {
	insertRowCount := 100
	for i := 0; i < 20; i++ {
		b.Run(fmt.Sprintf("%d", insertRowCount), func(b *testing.B) {
			db, err := sql.Open(
				"postgres",
				"postgresql://"+user+"@"+host+":26257/postgres?sslmode=disable")
			if err != nil {
				log.Fatalln(err)
			}
			defer func() {
				err := db.Close()
				if err != nil {
					log.Println(err)
				}
			}()

			for i := 0; i < b.N; i++ {
				issue4234(db, insertRowCount, 16)
			}
		})
	}
}
