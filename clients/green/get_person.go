package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetPersonRequest struct {
	requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (req *GetPersonRequest) Invoke(client *sdk.Client) (resp *GetPersonResponse, err error) {
	req.InitWithApiInfo("Green", "2017-08-25", "GetPerson", "/green/sface/getPerson", "green", "")
	req.Method = "POST"

	resp = &GetPersonResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetPersonResponse struct {
	responses.BaseResponse
}
