package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UnassociatePhysicalConnectionFromVirtualBorderRouterRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	PhysicalConnectionId string `position:"Query" name:"PhysicalConnectionId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	VbrId                string `position:"Query" name:"VbrId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r UnassociatePhysicalConnectionFromVirtualBorderRouterRequest) Invoke(client *sdk.Client) (response *UnassociatePhysicalConnectionFromVirtualBorderRouterResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UnassociatePhysicalConnectionFromVirtualBorderRouterRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "UnassociatePhysicalConnectionFromVirtualBorderRouter", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		UnassociatePhysicalConnectionFromVirtualBorderRouterResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UnassociatePhysicalConnectionFromVirtualBorderRouterResponse

	err = client.DoAction(&req, &resp)
	return
}

type UnassociatePhysicalConnectionFromVirtualBorderRouterResponse struct {
	RequestId string
}
