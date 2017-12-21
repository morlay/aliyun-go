package cds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RunJobRequest struct {
	JobName string `position:"Path" name:"JobName"`
}

func (r RunJobRequest) Invoke(client *sdk.Client) (response *RunJobResponse, err error) {
	req := struct {
		*requests.RoaRequest
		RunJobRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("Cds", "2017-09-25", "RunJob", "/v1/job/[JobName]/run", "", "")
	req.Method = "POST"

	resp := struct {
		*responses.BaseResponse
		RunJobResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RunJobResponse

	err = client.DoAction(&req, &resp)
	return
}

type RunJobResponse struct {
	RequestId string
}
