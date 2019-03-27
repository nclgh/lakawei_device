package handler

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/nclgh/lakawei_device/model"
	"github.com/nclgh/lakawei_scaffold/utils"
	"github.com/nclgh/lakawei_scaffold/rpc/device"
	"github.com/nclgh/lakawei_scaffold/rpc/common"
)

func tranDevice(v *model.Device) *device.Device {
	return &device.Device{
		Id:               v.Id,
		Code:             v.Code,
		Name:             v.Name,
		Model:            v.Model,
		Brand:            v.Brand,
		TagCode:          v.TagCode,
		DepartmentId:     v.DepartmentId,
		ManufacturerId:   v.ManufacturerId,
		ManufacturerDate: v.ManufacturerDate,
		Description:      v.Description,
	}
}

func batchTranDevice(vs []*model.Device) map[int64]*device.Device {
	ret := make(map[int64]*device.Device)
	for _, v := range vs {
		ret[v.Id] = tranDevice(v)
	}
	return ret
}

func rTranDevice(v *device.Device) *model.Device {
	return &model.Device{
		Id:               v.Id,
		Code:             v.Code,
		Name:             v.Name,
		Model:            v.Model,
		Brand:            v.Brand,
		TagCode:          v.TagCode,
		DepartmentId:     v.DepartmentId,
		ManufacturerId:   v.ManufacturerId,
		ManufacturerDate: v.ManufacturerDate,
		Description:      v.Description,
	}
}

func AddDevice(req *device.AddDeviceRequest) (rsp *device.AddDeviceResponse) {
	defer utils.RecoverPanic(func(err interface{}, stacks string) {
		logrus.Errorf("AddDeviceRequest panic: %v, stack: %v", err, stacks)
		rsp = getAddDeviceRequestResponse(common.CodeFailed, "panic")
	})
	err := model.InsertDevice(model.GetLakaweiDb(), rTranDevice(&req.Device))
	if err != nil {
		logrus.Errorf("insert device into mysql failed. code: %v, err: %v", req.Device.Code, err)
		return getAddDeviceRequestResponse(common.CodeFailed, fmt.Sprintf("err: %v", err))
	}
	rsp = getAddDeviceRequestResponse(common.CodeSuccess, "")
	return rsp
}

func getAddDeviceRequestResponse(code common.RspCode, msg string) *device.AddDeviceResponse {
	rsp := &device.AddDeviceResponse{
		Code: code,
		Msg:  msg,
	}
	return rsp
}

func DeleteDevice(req *device.DeleteDeviceRequest) (rsp *device.DeleteDeviceResponse) {
	defer utils.RecoverPanic(func(err interface{}, stacks string) {
		logrus.Errorf("DeleteDeviceRequest panic: %v, stack: %v", err, stacks)
		rsp = getDeleteDeviceRequestResponse(common.CodeFailed, "panic")
	})
	err := model.DeleteDevice(model.GetLakaweiDb(), req.Id)
	if err != nil {
		logrus.Errorf("delete device from mysql failed. err: %v", err)
		return getDeleteDeviceRequestResponse(common.CodeFailed, fmt.Sprintf("err: %v", err))
	}
	rsp = getDeleteDeviceRequestResponse(common.CodeSuccess, "")
	return rsp
}

func getDeleteDeviceRequestResponse(code common.RspCode, msg string) *device.DeleteDeviceResponse {
	rsp := &device.DeleteDeviceResponse{
		Code: code,
		Msg:  msg,
	}
	return rsp
}

func GetDeviceById(req *device.GetDeviceByIdRequest) (rsp *device.GetDeviceByIdResponse) {
	defer utils.RecoverPanic(func(err interface{}, stacks string) {
		logrus.Errorf("GetDeviceByIdRequest panic: %v, stack: %v", err, stacks)
		rsp = getGetDeviceByIdRequestResponse(common.CodeFailed, "panic")
	})
	ret, err := model.GetDeviceById(model.GetLakaweiDb(), req.Ids)
	if err != nil {
		logrus.Errorf("select device from mysql failed. err: %v", err)
		return getGetDeviceByIdRequestResponse(common.CodeFailed, fmt.Sprintf("err: %v", err))
	}
	rsp = getGetDeviceByIdRequestResponse(common.CodeSuccess, "")
	rsp.Devices = batchTranDevice(ret)
	return rsp
}

func getGetDeviceByIdRequestResponse(code common.RspCode, msg string) *device.GetDeviceByIdResponse {
	rsp := &device.GetDeviceByIdResponse{
		Code: code,
		Msg:  msg,
	}
	return rsp
}

func QueryDevice(req *device.QueryDeviceRequest) (rsp *device.QueryDeviceResponse) {
	defer utils.RecoverPanic(func(err interface{}, stacks string) {
		logrus.Errorf("QueryDeviceRequest panic: %v, stack: %v", err, stacks)
		rsp = getQueryDeviceResponse(common.CodeFailed, "panic")
	})
	ret, cnt, err := model.QueryDevice(model.GetLakaweiDb(), rTranDevice(req.Device), req.Page, req.PageSize, req.TimeFilter)
	if err != nil {
		logrus.Errorf("filter device from mysql failed. err: %v", err)
		return getQueryDeviceResponse(common.CodeFailed, fmt.Sprintf("err: %v", err))
	}
	rsp = getQueryDeviceResponse(common.CodeSuccess, "")
	rsp.Devices = batchTranDevice(ret)
	rsp.TotalCount = cnt
	return rsp
}

func getQueryDeviceResponse(code common.RspCode, msg string) *device.QueryDeviceResponse {
	rsp := &device.QueryDeviceResponse{
		Code: code,
		Msg:  msg,
	}
	return rsp
}
