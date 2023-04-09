package storage

import (
	"database/sql"
	"fmt"
)

const (
	psqlCreateProduct = `CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT product_id_pk PRIMARY KEY (id)         
	)`

	psqlCreateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers (
		id serial NOT NULL,
		client VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_id_pk PRIMARY KEY(id)
		)`

	psqlCreateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items(
		id serial NOT NULL,
		invoice_header_id INT NOT NULL,
		product_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_items_id_pk PRIMARY KEY(id),
		CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY
		(invoice_header_id) REFERENCES invoice_headers (id) ON UPDATE
		RESTRICT ON DELETE RESTRICT,
		CONSTRAINT invoice_items_product_id_fk FOREIGN KEY (product_id)
		REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
		)`
)

// PsqlProduct user for work with posgres
type PsqlProduct struct {
	db *sql.DB
}

// NewPsqlProduct return a new pointer of
func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}

// Migrate implement the interface product.Storage
func (p *PsqlProduct) Migrate() error {

	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()

	if err != nil {
		return err
	}

	fmt.Println("Migracion de producto ejecutada correctamente")
	return nil
}

// ****************************************************************************
// PsqlInvoice user for work with posgres
type PsqlInvoiceHeader struct {
	db *sql.DB
}

// NewPsqlInvoice return a new pointer of
func NewPsqlInvoiceHeader(db *sql.DB) *PsqlInvoiceHeader {
	return &PsqlInvoiceHeader{db}
}

// Migrate implement the interface invoice.Storage
func (p *PsqlInvoiceHeader) Migrate() error {

	stmt, err := p.db.Prepare(psqlCreateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()

	if err != nil {
		return err
	}

	fmt.Println("Migracion de INVOICEHEADERS ejecutada correctamente")
	return nil
}

// ****************************************************************************
// PsqlInvoiceItem user for work with posgres
type PsqlInvoiceItem struct {
	db *sql.DB
}

// NewPsqlInvoiceItem return a new pointer of
func NewPsqlInvoiceItem(db *sql.DB) *PsqlInvoiceItem {
	return &PsqlInvoiceItem{db}
}

// Migrate implement the interface invoiceitem.Storage
func (p *PsqlInvoiceItem) Migrate() error {

	stmt, err := p.db.Prepare(psqlCreateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()

	if err != nil {
		return err
	}

	fmt.Println("Migracion de INVOICEITEMS ejecutada correctamente")
	return nil
}
