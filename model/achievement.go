package model

import (
	"time"
	"github.com/jinzhu/gorm"
	"github.com/nclgh/lakawei_scaffold/rpc/device"
	"fmt"
)

const AchievementTableName = "achievement"

type Achievement struct {
	Id                     int64     `gorm:"primary_key;not null;auto_increment"`
	DeviceId               int64     `gorm:"type:bigint(20) unsigned;not null;default:0"`
	MemberId               int64     `gorm:"type:bigint(20) unsigned;not null;default:0"`
	DepartmentId           int64     `gorm:"type:bigint(20) unsigned;not null;default:0"`
	AchievementDate        time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	AchievementDescription string    `gorm:"type:text;not null"`
	AchievementRemark      string    `gorm:"type:text;not null"`
	PatentDescription      string    `gorm:"type:text;not null"`
	PaperDescription       string    `gorm:"type:text;not null"`
	CompetitionDescription string    `gorm:"type:text;not null"`
	CreateTime             time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdateTime             time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func InsertAchievement(db *gorm.DB, record *Achievement) error {
	record.CreateTime = time.Now()
	record.UpdateTime = time.Now()
	err := db.Table(AchievementTableName).Create(record).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteAchievement(db *gorm.DB, id int64) error {
	return db.Table(AchievementTableName).Delete(&Achievement{}, "id = ?", id).Error
}

func QueryAchievement(db *gorm.DB, m *Achievement, page, pageSize int64, timeFilter []*device.TimeFilter) ([]*Achievement, int64, error) {
	retDept := make([]*Achievement, 0)
	totalCnt := int64(0)
	query := db.Table(AchievementTableName).Where(m)
	tfTemplate := "%s >= ? AND %s <= ?"
	for _, v := range timeFilter {
		tf := fmt.Sprintf(tfTemplate, v.Field, v.Field)
		query = query.Where(tf, v.Start, v.End)
	}
	err := query.Offset(page * pageSize).Limit(pageSize).Find(&retDept).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Table(AchievementTableName).Where(m).Count(&totalCnt).Error
	if err != nil {
		return nil, 0, err
	}
	return retDept, totalCnt, nil
}

func GetAchievementById(db *gorm.DB, ids []int64) ([]*Achievement, error) {
	ret := make([]*Achievement, 0)
	err := db.Where("id in (?)", ids).Find(&ret).Error
	if err != nil {
		return nil, err
	}
	return ret, err
}
