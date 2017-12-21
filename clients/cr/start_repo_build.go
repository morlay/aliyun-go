package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StartRepoBuildRequest struct {
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
}

func (r StartRepoBuildRequest) Invoke(client *sdk.Client) (response *StartRepoBuildResponse, err error) {
	req := struct {
		*requests.RoaRequest
		StartRepoBuildRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("cr", "2016-06-07", "StartRepoBuild", "/repos/[RepoNamespace]/[RepoName]/build", "", "")
	req.Method = "PUT"

	resp := struct {
		*responses.BaseResponse
		StartRepoBuildResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.StartRepoBuildResponse

	err = client.DoAction(&req, &resp)
	return
}

type StartRepoBuildResponse struct {
}
