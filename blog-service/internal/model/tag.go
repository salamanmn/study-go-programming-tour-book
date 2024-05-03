package model

import (
	"gorm.io/gorm"
	"time"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

// 对模型操作进行封装
func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int64
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, values interface{}) error {
	db = db.Model(&Tag{}).Where("id = ? AND is_del = ?", t.ID, 0)
	return db.Updates(values).Error
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Delete(&t).Error
}

/*GORM v2版本剔除了gorm.Scope类型，
为了对公共字段created_on、modified_on、deleted_on、is_del不进行重复操作，
采用v2版本中的hook狗子函数。
*/

// BeforeCreate 新增行为的hook函数
func (t Tag) BeforeCreate(db *gorm.DB) (err error) {
	db.Statement.SetColumn("created_on", uint32(time.Now().Unix()))
	return nil
}

// BeforeUpdate 更新行为的hook函数
func (t Tag) BeforeUpdate(db *gorm.DB) (err error) {
	db.Statement.SetColumn("modified_on", uint32(time.Now().Unix()))
	return nil
}

// BeforeDelete 删除行为的hook函数
func (t Tag) BeforeDelete(db *gorm.DB) (err error) {
	db.Statement.SetColumn("deleted_on", uint32(time.Now().Unix()))
	db.Statement.SetColumn("is_del", 1)
	return nil
}
