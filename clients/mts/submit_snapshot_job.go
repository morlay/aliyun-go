package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubmitSnapshotJobRequest struct {
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SnapshotConfig       string `position:"Query" name:"SnapshotConfig"`
	PipelineId           string `position:"Query" name:"PipelineId"`
}

func (r SubmitSnapshotJobRequest) Invoke(client *sdk.Client) (response *SubmitSnapshotJobResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SubmitSnapshotJobRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "SubmitSnapshotJob", "", "")

	resp := struct {
		*responses.BaseResponse
		SubmitSnapshotJobResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SubmitSnapshotJobResponse

	err = client.DoAction(&req, &resp)
	return
}

type SubmitSnapshotJobResponse struct {
	RequestId   string
	SnapshotJob SubmitSnapshotJobSnapshotJob
}

type SubmitSnapshotJobSnapshotJob struct {
	Id               string
	UserData         string
	PipelineId       string
	State            string
	Code             string
	Count            string
	TileCount        string
	Message          string
	CreationTime     string
	Input            SubmitSnapshotJobInput
	SnapshotConfig   SubmitSnapshotJobSnapshotConfig
	MNSMessageResult SubmitSnapshotJobMNSMessageResult
}

type SubmitSnapshotJobInput struct {
	Bucket   string
	Location string
	Object   string
	RoleArn  string
}

type SubmitSnapshotJobSnapshotConfig struct {
	Time           string
	Interval       string
	Num            string
	Width          string
	Height         string
	FrameType      string
	OutputFile     SubmitSnapshotJobOutputFile
	TileOutputFile SubmitSnapshotJobTileOutputFile
	TileOut        SubmitSnapshotJobTileOut
}

type SubmitSnapshotJobOutputFile struct {
	Bucket   string
	Location string
	Object   string
	RoleArn  string
}

type SubmitSnapshotJobTileOutputFile struct {
	Bucket   string
	Location string
	Object   string
	RoleArn  string
}

type SubmitSnapshotJobTileOut struct {
	Lines         string
	Columns       string
	CellWidth     string
	CellHeight    string
	Margin        string
	Padding       string
	Color         string
	IsKeepCellPic string
}

type SubmitSnapshotJobMNSMessageResult struct {
	MessageId    string
	ErrorMessage string
	ErrorCode    string
}
