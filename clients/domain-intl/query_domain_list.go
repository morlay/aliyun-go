package domain_intl

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryDomainListRequest struct {
	requests.RpcRequest
	EndExpirationDate     int64  `position:"Query" name:"EndExpirationDate"`
	ProductDomainType     string `position:"Query" name:"ProductDomainType"`
	OrderKeyType          string `position:"Query" name:"OrderKeyType"`
	DomainName            string `position:"Query" name:"DomainName"`
	StartExpirationDate   int64  `position:"Query" name:"StartExpirationDate"`
	PageNum               int    `position:"Query" name:"PageNum"`
	OrderByType           string `position:"Query" name:"OrderByType"`
	EndRegistrationDate   int64  `position:"Query" name:"EndRegistrationDate"`
	UserClientIp          string `position:"Query" name:"UserClientIp"`
	PageSize              int    `position:"Query" name:"PageSize"`
	Lang                  string `position:"Query" name:"Lang"`
	QueryType             string `position:"Query" name:"QueryType"`
	StartRegistrationDate int64  `position:"Query" name:"StartRegistrationDate"`
}

func (req *QueryDomainListRequest) Invoke(client *sdk.Client) (resp *QueryDomainListResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "QueryDomainList", "domain", "")
	resp = &QueryDomainListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryDomainListResponse struct {
	responses.BaseResponse
	RequestId      common.String
	TotalItemNum   common.Integer
	CurrentPageNum common.Integer
	TotalPageNum   common.Integer
	PageSize       common.Integer
	PrePage        bool
	NextPage       bool
	Data           QueryDomainListDomainList
}

type QueryDomainListDomain struct {
	DomainName           common.String
	InstanceId           common.String
	ExpirationDate       common.String
	RegistrationDate     common.String
	DomainType           common.String
	DomainStatus         common.String
	ProductId            common.String
	ExpirationDateLong   common.Long
	RegistrationDateLong common.Long
	Premium              bool
}

type QueryDomainListDomainList []QueryDomainListDomain

func (list *QueryDomainListDomainList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryDomainListDomain)
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
