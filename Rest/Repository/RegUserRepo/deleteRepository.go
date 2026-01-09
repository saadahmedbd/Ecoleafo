package reguserrepo

func (r *regUserRepository) Delete(id uint) error {
	return r.db.Delete(id).Error
}
