package device

import (
	"time"
	"github.com/nclgh/lakawei_scaffold/rpc/common"
)

// Manufacturer
type Manufacturer struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type AddManufacturerRequest struct {
	Name string
}

type AddManufacturerResponse struct {
	Code common.RspCode
	Msg  string
}

type GetManufacturerByIdRequest struct {
	Ids []int64
}

type GetManufacturerByIdResponse struct {
	Manufacturers map[int64]*Manufacturer
	Code          common.RspCode
	Msg           string
}

type QueryManufacturerRequest struct {
	Manufacturer *Manufacturer
	Page         int64
	PageSize     int64
}

type QueryManufacturerResponse struct {
	Manufacturers map[int64]*Manufacturer
	TotalCount    int64
	Code          common.RspCode
	Msg           string
}

// Device
type Device struct {
	Code             string    `json:"code"`
	Name             string    `json:"name"`
	Model            string    `json:"model"`
	Brand            string    `json:"brand"`
	TagCode          string    `json:"tag_code"`
	DepartmentCode   string    `json:"department_code"`
	ManufacturerId   int64     `json:"manufacturer_id"`
	ManufacturerDate time.Time `json:"manufacturer_date"`
	RentStatus       int64     `json:"rent_status"`
	Description      string    `json:"description"`
}

type AddDeviceRequest struct {
	Device Device
}

type AddDeviceResponse struct {
	Code common.RspCode
	Msg  string
}

type DeleteDeviceRequest struct {
	Code string
}

type DeleteDeviceResponse struct {
	Code common.RspCode
	Msg  string
}

type GetDeviceByCodeRequest struct {
	Codes []string
}

type GetDeviceByCodeResponse struct {
	Devices map[string]*Device
	Code    common.RspCode
	Msg     string
}

type QueryDeviceRequest struct {
	Device   *Device
	Filter   *common.Filter
	Page     int64
	PageSize int64
}

type QueryDeviceResponse struct {
	Devices    map[string]*Device
	TotalCount int64
	Code       common.RspCode
	Msg        string
}

// Achievement
type Achievement struct {
	Id                     int64     `json:"id"`
	DeviceCode             string    `json:"device_code"`
	MemberCode             string    `json:"member_code"`
	DepartmentCode         string    `json:"department_code"`
	AchievementDate        time.Time `json:"achievement_date"`
	AchievementDescription string    `json:"achievement_description"`
	AchievementRemark      string    `json:"achievement_remark"`
	PatentDescription      string    `json:"patent_description"`
	PaperDescription       string    `json:"paper_description"`
	CompetitionDescription string    `json:"competition_description"`
}

type AddAchievementRequest struct {
	Achievement Achievement
}

type AddAchievementResponse struct {
	Code common.RspCode
	Msg  string
}

type DeleteAchievementRequest struct {
	Id int64
}

type DeleteAchievementResponse struct {
	Code common.RspCode
	Msg  string
}

type GetAchievementByIdRequest struct {
	Ids []int64
}

type GetAchievementByIdResponse struct {
	Achievements map[int64]*Achievement
	Code         common.RspCode
	Msg          string
}

type QueryAchievementRequest struct {
	Achievement *Achievement
	Filter      *common.Filter
	Page        int64
	PageSize    int64
}

type QueryAchievementResponse struct {
	Achievements map[int64]*Achievement
	TotalCount   int64
	Code         common.RspCode
	Msg          string
}

// Rent
type Rent struct {
	Id                 int64     `json:"id"`
	DeviceCode         string    `json:"device_code"`
	RentStatus         int64     `json:"rent_status"`
	BorrowerMemberCode string    `json:"borrower_member_code"`
	BorrowDate         time.Time `json:"borrow_date"`
	ExpectReturnDate   time.Time `json:"expect_return_date"`
	ReturnerMemberCode string    `json:"returner_member_code"`
	RealReturnDate     time.Time `json:"real_return_date"`
	BorrowRemark       string    `json:"borrow_remark"`
	ReturnRemark       string    `json:"return_remark"`
}

type AddRentRequest struct {
	Rent Rent
}

type AddRentResponse struct {
	Code common.RspCode
	Msg  string
}

type ReturnRentRequest struct {
	DeviceCode         string
	ReturnerMemberCode string
	ReturnRemark       string
}

type ReturnRentResponse struct {
	Code common.RspCode
	Msg  string
}

type QueryRentRequest struct {
	Rent     *Rent
	Filter   *common.Filter
	Page     int64
	PageSize int64
}

type QueryRentResponse struct {
	Rents      map[int64]*Rent
	TotalCount int64
	Code       common.RspCode
	Msg        string
}
