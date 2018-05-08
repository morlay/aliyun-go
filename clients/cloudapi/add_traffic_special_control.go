package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddTrafficSpecialControlRequest struct {
	requests.RpcRequest
	TrafficControlId string `position:"Query" name:"TrafficControlId"`
	SpecialType      string `position:"Query" name:"SpecialType"`
	SpecialKey       string `position:"Query" name:"SpecialKey"`
	TrafficValue     int    `position:"Query" name:"TrafficValue"`
}

func (req *AddTrafficSpecialControlRequest) Invoke(client *sdk.Client) (resp *AddTrafficSpecialControlResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "AddTrafficSpecialControl", "apigateway", "")
	resp = &AddTrafficSpecialControlResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddTrafficSpecialControlResponse struct {
	responses.BaseResponse
	RequestId common.String
}
