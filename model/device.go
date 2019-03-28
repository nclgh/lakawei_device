package model

import (
	"time"
	"github.com/jinzhu/gorm"
	"github.com/nclgh/lakawei_scaffold/rpc/common"
	"fmt"
)

const DeviceTableName = "device"

type Device struct {
	Id               int64     `gorm:"primary_key;not null;auto_increment"`
	Code             string    `gorm:"type:varchar(255);not null;default:''"`
	Name             string    `gorm:"type:varchar(255);not null;default:''"`
	Model            string    `gorm:"type:varchar(255);not null;default:''"`
	Brand            string    `gorm:"type:varchar(255);not null;default:''"`
	TagCode          string    `gorm:"type:varchar(255);not null;default:''"`
	DepartmentCode   string    `gorm:"type:varchar(255);not null;default:''"`
	ManufacturerId   int64     `gorm:"type:bigint(20) unsigned;not null;default:0"`
	ManufacturerDate time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	RentStatus       int64     `gorm:"type:tinyint(1) unsigned;not null;default:0"`
	Description      string    `gorm:"type:text;not null"`
	CreateTime       time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdateTime       time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func InsertDevice(db *gorm.DB, record *Device) error {
	record.RentStatus = DeviceRentStatusUnLend
	record.CreateTime = time.Now()
	record.UpdateTime = time.Now()
	err := db.Table(DeviceTableName).Create(record).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteDevice(db *gorm.DB, code string) error {
	return db.Table(DeviceTableName).Delete(&Device{}, "code = ?", code).Error
}

func QueryDevice(db *gorm.DB, m *Device, page, pageSize int64, filter *common.Filter) ([]*Device, int64, error) {
	retDept := make([]*Device, 0)
	totalCnt := int64(0)
	query := db.Table(DeviceTableName).Where(m)
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

func UpdateDeviceRentStatus(db *gorm.DB, deviceCode string, rentStatus int64) error {
	d := &Device{}
	err := db.Table(DeviceTableName).Where("code = ?", deviceCode).First(&d).Error
	if err != nil {
		return err
	}
	if d.RentStatus == rentStatus {
		return fmt.Errorf("device status is %v already", rentStatus)
	}
	err = db.Table(DeviceTableName).Where("code = ?", deviceCode).Update(&Device{RentStatus: rentStatus}).Error
	if err != nil {
		return err
	}
	return nil
}

func GetDeviceByCode(db *gorm.DB, codes []string) ([]*Device, error) {
	ret := make([]*Device, 0)
	err := db.Where("code in (?)", codes).Find(&ret).Error
	if err != nil {
		return nil, err
	}
	return ret, err
}
