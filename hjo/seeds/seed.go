package seeds

import (
	"errors"
	"hjo/database"
	"hjo/models"
	"hjo/logger"
	"go.uber.org/zap"

	"gorm.io/gorm"
)

func RunSeedsData(migrations database.Migrations) {
	InsertDummyUsersData(migrations)
}

func InsertDummyUsersData(migrations database.Migrations) {
	var userDummyData = []models.GoEasyCLITestUser{
		{
			Name:  "Fafa M",
			Email: "fm@fafa.com",
		},
		{
			Name:  "M Fafa",
			Email: "test@mfafa.com",
		},
		{
			Name:  "MF",
			Email: "hey@m.com",
		},
	}

	if migrations.DB.Migrator().HasTable(&models.GoEasyCLITestUser{}) {

		if err := migrations.DB.First(&models.GoEasyCLITestUser{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {

			logger.Info("dummy user seed data", zap.String("status", "start"))

			for _, v := range userDummyData {
				migrations.DB.Create(&models.GoEasyCLITestUser{
					Name:  v.Name,
					Email: v.Email,
				})
			}
			logger.Info("dummy user seed data", zap.String("status", "end"))
		}
	}
}
