package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ConfigDdosRequest struct {
	requests.RpcRequest
	InstanceId       string `position:"Query" name:"InstanceId"`
	InstanceType     string `position:"Query" name:"InstanceType"`
	FlowPosition     int    `position:"Query" name:"FlowPosition"`
	StrategyPosition int    `position:"Query" name:"StrategyPosition"`
	Level            int    `position:"Query" name:"Level"`
}

func (req *ConfigDdosRequest) Invoke(client *sdk.Client) (resp *ConfigDdosResponse, err error) {
	req.InitWithApiInfo("Yundun", "2015-04-16", "ConfigDdos", "", "")
	resp = &ConfigDdosResponse{}
	err = client.DoAction(req, resp)
	return
}

type ConfigDdosResponse struct {
	responses.BaseResponse
	RequestId common.String
}
