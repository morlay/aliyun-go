package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryDomainBySaleIdRequest struct {
	SaleId       string `position:"Query" name:"SaleId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

func (r QueryDomainBySaleIdRequest) Invoke(client *sdk.Client) (response *QueryDomainBySaleIdResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryDomainBySaleIdRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Domain", "2016-05-11", "QueryDomainBySaleId", "", "")

	resp := struct {
		*responses.BaseResponse
		QueryDomainBySaleIdResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryDomainBySaleIdResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryDomainBySaleIdResponse struct {
	UserId               string
	DomainName           string
	SaleId               string
	CreationDate         string
	ExpirationDate       string
	DomainRegType        string
	EnglishHolder        string
	ChineseHolder        string
	EnglishContactPerson string
	ChineseContactPerson string
	HolderEmail          string
	TransferOutStatus    string
	SafetyLock           string
	TransferLock         string
	WhoisProtected       bool
	Premium              bool
	Remark               string
	DnsList              QueryDomainBySaleIdDnsListList
}

type QueryDomainBySaleIdDnsListList []string

func (list *QueryDomainBySaleIdDnsListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
