package domain_intl

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SaveRegistrantProfileRequest struct {
	requests.RpcRequest
	Country                  string `position:"Query" name:"Country"`
	Address                  string `position:"Query" name:"Address"`
	TelArea                  string `position:"Query" name:"TelArea"`
	City                     string `position:"Query" name:"City"`
	RegistrantProfileId      int64  `position:"Query" name:"RegistrantProfileId"`
	Telephone                string `position:"Query" name:"Telephone"`
	DefaultRegistrantProfile string `position:"Query" name:"DefaultRegistrantProfile"`
	RegistrantOrganization   string `position:"Query" name:"RegistrantOrganization"`
	TelExt                   string `position:"Query" name:"TelExt"`
	Province                 string `position:"Query" name:"Province"`
	PostalCode               string `position:"Query" name:"PostalCode"`
	UserClientIp             string `position:"Query" name:"UserClientIp"`
	Lang                     string `position:"Query" name:"Lang"`
	Email                    string `position:"Query" name:"Email"`
	RegistrantName           string `position:"Query" name:"RegistrantName"`
}

func (req *SaveRegistrantProfileRequest) Invoke(client *sdk.Client) (resp *SaveRegistrantProfileResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveRegistrantProfile", "domain", "")
	resp = &SaveRegistrantProfileResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveRegistrantProfileResponse struct {
	responses.BaseResponse
	RequestId           common.String
	RegistrantProfileId common.Long
}
