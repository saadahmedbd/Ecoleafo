package categoryservice

import (
	"log"

	categorydto "github.com/saadahmedbd/Treestore/Rest/DTO/CategoryDTO"
)

func (s *categoryService) GetCategoriesForSeller() ([]categorydto.CategoryResponse, error) {
	categories, _, err := s.categoryRepo.FindAll(map[string]interface{}{}, 1, 1000)
	if err != nil {
		return nil, err
	}

	log.Printf("📦 Found %d categories for seller", len(categories))
	for _, cat := range categories {
		log.Printf("  - ID: %d, Name: %s, Active: %v", cat.ID, cat.Name, cat.IsActive)
	}

	responses := make([]categorydto.CategoryResponse, len(categories))
	for i, cat := range categories {
		resp, _ := s.mapToCategoryResponse(&cat)
		responses[i] = *resp
	}

	return responses, nil
}
