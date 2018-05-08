package cds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	JobName             common.String
	SuccessRate         common.Float
	TotalBuilds         common.Integer
	RequestId           common.String
	LastSuccessfulBuild GetJobLastSuccessfulBuild
	LastFailedBuild     GetJobLastFailedBuild
}

type GetJobLastSuccessfulBuild struct {
	BuildEnv    common.String
	BuildNumber common.String
	Duration    common.Integer
	Log         common.String
	StartTime   common.Long
}

type GetJobLastFailedBuild struct {
	BuildEnv    common.String
	BuildNumber common.String
	Duration    common.Integer
	Log         common.String
	StartTime   common.Long
}
