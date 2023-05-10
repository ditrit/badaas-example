package main

import (
	"github.com/ditrit/badaas/persistence/models"
	"github.com/ditrit/badaas/services"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// TODO volver a poner los errores pero que ande cuando ya esten en la base
func CreateEAVCRUDObjects(logger *zap.Logger, db *gorm.DB) {
	logger.Sugar().Info("Setting up Posts EAV example")

	userID := "wowASuperCoolUserID"

	// creation of Profile type and associated attributes
	profileType := &models.EntityType{
		Name: "profile",
	}
	displayNameAttr := &models.Attribute{
		EntityTypeID: profileType.ID,
		Name:         "displayName",
		ValueType:    models.StringValueType,
		Required:     true,
	}
	urlPicAttr := &models.Attribute{
		EntityTypeID:  profileType.ID,
		Name:          "urlPic",
		ValueType:     models.StringValueType,
		Required:      false,
		Default:       true,
		DefaultString: "https://www.startpage.com/av/proxy-image?piurl=https%3A%2F%2Fimg.favpng.com%2F17%2F19%2F1%2Fbusiness-google-account-organization-service-png-favpng-sUuKmS4aDNRzxDKx8kJciXdFp.jpg&sp=1672915826Tc106d9b5cab08d9d380ce6fdc9564b199a49e494a069e1923c21aa202ba3ed73", //nolint:lll
	}
	userIDAttr := &models.Attribute{
		EntityTypeID: profileType.ID,
		Name:         "userId",
		ValueType:    models.StringValueType,
		Required:     true,
	}
	profileType.Attributes = append(profileType.Attributes,
		displayNameAttr,
		urlPicAttr,
		userIDAttr,
	)

	// instantiation of a Profile
	adminProfile := &models.Entity{
		EntityTypeID: profileType.ID,
		EntityType:   profileType,
	}
	displayNameVal, _ := models.NewStringValue(displayNameAttr, "The Super Admin")
	userPicVal, _ := models.NewNullValue(urlPicAttr)
	userIDVal, _ := models.NewStringValue(userIDAttr, userID)
	adminProfile.Fields = append(adminProfile.Fields,
		displayNameVal,
		userPicVal,
		userIDVal,
	)

	_ = db.Create(adminProfile).Error

	// creation of Post type and associated attributes
	postType := &models.EntityType{
		Name: "post",
	}
	titleAttr := &models.Attribute{
		EntityTypeID: postType.ID,
		Name:         "title",
		ValueType:    models.StringValueType,
		Required:     true,
	}
	bodyAttr := &models.Attribute{
		Name:          "body",
		ValueType:     models.StringValueType,
		Default:       false,
		DefaultString: "empty",
	}
	ownerAttr := &models.Attribute{
		Name:      "ownerID",
		ValueType: models.StringValueType,
		Required:  true,
	}

	postType.Attributes = append(
		postType.Attributes, titleAttr, bodyAttr, ownerAttr,
	)

	// instantiation of a Post
	whyCatsLikeMice := &models.Entity{
		EntityTypeID: postType.ID,
		EntityType:   postType,
	}
	titleVal, _ := models.NewStringValue(titleAttr, "Why cats like mice?")
	bodyVal, _ := models.NewStringValue(bodyAttr,
		`Lorem ipsum dolor sit amet, consectetur adipiscing elit.

		In consectetur, ex at hendrerit lobortis, tellus lorem blandit eros, vel ornare odio lorem eget nisi.

		In erat mi, pharetra ut lacinia at, facilisis vitae nunc.
	`)
	ownerVal, _ := models.NewStringValue(ownerAttr, userID)

	whyCatsLikeMice.Fields = append(whyCatsLikeMice.Fields,
		titleVal, bodyVal, ownerVal,
	)

	_ = db.Create(whyCatsLikeMice).Error
	logger.Sugar().Info("Finished creating Posts EAV example")
}

func CreateCRUDObjects(
	logger *zap.Logger,
	db *gorm.DB,
	crudProductService services.CRUDService[models.Product, uuid.UUID],
) {
	logger.Sugar().Info("Setting up CRUD example")

	product1, _ := crudProductService.CreateEntity(map[string]any{
		"int": 1,
	})

	product2, _ := crudProductService.CreateEntity(map[string]any{
		"int": 2,
	})

	company1 := &models.Company{
		Name: "ditrit",
	}
	_ = db.Create(company1).Error
	company2 := &models.Company{
		Name: "orness",
	}
	_ = db.Create(company2).Error

	seller1 := &models.Seller{
		Name:      "franco",
		CompanyID: &company1.ID,
	}
	_ = db.Create(seller1).Error
	seller2 := &models.Seller{
		Name:      "agustin",
		CompanyID: &company2.ID,
	}
	_ = db.Create(seller2).Error

	sale1 := &models.Sale{
		Product: product1,
		Seller:  seller1,
	}
	_ = db.Create(sale1).Error
	sale2 := &models.Sale{
		Product: product2,
		Seller:  seller2,
	}
	_ = db.Create(sale2).Error

	logger.Sugar().Info("Finished creating CRUD example")
}
