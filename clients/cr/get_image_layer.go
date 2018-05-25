package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetImageLayerRequest struct {
	requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
	Tag           string `position:"Path" name:"Tag"`
}

func (req *GetImageLayerRequest) Invoke(client *sdk.Client) (resp *GetImageLayerResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "GetImageLayer", "/repos/[RepoNamespace]/[RepoName]/tags/[Tag]/layers", "", "")
	req.Method = "GET"

	resp = &GetImageLayerResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetImageLayerResponse struct {
	responses.BaseResponse
}
