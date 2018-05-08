package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetPasswordPolicyRequest struct {
	requests.RpcRequest
	RequireNumbers             string `position:"Query" name:"RequireNumbers"`
	PasswordReusePrevention    int    `position:"Query" name:"PasswordReusePrevention"`
	RequireUppercaseCharacters string `position:"Query" name:"RequireUppercaseCharacters"`
	MaxPasswordAge             int    `position:"Query" name:"MaxPasswordAge"`
	MaxLoginAttemps            int    `position:"Query" name:"MaxLoginAttemps"`
	HardExpiry                 string `position:"Query" name:"HardExpiry"`
	MinimumPasswordLength      int    `position:"Query" name:"MinimumPasswordLength"`
	RequireLowercaseCharacters string `position:"Query" name:"RequireLowercaseCharacters"`
	RequireSymbols             string `position:"Query" name:"RequireSymbols"`
}

func (req *SetPasswordPolicyRequest) Invoke(client *sdk.Client) (resp *SetPasswordPolicyResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "SetPasswordPolicy", "", "")
	resp = &SetPasswordPolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetPasswordPolicyResponse struct {
	responses.BaseResponse
	RequestId      common.String
	PasswordPolicy SetPasswordPolicyPasswordPolicy
}

type SetPasswordPolicyPasswordPolicy struct {
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
