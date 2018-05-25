package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteImageRequest struct {
	requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
	Tag           string `position:"Path" name:"Tag"`
}

func (req *DeleteImageRequest) Invoke(client *sdk.Client) (resp *DeleteImageResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "DeleteImage", "/repos/[RepoNamespace]/[RepoName]/tags/[Tag]", "", "")
	req.Method = "DELETE"

	resp = &DeleteImageResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteImageResponse struct {
	responses.BaseResponse
}
