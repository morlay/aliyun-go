package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type MoveResourceGroupRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	NewResourceGroupId   string `position:"Query" name:"NewResourceGroupId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceType         string `position:"Query" name:"ResourceType"`
}

func (req *MoveResourceGroupRequest) Invoke(client *sdk.Client) (resp *MoveResourceGroupResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "MoveResourceGroup", "vpc", "")
	resp = &MoveResourceGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type MoveResourceGroupResponse struct {
	responses.BaseResponse
	RequestId string
}
