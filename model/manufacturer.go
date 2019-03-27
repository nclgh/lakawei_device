package model

import (
	"time"
	"github.com/jinzhu/gorm"
)

const ManufacturerTableName = "manufacturer"

type Manufacturer struct {
	Id           int64     `gorm:"primary_key;not null;auto_increment"`
	Name         string    `gorm:"type:varchar(255);not null;default:''"`
	CreateTime   time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdateTime   time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func InsertManufacturer(db *gorm.DB, name string) error {
	record := &Manufacturer{
		Name:         name,
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
	}
	err := db.Table(ManufacturerTableName).Create(record).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteManufacturer(db *gorm.DB, id int64) error {
	return db.Table(ManufacturerTableName).Delete(&Manufacturer{}, "id = ?", id).Error
}

func QueryManufacturer(db *gorm.DB, m *Manufacturer, page, pageSize int64) ([]*Manufacturer, int64, error) {
	retDept := make([]*Manufacturer, 0)
	totalCnt := int64(0)
	err := db.Table(ManufacturerTableName).Where(m).Offset(page * pageSize).Limit(pageSize).Find(&retDept).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Table(ManufacturerTableName).Where(m).Count(&totalCnt).Error
	if err != nil {
		return nil, 0, err
	}
	return retDept, totalCnt, nil
}

func GetManufacturerById(db *gorm.DB, ids []int64) ([]*Manufacturer, error) {
	ret := make([]*Manufacturer, 0)
	err := db.Where("id in (?)", ids).Find(&ret).Error
	if err != nil {
		return nil, err
	}
	return ret, err
}
