package main

import (
	"github.com/ditrit/badaas/persistence/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func StartExample(logger *zap.Logger, db *gorm.DB) error {
	logger.Sugar().Info("Setting up Posts example")

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

	err := db.Create(adminProfile).Error
	if err != nil {
		return err
	}

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

	err = db.Create(whyCatsLikeMice).Error
	if err != nil {
		return err
	}

	logger.Sugar().Info("Finished populating the database")

	return nil
}
