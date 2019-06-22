package benchmark

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/hueypark/swarm/model"
	"github.com/jinzhu/gorm"
)

// Issue4234 is benchmark for issue 4234
// https://github.com/cockroachdb/docs/issues/4234
func Issue4234(db *gorm.DB) {
	startTime := time.Now()

	var counter int64

	db.AutoMigrate(&model.User{})

	var wg sync.WaitGroup

	for i := 0; i < runtime.NumCPU()*10; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 10000; j++ {
				db.Create(&model.User{Name: "foo"})
				atomic.AddInt64(&counter, 1)

				log.Println(fmt.Sprintf("time: %v, counter: %v", time.Now().Sub(startTime), counter))
			}

			wg.Done()
		}()
	}

	wg.Wait()
}
