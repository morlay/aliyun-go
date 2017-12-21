package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyVRouterAttributeRequest struct {
	VRouterName          string `position:"Query" name:"VRouterName"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	VRouterId            string `position:"Query" name:"VRouterId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyVRouterAttributeRequest) Invoke(client *sdk.Client) (response *ModifyVRouterAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyVRouterAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyVRouterAttribute", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		ModifyVRouterAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyVRouterAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyVRouterAttributeResponse struct {
	RequestId string
}
