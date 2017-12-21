package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteTrafficSpecialControlRequest struct {
	TrafficControlId string `position:"Query" name:"TrafficControlId"`
	SpecialType      string `position:"Query" name:"SpecialType"`
	SpecialKey       string `position:"Query" name:"SpecialKey"`
}

func (r DeleteTrafficSpecialControlRequest) Invoke(client *sdk.Client) (response *DeleteTrafficSpecialControlResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteTrafficSpecialControlRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteTrafficSpecialControl", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DeleteTrafficSpecialControlResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteTrafficSpecialControlResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteTrafficSpecialControlResponse struct {
	RequestId string
}
