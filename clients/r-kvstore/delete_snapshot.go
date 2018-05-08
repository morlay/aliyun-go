package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteSnapshotRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Domain" name:"InstanceId"`
	SnapshotId           string `position:"Domain" name:"SnapshotId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteSnapshotRequest) Invoke(client *sdk.Client) (resp *DeleteSnapshotResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DeleteSnapshot", "redisa", "")
	resp = &DeleteSnapshotResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteSnapshotResponse struct {
	responses.BaseResponse
	RequestId common.String
}
