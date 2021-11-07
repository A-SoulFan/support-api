package model

import "gorm.io/gorm"

const (
	keyValueTableName = "key_val"
)

type (
	KeyValueModel interface {
		FindAllByKey(key string) ([]*KeyValue, error)
		FindOneByKey(key string) (*KeyValue, error)
	}

	defaultKeyValueModel struct {
		conn *gorm.DB
	}

	KeyValue struct {
		Id        uint   `json:"id"`
		Name      string `json:"name"`
		Key       string `json:"key"`
		Value     []byte `json:"value"`
		Sort      uint   `json:"sort"`
		DeletedAt uint   `json:"-"`
	}
)

func NewDefaultKeyValueModel(conn *gorm.DB) *defaultKeyValueModel {
	return &defaultKeyValueModel{conn: conn}
}

func (m *defaultKeyValueModel) FindAllByKey(key string) ([]*KeyValue, error) {
	var list []*KeyValue
	result := m.conn.Table(keyValueTableName).Where("key = ? AND deleted_at = 0", key).
		Order("sort").
		Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}

	return list, nil
}

func (m *defaultKeyValueModel) FindOneByKey(key string) (*KeyValue, error) {
	var data *KeyValue
	result := m.conn.Table(keyValueTableName).Where("`key` = ? AND deleted_at = 0", key).
		Order("sort").
		Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}
