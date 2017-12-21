package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRepoListByNamespaceRequest struct {
	requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	Page          int    `position:"Query" name:"Page"`
	PageSize      int    `position:"Query" name:"PageSize"`
}

func (req *GetRepoListByNamespaceRequest) Invoke(client *sdk.Client) (resp *GetRepoListByNamespaceResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "GetRepoListByNamespace", "/repos/[RepoNamespace]", "", "")
	req.Method = "GET"

	resp = &GetRepoListByNamespaceResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetRepoListByNamespaceResponse struct {
	responses.BaseResponse
}
