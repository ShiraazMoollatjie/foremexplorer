package db

import "github.com/jinzhu/gorm"

type Cursor struct {
	gorm.Model
	Name  string `gorm:"uniqueIndex:uniq_name"`
	Value int
}

// SetCursor saves a cursor to keep track of iterations in the db. This is not thread safe.
func SetCursor(db *gorm.DB, name string, value int) error {
	c, err := GetCursor(db, name)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err
	}

	if gorm.IsRecordNotFoundError(err) {
		return db.Create(&Cursor{
			Name:  name,
			Value: value,
		}).Error
	} else {
		return db.Model(&c).Update("value", value).Error
	}
}

func GetCursor(db *gorm.DB, name string) (*Cursor, error) {
	var c Cursor
	r := db.Where(map[string]interface{}{"Name": name}).First(&c)

	return &c, r.Error
}
