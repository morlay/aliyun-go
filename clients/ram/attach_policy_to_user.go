package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AttachPolicyToUserRequest struct {
	requests.RpcRequest
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
	UserName   string `position:"Query" name:"UserName"`
}

func (req *AttachPolicyToUserRequest) Invoke(client *sdk.Client) (resp *AttachPolicyToUserResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "AttachPolicyToUser", "", "")
	resp = &AttachPolicyToUserResponse{}
	err = client.DoAction(req, resp)
	return
}

type AttachPolicyToUserResponse struct {
	responses.BaseResponse
	RequestId string
}
