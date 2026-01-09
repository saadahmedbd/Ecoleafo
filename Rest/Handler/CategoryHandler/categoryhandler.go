package categoryHandler

import categoryservice "github.com/saadahmedbd/Treestore/Rest/Service/categoryService"

type CategoryHandler struct {
	categoryService categoryservice.CategoryService
}

// NewCategoryHandler creates a new category handler instance
func NewCategoryHandler(categoryService categoryservice.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}
