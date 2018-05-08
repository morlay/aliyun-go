package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetPolicyVersionRequest struct {
	requests.RpcRequest
	VersionId  string `position:"Query" name:"VersionId"`
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
}

func (req *GetPolicyVersionRequest) Invoke(client *sdk.Client) (resp *GetPolicyVersionResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "GetPolicyVersion", "", "")
	resp = &GetPolicyVersionResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetPolicyVersionResponse struct {
	responses.BaseResponse
	RequestId     common.String
	PolicyVersion GetPolicyVersionPolicyVersion
}

type GetPolicyVersionPolicyVersion struct {
	VersionId        common.String
	IsDefaultVersion bool
	PolicyDocument   common.String
	CreateDate       common.String
}
