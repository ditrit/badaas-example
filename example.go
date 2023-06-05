package main

import (
	"github.com/ditrit/badaas-example/models"
	"github.com/ditrit/badaas/badorm"
	badaasModels "github.com/ditrit/badaas/persistence/models"
	"github.com/ditrit/badaas/services"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func CreateEAVCRUDObjects(
	logger *zap.Logger,
	db *gorm.DB,
	eavService services.EAVService,
) error {
	_, err := eavService.GetEntities("profile", map[string]any{})
	if err != nil {
		logger.Sugar().Info("Setting up Posts EAV example")

		userID := "wowASuperCoolUserID"

		// creation of Profile type and associated attributes
		profileType := &badaasModels.EntityType{
			Name: "profile",
		}
		displayNameAttr := &badaasModels.Attribute{
			EntityTypeID: profileType.ID,
			Name:         "displayName",
			ValueType:    badaasModels.StringValueType,
			Required:     true,
		}
		urlPicAttr := &badaasModels.Attribute{
			EntityTypeID:  profileType.ID,
			Name:          "urlPic",
			ValueType:     badaasModels.StringValueType,
			Required:      false,
			Default:       true,
			DefaultString: "https://www.startpage.com/av/proxy-image?piurl=https%3A%2F%2Fimg.favpng.com%2F17%2F19%2F1%2Fbusiness-google-account-organization-service-png-favpng-sUuKmS4aDNRzxDKx8kJciXdFp.jpg&sp=1672915826Tc106d9b5cab08d9d380ce6fdc9564b199a49e494a069e1923c21aa202ba3ed73", //nolint:lll
		}
		userIDAttr := &badaasModels.Attribute{
			EntityTypeID: profileType.ID,
			Name:         "userId",
			ValueType:    badaasModels.StringValueType,
			Required:     true,
		}
		profileType.Attributes = append(profileType.Attributes,
			displayNameAttr,
			urlPicAttr,
			userIDAttr,
		)

		// instantiation of a Profile
		adminProfile := &badaasModels.Entity{
			EntityTypeID: profileType.ID,
			EntityType:   profileType,
		}
		displayNameVal, _ := badaasModels.NewStringValue(displayNameAttr, "The Super Admin")
		userPicVal, _ := badaasModels.NewNullValue(urlPicAttr)
		userIDVal, _ := badaasModels.NewStringValue(userIDAttr, userID)
		adminProfile.Fields = append(adminProfile.Fields,
			displayNameVal,
			userPicVal,
			userIDVal,
		)

		err = db.Create(adminProfile).Error
		if err != nil {
			return err
		}

		// creation of Post type and associated attributes
		postType := &badaasModels.EntityType{
			Name: "post",
		}
		titleAttr := &badaasModels.Attribute{
			EntityTypeID: postType.ID,
			Name:         "title",
			ValueType:    badaasModels.StringValueType,
			Required:     true,
		}
		bodyAttr := &badaasModels.Attribute{
			Name:          "body",
			ValueType:     badaasModels.StringValueType,
			Default:       false,
			DefaultString: "empty",
		}
		ownerAttr := &badaasModels.Attribute{
			Name:      "ownerID",
			ValueType: badaasModels.StringValueType,
			Required:  true,
		}

		postType.Attributes = append(
			postType.Attributes, titleAttr, bodyAttr, ownerAttr,
		)

		// instantiation of a Post
		whyCatsLikeMice := &badaasModels.Entity{
			EntityTypeID: postType.ID,
			EntityType:   postType,
		}
		titleVal, _ := badaasModels.NewStringValue(titleAttr, "Why cats like mice?")
		bodyVal, _ := badaasModels.NewStringValue(bodyAttr,
			`Lorem ipsum dolor sit amet, consectetur adipiscing elit.

		In consectetur, ex at hendrerit lobortis, tellus lorem blandit eros, vel ornare odio lorem eget nisi.

		In erat mi, pharetra ut lacinia at, facilisis vitae nunc.
	`)
		ownerVal, _ := badaasModels.NewStringValue(ownerAttr, userID)

		whyCatsLikeMice.Fields = append(whyCatsLikeMice.Fields,
			titleVal, bodyVal, ownerVal,
		)

		err = db.Create(whyCatsLikeMice).Error
		if err != nil {
			return err
		}
		logger.Sugar().Info("Finished creating Posts EAV example")
	}

	return nil
}

func CreateCRUDObjects(
	logger *zap.Logger,
	db *gorm.DB,
	crudProductRepository badorm.CRUDRepository[models.Product, badorm.UUID],
) error {
	products, err := crudProductRepository.GetAll(db)
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
