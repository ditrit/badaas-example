package models

import (
	"github.com/ditrit/badaas/badorm"
	"github.com/google/uuid"
)

type Company struct {
	badorm.UUIDModel

	Name    string
	Sellers []Seller
}

type Product struct {
	badorm.UUIDModel

	String string
	Int    int
	Float  float64
	Bool   bool
}

type Seller struct {
	badorm.UUIDModel

	Name      string
	CompanyID *uuid.UUID
}

type Sale struct {
	badorm.UUIDModel

	// belongsTo Product
	Product   *Product
	ProductID uuid.UUID

	// belongsTo Seller
	Seller   *Seller
	SellerID uuid.UUID
}
