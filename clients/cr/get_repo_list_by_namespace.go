package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRepoListByNamespaceRequest struct {
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	Page          int    `position:"Query" name:"Page"`
	PageSize      int    `position:"Query" name:"PageSize"`
}

func (r GetRepoListByNamespaceRequest) Invoke(client *sdk.Client) (response *GetRepoListByNamespaceResponse, err error) {
	req := struct {
		*requests.RoaRequest
		GetRepoListByNamespaceRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("cr", "2016-06-07", "GetRepoListByNamespace", "/repos/[RepoNamespace]", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		GetRepoListByNamespaceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetRepoListByNamespaceResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetRepoListByNamespaceResponse struct {
}
