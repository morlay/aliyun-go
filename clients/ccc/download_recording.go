package ccc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DownloadRecordingRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	FileName   string `position:"Query" name:"FileName"`
	Channel    string `position:"Query" name:"Channel"`
}

func (req *DownloadRecordingRequest) Invoke(client *sdk.Client) (resp *DownloadRecordingResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "DownloadRecording", "ccc", "")
	resp = &DownloadRecordingResponse{}
	err = client.DoAction(req, resp)
	return
}

type DownloadRecordingResponse struct {
	responses.BaseResponse
	RequestId          common.String
	Success            bool
	Code               common.String
	Message            common.String
	HttpStatusCode     common.Integer
	MediaDownloadParam DownloadRecordingMediaDownloadParam
}

type DownloadRecordingMediaDownloadParam struct {
	SignatureUrl common.String
	FileName     common.String
}
