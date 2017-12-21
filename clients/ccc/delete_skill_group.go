package ccc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteSkillGroupRequest struct {
	InstanceId   string `position:"Query" name:"InstanceId"`
	SkillGroupId string `position:"Query" name:"SkillGroupId"`
}

func (r DeleteSkillGroupRequest) Invoke(client *sdk.Client) (response *DeleteSkillGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteSkillGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CCC", "2017-07-05", "DeleteSkillGroup", "ccc", "")

	resp := struct {
		*responses.BaseResponse
		DeleteSkillGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteSkillGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteSkillGroupResponse struct {
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
}
