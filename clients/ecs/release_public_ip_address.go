package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReleasePublicIpAddressRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	PublicIpAddress      string `position:"Query" name:"PublicIpAddress"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ReleasePublicIpAddressRequest) Invoke(client *sdk.Client) (response *ReleasePublicIpAddressResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ReleasePublicIpAddressRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ReleasePublicIpAddress", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ReleasePublicIpAddressResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ReleasePublicIpAddressResponse

	err = client.DoAction(&req, &resp)
	return
}

type ReleasePublicIpAddressResponse struct {
	RequestId string
}
