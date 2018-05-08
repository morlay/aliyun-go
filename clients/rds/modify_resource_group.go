package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyResourceGroupRequest struct {
	requests.RpcRequest
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyResourceGroupRequest) Invoke(client *sdk.Client) (resp *ModifyResourceGroupResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyResourceGroup", "rds", "")
	resp = &ModifyResourceGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyResourceGroupResponse struct {
	responses.BaseResponse
	RequestId common.String
}
