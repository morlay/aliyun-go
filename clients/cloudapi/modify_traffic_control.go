package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyTrafficControlRequest struct {
	requests.RpcRequest
	TrafficControlId   string `position:"Query" name:"TrafficControlId"`
	TrafficControlName string `position:"Query" name:"TrafficControlName"`
	TrafficControlUnit string `position:"Query" name:"TrafficControlUnit"`
	ApiDefault         int    `position:"Query" name:"ApiDefault"`
	UserDefault        int    `position:"Query" name:"UserDefault"`
	AppDefault         int    `position:"Query" name:"AppDefault"`
	Description        string `position:"Query" name:"Description"`
}

func (req *ModifyTrafficControlRequest) Invoke(client *sdk.Client) (resp *ModifyTrafficControlResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "ModifyTrafficControl", "apigateway", "")
	resp = &ModifyTrafficControlResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyTrafficControlResponse struct {
	responses.BaseResponse
	RequestId common.String
}
