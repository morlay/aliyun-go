package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateMyGroupMembersRequest struct {
	requests.RpcRequest
	GroupId int64  `position:"Query" name:"GroupId"`
	Masters string `position:"Query" name:"Masters"`
}

func (req *UpdateMyGroupMembersRequest) Invoke(client *sdk.Client) (resp *UpdateMyGroupMembersResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "UpdateMyGroupMembers", "cms", "")
	resp = &UpdateMyGroupMembersResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateMyGroupMembersResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorCode    int
	ErrorMessage string
}
