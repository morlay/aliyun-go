package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeviceDeleteRequest struct {
	Did int64  `position:"Query" name:"Did"`
	Mac string `position:"Query" name:"Mac"`
}

func (r DeviceDeleteRequest) Invoke(client *sdk.Client) (response *DeviceDeleteResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeviceDeleteRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "DeviceDelete", "", "")

	resp := struct {
		*responses.BaseResponse
		DeviceDeleteResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeviceDeleteResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeviceDeleteResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
