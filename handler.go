package main

import (
	"github.com/nclgh/lakawei_device/handler"
	"github.com/nclgh/lakawei_scaffold/rpc/device"
)

type ServiceDevice struct {
}

func (server *ServiceDevice) AddManufacturer(req device.AddManufacturerRequest, res *device.AddManufacturerResponse) error {
	resp := handler.AddManufacturer(&req)
	*res = *resp
	return nil
}

func (server *ServiceDevice) GetManufacturerById(req device.GetManufacturerByIdRequest, res *device.GetManufacturerByIdResponse) error {
	resp := handler.GetManufacturerById(&req)
	*res = *resp
	return nil
}

func (server *ServiceDevice) QueryManufacturer(req device.QueryManufacturerRequest, res *device.QueryManufacturerResponse) error {
	resp := handler.QueryManufacturer(&req)
	*res = *resp
	return nil
}

func (server *ServiceDevice) AddDevice(req device.AddDeviceRequest, res *device.AddDeviceResponse) error {
	resp := handler.AddDevice(&req)
	*res = *resp
	return nil
}

func (server *ServiceDevice) DeleteDevice(req device.DeleteDeviceRequest, res *device.DeleteDeviceResponse) error {
	resp := handler.DeleteDevice(&req)
	*res = *resp
	return nil
}

func (server *ServiceDevice) GetDeviceByCode(req device.GetDeviceByCodeRequest, res *device.GetDeviceByCodeResponse) error {
	resp := handler.GetDeviceByCode(&req)
	*res = *resp
	return nil
}

func (server *ServiceDevice) QueryDevice(req device.QueryDeviceRequest, res *device.QueryDeviceResponse) error {
	resp := handler.QueryDevice(&req)
	*res = *resp
	return nil
}

func (server *ServiceDevice) AddAchievement(req device.AddAchievementRequest, res *device.AddAchievementResponse) error {
	resp := handler.AddAchievement(&req)
	*res = *resp
	return nil
}

func (server *ServiceDevice) DeleteAchievement(req device.DeleteAchievementRequest, res *device.DeleteAchievementResponse) error {
	resp := handler.DeleteAchievement(&req)
	*res = *resp
	return nil
}

func (server *ServiceDevice) GetAchievementById(req device.GetAchievementByIdRequest, res *device.GetAchievementByIdResponse) error {
	resp := handler.GetAchievementById(&req)
	*res = *resp
	return nil
}

func (server *ServiceDevice) QueryAchievement(req device.QueryAchievementRequest, res *device.QueryAchievementResponse) error {
	resp := handler.QueryAchievement(&req)
	*res = *resp
	return nil
}


func (server *ServiceDevice) AddRent(req device.AddRentRequest, res *device.AddRentResponse) error {
	resp := handler.AddRent(&req)
	*res = *resp
	return nil
}

func (server *ServiceDevice) ReturnRent(req device.ReturnRentRequest, res *device.ReturnRentResponse) error {
	resp := handler.ReturnRent(&req)
	*res = *resp
	return nil
}

func (server *ServiceDevice) QueryRent(req device.QueryRentRequest, res *device.QueryRentResponse) error {
	resp := handler.QueryRent(&req)
	*res = *resp
	return nil
}