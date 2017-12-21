package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopCreateRequest struct {
	requests.RpcRequest
	ShopCoordinate    string `position:"Query" name:"ShopCoordinate"`
	ShopProvince      string `position:"Query" name:"ShopProvince"`
	ShopTopType       int    `position:"Query" name:"ShopTopType"`
	ShopAddress       string `position:"Query" name:"ShopAddress"`
	ShopType          int    `position:"Query" name:"ShopType"`
	WarnEmail         string `position:"Query" name:"WarnEmail"`
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
	Bid               int64  `position:"Query" name:"Bid"`
	ShopManager       string `position:"Query" name:"ShopManager"`
	ShopBusinessHours string `position:"Query" name:"ShopBusinessHours"`
}

func (req *ShopCreateRequest) Invoke(client *sdk.Client) (resp *ShopCreateResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopCreate", "", "")
	resp = &ShopCreateResponse{}
	err = client.DoAction(req, resp)
	return
}

type ShopCreateResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
