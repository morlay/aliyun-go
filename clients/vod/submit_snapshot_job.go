package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubmitSnapshotJobRequest struct {
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

func (r SubmitSnapshotJobRequest) Invoke(client *sdk.Client) (response *SubmitSnapshotJobResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SubmitSnapshotJobRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "SubmitSnapshotJob", "", "")

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
	JobId string
}
