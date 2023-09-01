package services

import "codeid.northwind/repositories"

type ServiceManager struct {
	CategoryService
	// tabel lain
}

// Constructor
func NewServiceManager(repoMgr *repositories.RepositoryManager) *ServiceManager {
	return &ServiceManager{
		CategoryService: *NewCategoryService(&repoMgr.CategoryRepository),
		// tabel lain
	}
}
