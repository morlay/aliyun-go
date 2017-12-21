
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type DeleteSnapshotSettingsRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
InstanceId string `position:"Domain" name:"InstanceId"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r DeleteSnapshotSettingsRequest) Invoke(client *sdk.Client) (response *DeleteSnapshotSettingsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteSnapshotSettingsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DeleteSnapshotSettings", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		DeleteSnapshotSettingsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.DeleteSnapshotSettingsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteSnapshotSettingsResponse struct {
RequestId string
}

