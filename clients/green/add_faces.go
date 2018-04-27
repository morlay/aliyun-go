package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddFacesRequest struct {
	requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (req *AddFacesRequest) Invoke(client *sdk.Client) (resp *AddFacesResponse, err error) {
	req.InitWithApiInfo("Green", "2017-08-25", "AddFaces", "/green/sface/addFaces", "green", "")
	req.Method = "POST"

	resp = &AddFacesResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddFacesResponse struct {
	responses.BaseResponse
}
