package categoryservice

// ReorderCategories updates sort order for multiple categories
func (s *categoryService) ReorderCategories(updates []struct {
	ID        uint
	SortOrder int
}) error {
	for _, update := range updates {
		if err := s.categoryRepo.UpdateSortOrder(update.ID, update.SortOrder); err != nil {
			return err
		}
	}
	return nil
}
