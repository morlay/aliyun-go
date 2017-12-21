package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopUpdateRequest struct {
	ShopCoordinate    string `position:"Query" name:"ShopCoordinate"`
	ShopProvince      string `position:"Query" name:"ShopProvince"`
	ShopTopType       int    `position:"Query" name:"ShopTopType"`
	ShopAddress       string `position:"Query" name:"ShopAddress"`
	ShopType          int    `position:"Query" name:"ShopType"`
	WarnEmail         string `position:"Query" name:"WarnEmail"`
	Sid               int64  `position:"Query" name:"Sid"`
	ShopTel           string `position:"Query" name:"ShopTel"`
	WarnpHone         string `position:"Query" name:"WarnpHone"`
	Warn              int    `position:"Query" name:"Warn"`
	ShopArea          int    `position:"Query" name:"ShopArea"`
	ShopRemarks       string `position:"Query" name:"ShopRemarks"`
	ShopCity          string `position:"Query" name:"ShopCity"`
	ShopSubtype       int    `position:"Query" name:"ShopSubtype"`
	ShopBrand         string `position:"Query" name:"ShopBrand"`
	ShopName          string `position:"Query" name:"ShopName"`
	ShopCloseWarn     int    `position:"Query" name:"ShopCloseWarn"`
	ShopManager       string `position:"Query" name:"ShopManager"`
	ShopBusinessHours string `position:"Query" name:"ShopBusinessHours"`
}

func (r ShopUpdateRequest) Invoke(client *sdk.Client) (response *ShopUpdateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ShopUpdateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopUpdate", "", "")

	resp := struct {
		*responses.BaseResponse
		ShopUpdateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ShopUpdateResponse

	err = client.DoAction(&req, &resp)
	return
}

type ShopUpdateResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
