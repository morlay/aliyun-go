package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetPasswordPolicyRequest struct {
}

func (r GetPasswordPolicyRequest) Invoke(client *sdk.Client) (response *GetPasswordPolicyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetPasswordPolicyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "GetPasswordPolicy", "", "")

	resp := struct {
		*responses.BaseResponse
		GetPasswordPolicyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetPasswordPolicyResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetPasswordPolicyResponse struct {
	RequestId      string
	PasswordPolicy GetPasswordPolicyPasswordPolicy
}

type GetPasswordPolicyPasswordPolicy struct {
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
