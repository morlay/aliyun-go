package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteTrafficSpecialControlRequest struct {
	requests.RpcRequest
	TrafficControlId string `position:"Query" name:"TrafficControlId"`
	SpecialType      string `position:"Query" name:"SpecialType"`
	SpecialKey       string `position:"Query" name:"SpecialKey"`
}

func (req *DeleteTrafficSpecialControlRequest) Invoke(client *sdk.Client) (resp *DeleteTrafficSpecialControlResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteTrafficSpecialControl", "apigateway", "")
	resp = &DeleteTrafficSpecialControlResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteTrafficSpecialControlResponse struct {
	responses.BaseResponse
	RequestId common.String
}
