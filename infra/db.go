package infra

import (
	"github.com/PolarPanda611/trinitygo/example/http/domain/object"
	"github.com/jinzhu/gorm"
)

// instantiate DB globally
var DB *gorm.DB

// Migrate structs defined under the entity folder
func Migrate() {
	DB.AutoMigrate(&object.User{})
}
