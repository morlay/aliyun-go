package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateVpcRequest struct {
	requests.RpcRequest
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

func (req *CreateVpcRequest) Invoke(client *sdk.Client) (resp *CreateVpcResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "CreateVpc", "ecs", "")
	resp = &CreateVpcResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateVpcResponse struct {
	responses.BaseResponse
	RequestId    common.String
	VpcId        common.String
	VRouterId    common.String
	RouteTableId common.String
}
