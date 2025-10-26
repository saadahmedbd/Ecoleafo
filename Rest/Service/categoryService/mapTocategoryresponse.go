package categoryservice

import (
	models "github.com/saadahmedbd/Treestore/Models"
	categorydto "github.com/saadahmedbd/Treestore/Rest/DTO/CategoryDTO"
)

// mapToCategoryResponse converts model to response DTO
func (s *categoryService) mapToCategoryResponse(category *models.Category) (*categorydto.CategoryResponse, error) {
	// Get product count
	productCount, _ := s.categoryRepo.GetProductCount(category.ID)

	resp := &categorydto.CategoryResponse{
		ID:              category.ID,
		Name:            category.Name,
		Slug:            category.Slug,
		Description:     category.Description,
		Image:           category.Image,
		Icon:            category.Icon,
		ParentID:        category.ParentID,
		SortOrder:       category.SortOrder,
		IsFeatured:      category.IsFeatured,
		IsActive:        category.IsActive,
		ProductCount:    productCount,
		MetaTitle:       category.MetaTitle,
		MetaDescription: category.MetaDescription,
		MetaKeywords:    category.MetaKeywords,
		CreatedAt:       category.CreatedAt,
		UpdatedAt:       category.UpdatedAt,
	}

	// Map parent if exists
	if category.Parent != nil {
		resp.Parent = &categorydto.CategoryBriefResponse{
			ID:   category.Parent.ID,
			Name: category.Parent.Name,
			Slug: category.Parent.Slug,
			Icon: category.Parent.Icon,
		}
	}

	// Map children
	if len(category.Children) > 0 {
		children := make([]categorydto.CategoryBriefResponse, len(category.Children))
		for i, child := range category.Children {
			children[i] = categorydto.CategoryBriefResponse{
				ID:         child.ID,
				Name:       child.Name,
				Slug:       child.Slug,
				Image:      child.Image,
				Icon:       child.Icon,
				IsActive:   child.IsActive,
				IsFeatured: child.IsFeatured,
			}
		}
		resp.Children = children
	}

	return resp, nil
}
