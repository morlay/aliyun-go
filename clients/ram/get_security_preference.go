package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetSecurityPreferenceRequest struct {
}

func (r GetSecurityPreferenceRequest) Invoke(client *sdk.Client) (response *GetSecurityPreferenceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetSecurityPreferenceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "GetSecurityPreference", "", "")

	resp := struct {
		*responses.BaseResponse
		GetSecurityPreferenceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetSecurityPreferenceResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetSecurityPreferenceResponse struct {
	RequestId          string
	SecurityPreference GetSecurityPreferenceSecurityPreference
}

type GetSecurityPreferenceSecurityPreference struct {
	LoginProfilePreference GetSecurityPreferenceLoginProfilePreference
	AccessKeyPreference    GetSecurityPreferenceAccessKeyPreference
	PublicKeyPreference    GetSecurityPreferencePublicKeyPreference
	MFAPreference          GetSecurityPreferenceMFAPreference
}

type GetSecurityPreferenceLoginProfilePreference struct {
	EnableSaveMFATicket       bool
	AllowUserToChangePassword bool
}

type GetSecurityPreferenceAccessKeyPreference struct {
	AllowUserToManageAccessKeys bool
}

type GetSecurityPreferencePublicKeyPreference struct {
	AllowUserToManagePublicKeys bool
}

type GetSecurityPreferenceMFAPreference struct {
	AllowUserToManageMFADevices bool
}
