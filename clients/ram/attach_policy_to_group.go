package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AttachPolicyToGroupRequest struct {
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
	GroupName  string `position:"Query" name:"GroupName"`
}

func (r AttachPolicyToGroupRequest) Invoke(client *sdk.Client) (response *AttachPolicyToGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AttachPolicyToGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "AttachPolicyToGroup", "", "")

	resp := struct {
		*responses.BaseResponse
		AttachPolicyToGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AttachPolicyToGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type AttachPolicyToGroupResponse struct {
	RequestId string
}
