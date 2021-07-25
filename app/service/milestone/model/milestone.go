package model

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type (
	MilestoneModel interface {
		Insert(data *Milestone) error
		Delete(primaryKey uint) error
		Update(data *Milestone) error
		SearchTitles(keyword string, limit uint) ([]*Milestone, error)
		FindOne(primaryKey uint) (*Milestone, error)
		FindAllByTimestamp(startTimestamp, limit uint, order string) ([]*Milestone, error)
	}

	defaultMilestoneModel struct {
		conn *gorm.DB
	}

	Milestone struct {
		Id        uint   `json:"id"`
		Title     string `json:"title"`
		Subtitled string `json:"subtitled"`
		Type      string `json:"type"`
		Content   string `json:"content"`
		TargetUrl string `json:"target_url"`
		Timestamp uint   `json:"timestamp"`
		CreatedAt uint   `json:"created_at" gorm:"autoCreateTime:milli"`
		UpdatedAt uint   `json:"updated_at" gorm:"autoUpdateTime:milli"`
		DeletedAt uint   `json:"deleted_at" gorm:"index:idx_deleted"`
	}
)

func NewMilestoneModel(conn *gorm.DB) MilestoneModel {
	return &defaultMilestoneModel{conn: conn}
}

func (m *defaultMilestoneModel) Insert(data *Milestone) error {
	result := m.conn.Table("milestones").Create(data)
	return result.Error
}

func (m *defaultMilestoneModel) Delete(primaryKey uint) error {
	result := m.conn.Exec("UPDATE milestones SET deleted_at = ? WHERE id = ? AND deleted_at = 0", time.Now().UnixNano()/1e6, primaryKey)
	return result.Error
}

func (m *defaultMilestoneModel) Update(data *Milestone) error {
	result := m.conn.Table("milestones").Updates(data)
	return result.Error
}

func (m *defaultMilestoneModel) SearchTitles(keyword string, limit uint) ([]*Milestone, error) {
	var list []*Milestone
	result := m.conn.Raw("SELECT * FROM milestones WHERE Title LIKE ? AND deleted_at = 0 LIMIT 0, ?", keyword+"%", limit).Find(&list)
	return list, result.Error
}

func (m *defaultMilestoneModel) FindOne(primaryKey uint) (*Milestone, error) {
	milestone := &Milestone{}
	result := m.conn.Raw("SELECT * FROM milestones WHERE id = ? AND deleted_at = 0", primaryKey).First(milestone)
	if result.Error != nil {
		return nil, result.Error
	}
	return milestone, nil
}

func (m *defaultMilestoneModel) FindAllByTimestamp(startTimestamp, limit uint, order string) ([]*Milestone, error) {
	var list []*Milestone
	result := m.conn.Raw(fmt.Sprintf("SELECT * FROM milestones WHERE timestamp < ? AND deleted_at = 0 ORDER BY timestamp %s LIMIT 0, ?", order), startTimestamp, limit).Find(&list)
	return list, result.Error
}
