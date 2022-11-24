package models

import "gorm.io/gorm"

type OSSFile struct {
	CreatorBase
	SaveName string `json:"saveName" gorm:"column:save_name;comment: 文件名"`
	SavePath string `json:"savePath" gorm:"column:save_path;comment: 文件路径"`
	FileType string `json:"fileType" gorm:"column:file_type;comment: 文件类型"`
}

func (t *OSSFile) TableName() string {
	return "oss_files"
}

func (t *OSSFile) Insert(db *gorm.DB) error {
	return db.Model(&t).Create(&t).Error
}
