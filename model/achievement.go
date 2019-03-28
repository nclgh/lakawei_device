package model

import (
	"time"
	"github.com/jinzhu/gorm"
	"github.com/nclgh/lakawei_scaffold/rpc/common"
)

const AchievementTableName = "achievement"

type Achievement struct {
	Id                     int64     `gorm:"primary_key;not null;auto_increment"`
	DeviceCode             string    `gorm:"type:varchar(255);not null;default:''"`
	MemberCode             string    `gorm:"type:varchar(255);not null;default:''"`
	DepartmentCode         string    `gorm:"type:varchar(255);not null;default:''"`
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

func QueryAchievement(db *gorm.DB, m *Achievement, page, pageSize int64, filter *common.Filter) ([]*Achievement, int64, error) {
	retDept := make([]*Achievement, 0)
	totalCnt := int64(0)
	query := db.Table(AchievementTableName).Where(m)
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

func GetAchievementById(db *gorm.DB, ids []int64) ([]*Achievement, error) {
	ret := make([]*Achievement, 0)
	err := db.Where("id in (?)", ids).Find(&ret).Error
	if err != nil {
		return nil, err
	}
	return ret, err
}
