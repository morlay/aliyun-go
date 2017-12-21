package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryDomainListRequest struct {
	requests.RpcRequest
	ProductDomainType string `position:"Query" name:"ProductDomainType"`
	RegStartDate      int64  `position:"Query" name:"RegStartDate"`
	OrderKeyType      string `position:"Query" name:"OrderKeyType"`
	GroupId           string `position:"Query" name:"GroupId"`
	DeadEndDate       int64  `position:"Query" name:"DeadEndDate"`
	DomainName        string `position:"Query" name:"DomainName"`
	StartDate         string `position:"Query" name:"StartDate"`
	PageNum           int    `position:"Query" name:"PageNum"`
	OrderByType       string `position:"Query" name:"OrderByType"`
	RegEndDate        int64  `position:"Query" name:"RegEndDate"`
	EndDate           string `position:"Query" name:"EndDate"`
	DomainType        string `position:"Query" name:"DomainType"`
	DeadStartDate     int64  `position:"Query" name:"DeadStartDate"`
	UserClientIp      string `position:"Query" name:"UserClientIp"`
	PageSize          int    `position:"Query" name:"PageSize"`
	Lang              string `position:"Query" name:"Lang"`
	QueryType         string `position:"Query" name:"QueryType"`
}

func (req *QueryDomainListRequest) Invoke(client *sdk.Client) (resp *QueryDomainListResponse, err error) {
	req.InitWithApiInfo("Domain", "2016-05-11", "QueryDomainList", "", "")
	resp = &QueryDomainListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryDomainListResponse struct {
	responses.BaseResponse
	RequestId      string
	TotalItemNum   int
	CurrentPageNum int
	TotalPageNum   int
	PageSize       int
	PrePage        bool
	NextPage       bool
	Data           QueryDomainListDomainList
}

type QueryDomainListDomain struct {
	DomainName        string
	SaleId            string
	DeadDate          string
	RegDate           string
	DomainAuditStatus string
	DomainRegType     string
	GroupId           string
	DomainType        string
	DomainStatus      string
	DeadDateStatus    string
	ProductId         string
	DeadDateLong      int64
	RegDateLong       int64
	Remark            string
	Premium           bool
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
