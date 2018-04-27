package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetPersonsRequest struct {
	requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (req *GetPersonsRequest) Invoke(client *sdk.Client) (resp *GetPersonsResponse, err error) {
	req.InitWithApiInfo("Green", "2017-08-25", "GetPersons", "/green/sface/getPersons", "green", "")
	req.Method = "POST"

	resp = &GetPersonsResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetPersonsResponse struct {
	responses.BaseResponse
}
