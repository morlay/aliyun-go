
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type RestoreSnapshotRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
InstanceId string `position:"Domain" name:"InstanceId"`
SnapshotId string `position:"Domain" name:"SnapshotId"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r RestoreSnapshotRequest) Invoke(client *sdk.Client) (response *RestoreSnapshotResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RestoreSnapshotRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "RestoreSnapshot", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		RestoreSnapshotResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.RestoreSnapshotResponse

	err = client.DoAction(&req, &resp)
	return
}

type RestoreSnapshotResponse struct {
RequestId string
}

