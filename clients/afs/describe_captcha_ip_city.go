package afs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCaptchaIpCityRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	ConfigName      string `position:"Query" name:"ConfigName"`
	Time            string `position:"Query" name:"Time"`
	Type            string `position:"Query" name:"Type"`
}

func (req *DescribeCaptchaIpCityRequest) Invoke(client *sdk.Client) (resp *DescribeCaptchaIpCityResponse, err error) {
	req.InitWithApiInfo("afs", "2018-01-12", "DescribeCaptchaIpCity", "", "")
	resp = &DescribeCaptchaIpCityResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCaptchaIpCityResponse struct {
	responses.BaseResponse
	RequestId     common.String
	BizCode       common.String
	HasData       bool
	CaptchaCities DescribeCaptchaIpCityCaptchaCitieList
	CaptchaIps    DescribeCaptchaIpCityCaptchaIpList
}

type DescribeCaptchaIpCityCaptchaCitie struct {
	Location common.String
	Lat      common.String
	Lng      common.String
	Pv       common.Integer
}

type DescribeCaptchaIpCityCaptchaIp struct {
	Ip    common.String
	Value common.Integer
}

type DescribeCaptchaIpCityCaptchaCitieList []DescribeCaptchaIpCityCaptchaCitie

func (list *DescribeCaptchaIpCityCaptchaCitieList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCaptchaIpCityCaptchaCitie)
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

type DescribeCaptchaIpCityCaptchaIpList []DescribeCaptchaIpCityCaptchaIp

func (list *DescribeCaptchaIpCityCaptchaIpList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCaptchaIpCityCaptchaIp)
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
