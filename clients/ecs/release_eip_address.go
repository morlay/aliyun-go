package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReleaseEipAddressRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AllocationId         string `position:"Query" name:"AllocationId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ReleaseEipAddressRequest) Invoke(client *sdk.Client) (response *ReleaseEipAddressResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ReleaseEipAddressRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ReleaseEipAddress", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ReleaseEipAddressResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ReleaseEipAddressResponse

	err = client.DoAction(&req, &resp)
	return
}

type ReleaseEipAddressResponse struct {
	RequestId string
}
