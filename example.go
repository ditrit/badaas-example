package main

import (
	"github.com/ditrit/badaas-example/models"
	"github.com/ditrit/badaas/badorm"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func CreateCRUDObjects(
	logger *zap.Logger,
	db *gorm.DB,
	crudProductRepository badorm.CRUDRepository[models.Product, badorm.UUID],
) error {
	products, err := crudProductRepository.Query(db)
	if err != nil {
		return err
	}

	if len(products) == 0 {
		logger.Sugar().Info("Setting up CRUD example")

		product1 := &models.Product{
			Int: 1,
		}
		err = crudProductRepository.Create(db, product1)
		if err != nil {
			return err
		}

		product2 := &models.Product{
			Int: 2,
		}
		err = crudProductRepository.Create(db, product2)
		if err != nil {
			return err
		}

		company1 := &models.Company{
			Name: "ditrit",
		}
		err = db.Create(company1).Error
		if err != nil {
			return err
		}
		company2 := &models.Company{
			Name: "orness",
		}
		err = db.Create(company2).Error
		if err != nil {
			return err
		}

		seller1 := &models.Seller{
			Name:      "franco",
			CompanyID: &company1.ID,
		}
		err = db.Create(seller1).Error
		if err != nil {
			return err
		}
		seller2 := &models.Seller{
			Name:      "agustin",
			CompanyID: &company2.ID,
		}
		err = db.Create(seller2).Error
		if err != nil {
			return err
		}

		sale1 := &models.Sale{
			Product: product1,
			Seller:  seller1,
		}
		err = db.Create(sale1).Error
		if err != nil {
			return err
		}
		sale2 := &models.Sale{
			Product: product2,
			Seller:  seller2,
		}
		err = db.Create(sale2).Error
		if err != nil {
			return err
		}

		logger.Sugar().Info("Finished creating CRUD example")
	}

	return nil
}
