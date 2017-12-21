package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RestoreSnapshotRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Domain" name:"InstanceId"`
	SnapshotId           string `position:"Domain" name:"SnapshotId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *RestoreSnapshotRequest) Invoke(client *sdk.Client) (resp *RestoreSnapshotResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "RestoreSnapshot", "redisa", "")
	resp = &RestoreSnapshotResponse{}
	err = client.DoAction(req, resp)
	return
}

type RestoreSnapshotResponse struct {
	responses.BaseResponse
	RequestId string
}
