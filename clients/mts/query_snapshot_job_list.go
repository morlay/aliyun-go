package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QuerySnapshotJobListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	SnapshotJobIds       string `position:"Query" name:"SnapshotJobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r QuerySnapshotJobListRequest) Invoke(client *sdk.Client) (response *QuerySnapshotJobListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QuerySnapshotJobListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "QuerySnapshotJobList", "", "")

	resp := struct {
		*responses.BaseResponse
		QuerySnapshotJobListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QuerySnapshotJobListResponse

	err = client.DoAction(&req, &resp)
	return
}

type QuerySnapshotJobListResponse struct {
	RequestId              string
	SnapshotJobList        QuerySnapshotJobListSnapshotJobList
	NonExistSnapshotJobIds QuerySnapshotJobListNonExistSnapshotJobIdList
}

type QuerySnapshotJobListSnapshotJob struct {
	Id               string
	UserData         string
	PipelineId       string
	State            string
	Code             string
	Count            string
	TileCount        string
	Message          string
	CreationTime     string
	Input            QuerySnapshotJobListInput
	SnapshotConfig   QuerySnapshotJobListSnapshotConfig
	MNSMessageResult QuerySnapshotJobListMNSMessageResult
}

type QuerySnapshotJobListInput struct {
	Bucket   string
	Location string
	Object   string
	RoleArn  string
}

type QuerySnapshotJobListSnapshotConfig struct {
	Time           string
	Interval       string
	Num            string
	Width          string
	Height         string
	FrameType      string
	OutputFile     QuerySnapshotJobListOutputFile
	TileOutputFile QuerySnapshotJobListTileOutputFile
	TileOut        QuerySnapshotJobListTileOut
}

type QuerySnapshotJobListOutputFile struct {
	Bucket   string
	Location string
	Object   string
	RoleArn  string
}

type QuerySnapshotJobListTileOutputFile struct {
	Bucket   string
	Location string
	Object   string
	RoleArn  string
}

type QuerySnapshotJobListTileOut struct {
	Lines         string
	Columns       string
	CellWidth     string
	CellHeight    string
	Margin        string
	Padding       string
	Color         string
	IsKeepCellPic string
}

type QuerySnapshotJobListMNSMessageResult struct {
	MessageId    string
	ErrorMessage string
	ErrorCode    string
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

type QuerySnapshotJobListNonExistSnapshotJobIdList []string

func (list *QuerySnapshotJobListNonExistSnapshotJobIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
