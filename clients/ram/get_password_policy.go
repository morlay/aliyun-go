package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetPasswordPolicyRequest struct {
	requests.RpcRequest
}

func (req *GetPasswordPolicyRequest) Invoke(client *sdk.Client) (resp *GetPasswordPolicyResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "GetPasswordPolicy", "", "")
	resp = &GetPasswordPolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetPasswordPolicyResponse struct {
	responses.BaseResponse
	RequestId      common.String
	PasswordPolicy GetPasswordPolicyPasswordPolicy
}

type GetPasswordPolicyPasswordPolicy struct {
	MinimumPasswordLength      common.Integer
	RequireLowercaseCharacters bool
	RequireUppercaseCharacters bool
	RequireNumbers             bool
	RequireSymbols             bool
	HardExpiry                 bool
	MaxPasswordAge             common.Integer
	PasswordReusePrevention    common.Integer
	MaxLoginAttemps            common.Integer
}
