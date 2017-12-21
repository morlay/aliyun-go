
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type GetSnapshotSettingsRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
InstanceId string `position:"Query" name:"InstanceId"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r GetSnapshotSettingsRequest) Invoke(client *sdk.Client) (response *GetSnapshotSettingsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetSnapshotSettingsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "GetSnapshotSettings", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		GetSnapshotSettingsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.GetSnapshotSettingsResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetSnapshotSettingsResponse struct {
RequestId string
InstanceId string
BeginHour int
EndHour int
RetentionDay int
MaxAutoSnapshots int
MaxManualSnapshots int
DayList int
NextTime string
}

