package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SubmitPreprocessJobsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VideoId              string `position:"Query" name:"VideoId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PreprocessType       string `position:"Query" name:"PreprocessType"`
}

func (req *SubmitPreprocessJobsRequest) Invoke(client *sdk.Client) (resp *SubmitPreprocessJobsResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "SubmitPreprocessJobs", "vod", "")
	resp = &SubmitPreprocessJobsResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitPreprocessJobsResponse struct {
	responses.BaseResponse
	RequestId      common.String
	PreprocessJobs SubmitPreprocessJobsPreprocessJobList
}

type SubmitPreprocessJobsPreprocessJob struct {
	JobId common.String
}

type SubmitPreprocessJobsPreprocessJobList []SubmitPreprocessJobsPreprocessJob

func (list *SubmitPreprocessJobsPreprocessJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitPreprocessJobsPreprocessJob)
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
