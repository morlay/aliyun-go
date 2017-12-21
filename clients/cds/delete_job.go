package cds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteJobRequest struct {
	requests.RoaRequest
	JobName string `position:"Path" name:"JobName"`
}

func (req *DeleteJobRequest) Invoke(client *sdk.Client) (resp *DeleteJobResponse, err error) {
	req.InitWithApiInfo("Cds", "2017-09-25", "DeleteJob", "/v1/job/[JobName]", "", "")
	req.Method = "DELETE"

	resp = &DeleteJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteJobResponse struct {
	responses.BaseResponse
	RequestId string
}
