package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveRegistrantProfileRequest struct {
	requests.RpcRequest
	Country                  string `position:"Query" name:"Country"`
	Address                  string `position:"Query" name:"Address"`
	TelArea                  string `position:"Query" name:"TelArea"`
	City                     string `position:"Query" name:"City"`
	RegistrantProfileId      int64  `position:"Query" name:"RegistrantProfileId"`
	ZhAddress                string `position:"Query" name:"ZhAddress"`
	RegistrantType           string `position:"Query" name:"RegistrantType"`
	Telephone                string `position:"Query" name:"Telephone"`
	DefaultRegistrantProfile string `position:"Query" name:"DefaultRegistrantProfile"`
	ZhCity                   string `position:"Query" name:"ZhCity"`
	ZhProvince               string `position:"Query" name:"ZhProvince"`
	RegistrantOrganization   string `position:"Query" name:"RegistrantOrganization"`
	TelExt                   string `position:"Query" name:"TelExt"`
	Province                 string `position:"Query" name:"Province"`
	ZhRegistrantName         string `position:"Query" name:"ZhRegistrantName"`
	PostalCode               string `position:"Query" name:"PostalCode"`
	UserClientIp             string `position:"Query" name:"UserClientIp"`
	Lang                     string `position:"Query" name:"Lang"`
	Email                    string `position:"Query" name:"Email"`
	RegistrantName           string `position:"Query" name:"RegistrantName"`
	ZhRegistrantOrganization string `position:"Query" name:"ZhRegistrantOrganization"`
}

func (req *SaveRegistrantProfileRequest) Invoke(client *sdk.Client) (resp *SaveRegistrantProfileResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "SaveRegistrantProfile", "", "")
	resp = &SaveRegistrantProfileResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveRegistrantProfileResponse struct {
	responses.BaseResponse
	RequestId           string
	RegistrantProfileId int64
}
