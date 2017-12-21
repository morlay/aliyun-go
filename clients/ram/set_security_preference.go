package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetSecurityPreferenceRequest struct {
	requests.RpcRequest
	AllowUserToManageAccessKeys string `position:"Query" name:"AllowUserToManageAccessKeys"`
	AllowUserToManageMFADevices string `position:"Query" name:"AllowUserToManageMFADevices"`
	AllowUserToManagePublicKeys string `position:"Query" name:"AllowUserToManagePublicKeys"`
	EnableSaveMFATicket         string `position:"Query" name:"EnableSaveMFATicket"`
	AllowUserToChangePassword   string `position:"Query" name:"AllowUserToChangePassword"`
}

func (req *SetSecurityPreferenceRequest) Invoke(client *sdk.Client) (resp *SetSecurityPreferenceResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "SetSecurityPreference", "", "")
	resp = &SetSecurityPreferenceResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetSecurityPreferenceResponse struct {
	responses.BaseResponse
	RequestId          string
	SecurityPreference SetSecurityPreferenceSecurityPreference
}

type SetSecurityPreferenceSecurityPreference struct {
	LoginProfilePreference SetSecurityPreferenceLoginProfilePreference
	AccessKeyPreference    SetSecurityPreferenceAccessKeyPreference
	PublicKeyPreference    SetSecurityPreferencePublicKeyPreference
	MFAPreference          SetSecurityPreferenceMFAPreference
}

type SetSecurityPreferenceLoginProfilePreference struct {
	EnableSaveMFATicket       bool
	AllowUserToChangePassword bool
}

type SetSecurityPreferenceAccessKeyPreference struct {
	AllowUserToManageAccessKeys bool
}

type SetSecurityPreferencePublicKeyPreference struct {
	AllowUserToManagePublicKeys bool
}

type SetSecurityPreferenceMFAPreference struct {
	AllowUserToManageMFADevices bool
}
