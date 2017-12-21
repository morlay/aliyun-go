package cds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetJobRequest struct {
	JobName string `position:"Path" name:"JobName"`
}

func (r GetJobRequest) Invoke(client *sdk.Client) (response *GetJobResponse, err error) {
	req := struct {
		*requests.RoaRequest
		GetJobRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("Cds", "2017-09-25", "GetJob", "/v1/job/[JobName]", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		GetJobResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetJobResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetJobResponse struct {
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
