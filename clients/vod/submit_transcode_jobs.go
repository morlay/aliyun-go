package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SubmitTranscodeJobsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	TemplateGroupId      string `position:"Query" name:"TemplateGroupId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VideoId              string `position:"Query" name:"VideoId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	EncryptConfig        string `position:"Query" name:"EncryptConfig"`
	PipelineId           string `position:"Query" name:"PipelineId"`
}

func (req *SubmitTranscodeJobsRequest) Invoke(client *sdk.Client) (resp *SubmitTranscodeJobsResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "SubmitTranscodeJobs", "vod", "")
	resp = &SubmitTranscodeJobsResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitTranscodeJobsResponse struct {
	responses.BaseResponse
	RequestId     common.String
	TranscodeJobs SubmitTranscodeJobsTranscodeJobList
}

type SubmitTranscodeJobsTranscodeJob struct {
	JobId common.String
}

type SubmitTranscodeJobsTranscodeJobList []SubmitTranscodeJobsTranscodeJob

func (list *SubmitTranscodeJobsTranscodeJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitTranscodeJobsTranscodeJob)
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
