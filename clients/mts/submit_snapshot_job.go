package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SubmitSnapshotJobRequest struct {
	requests.RpcRequest
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SnapshotConfig       string `position:"Query" name:"SnapshotConfig"`
	PipelineId           string `position:"Query" name:"PipelineId"`
}

func (req *SubmitSnapshotJobRequest) Invoke(client *sdk.Client) (resp *SubmitSnapshotJobResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "SubmitSnapshotJob", "mts", "")
	resp = &SubmitSnapshotJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitSnapshotJobResponse struct {
	responses.BaseResponse
	RequestId   common.String
	SnapshotJob SubmitSnapshotJobSnapshotJob
}

type SubmitSnapshotJobSnapshotJob struct {
	Id               common.String
	UserData         common.String
	PipelineId       common.String
	State            common.String
	Code             common.String
	Count            common.String
	TileCount        common.String
	Message          common.String
	CreationTime     common.String
	Input            SubmitSnapshotJobInput
	SnapshotConfig   SubmitSnapshotJobSnapshotConfig
	MNSMessageResult SubmitSnapshotJobMNSMessageResult
}

type SubmitSnapshotJobInput struct {
	Bucket   common.String
	Location common.String
	Object   common.String
	RoleArn  common.String
}

type SubmitSnapshotJobSnapshotConfig struct {
	Time           common.String
	Interval       common.String
	Num            common.String
	Width          common.String
	Height         common.String
	FrameType      common.String
	OutputFile     SubmitSnapshotJobOutputFile
	TileOutputFile SubmitSnapshotJobTileOutputFile
	TileOut        SubmitSnapshotJobTileOut
}

type SubmitSnapshotJobOutputFile struct {
	Bucket   common.String
	Location common.String
	Object   common.String
	RoleArn  common.String
}

type SubmitSnapshotJobTileOutputFile struct {
	Bucket   common.String
	Location common.String
	Object   common.String
	RoleArn  common.String
}

type SubmitSnapshotJobTileOut struct {
	Lines         common.String
	Columns       common.String
	CellWidth     common.String
	CellHeight    common.String
	Margin        common.String
	Padding       common.String
	Color         common.String
	IsKeepCellPic common.String
}

type SubmitSnapshotJobMNSMessageResult struct {
	MessageId    common.String
	ErrorMessage common.String
	ErrorCode    common.String
}
