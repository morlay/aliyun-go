package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetFacesRequest struct {
	requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (req *GetFacesRequest) Invoke(client *sdk.Client) (resp *GetFacesResponse, err error) {
	req.InitWithApiInfo("Green", "2017-08-25", "GetFaces", "/green/sface/getFaces", "green", "")
	req.Method = "POST"

	resp = &GetFacesResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetFacesResponse struct {
	responses.BaseResponse
}
