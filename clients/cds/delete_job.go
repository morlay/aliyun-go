package cds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteJobRequest struct {
	JobName string `position:"Path" name:"JobName"`
}

func (r DeleteJobRequest) Invoke(client *sdk.Client) (response *DeleteJobResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DeleteJobRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("Cds", "2017-09-25", "DeleteJob", "/v1/job/[JobName]", "", "")
	req.Method = "DELETE"

	resp := struct {
		*responses.BaseResponse
		DeleteJobResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteJobResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteJobResponse struct {
	RequestId string
}
