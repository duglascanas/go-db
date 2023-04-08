package main

import (
	"log"

	"github.com/duglascanas/go-db/pkg/product"
	"github.com/duglascanas/go-db/pkg/storage"
)

func main() {
	storage.NewPostgresDB()
	//Product
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}
	//Invoice
	storageInvoice := storage.NewPsqlInvoice(storage.Pool())
	serviceInvoice := invoice.NewService(storageInvoice)

	if err := serviceInvoice.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}
	//Invoice
	storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
	serviceInvoiceItem := invoiceItem.NewService(storageInvoiceItem)

	if err := serviceInvoiceItem.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}
}
