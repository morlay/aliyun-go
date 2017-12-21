package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateRepoRequest struct {
}

func (r CreateRepoRequest) Invoke(client *sdk.Client) (response *CreateRepoResponse, err error) {
	req := struct {
		*requests.RoaRequest
		CreateRepoRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("cr", "2016-06-07", "CreateRepo", "/repos", "", "")
	req.Method = "PUT"

	resp := struct {
		*responses.BaseResponse
		CreateRepoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateRepoResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateRepoResponse struct {
}
