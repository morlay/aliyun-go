package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ConvertNatPublicIpToEipRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
}

func (r ConvertNatPublicIpToEipRequest) Invoke(client *sdk.Client) (response *ConvertNatPublicIpToEipResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ConvertNatPublicIpToEipRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ConvertNatPublicIpToEip", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ConvertNatPublicIpToEipResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ConvertNatPublicIpToEipResponse

	err = client.DoAction(&req, &resp)
	return
}

type ConvertNatPublicIpToEipResponse struct {
	RequestId string
}
