package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type TerminateVirtualBorderRouterRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	VbrId                string `position:"Query" name:"VbrId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r TerminateVirtualBorderRouterRequest) Invoke(client *sdk.Client) (response *TerminateVirtualBorderRouterResponse, err error) {
	req := struct {
		*requests.RpcRequest
		TerminateVirtualBorderRouterRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "TerminateVirtualBorderRouter", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		TerminateVirtualBorderRouterResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.TerminateVirtualBorderRouterResponse

	err = client.DoAction(&req, &resp)
	return
}

type TerminateVirtualBorderRouterResponse struct {
	RequestId string
}
