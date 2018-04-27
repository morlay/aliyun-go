package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddPersonRequest struct {
	requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (req *AddPersonRequest) Invoke(client *sdk.Client) (resp *AddPersonResponse, err error) {
	req.InitWithApiInfo("Green", "2017-08-25", "AddPerson", "/green/sface/addPerson", "green", "")
	req.Method = "POST"

	resp = &AddPersonResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddPersonResponse struct {
	responses.BaseResponse
}
