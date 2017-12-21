
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type DeleteSnapshotRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
InstanceId string `position:"Domain" name:"InstanceId"`
SnapshotId string `position:"Domain" name:"SnapshotId"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r DeleteSnapshotRequest) Invoke(client *sdk.Client) (response *DeleteSnapshotResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteSnapshotRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DeleteSnapshot", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		DeleteSnapshotResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.DeleteSnapshotResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteSnapshotResponse struct {
RequestId string
}

