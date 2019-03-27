package handler

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/nclgh/lakawei_device/model"
	"github.com/nclgh/lakawei_scaffold/utils"
	"github.com/nclgh/lakawei_scaffold/rpc/device"
	"github.com/nclgh/lakawei_scaffold/rpc/common"
)

func tranManufacturer(v *model.Manufacturer) *device.Manufacturer {
	return &device.Manufacturer{
		Id:   v.Id,
		Name: v.Name,
	}
}

func batchTranManufacturer(vs []*model.Manufacturer) map[int64]*device.Manufacturer {
	ret := make(map[int64]*device.Manufacturer)
	for _, v := range vs {
		ret[v.Id] = tranManufacturer(v)
	}
	return ret
}

func rTranManufacturer(v *device.Manufacturer) *model.Manufacturer {
	return &model.Manufacturer{
		Id:   v.Id,
		Name: v.Name,
	}
}

func AddManufacturer(req *device.AddManufacturerRequest) (rsp *device.AddManufacturerResponse) {
	defer utils.RecoverPanic(func(err interface{}, stacks string) {
		logrus.Errorf("AddManufacturerRequest panic: %v, stack: %v", err, stacks)
		rsp = getAddManufacturerRequestResponse(common.CodeFailed, "panic")
	})
	err := model.InsertManufacturer(model.GetLakaweiDb(), req.Name)
	if err != nil {
		logrus.Errorf("insert manufacturer into mysql failed. name: %v, err: %v", req.Name, err)
		return getAddManufacturerRequestResponse(common.CodeFailed, fmt.Sprintf("err: %v", err))
	}
	rsp = getAddManufacturerRequestResponse(common.CodeSuccess, "")
	return rsp
}

func getAddManufacturerRequestResponse(code common.RspCode, msg string) *device.AddManufacturerResponse {
	rsp := &device.AddManufacturerResponse{
		Code: code,
		Msg:  msg,
	}
	return rsp
}

func GetManufacturerById(req *device.GetManufacturerByIdRequest) (rsp *device.GetManufacturerByIdResponse) {
	defer utils.RecoverPanic(func(err interface{}, stacks string) {
		logrus.Errorf("GetManufacturerByIdRequest panic: %v, stack: %v", err, stacks)
		rsp = getGetManufacturerByIdRequestResponse(common.CodeFailed, "panic")
	})
	ret, err := model.GetManufacturerById(model.GetLakaweiDb(), req.Ids)
	if err != nil {
		logrus.Errorf("select manufacturer from mysql failed. err: %v", err)
		return getGetManufacturerByIdRequestResponse(common.CodeFailed, fmt.Sprintf("err: %v", err))
	}
	rsp = getGetManufacturerByIdRequestResponse(common.CodeSuccess, "")
	rsp.Manufacturers = batchTranManufacturer(ret)
	return rsp
}

func getGetManufacturerByIdRequestResponse(code common.RspCode, msg string) *device.GetManufacturerByIdResponse {
	rsp := &device.GetManufacturerByIdResponse{
		Code: code,
		Msg:  msg,
	}
	return rsp
}

func QueryManufacturer(req *device.QueryManufacturerRequest) (rsp *device.QueryManufacturerResponse) {
	defer utils.RecoverPanic(func(err interface{}, stacks string) {
		logrus.Errorf("QueryManufacturerRequest panic: %v, stack: %v", err, stacks)
		rsp = getQueryManufacturerResponse(common.CodeFailed, "panic")
	})
	ret, cnt, err := model.QueryManufacturer(model.GetLakaweiDb(), rTranManufacturer(req.Manufacturer), req.Page, req.PageSize)
	if err != nil {
		logrus.Errorf("filter manufacturer from mysql failed. err: %v", err)
		return getQueryManufacturerResponse(common.CodeFailed, fmt.Sprintf("err: %v", err))
	}
	rsp = getQueryManufacturerResponse(common.CodeSuccess, "")
	rsp.Manufacturers = batchTranManufacturer(ret)
	rsp.TotalCount = cnt
	return rsp
}

func getQueryManufacturerResponse(code common.RspCode, msg string) *device.QueryManufacturerResponse {
	rsp := &device.QueryManufacturerResponse{
		Code: code,
		Msg:  msg,
	}
	return rsp
}
