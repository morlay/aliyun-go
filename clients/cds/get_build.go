package cds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetBuildRequest struct {
	requests.RoaRequest
	BuildNumber int    `position:"Path" name:"BuildNumber"`
	JobName     string `position:"Path" name:"JobName"`
}

func (req *GetBuildRequest) Invoke(client *sdk.Client) (resp *GetBuildResponse, err error) {
	req.InitWithApiInfo("Cds", "2017-09-25", "GetBuild", "/v1/job/[JobName]/build/[BuildNumber]", "", "")
	req.Method = "GET"

	resp = &GetBuildResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetBuildResponse struct {
	responses.BaseResponse
	BuildEnv    string
	BuildNumber int
	Duration    int
	Log         string
	StartTime   int64
	RequestId   string
}
