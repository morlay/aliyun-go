package cds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StopBuildRequest struct {
	BuildNumber int    `position:"Path" name:"BuildNumber"`
	JobName     string `position:"Path" name:"JobName"`
}

func (r StopBuildRequest) Invoke(client *sdk.Client) (response *StopBuildResponse, err error) {
	req := struct {
		*requests.RoaRequest
		StopBuildRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("Cds", "2017-09-25", "StopBuild", "/v1/job/[JobName]/build/[BuildNumber]/stop", "", "")
	req.Method = "POST"

	resp := struct {
		*responses.BaseResponse
		StopBuildResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.StopBuildResponse

	err = client.DoAction(&req, &resp)
	return
}

type StopBuildResponse struct {
	RequestId string
}
