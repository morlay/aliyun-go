package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryRegistrantProfilesRequest struct {
	requests.RpcRequest
	RegistrantOrganization   string `position:"Query" name:"RegistrantOrganization"`
	UserClientIp             string `position:"Query" name:"UserClientIp"`
	RegistrantProfileId      int64  `position:"Query" name:"RegistrantProfileId"`
	PageSize                 int    `position:"Query" name:"PageSize"`
	RegistrantType           string `position:"Query" name:"RegistrantType"`
	RealNameStatus           string `position:"Query" name:"RealNameStatus"`
	Lang                     string `position:"Query" name:"Lang"`
	PageNum                  int    `position:"Query" name:"PageNum"`
	DefaultRegistrantProfile string `position:"Query" name:"DefaultRegistrantProfile"`
	ZhRegistrantOrganization string `position:"Query" name:"ZhRegistrantOrganization"`
}

func (req *QueryRegistrantProfilesRequest) Invoke(client *sdk.Client) (resp *QueryRegistrantProfilesResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "QueryRegistrantProfiles", "", "")
	resp = &QueryRegistrantProfilesResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryRegistrantProfilesResponse struct {
	responses.BaseResponse
	RequestId          common.String
	TotalItemNum       common.Integer
	CurrentPageNum     common.Integer
	TotalPageNum       common.Integer
	PageSize           common.Integer
	PrePage            bool
	NextPage           bool
	RegistrantProfiles QueryRegistrantProfilesRegistrantProfileList
}

type QueryRegistrantProfilesRegistrantProfile struct {
	RegistrantProfileId      common.Long
	CreateTime               common.String
	UpdateTime               common.String
	DefaultRegistrantProfile bool
	RegistrantName           common.String
	RegistrantOrganization   common.String
	Country                  common.String
	Province                 common.String
	City                     common.String
	Address                  common.String
	Email                    common.String
	PostalCode               common.String
	TelArea                  common.String
	Telephone                common.String
	TelExt                   common.String
	EmailVerificationStatus  common.Integer
	ZhRegistrantName         common.String
	ZhRegistrantOrganization common.String
	ZhProvince               common.String
	ZhCity                   common.String
	ZhAddress                common.String
	RegistrantType           common.String
	RealNameStatus           common.String
}

type QueryRegistrantProfilesRegistrantProfileList []QueryRegistrantProfilesRegistrantProfile

func (list *QueryRegistrantProfilesRegistrantProfileList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryRegistrantProfilesRegistrantProfile)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
