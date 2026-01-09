package categoryservice

import (
	models "github.com/saadahmedbd/Treestore/Models"
	categorydto "github.com/saadahmedbd/Treestore/Rest/DTO/CategoryDTO"
)

// buildTree recursively builds category tree
func (s *categoryService) buildTree(categories []models.Category, parentID *uint) []categorydto.CategoryTreeResponse {
	var tree []categorydto.CategoryTreeResponse

	for _, cat := range categories {
		// Check if this category matches the parent
		if (parentID == nil && cat.ParentID == nil) ||
			(parentID != nil && cat.ParentID != nil && *cat.ParentID == *parentID) {
			node := categorydto.CategoryTreeResponse{
				ID:   cat.ID,
				Name: cat.Name,
				Slug: cat.Slug,
				Icon: cat.Icon,
			}

			// Recursively build children
			node.Children = s.buildTree(categories, &cat.ID)
			tree = append(tree, node)
		}
	}

	return tree
}
