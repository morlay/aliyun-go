package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteMasterSlaveServerGroupRequest struct {
	requests.RpcRequest
	Access_key_id            string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId          int64  `position:"Query" name:"ResourceOwnerId"`
	MasterSlaveServerGroupId string `position:"Query" name:"MasterSlaveServerGroupId"`
	ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount             string `position:"Query" name:"OwnerAccount"`
	OwnerId                  int64  `position:"Query" name:"OwnerId"`
	Tags                     string `position:"Query" name:"Tags"`
}

func (req *DeleteMasterSlaveServerGroupRequest) Invoke(client *sdk.Client) (resp *DeleteMasterSlaveServerGroupResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "DeleteMasterSlaveServerGroup", "slb", "")
	resp = &DeleteMasterSlaveServerGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteMasterSlaveServerGroupResponse struct {
	responses.BaseResponse
	RequestId common.String
}
