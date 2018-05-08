package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryAdvancedDomainListRequest struct {
	requests.RpcRequest
	ProductDomainType     string `position:"Query" name:"ProductDomainType"`
	PageNum               int    `position:"Query" name:"PageNum"`
	Excluded              string `position:"Query" name:"Excluded"`
	StartLength           int    `position:"Query" name:"StartLength"`
	ExcludedSuffix        string `position:"Query" name:"ExcludedSuffix"`
	PageSize              int    `position:"Query" name:"PageSize"`
	Lang                  string `position:"Query" name:"Lang"`
	ExcludedPrefix        string `position:"Query" name:"ExcludedPrefix"`
	KeyWord               string `position:"Query" name:"KeyWord"`
	ProductDomainTypeSort string `position:"Query" name:"ProductDomainTypeSort"`
	EndExpirationDate     int64  `position:"Query" name:"EndExpirationDate"`
	Suffixs               string `position:"Query" name:"Suffixs"`
	DomainNameSort        string `position:"Query" name:"DomainNameSort"`
	ExpirationDateSort    string `position:"Query" name:"ExpirationDateSort"`
	StartExpirationDate   int64  `position:"Query" name:"StartExpirationDate"`
	DomainStatus          int    `position:"Query" name:"DomainStatus"`
	DomainGroupId         int64  `position:"Query" name:"DomainGroupId"`
	KeyWordSuffix         string `position:"Query" name:"KeyWordSuffix"`
	KeyWordPrefix         string `position:"Query" name:"KeyWordPrefix"`
	TradeType             int    `position:"Query" name:"TradeType"`
	EndRegistrationDate   int64  `position:"Query" name:"EndRegistrationDate"`
	Form                  int    `position:"Query" name:"Form"`
	UserClientIp          string `position:"Query" name:"UserClientIp"`
	RegistrationDateSort  string `position:"Query" name:"RegistrationDateSort"`
	StartRegistrationDate int64  `position:"Query" name:"StartRegistrationDate"`
	EndLength             int    `position:"Query" name:"EndLength"`
}

func (req *QueryAdvancedDomainListRequest) Invoke(client *sdk.Client) (resp *QueryAdvancedDomainListResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "QueryAdvancedDomainList", "", "")
	resp = &QueryAdvancedDomainListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryAdvancedDomainListResponse struct {
	responses.BaseResponse
	RequestId      common.String
	TotalItemNum   common.Integer
	CurrentPageNum common.Integer
	TotalPageNum   common.Integer
	PageSize       common.Integer
	PrePage        bool
	NextPage       bool
	Data           QueryAdvancedDomainListDomainList
}

type QueryAdvancedDomainListDomain struct {
	DomainName             common.String
	InstanceId             common.String
	ExpirationDate         common.String
	RegistrationDate       common.String
	DomainType             common.String
	DomainStatus           common.String
	ProductId              common.String
	ExpirationDateLong     common.Long
	RegistrationDateLong   common.Long
	Premium                bool
	DomainAuditStatus      common.String
	ExpirationDateStatus   common.String
	RegistrantType         common.String
	DomainGroupId          common.String
	Remark                 common.String
	DomainGroupName        common.String
	ExpirationCurrDateDiff common.Integer
}

type QueryAdvancedDomainListDomainList []QueryAdvancedDomainListDomain

func (list *QueryAdvancedDomainListDomainList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAdvancedDomainListDomain)
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
