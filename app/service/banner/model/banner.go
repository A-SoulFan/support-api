package model

import "gorm.io/gorm"

const (
	bannerTableName = "banner"
)

type (
	BannerModel interface {
		FindAllByType(t string) ([]*Banner, error)
	}

	defaultBannerModel struct {
		conn *gorm.DB
	}

	Banner struct {
		Id        uint   `json:"id"`
		Type      string `json:"type"`
		Sort      uint   `json:"sort"`
		Url       string `json:"url"`
		Title     string `json:"title"`
		Desc      string `json:"desc"`
		Content   string `json:"content"`
		DeletedAt uint   `json:"-"`
	}
)

func NewDefaultBannerModel(conn *gorm.DB) BannerModel {
	return &defaultBannerModel{conn: conn}
}

func (m *defaultBannerModel) FindAllByType(t string) ([]*Banner, error) {
	var list []*Banner
	result := m.conn.Table(bannerTableName).Where("type = ? AND deleted_at = 0", t).Order("sort").Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}
