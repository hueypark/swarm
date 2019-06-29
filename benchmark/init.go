package benchmark

import (
	"flag"
)

var (
	user     string
	host     string
	port     string
	database string
)

func init() {
	flag.StringVar(&user, "user", "root", "user")
	flag.StringVar(&port, "host", "localhost", "host")
	flag.StringVar(&port, "port", "26257", "port")
	flag.StringVar(&database, "database", "postgres", "database")
	flag.Parse()
}
