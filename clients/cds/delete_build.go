package cds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteBuildRequest struct {
	BuildNumber int    `position:"Path" name:"BuildNumber"`
	JobName     string `position:"Path" name:"JobName"`
}

func (r DeleteBuildRequest) Invoke(client *sdk.Client) (response *DeleteBuildResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DeleteBuildRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("Cds", "2017-09-25", "DeleteBuild", "/v1/job/[JobName]/build/[BuildNumber]", "", "")
	req.Method = "DELETE"

	resp := struct {
		*responses.BaseResponse
		DeleteBuildResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteBuildResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteBuildResponse struct {
	RequestId string
}
