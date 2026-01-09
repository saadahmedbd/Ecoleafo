package productservice

import (
	models "github.com/saadahmedbd/Treestore/Models"
	"gorm.io/gorm"
)

func (s *ProductService) IncrementViewCount(id uint) error {
	return s.db.Model(&models.Product{}).Where("id = ?", id).Update("view_count", gorm.Expr("view_count + 1")).Error
}
