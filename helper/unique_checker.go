package helper

import (
	"fmt"
	"gorm.io/gorm"
)

func IsUnique(model interface{}, fieldName string, value interface{}, con *gorm.DB) (bool, error) {
	var count int64
	query := fmt.Sprintf("%s = ?", fieldName)
	con.Model(model).Where(query, value).Count(&count)
	if count == 0 {
		return false, nil
	} else {
		err := fmt.Errorf(fieldName + " already exists")
		return true, err
	}
}
