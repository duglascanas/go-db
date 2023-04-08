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

	psqlCreateInvoice = `CREATE TABLE IF NOT EXISTS invoices(
		id serial NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_id_pk PRIMARY KEY(id)
		)`

	psqlCreateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoiceitems(
		id serial NOT NULL,
		invoice_id NOT NULL,
		product_id NOT NULL,
		updated_at TIMESTAMP,
		CONSTRAINT invoiceitem_id_pk PRIMARY KEY(id)
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
type PsqlInvoice struct {
	db *sql.DB
}

// NewPsqlInvoice return a new pointer of
func NewPsqlInvoice(db *sql.DB) *PsqlInvoice {
	return &PsqlInvoice{db}
}

// Migrate implement the interface invoice.Storage
func (p *PsqlInvoice) Migrate() error {

	stmt, err := p.db.Prepare(psqlCreateInvoice)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()

	if err != nil {
		return err
	}

	fmt.Println("Migracion de INVOICE ejecutada correctamente")
	return nil
}

// ****************************************************************************
// PsqlInvoiceItem user for work with posgres
type PsqlInvoiceItem struct {
	db *sql.DB
}

// NewPsqlProduct return a new pointer of
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

	fmt.Println("Migracion de INVOICEITEM ejecutada correctamente")
	return nil
}
