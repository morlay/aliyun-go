package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateVpcRequest struct {
	VpcName              string `position:"Query" name:"VpcName"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	CidrBlock            string `position:"Query" name:"CidrBlock"`
	Description          string `position:"Query" name:"Description"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r CreateVpcRequest) Invoke(client *sdk.Client) (response *CreateVpcResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateVpcRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "CreateVpc", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		CreateVpcResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateVpcResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateVpcResponse struct {
	RequestId    string
	VpcId        string
	VRouterId    string
	RouteTableId string
}
