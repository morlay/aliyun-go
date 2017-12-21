package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopDataAlarmRequest struct {
	WarnPhone string `position:"Query" name:"WarnPhone"`
	Warn      int    `position:"Query" name:"Warn"`
	CloseWarn int    `position:"Query" name:"CloseWarn"`
	WarnEmail string `position:"Query" name:"WarnEmail"`
	Sid       int64  `position:"Query" name:"Sid"`
}

func (r ShopDataAlarmRequest) Invoke(client *sdk.Client) (response *ShopDataAlarmResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ShopDataAlarmRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopDataAlarm", "", "")

	resp := struct {
		*responses.BaseResponse
		ShopDataAlarmResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ShopDataAlarmResponse

	err = client.DoAction(&req, &resp)
	return
}

type ShopDataAlarmResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
