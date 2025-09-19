package database

import (
	"fmt"

	models "github.com/saadahmedbd/Treestore/Models"
	util "github.com/saadahmedbd/Treestore/Util"
	"gorm.io/gorm"
)

// intial data(seed data)
func SeedData(db *gorm.DB) error {
	fmt.Println("Seeding database with initial data...")
	//hash password
	adminPassword, _ := util.HashPassword("admin123")
	sellerPassword, _ := util.HashPassword("seller123")
	buyerPassword, _ := util.HashPassword("buyer123")

	//create roles

	// Create roles
	roles := []models.Role{
		{Name: "admin", Description: "Administrator with full access"},
		{Name: "seller", Description: "Sellers who can manage products"},
		{Name: "buyer", Description: "Customers who can place orders"},
	}
	for _, role := range roles {
		var existingRole models.Role
		result := db.Where("name = ?", role.Name).First(&existingRole)
		if result.Error != nil {
			if err := db.Create(&role).Error; err != nil {
				return fmt.Errorf("failed to create role %s: %v", role.Name, err)
			}
			fmt.Printf(" Created role: %s\n", role.Name)
		} else {
			fmt.Printf(" Role already exists: %s\n", role.Name)
		}
	}
	// Get role IDs
	var adminRole, sellerRole, buyerRole models.Role
	db.Where("name = ?", "admin").First(&adminRole)
	db.Where("name = ?", "seller").First(&sellerRole)
	db.Where("name = ?", "buyer").First(&buyerRole)

	// Create admin user
	adminUser := models.User{
		FirstName:  "Admin",
		LastName:   "User",
		Email:      "admin@example.com",
		Password:   adminPassword,
		Phone:      "+1234567890",
		StoreName:  "Admin Store",
		StoreDesc:  "Administrative store",
		IsVerified: true,
	}
	var existingAdmin models.User
	result := db.Where("email = ?", adminUser.Email).First(&existingAdmin)
	if result.Error != nil {
		if err := db.Create(&adminUser).Error; err != nil {
			return fmt.Errorf("failed to create admin user: %v", err)
		}
		fmt.Printf("✅ Created admin user: %s\n", adminUser.Email)
	} else {
		fmt.Printf("ℹ️  Admin user already exists: %s\n", adminUser.Email)
	}

	// Create sample sellers
	sellers := []models.User{
		{
			FirstName: "John",
			LastName:  "Electronics",
			Email:     "john@electronics.com",
			Password:  sellerPassword,
			Phone:     "+1234567891",
			StoreName: "John's Electronics",
			StoreDesc: "Best electronics and gadgets",

			IsVerified: true,
		},
		{
			FirstName: "Sarah",
			LastName:  "Fashion",
			Email:     "sarah@fashion.com",
			Password:  sellerPassword,
			Phone:     "+1234567892",
			StoreName: "Sarah's Fashion",
			StoreDesc: "Trendy clothes and accessories",

			IsVerified: true,
		},
	}
	for _, seller := range sellers {
		var existingSeller models.User
		result := db.Where("email = ?", seller.Email).First(&existingSeller)
		if result.Error != nil {
			if err := db.Create(&seller).Error; err != nil {
				return fmt.Errorf("failed to create seller %s: %v", seller.Email, err)
			}
			fmt.Printf("✅ Created seller: %s\n", seller.Email)
		} else {
			fmt.Printf("ℹ️  Seller already exists: %s\n", seller.Email)
		}
	}
	// Create sample buyers
	buyers := []models.Buyer{
		{
			FirstName:      "Jane",
			LastName:       "Customer",
			Email:          "jane@customer.com",
			Password:       buyerPassword,
			Phone:          "+1234567893",
			DefaultAddress: "123 Main St, City, State 12345",
			EmailVerified:  true,
		},
	}

	for _, buyer := range buyers {
		var existingBuyer models.Buyer
		result := db.Where("email = ?", buyer.Email).First(&existingBuyer)
		if result.Error != nil {
			if err := db.Create(&buyer).Error; err != nil {
				return fmt.Errorf("failed to create buyer %s: %v", buyer.Email, err)
			}
			fmt.Printf("✅ Created buyer: %s\n", buyer.Email)
		} else {
			fmt.Printf("ℹ️  Buyer already exists: %s\n", buyer.Email)
		}
	}

	// Create categories
	categories := []models.Category{
		{Name: "Electronics", Slug: "electronics", IsActive: true},
		{Name: "Clothing", Slug: "clothing", IsActive: true},
		{Name: "Books", Slug: "books", IsActive: true},
		{Name: "Home & Garden", Slug: "home-garden", IsActive: true},
		{Name: "Sports", Slug: "sports", IsActive: true},
	}

	for _, category := range categories {
		var existingCategory models.Category
		result := db.Where("slug = ?", category.Slug).First(&existingCategory)
		if result.Error != nil {
			if err := db.Create(&category).Error; err != nil {
				return fmt.Errorf("failed to create category %s: %v", category.Name, err)
			}
			fmt.Printf("✅ Created category: %s\n", category.Name)
		} else {
			fmt.Printf("ℹ️  Category already exists: %s\n", category.Name)
		}
	}
	// Get seller and category IDs for products
	var johnSeller, sarahSeller models.User
	var electronicsCategory, clothingCategory models.Category

	db.Where("email = ?", "john@electronics.com").First(&johnSeller)
	db.Where("email = ?", "sarah@fashion.com").First(&sarahSeller)
	db.Where("slug = ?", "electronics").First(&electronicsCategory)
	db.Where("slug = ?", "clothing").First(&clothingCategory)

	// Create sample products
	products := []models.Product{
		{
			SellerID:    johnSeller.ID,
			Name:        "iPhone 15 Pro",
			Slug:        "iphone-15-pro",
			Description: "Latest iPhone with advanced features",
			SKU:         "IP15PRO001",
			CategoryID:  electronicsCategory.ID,
			Price:       999.99,
			Quantity:    50,
			Image:       "https://example.com/iphone15pro.jpg",
			IsActive:    true,
			IsApproved:  true,
		},
		{
			SellerID:    johnSeller.ID,
			Name:        "MacBook Air M3",
			Slug:        "macbook-air-m3",
			Description: "Powerful laptop for professionals",
			SKU:         "MBA-M3-001",
			CategoryID:  electronicsCategory.ID,
			Price:       1299.99,
			Quantity:    25,
			Image:       "https://example.com/macbook-air.jpg",
			IsActive:    true,
			IsApproved:  true,
		},
		{
			SellerID:    sarahSeller.ID,
			Name:        "Designer T-Shirt",
			Slug:        "designer-t-shirt",
			Description: "Premium quality cotton t-shirt",
			SKU:         "TSH-DES-001",
			CategoryID:  clothingCategory.ID,
			Price:       49.99,
			Quantity:    100,
			Image:       "https://example.com/t-shirt.jpg",
			IsActive:    true,
			IsApproved:  true,
		},
	}

	for _, product := range products {
		var existingProduct models.Product
		result := db.Where("sku = ?", product.SKU).First(&existingProduct)
		if result.Error != nil {
			if err := db.Create(&product).Error; err != nil {
				return fmt.Errorf("failed to create product %s: %v", product.Name, err)
			}
			fmt.Printf("✅ Created product: %s\n", product.Name)
		} else {
			fmt.Printf("ℹ️  Product already exists: %s\n", product.Name)
		}
	}

	fmt.Println("✅ Database seeded successfully!")
	return nil
}
