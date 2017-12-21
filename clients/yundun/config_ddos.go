package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ConfigDdosRequest struct {
	InstanceId       string `position:"Query" name:"InstanceId"`
	InstanceType     string `position:"Query" name:"InstanceType"`
	FlowPosition     int    `position:"Query" name:"FlowPosition"`
	StrategyPosition int    `position:"Query" name:"StrategyPosition"`
	Level            int    `position:"Query" name:"Level"`
}

func (r ConfigDdosRequest) Invoke(client *sdk.Client) (response *ConfigDdosResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ConfigDdosRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Yundun", "2015-04-16", "ConfigDdos", "", "")

	resp := struct {
		*responses.BaseResponse
		ConfigDdosResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ConfigDdosResponse

	err = client.DoAction(&req, &resp)
	return
}

type ConfigDdosResponse struct {
	RequestId string
}
