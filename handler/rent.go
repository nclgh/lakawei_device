package handler

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/nclgh/lakawei_device/model"
	"github.com/nclgh/lakawei_scaffold/utils"
	"github.com/nclgh/lakawei_scaffold/rpc/device"
	"github.com/nclgh/lakawei_scaffold/rpc/common"
)

func tranRent(v *model.Rent) *device.Rent {
	return &device.Rent{
		Id:                 v.Id,
		DeviceCode:         v.DeviceCode,
		RentStatus:         v.RentStatus,
		BorrowerMemberCode: v.ReturnerMemberCode,
		BorrowDate:         v.BorrowDate,
		BorrowRemark:       v.BorrowRemark,
		ExpectReturnDate:   v.ExpectReturnDate,
		ReturnerMemberCode: v.ReturnerMemberCode,
		RealReturnDate:     v.RealReturnDate,
		ReturnRemark:       v.ReturnRemark,
	}
}

func batchTranRent(vs []*model.Rent) map[int64]*device.Rent {
	ret := make(map[int64]*device.Rent)
	for _, v := range vs {
		ret[v.Id] = tranRent(v)
	}
	return ret
}

func rTranRent(v *device.Rent) *model.Rent {
	return &model.Rent{
		Id:                 v.Id,
		DeviceCode:         v.DeviceCode,
		RentStatus:         v.RentStatus,
		BorrowerMemberCode: v.BorrowerMemberCode,
		BorrowDate:         v.BorrowDate,
		BorrowRemark:       v.BorrowRemark,
		ExpectReturnDate:   v.ExpectReturnDate,
		ReturnerMemberCode: v.ReturnerMemberCode,
		RealReturnDate:     v.RealReturnDate,
		ReturnRemark:       v.ReturnRemark,
	}
}

func AddRent(req *device.AddRentRequest) (rsp *device.AddRentResponse) {
	defer utils.RecoverPanic(func(err interface{}, stacks string) {
		logrus.Errorf("AddRentRequest panic: %v, stack: %v", err, stacks)
		rsp = getAddRentRequestResponse(common.CodeFailed, "panic")
	})
	db := model.GetLakaweiDb().Begin()
	defer db.Rollback()
	err := model.InsertRent(db, rTranRent(&req.Rent))
	if err != nil {
		logrus.Errorf("insert rent into mysql failed. device_id: %v, err: %v", req.Rent.DeviceCode, err)
		return getAddRentRequestResponse(common.CodeFailed, fmt.Sprintf("err: %v", err))
	}
	db.Commit()
	rsp = getAddRentRequestResponse(common.CodeSuccess, "")
	return rsp
}

func getAddRentRequestResponse(code common.RspCode, msg string) *device.AddRentResponse {
	rsp := &device.AddRentResponse{
		Code: code,
		Msg:  msg,
	}
	return rsp
}

func ReturnRent(req *device.ReturnRentRequest) (rsp *device.ReturnRentResponse) {
	defer utils.RecoverPanic(func(err interface{}, stacks string) {
		logrus.Errorf("ReturnRentRequest panic: %v, stack: %v", err, stacks)
		rsp = getReturnRentRequestResponse(common.CodeFailed, "panic")
	})
	db := model.GetLakaweiDb().Begin()
	defer db.Rollback()
	err := model.ReturnRent(db, &model.Rent{
		DeviceCode:         req.DeviceCode,
		ReturnerMemberCode: req.ReturnerMemberCode,
		ReturnRemark:     req.ReturnRemark,
	})
	if err != nil {
		logrus.Errorf("return rent from mysql failed. err: %v", err)
		return getReturnRentRequestResponse(common.CodeFailed, fmt.Sprintf("err: %v", err))
	}
	db.Commit()
	rsp = getReturnRentRequestResponse(common.CodeSuccess, "")
	return rsp
}

func getReturnRentRequestResponse(code common.RspCode, msg string) *device.ReturnRentResponse {
	rsp := &device.ReturnRentResponse{
		Code: code,
		Msg:  msg,
	}
	return rsp
}

func QueryRent(req *device.QueryRentRequest) (rsp *device.QueryRentResponse) {
	defer utils.RecoverPanic(func(err interface{}, stacks string) {
		logrus.Errorf("QueryRentRequest panic: %v, stack: %v", err, stacks)
		rsp = getQueryRentResponse(common.CodeFailed, "panic")
	})
	ret, cnt, err := model.QueryRent(model.GetLakaweiDb(), rTranRent(req.Rent), req.Page, req.PageSize, req.Filter)
	if err != nil {
		logrus.Errorf("filter rent from mysql failed. err: %v", err)
		return getQueryRentResponse(common.CodeFailed, fmt.Sprintf("err: %v", err))
	}
	rsp = getQueryRentResponse(common.CodeSuccess, "")
	rsp.Rents = batchTranRent(ret)
	rsp.TotalCount = cnt
	return rsp
}

func getQueryRentResponse(code common.RspCode, msg string) *device.QueryRentResponse {
	rsp := &device.QueryRentResponse{
		Code: code,
		Msg:  msg,
	}
	return rsp
}
