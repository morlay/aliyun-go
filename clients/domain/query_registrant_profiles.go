package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId          string
	TotalItemNum       int
	CurrentPageNum     int
	TotalPageNum       int
	PageSize           int
	PrePage            bool
	NextPage           bool
	RegistrantProfiles QueryRegistrantProfilesRegistrantProfileList
}

type QueryRegistrantProfilesRegistrantProfile struct {
	RegistrantProfileId      int64
	CreateTime               string
	UpdateTime               string
	DefaultRegistrantProfile bool
	RegistrantName           string
	RegistrantOrganization   string
	Country                  string
	Province                 string
	City                     string
	Address                  string
	Email                    string
	PostalCode               string
	TelArea                  string
	Telephone                string
	TelExt                   string
	EmailVerificationStatus  int
	ZhRegistrantName         string
	ZhRegistrantOrganization string
	ZhProvince               string
	ZhCity                   string
	ZhAddress                string
	RegistrantType           string
	RealNameStatus           string
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
