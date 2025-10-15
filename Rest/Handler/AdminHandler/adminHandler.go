package adminhandler

import adminservice "github.com/saadahmedbd/Treestore/Rest/Service/AdminService"

type Adminhandler struct {
	adminService adminservice.Adminservice
}

func NewAdminHandler(adminService adminservice.Adminservice) *Adminhandler {
	return &Adminhandler{
		adminService: adminService,
	}
}
