package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeletePersonRequest struct {
	requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (req *DeletePersonRequest) Invoke(client *sdk.Client) (resp *DeletePersonResponse, err error) {
	req.InitWithApiInfo("Green", "2017-08-25", "DeletePerson", "/green/sface/deletePerson", "green", "")
	req.Method = "POST"

	resp = &DeletePersonResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeletePersonResponse struct {
	responses.BaseResponse
}
