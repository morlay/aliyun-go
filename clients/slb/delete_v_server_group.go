package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteVServerGroupRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	VServerGroupId       string `position:"Query" name:"VServerGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *DeleteVServerGroupRequest) Invoke(client *sdk.Client) (resp *DeleteVServerGroupResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "DeleteVServerGroup", "slb", "")
	resp = &DeleteVServerGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteVServerGroupResponse struct {
	responses.BaseResponse
	RequestId string
}
