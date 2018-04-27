package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateVpcRequest struct {
	requests.RpcRequest
	VpcName              string `position:"Query" name:"VpcName"`
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
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
	req.InitWithApiInfo("Vpc", "2016-04-28", "CreateVpc", "vpc", "")
	resp = &CreateVpcResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateVpcResponse struct {
	responses.BaseResponse
	RequestId       string
	VpcId           string
	VRouterId       string
	RouteTableId    string
	ResourceGroupId string
}
