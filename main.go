package main

import "github.com/duglascanas/go-db/pkg/storage"

func main() {
	storage.NewPostgresDB()
	storage.NewPostgresDB()
	storage.NewPostgresDB()
	storage.NewPostgresDB()

}
