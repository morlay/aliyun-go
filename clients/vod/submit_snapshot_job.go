package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SubmitSnapshotJobRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SpecifiedOffsetTime  int64  `position:"Query" name:"SpecifiedOffsetTime"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Width                string `position:"Query" name:"Width"`
	Count                int64  `position:"Query" name:"Count"`
	VideoId              string `position:"Query" name:"VideoId"`
	Interval             int64  `position:"Query" name:"Interval"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SpriteSnapshotConfig string `position:"Query" name:"SpriteSnapshotConfig"`
	Height               string `position:"Query" name:"Height"`
}

func (req *SubmitSnapshotJobRequest) Invoke(client *sdk.Client) (resp *SubmitSnapshotJobResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "SubmitSnapshotJob", "vod", "")
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
	JobId common.String
}
