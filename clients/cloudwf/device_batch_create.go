package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeviceBatchCreateRequest struct {
	Sn         string `position:"Query" name:"Sn"`
	DeviceType int    `position:"Query" name:"DeviceType"`
}

func (r DeviceBatchCreateRequest) Invoke(client *sdk.Client) (response *DeviceBatchCreateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeviceBatchCreateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "DeviceBatchCreate", "", "")

	resp := struct {
		*responses.BaseResponse
		DeviceBatchCreateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeviceBatchCreateResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeviceBatchCreateResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
