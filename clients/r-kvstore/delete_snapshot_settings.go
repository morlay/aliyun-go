package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteSnapshotSettingsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Domain" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteSnapshotSettingsRequest) Invoke(client *sdk.Client) (resp *DeleteSnapshotSettingsResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DeleteSnapshotSettings", "redisa", "")
	resp = &DeleteSnapshotSettingsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteSnapshotSettingsResponse struct {
	responses.BaseResponse
	RequestId string
}
