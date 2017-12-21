package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AttachPolicyToGroupRequest struct {
	requests.RpcRequest
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
	GroupName  string `position:"Query" name:"GroupName"`
}

func (req *AttachPolicyToGroupRequest) Invoke(client *sdk.Client) (resp *AttachPolicyToGroupResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "AttachPolicyToGroup", "", "")
	resp = &AttachPolicyToGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type AttachPolicyToGroupResponse struct {
	responses.BaseResponse
	RequestId string
}
