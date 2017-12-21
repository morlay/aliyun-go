package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteTrafficControlRequest struct {
	requests.RpcRequest
	TrafficControlId string `position:"Query" name:"TrafficControlId"`
}

func (req *DeleteTrafficControlRequest) Invoke(client *sdk.Client) (resp *DeleteTrafficControlResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteTrafficControl", "apigateway", "")
	resp = &DeleteTrafficControlResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteTrafficControlResponse struct {
	responses.BaseResponse
	RequestId string
}
