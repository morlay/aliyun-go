package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetPersonRequest struct {
	requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (req *SetPersonRequest) Invoke(client *sdk.Client) (resp *SetPersonResponse, err error) {
	req.InitWithApiInfo("Green", "2017-08-25", "SetPerson", "/green/sface/setPerson", "green", "")
	req.Method = "POST"

	resp = &SetPersonResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetPersonResponse struct {
	responses.BaseResponse
}
