package models_test

import (
	"github.com/hublabs/stock-base-info-agent/models"

	"github.com/go-xorm/xorm"
)

var (
	committed = models.Committed{
		CreatedBy: "li.dongxun",
	}
	SeedBatchJob = []models.MngBatchJob{
		{
			TenantCode: "hublabs",
			BatchName:  string(models.Migration),
			JobName:    string(models.LocationStore),
			IsUsed:     false,
			Committed:  committed,
		},
	}
)

func CreateSeedData(xormEngine *xorm.Engine) error {
	if _, err := xormEngine.Table(models.TableNameForMngBatchJob).Insert(&SeedBatchJob); err != nil {
		return err
	}
	return nil
}
