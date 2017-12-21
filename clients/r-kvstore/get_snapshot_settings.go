package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetSnapshotSettingsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *GetSnapshotSettingsRequest) Invoke(client *sdk.Client) (resp *GetSnapshotSettingsResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "GetSnapshotSettings", "redisa", "")
	resp = &GetSnapshotSettingsResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetSnapshotSettingsResponse struct {
	responses.BaseResponse
	RequestId          string
	InstanceId         string
	BeginHour          int
	EndHour            int
	RetentionDay       int
	MaxAutoSnapshots   int
	MaxManualSnapshots int
	DayList            int
	NextTime           string
}
