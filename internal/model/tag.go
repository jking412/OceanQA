package model

type Tag struct {
	Id   uint64 `gorm:"column:id;primaryKey;autoIncrement"`
	Name string `gorm:"column:name;type:varchar(255)"`
}

func (t *Tag) TableName() string {
	return "tags"
}
