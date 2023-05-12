package models

import (
	"github.com/ditrit/badaas/persistence/models"
	"github.com/google/uuid"
)

type Company struct {
	models.BaseModel

	Name    string
	Sellers []Seller
}

type Product struct {
	models.BaseModel

	String string
	Int    int
	Float  float64
	Bool   bool
}

type Seller struct {
	models.BaseModel

	Name      string
	CompanyID *uuid.UUID
}

type Sale struct {
	models.BaseModel

	// belongsTo Product
	Product   *Product
	ProductID uuid.UUID

	// belongsTo Seller
	Seller   *Seller
	SellerID uuid.UUID
}
