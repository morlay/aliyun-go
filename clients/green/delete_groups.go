package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteGroupsRequest struct {
	requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (req *DeleteGroupsRequest) Invoke(client *sdk.Client) (resp *DeleteGroupsResponse, err error) {
	req.InitWithApiInfo("Green", "2017-08-25", "DeleteGroups", "/green/sface/deleteGroupsOfPerson", "green", "")
	req.Method = "POST"

	resp = &DeleteGroupsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteGroupsResponse struct {
	responses.BaseResponse
}
