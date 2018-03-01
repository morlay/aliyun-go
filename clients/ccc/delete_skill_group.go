package ccc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteSkillGroupRequest struct {
	requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	SkillGroupId string `position:"Query" name:"SkillGroupId"`
}

func (req *DeleteSkillGroupRequest) Invoke(client *sdk.Client) (resp *DeleteSkillGroupResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "DeleteSkillGroup", "CCC", "")
	resp = &DeleteSkillGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteSkillGroupResponse struct {
	responses.BaseResponse
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
}
