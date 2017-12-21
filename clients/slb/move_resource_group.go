package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type MoveResourceGroupRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceType         string `position:"Query" name:"ResourceType"`
	Tags                 string `position:"Query" name:"Tags"`
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	NewResourceGroupId   string `position:"Query" name:"NewResourceGroupId"`
}

func (req *MoveResourceGroupRequest) Invoke(client *sdk.Client) (resp *MoveResourceGroupResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "MoveResourceGroup", "slb", "")
	resp = &MoveResourceGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type MoveResourceGroupResponse struct {
	responses.BaseResponse
	RequestId string
}
