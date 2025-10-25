package Config

import cloudniaryservice "github.com/saadahmedbd/Treestore/Rest/Service/CloudniaryService"

func InitializeCloudinary() *cloudniaryservice.CloudniaryService {
	cfg := GetConfig()

	config := cloudniaryservice.CloudinaryConfig{
		CloudName: cfg.CloudinaryCloudName,
		APIKey:    cfg.CloudinaryAPIKey,
		APISecret: cfg.CloudinaryAPISecret,
	}

	return cloudniaryservice.NewCloudinaryService(config)
}
