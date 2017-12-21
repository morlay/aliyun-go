package cds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetBuildRequest struct {
	BuildNumber int    `position:"Path" name:"BuildNumber"`
	JobName     string `position:"Path" name:"JobName"`
}

func (r GetBuildRequest) Invoke(client *sdk.Client) (response *GetBuildResponse, err error) {
	req := struct {
		*requests.RoaRequest
		GetBuildRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("Cds", "2017-09-25", "GetBuild", "/v1/job/[JobName]/build/[BuildNumber]", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		GetBuildResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetBuildResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetBuildResponse struct {
	BuildEnv    string
	BuildNumber int
	Duration    int
	Log         string
	StartTime   int64
	RequestId   string
}
