package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifySnapshotAttributeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SnapshotId           string `position:"Query" name:"SnapshotId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	SnapshotName         string `position:"Query" name:"SnapshotName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifySnapshotAttributeRequest) Invoke(client *sdk.Client) (resp *ModifySnapshotAttributeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifySnapshotAttribute", "ecs", "")
	resp = &ModifySnapshotAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifySnapshotAttributeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
