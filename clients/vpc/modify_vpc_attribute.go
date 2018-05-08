package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyVpcAttributeRequest struct {
	requests.RpcRequest
	VpcName              string `position:"Query" name:"VpcName"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyVpcAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyVpcAttributeResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyVpcAttribute", "vpc", "")
	resp = &ModifyVpcAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyVpcAttributeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
