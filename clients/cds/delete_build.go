package cds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteBuildRequest struct {
	requests.RoaRequest
	BuildNumber int    `position:"Path" name:"BuildNumber"`
	JobName     string `position:"Path" name:"JobName"`
}

func (req *DeleteBuildRequest) Invoke(client *sdk.Client) (resp *DeleteBuildResponse, err error) {
	req.InitWithApiInfo("Cds", "2017-09-25", "DeleteBuild", "/v1/job/[JobName]/build/[BuildNumber]", "", "")
	req.Method = "DELETE"

	resp = &DeleteBuildResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteBuildResponse struct {
	responses.BaseResponse
	RequestId common.String
}
