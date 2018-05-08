package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RemoveThingTopoRequest struct {
	requests.RpcRequest
	IotId      string `position:"Query" name:"IotId"`
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
}

func (req *RemoveThingTopoRequest) Invoke(client *sdk.Client) (resp *RemoveThingTopoResponse, err error) {
	req.InitWithApiInfo("Iot", "2018-01-20", "RemoveThingTopo", "", "")
	resp = &RemoveThingTopoResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveThingTopoResponse struct {
	responses.BaseResponse
	RequestId    common.String
	Success      bool
	ErrorMessage common.String
	Data         bool
}
