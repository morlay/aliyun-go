package cds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type StopBuildRequest struct {
	requests.RoaRequest
	BuildNumber int    `position:"Path" name:"BuildNumber"`
	JobName     string `position:"Path" name:"JobName"`
}

func (req *StopBuildRequest) Invoke(client *sdk.Client) (resp *StopBuildResponse, err error) {
	req.InitWithApiInfo("Cds", "2017-09-25", "StopBuild", "/v1/job/[JobName]/build/[BuildNumber]/stop", "", "")
	req.Method = "POST"

	resp = &StopBuildResponse{}
	err = client.DoAction(req, resp)
	return
}

type StopBuildResponse struct {
	responses.BaseResponse
	RequestId common.String
}
