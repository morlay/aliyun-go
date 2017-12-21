package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetSecurityPreferenceRequest struct {
	requests.RpcRequest
}

func (req *GetSecurityPreferenceRequest) Invoke(client *sdk.Client) (resp *GetSecurityPreferenceResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "GetSecurityPreference", "", "")
	resp = &GetSecurityPreferenceResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetSecurityPreferenceResponse struct {
	responses.BaseResponse
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
