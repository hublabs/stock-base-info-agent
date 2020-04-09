package worker

import (
	"context"
	"fmt"

	"github.com/pangpanglabs/goetl"
)

type MigrationLocationStoreWorker struct{}

func (MigrationLocationStoreWorker) New() *goetl.ETL {
	return goetl.New(MigrationLocationStoreWorker{})
}

func (etl MigrationLocationStoreWorker) Extract(ctx context.Context) (interface{}, error) {
	fmt.Println("extract from api or legacy database")
	return nil, nil
}

func (etl MigrationLocationStoreWorker) Transform(ctx context.Context, source interface{}) (interface{}, error) {
	fmt.Println("transform structure")
	return nil, nil

}

func (etl MigrationLocationStoreWorker) Load(ctx context.Context, source interface{}) error {
	fmt.Println("load to hublabs-delivery")
	return nil
}
