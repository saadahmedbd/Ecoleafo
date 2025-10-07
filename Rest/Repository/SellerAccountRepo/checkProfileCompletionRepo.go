package selleraccountrepo

import models "github.com/saadahmedbd/Treestore/Models"

func (r *sellerRegistrationRepository) CheckProfileCompletion(sellerID uint) (bool, []string, error) {
	var seller models.User
	err := r.db.Where("id = ?", sellerID).First(&seller).Error
	if err != nil {
		return false, nil, err
	}
	var missingFields []string

	//check business info
	if seller.Phone == "" {
		missingFields = append(missingFields, "Phone")
	}
	if seller.BusinessEmail == "" {
		missingFields = append(missingFields, "Business Email")
	}
	if seller.Address == "" {
		missingFields = append(missingFields, "Address")
	}
	if seller.City == "" {
		missingFields = append(missingFields, "City")
	}
	if seller.State == "" {
		missingFields = append(missingFields, "state")
	}
	if seller.PostalCode == "" {
		missingFields = append(missingFields, "postal_code")
	}

	// check payment method
	var paymentMethodsCount int64
	err = r.db.Model(&models.SellerPaymentMethod{}).
		Where("seller_id = ? AND is_active = ?", sellerID, true).
		Count(&paymentMethodsCount).Error
	if err != nil {
		return false, nil, err
	}
	if paymentMethodsCount == 0 {
		missingFields = append(missingFields, "At least one active payment method")
	}
	return len(missingFields) == 0, missingFields, nil

}
