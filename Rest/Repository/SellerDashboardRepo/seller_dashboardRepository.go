package sellerdashboardrepo

import "gorm.io/gorm"

type DashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) *DashboardRepository {
	return &DashboardRepository{
		db: db,
	}
}
