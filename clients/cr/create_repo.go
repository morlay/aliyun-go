package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateRepoRequest struct {
	requests.RoaRequest
}

func (req *CreateRepoRequest) Invoke(client *sdk.Client) (resp *CreateRepoResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "CreateRepo", "/repos", "", "")
	req.Method = "PUT"

	resp = &CreateRepoResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateRepoResponse struct {
	responses.BaseResponse
}
