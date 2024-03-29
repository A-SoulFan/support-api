package model

import (
	"gorm.io/gorm"
	"time"
)

type (
	StrollModel interface {
		Insert(data *Stroll) error
		Delete(primaryKey uint) error
		Update(data *Stroll) error
		UpdateCover(bvId, cover string) error
		FindOne(primaryKey uint) (*Stroll, error)
		FindAllByIds(primaryKeyList []uint) ([]*Stroll, error)
		FindMaxId() (uint, error)
		FindLastUpdateTime() (uint, error)
	}

	defaultStrollModel struct {
		conn *gorm.DB
	}

	Stroll struct {
		Id        uint   `json:"id" gorm:"primaryKey"`
		Title     string `json:"title"`
		Cover     string `json:"cover"`
		BV        string `json:"bv" gorm:"uniqueIndex:uq_bv"`
		TargetUrl string `json:"target_url"`
		Play      string `json:"play"`
		CreatedAt uint   `json:"created_at" gorm:"autoCreateTime:milli"`
		UpdatedAt uint   `json:"updated_at" gorm:"autoUpdateTime:milli"`
		DeletedAt uint   `json:"deleted_at" gorm:"index:idx_deleted,uniqueIndex:uq_bv"`
	}
)

func NewStrollModel(conn *gorm.DB) StrollModel {
	return &defaultStrollModel{conn: conn}
}

func (m *defaultStrollModel) Insert(data *Stroll) error {
	result := m.conn.Table("stroll").Create(data)
	return result.Error
}

func (m *defaultStrollModel) Delete(primaryKey uint) error {
	result := m.conn.Exec("UPDATE stroll SET deleted_at = ? WHERE id = ? AND deleted_at = 0", time.Now().UnixNano()/1e6, primaryKey)
	return result.Error
}

func (m *defaultStrollModel) Update(data *Stroll) error {
	result := m.conn.Table("stroll").Updates(data)
	return result.Error
}

func (m *defaultStrollModel) UpdateCover(bvId, cover string) error {
	result := m.conn.Table("stroll").Where("bv = ?", bvId).Update("cover", cover)
	return result.Error
}

func (m *defaultStrollModel) FindOne(primaryKey uint) (*Stroll, error) {
	stroll := &Stroll{}
	result := m.conn.Raw("SELECT * FROM stroll WHERE id = ? AND deleted_at = 0", primaryKey).First(stroll)
	if result.Error != nil {
		return nil, result.Error
	}
	return stroll, nil
}

func (m *defaultStrollModel) FindAllByIds(primaryKeyList []uint) ([]*Stroll, error) {
	var strollList []*Stroll
	result := m.conn.Raw("SELECT * FROM stroll WHERE id IN (?) AND deleted_at = 0", primaryKeyList).Find(&strollList)
	return strollList, result.Error
}

func (m *defaultStrollModel) FindMaxId() (uint, error) {
	stroll := &Stroll{}
	result := m.conn.Raw("SELECT id FROM stroll WHERE deleted_at = 0 ORDER BY id DESC LIMIT 0, 1").Scan(&stroll)
	if result.Error != nil {
		return 0, result.Error
	}

	return stroll.Id, nil
}

func (m *defaultStrollModel) FindLastUpdateTime() (uint, error) {
	var stroll *Stroll
	result := m.conn.Raw("SELECT * FROM stroll WHERE deleted_at = 0 ORDER BY id DESC LIMIT 0, 1").Find(&stroll)
	if result.Error != nil {
		return 0, result.Error
	}

	return stroll.CreatedAt, nil
}
