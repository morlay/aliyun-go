package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetImageManifestRequest struct {
	requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
	Tag           string `position:"Path" name:"Tag"`
	SchemaVersion int    `position:"Query" name:"SchemaVersion"`
}

func (req *GetImageManifestRequest) Invoke(client *sdk.Client) (resp *GetImageManifestResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "GetImageManifest", "/repos/[RepoNamespace]/[RepoName]/tags/[Tag]/manifest", "", "")
	req.Method = "GET"

	resp = &GetImageManifestResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetImageManifestResponse struct {
	responses.BaseResponse
}
