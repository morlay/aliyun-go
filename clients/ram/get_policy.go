package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Policy    GetPolicyPolicy
}

type GetPolicyPolicy struct {
	PolicyName      common.String
	PolicyType      common.String
	Description     common.String
	DefaultVersion  common.String
	PolicyDocument  common.String
	CreateDate      common.String
	UpdateDate      common.String
	AttachmentCount common.Integer
}
