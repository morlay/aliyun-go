package cds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RunJobRequest struct {
	requests.RoaRequest
	JobName string `position:"Path" name:"JobName"`
}

func (req *RunJobRequest) Invoke(client *sdk.Client) (resp *RunJobResponse, err error) {
	req.InitWithApiInfo("Cds", "2017-09-25", "RunJob", "/v1/job/[JobName]/run", "", "")
	req.Method = "POST"

	resp = &RunJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type RunJobResponse struct {
	responses.BaseResponse
	RequestId common.String
}
