package handlers

import "go_fiber/controller"


func (h *MyHandler) CheckHandler() controller.ICheckController {
	
	checkController := controller.NewCheckController()
	return checkController
}