package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type NotifyAddThingTopoRequest struct {
	requests.RpcRequest
	GwProductKey  string `position:"Query" name:"GwProductKey"`
	GwDeviceName  string `position:"Query" name:"GwDeviceName"`
	GwIotId       string `position:"Query" name:"GwIotId"`
	DeviceListStr string `position:"Query" name:"DeviceListStr"`
}

func (req *NotifyAddThingTopoRequest) Invoke(client *sdk.Client) (resp *NotifyAddThingTopoResponse, err error) {
	req.InitWithApiInfo("Iot", "2018-01-20", "NotifyAddThingTopo", "", "")
	resp = &NotifyAddThingTopoResponse{}
	err = client.DoAction(req, resp)
	return
}

type NotifyAddThingTopoResponse struct {
	responses.BaseResponse
	RequestId    common.String
	Success      bool
	ErrorMessage common.String
}
