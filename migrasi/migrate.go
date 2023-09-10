package migrasi

import (
	"backend/domain/model"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&model.Mahasiswa{})
	db.AutoMigrate(&model.Dosen{})
	db.AutoMigrate(&model.MataKuliah{})
	db.AutoMigrate(&model.Kuliah{})

}
