package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddGroupsRequest struct {
	requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (req *AddGroupsRequest) Invoke(client *sdk.Client) (resp *AddGroupsResponse, err error) {
	req.InitWithApiInfo("Green", "2017-08-25", "AddGroups", "/green/sface/addGroupsOfPerson", "green", "")
	req.Method = "POST"

	resp = &AddGroupsResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddGroupsResponse struct {
	responses.BaseResponse
}
