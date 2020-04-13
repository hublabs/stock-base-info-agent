package models

import (
	"context"
	"fmt"

	"github.com/hublabs/stock-base-info-agent/factory"
)

type MngBatchJob struct {
	Id         int64  `json:"id"`
	TenantCode string `json:"tenant_code" xorm:"index unique(batch)"`
	BatchName  string `json:"batch_name" xorm:"index unique(batch)"`
	JobName    string `json:"job_name " xorm:"index unique(batch)"`
	IsUsed     bool   `json:"is_used "`
	Committed  `xorm:"extends"`
}

const TableNameForMngBatchJob = "mng_batch_job"

func (MngBatchJob) TableName() string {
	return TableNameForMngBatchJob
}

func (MngBatchJob) Find() ([]MngBatchJob, error) {
	var results []MngBatchJob
	if err := factory.
		GetDefaultMysqlEngine().
		Table(TableNameForMngBatchJob).
		Find(&results); err != nil {
		return nil, err
	}
	return results, nil
}

func (MngBatchJob) GetByBatchAndJobName(batch, job string) (MngBatchJob, error) {
	var batchJobInfo MngBatchJob
	if _, err := factory.
		GetDefaultMysqlEngine().
		Table(TableNameForMngBatchJob).
		Where("batch_name = ? and job_name = ?", batch, job).
		Get(&batchJobInfo); err != nil {
		return MngBatchJob{}, err
	}
	return batchJobInfo, nil
}

func (MngBatchJob) getById(id int64) (MngBatchJob, error) {
	var batchJobInfo MngBatchJob
	if _, err := factory.
		GetDefaultMysqlEngine().
		Table(TableNameForMngBatchJob).
		Where("id = ?", id).
		Get(&batchJobInfo); err != nil {
		return MngBatchJob{}, err
	}
	return batchJobInfo, nil
}

func (batchJob *MngBatchJob) Create(ctx context.Context) (*MngBatchJob, error) {
	if _, err := factory.
		GetDefaultMysqlEngine().
		Insert(batchJob); err != nil {
		return nil, err
	}
	return batchJob, nil
}

func (MngBatchJob) ToggleJobById(id int64, updatedBy string) (MngBatchJob, error) {
	batchJobInfo, err := MngBatchJob{}.getById(id)
	if err != nil {
		return MngBatchJob{}, nil
	}
	var used int64
	if batchJobInfo.IsUsed {
		used = 0
		batchJobInfo.IsUsed = false
	} else {
		used = 1
		batchJobInfo.IsUsed = true
	}

	sql := fmt.Sprintf(`update mng_batch_job_info set is_used = '%v', updated_by = '%s' where id = '%v'`,
		used, updatedBy, id)
	if _, err := factory.GetDefaultMysqlEngine().Exec(sql); err != nil {
		return MngBatchJob{}, err
	}

	return batchJobInfo, nil
}

func (MngBatchJob) DeleteById(id int64) (MngBatchJob, error) {
	batchJobInfo, err := MngBatchJob{}.getById(id)
	if err != nil {
		return MngBatchJob{}, nil
	}
	if _, err := factory.
		GetDefaultMysqlEngine().
		Where("id = ?", id).
		Delete(batchJobInfo); err != nil {
		return MngBatchJob{}, nil
	}
	return batchJobInfo, nil
}
