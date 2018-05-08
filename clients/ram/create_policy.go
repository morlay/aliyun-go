package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreatePolicyRequest struct {
	requests.RpcRequest
	Description    string `position:"Query" name:"Description"`
	PolicyName     string `position:"Query" name:"PolicyName"`
	PolicyDocument string `position:"Query" name:"PolicyDocument"`
}

func (req *CreatePolicyRequest) Invoke(client *sdk.Client) (resp *CreatePolicyResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "CreatePolicy", "", "")
	resp = &CreatePolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreatePolicyResponse struct {
	responses.BaseResponse
	RequestId common.String
	Policy    CreatePolicyPolicy
}

type CreatePolicyPolicy struct {
	PolicyName     common.String
	PolicyType     common.String
	Description    common.String
	DefaultVersion common.String
	CreateDate     common.String
}
