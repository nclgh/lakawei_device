package handler

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/nclgh/lakawei_device/model"
	"github.com/nclgh/lakawei_scaffold/utils"
	"github.com/nclgh/lakawei_scaffold/rpc/device"
	"github.com/nclgh/lakawei_scaffold/rpc/common"
)

func tranAchievement(v *model.Achievement) *device.Achievement {
	return &device.Achievement{
		Id:                     v.Id,
		DeviceId:               v.DeviceId,
		MemberId:               v.MemberId,
		DepartmentId:           v.DepartmentId,
		AchievementDate:        v.AchievementDate,
		AchievementDescription: v.AchievementDescription,
		AchievementRemark:      v.AchievementRemark,
		PatentDescription:      v.PatentDescription,
		PaperDescription:       v.PaperDescription,
		CompetitionDescription: v.CompetitionDescription,
	}
}

func batchTranAchievement(vs []*model.Achievement) map[int64]*device.Achievement {
	ret := make(map[int64]*device.Achievement)
	for _, v := range vs {
		ret[v.Id] = tranAchievement(v)
	}
	return ret
}

func rTranAchievement(v *device.Achievement) *model.Achievement {
	return &model.Achievement{
		Id:                     v.Id,
		DeviceId:               v.DeviceId,
		MemberId:               v.MemberId,
		DepartmentId:           v.DepartmentId,
		AchievementDate:        v.AchievementDate,
		AchievementDescription: v.AchievementDescription,
		AchievementRemark:      v.AchievementRemark,
		PatentDescription:      v.PatentDescription,
		PaperDescription:       v.PaperDescription,
		CompetitionDescription: v.CompetitionDescription,
	}
}

func AddAchievement(req *device.AddAchievementRequest) (rsp *device.AddAchievementResponse) {
	defer utils.RecoverPanic(func(err interface{}, stacks string) {
		logrus.Errorf("AddAchievementRequest panic: %v, stack: %v", err, stacks)
		rsp = getAddAchievementRequestResponse(common.CodeFailed, "panic")
	})
	err := model.InsertAchievement(model.GetLakaweiDb(), rTranAchievement(&req.Achievement))
	if err != nil {
		logrus.Errorf("insert achievement into mysql failed. deviceId: %v, err: %v", req.Achievement.DeviceId, err)
		return getAddAchievementRequestResponse(common.CodeFailed, fmt.Sprintf("err: %v", err))
	}
	rsp = getAddAchievementRequestResponse(common.CodeSuccess, "")
	return rsp
}

func getAddAchievementRequestResponse(code common.RspCode, msg string) *device.AddAchievementResponse {
	rsp := &device.AddAchievementResponse{
		Code: code,
		Msg:  msg,
	}
	return rsp
}

func DeleteAchievement(req *device.DeleteAchievementRequest) (rsp *device.DeleteAchievementResponse) {
	defer utils.RecoverPanic(func(err interface{}, stacks string) {
		logrus.Errorf("DeleteAchievementRequest panic: %v, stack: %v", err, stacks)
		rsp = getDeleteAchievementRequestResponse(common.CodeFailed, "panic")
	})
	err := model.DeleteAchievement(model.GetLakaweiDb(), req.Id)
	if err != nil {
		logrus.Errorf("delete achievement from mysql failed. err: %v", err)
		return getDeleteAchievementRequestResponse(common.CodeFailed, fmt.Sprintf("err: %v", err))
	}
	rsp = getDeleteAchievementRequestResponse(common.CodeSuccess, "")
	return rsp
}

func getDeleteAchievementRequestResponse(code common.RspCode, msg string) *device.DeleteAchievementResponse {
	rsp := &device.DeleteAchievementResponse{
		Code: code,
		Msg:  msg,
	}
	return rsp
}

func GetAchievementById(req *device.GetAchievementByIdRequest) (rsp *device.GetAchievementByIdResponse) {
	defer utils.RecoverPanic(func(err interface{}, stacks string) {
		logrus.Errorf("GetAchievementByIdRequest panic: %v, stack: %v", err, stacks)
		rsp = getGetAchievementByIdRequestResponse(common.CodeFailed, "panic")
	})
	ret, err := model.GetAchievementById(model.GetLakaweiDb(), req.Ids)
	if err != nil {
		logrus.Errorf("select achievement from mysql failed. err: %v", err)
		return getGetAchievementByIdRequestResponse(common.CodeFailed, fmt.Sprintf("err: %v", err))
	}
	rsp = getGetAchievementByIdRequestResponse(common.CodeSuccess, "")
	rsp.Achievements = batchTranAchievement(ret)
	return rsp
}

func getGetAchievementByIdRequestResponse(code common.RspCode, msg string) *device.GetAchievementByIdResponse {
	rsp := &device.GetAchievementByIdResponse{
		Code: code,
		Msg:  msg,
	}
	return rsp
}

func QueryAchievement(req *device.QueryAchievementRequest) (rsp *device.QueryAchievementResponse) {
	defer utils.RecoverPanic(func(err interface{}, stacks string) {
		logrus.Errorf("QueryAchievementRequest panic: %v, stack: %v", err, stacks)
		rsp = getQueryAchievementResponse(common.CodeFailed, "panic")
	})
	ret, cnt, err := model.QueryAchievement(model.GetLakaweiDb(), rTranAchievement(req.Achievement), req.Page, req.PageSize, req.TimeFilter)
	if err != nil {
		logrus.Errorf("filter achievement from mysql failed. err: %v", err)
		return getQueryAchievementResponse(common.CodeFailed, fmt.Sprintf("err: %v", err))
	}
	rsp = getQueryAchievementResponse(common.CodeSuccess, "")
	rsp.Achievements = batchTranAchievement(ret)
	rsp.TotalCount = cnt
	return rsp
}

func getQueryAchievementResponse(code common.RspCode, msg string) *device.QueryAchievementResponse {
	rsp := &device.QueryAchievementResponse{
		Code: code,
		Msg:  msg,
	}
	return rsp
}
