
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type CreateSnapshotRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
InstanceId string `position:"Domain" name:"InstanceId"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
SnapshotName string `position:"Domain" name:"SnapshotName"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r CreateSnapshotRequest) Invoke(client *sdk.Client) (response *CreateSnapshotResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateSnapshotRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "CreateSnapshot", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		CreateSnapshotResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.CreateSnapshotResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateSnapshotResponse struct {
RequestId string
SnapshotId string
Status string
}

