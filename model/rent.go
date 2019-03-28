package model

import (
	"time"
	"github.com/pkg/errors"
	"github.com/jinzhu/gorm"
	"github.com/nclgh/lakawei_scaffold/rpc/common"
)

const RentTableName = "rent"

type Rent struct {
	Id                 int64     `gorm:"primary_key;not null;auto_increment"`
	DeviceCode         string    `gorm:"type:varchar(255);not null;default:''"`
	RentStatus         int64     `gorm:"type:bigint(20) unsigned;not null;default:0"`
	BorrowerMemberCode string    `gorm:"type:varchar(255);not null;default:''"`
	BorrowDate         time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	ExpectReturnDate   time.Time `gorm:"type:bigint(20) unsigned;not null;default:0"`
	ReturnerMemberCode string    `gorm:"type:varchar(255);not null;default:''"`
	RealReturnDate     time.Time `gorm:"type:bigint(20) unsigned;not null;default:0"`
	BorrowRemark       string    `gorm:"type:text;not null"`
	ReturnRemark       string    `gorm:"type:text;not null"`
	CreateTime         time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdateTime         time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func InsertRent(db *gorm.DB, record *Rent) error {
	record.RentStatus = DeviceRentStatusLend
	record.BorrowDate = time.Now()
	record.ReturnRemark = ""
	record.ReturnerMemberCode = ""
	record.RealReturnDate = time.Unix(0, 0)
	record.CreateTime = time.Now()
	record.UpdateTime = time.Now()
	err := db.Table(RentTableName).Create(record).Error
	if err != nil {
		return err
	}
	return UpdateDeviceRentStatus(db, record.DeviceCode, DeviceRentStatusLend)
}

func ReturnRent(db *gorm.DB, record *Rent) error {
	record.RentStatus = DeviceRentStatusUnLend
	record.RealReturnDate = time.Now()
	up := db.Table(RentTableName).Where("device_code = ? AND rent_status = ?", record.DeviceCode, DeviceRentStatusLend).Update(record)
	err := up.Error
	if err != nil {
		return err
	}
	rowEffect := up.RowsAffected
	if rowEffect <= 0 {
		return errors.New("device rent record not found")
	}
	return UpdateDeviceRentStatus(db, record.DeviceCode, DeviceRentStatusUnLend)
}

func QueryRent(db *gorm.DB, m *Rent, page, pageSize int64, filter *common.Filter) ([]*Rent, int64, error) {
	retDept := make([]*Rent, 0)
	totalCnt := int64(0)
	query := db.Table(RentTableName).Where(m)
	query = filter.Filter(query)
	err := query.Offset(page * pageSize).Limit(pageSize).Find(&retDept).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Count(&totalCnt).Error
	if err != nil {
		return nil, 0, err
	}
	return retDept, totalCnt, nil
}
