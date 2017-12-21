package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetPasswordPolicyRequest struct {
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

func (r SetPasswordPolicyRequest) Invoke(client *sdk.Client) (response *SetPasswordPolicyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetPasswordPolicyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "SetPasswordPolicy", "", "")

	resp := struct {
		*responses.BaseResponse
		SetPasswordPolicyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetPasswordPolicyResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetPasswordPolicyResponse struct {
	RequestId      string
	PasswordPolicy SetPasswordPolicyPasswordPolicy
}

type SetPasswordPolicyPasswordPolicy struct {
	MinimumPasswordLength      int
	RequireLowercaseCharacters bool
	RequireUppercaseCharacters bool
	RequireNumbers             bool
	RequireSymbols             bool
	HardExpiry                 bool
	MaxPasswordAge             int
	PasswordReusePrevention    int
	MaxLoginAttemps            int
}
