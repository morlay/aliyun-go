package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetSnapshotSettingsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	EndHour              int    `position:"Domain" name:"EndHour"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DayList              int    `position:"Domain" name:"DayList"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InstanceId           string `position:"Domain" name:"InstanceId"`
	RetentionDay         int    `position:"Domain" name:"RetentionDay"`
	MaxManualSnapshots   int    `position:"Domain" name:"MaxManualSnapshots"`
	MaxAutoSnapshots     int    `position:"Domain" name:"MaxAutoSnapshots"`
	BeginHour            int    `position:"Domain" name:"BeginHour"`
}

func (req *SetSnapshotSettingsRequest) Invoke(client *sdk.Client) (resp *SetSnapshotSettingsResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "SetSnapshotSettings", "redisa", "")
	resp = &SetSnapshotSettingsResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetSnapshotSettingsResponse struct {
	responses.BaseResponse
	RequestId string
}
