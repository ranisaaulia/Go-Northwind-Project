package controllers

import "codeid.northwind/services"

type ControllerManager struct {
	CategoryController
}

// Constructor

func NewControllerManager(serviceMgr *services.ServiceManager) *ControllerManager {
	return &ControllerManager{
		*NewCategoryController(&serviceMgr.CategoryService),
	}
}
