package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeviceShowListRequest struct {
	Dirc       string `position:"Query" name:"Dirc"`
	Page       int    `position:"Query" name:"Page"`
	Per        int    `position:"Query" name:"Per"`
	DeviceType int    `position:"Query" name:"DeviceType"`
	Sid        int64  `position:"Query" name:"Sid"`
}

func (r DeviceShowListRequest) Invoke(client *sdk.Client) (response *DeviceShowListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeviceShowListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "DeviceShowList", "", "")

	resp := struct {
		*responses.BaseResponse
		DeviceShowListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeviceShowListResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeviceShowListResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
