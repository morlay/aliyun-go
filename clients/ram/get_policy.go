package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetPolicyRequest struct {
	requests.RpcRequest
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
}

func (req *GetPolicyRequest) Invoke(client *sdk.Client) (resp *GetPolicyResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "GetPolicy", "", "")
	resp = &GetPolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetPolicyResponse struct {
	responses.BaseResponse
	RequestId string
	Policy    GetPolicyPolicy
}

type GetPolicyPolicy struct {
	PolicyName      string
	PolicyType      string
	Description     string
	DefaultVersion  string
	PolicyDocument  string
	CreateDate      string
	UpdateDate      string
	AttachmentCount int
}
