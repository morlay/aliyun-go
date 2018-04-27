package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetGroupsRequest struct {
	requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (req *GetGroupsRequest) Invoke(client *sdk.Client) (resp *GetGroupsResponse, err error) {
	req.InitWithApiInfo("Green", "2017-08-25", "GetGroups", "/green/sface/getGroups", "green", "")
	req.Method = "POST"

	resp = &GetGroupsResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetGroupsResponse struct {
	responses.BaseResponse
}
