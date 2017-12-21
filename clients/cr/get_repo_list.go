package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRepoListRequest struct {
	requests.RoaRequest
	Page     int `position:"Query" name:"Page"`
	PageSize int `position:"Query" name:"PageSize"`
}

func (req *GetRepoListRequest) Invoke(client *sdk.Client) (resp *GetRepoListResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "GetRepoList", "/repos", "", "")
	req.Method = "GET"

	resp = &GetRepoListResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetRepoListResponse struct {
	responses.BaseResponse
}
