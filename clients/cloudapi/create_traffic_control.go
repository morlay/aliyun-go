package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateTrafficControlRequest struct {
	requests.RpcRequest
	TrafficControlName string `position:"Query" name:"TrafficControlName"`
	TrafficControlUnit string `position:"Query" name:"TrafficControlUnit"`
	ApiDefault         int    `position:"Query" name:"ApiDefault"`
	UserDefault        int    `position:"Query" name:"UserDefault"`
	AppDefault         int    `position:"Query" name:"AppDefault"`
	Description        string `position:"Query" name:"Description"`
}

func (req *CreateTrafficControlRequest) Invoke(client *sdk.Client) (resp *CreateTrafficControlResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "CreateTrafficControl", "apigateway", "")
	resp = &CreateTrafficControlResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateTrafficControlResponse struct {
	responses.BaseResponse
	RequestId        common.String
	TrafficControlId common.String
}
