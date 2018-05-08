package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateSnapshotRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Domain" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SnapshotName         string `position:"Domain" name:"SnapshotName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CreateSnapshotRequest) Invoke(client *sdk.Client) (resp *CreateSnapshotResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "CreateSnapshot", "redisa", "")
	resp = &CreateSnapshotResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateSnapshotResponse struct {
	responses.BaseResponse
	RequestId  common.String
	SnapshotId common.String
	Status     common.String
}
