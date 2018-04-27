package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteFacesRequest struct {
	requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (req *DeleteFacesRequest) Invoke(client *sdk.Client) (resp *DeleteFacesResponse, err error) {
	req.InitWithApiInfo("Green", "2017-08-25", "DeleteFaces", "/green/sface/deleteFaces", "green", "")
	req.Method = "POST"

	resp = &DeleteFacesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteFacesResponse struct {
	responses.BaseResponse
}
