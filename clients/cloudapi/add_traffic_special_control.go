package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddTrafficSpecialControlRequest struct {
	TrafficControlId string `position:"Query" name:"TrafficControlId"`
	SpecialType      string `position:"Query" name:"SpecialType"`
	SpecialKey       string `position:"Query" name:"SpecialKey"`
	TrafficValue     int    `position:"Query" name:"TrafficValue"`
}

func (r AddTrafficSpecialControlRequest) Invoke(client *sdk.Client) (response *AddTrafficSpecialControlResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddTrafficSpecialControlRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "AddTrafficSpecialControl", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		AddTrafficSpecialControlResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddTrafficSpecialControlResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddTrafficSpecialControlResponse struct {
	RequestId string
}
