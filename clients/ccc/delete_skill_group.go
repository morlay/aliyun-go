package ccc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteSkillGroupRequest struct {
	requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	SkillGroupId string `position:"Query" name:"SkillGroupId"`
}

func (req *DeleteSkillGroupRequest) Invoke(client *sdk.Client) (resp *DeleteSkillGroupResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "DeleteSkillGroup", "ccc", "")
	resp = &DeleteSkillGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteSkillGroupResponse struct {
	responses.BaseResponse
	RequestId      common.String
	Success        bool
	Code           common.String
	Message        common.String
	HttpStatusCode common.Integer
}
