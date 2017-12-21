package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteAllTrafficSpecialControlRequest struct {
	requests.RpcRequest
	TrafficControlId string `position:"Query" name:"TrafficControlId"`
}

func (req *DeleteAllTrafficSpecialControlRequest) Invoke(client *sdk.Client) (resp *DeleteAllTrafficSpecialControlResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteAllTrafficSpecialControl", "apigateway", "")
	resp = &DeleteAllTrafficSpecialControlResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteAllTrafficSpecialControlResponse struct {
	responses.BaseResponse
	RequestId string
}
