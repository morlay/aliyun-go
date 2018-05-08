package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QuerySnapshotJobListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	SnapshotJobIds       string `position:"Query" name:"SnapshotJobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QuerySnapshotJobListRequest) Invoke(client *sdk.Client) (resp *QuerySnapshotJobListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QuerySnapshotJobList", "mts", "")
	resp = &QuerySnapshotJobListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QuerySnapshotJobListResponse struct {
	responses.BaseResponse
	RequestId              common.String
	SnapshotJobList        QuerySnapshotJobListSnapshotJobList
	NonExistSnapshotJobIds QuerySnapshotJobListNonExistSnapshotJobIdList
}

type QuerySnapshotJobListSnapshotJob struct {
	Id               common.String
	UserData         common.String
	PipelineId       common.String
	State            common.String
	Code             common.String
	Count            common.String
	TileCount        common.String
	Message          common.String
	CreationTime     common.String
	Input            QuerySnapshotJobListInput
	SnapshotConfig   QuerySnapshotJobListSnapshotConfig
	MNSMessageResult QuerySnapshotJobListMNSMessageResult
}

type QuerySnapshotJobListInput struct {
	Bucket   common.String
	Location common.String
	Object   common.String
	RoleArn  common.String
}

type QuerySnapshotJobListSnapshotConfig struct {
	Time           common.String
	Interval       common.String
	Num            common.String
	Width          common.String
	Height         common.String
	FrameType      common.String
	OutputFile     QuerySnapshotJobListOutputFile
	TileOutputFile QuerySnapshotJobListTileOutputFile
	TileOut        QuerySnapshotJobListTileOut
}

type QuerySnapshotJobListOutputFile struct {
	Bucket   common.String
	Location common.String
	Object   common.String
	RoleArn  common.String
}

type QuerySnapshotJobListTileOutputFile struct {
	Bucket   common.String
	Location common.String
	Object   common.String
	RoleArn  common.String
}

type QuerySnapshotJobListTileOut struct {
	Lines         common.String
	Columns       common.String
	CellWidth     common.String
	CellHeight    common.String
	Margin        common.String
	Padding       common.String
	Color         common.String
	IsKeepCellPic common.String
}

type QuerySnapshotJobListMNSMessageResult struct {
	MessageId    common.String
	ErrorMessage common.String
	ErrorCode    common.String
}

type QuerySnapshotJobListSnapshotJobList []QuerySnapshotJobListSnapshotJob

func (list *QuerySnapshotJobListSnapshotJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QuerySnapshotJobListSnapshotJob)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type QuerySnapshotJobListNonExistSnapshotJobIdList []common.String

func (list *QuerySnapshotJobListNonExistSnapshotJobIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
