package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type TerminatePhysicalConnectionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	PhysicalConnectionId string `position:"Query" name:"PhysicalConnectionId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r TerminatePhysicalConnectionRequest) Invoke(client *sdk.Client) (response *TerminatePhysicalConnectionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		TerminatePhysicalConnectionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "TerminatePhysicalConnection", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		TerminatePhysicalConnectionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.TerminatePhysicalConnectionResponse

	err = client.DoAction(&req, &resp)
	return
}

type TerminatePhysicalConnectionResponse struct {
	RequestId string
}
