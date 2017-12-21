package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRepoListRequest struct {
	Page     int `position:"Query" name:"Page"`
	PageSize int `position:"Query" name:"PageSize"`
}

func (r GetRepoListRequest) Invoke(client *sdk.Client) (response *GetRepoListResponse, err error) {
	req := struct {
		*requests.RoaRequest
		GetRepoListRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("cr", "2016-06-07", "GetRepoList", "/repos", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		GetRepoListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetRepoListResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetRepoListResponse struct {
}
