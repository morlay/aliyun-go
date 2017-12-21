package cds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetJobRequest struct {
	requests.RoaRequest
	JobName string `position:"Path" name:"JobName"`
}

func (req *GetJobRequest) Invoke(client *sdk.Client) (resp *GetJobResponse, err error) {
	req.InitWithApiInfo("Cds", "2017-09-25", "GetJob", "/v1/job/[JobName]", "", "")
	req.Method = "GET"

	resp = &GetJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetJobResponse struct {
	responses.BaseResponse
	JobName             string
	SuccessRate         float32
	TotalBuilds         int
	RequestId           string
	LastSuccessfulBuild GetJobLastSuccessfulBuild
	LastFailedBuild     GetJobLastFailedBuild
}

type GetJobLastSuccessfulBuild struct {
	BuildEnv    string
	BuildNumber string
	Duration    int
	Log         string
	StartTime   int64
}

type GetJobLastFailedBuild struct {
	BuildEnv    string
	BuildNumber string
	Duration    int
	Log         string
	StartTime   int64
}
