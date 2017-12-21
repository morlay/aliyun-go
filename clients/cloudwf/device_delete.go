package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeviceDeleteRequest struct {
	requests.RpcRequest
	Did int64  `position:"Query" name:"Did"`
	Mac string `position:"Query" name:"Mac"`
}

func (req *DeviceDeleteRequest) Invoke(client *sdk.Client) (resp *DeviceDeleteResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "DeviceDelete", "", "")
	resp = &DeviceDeleteResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeviceDeleteResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
