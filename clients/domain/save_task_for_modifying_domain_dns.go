package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveTaskForModifyingDomainDnsRequest struct {
	SaleId       string                                    `position:"Query" name:"SaleId"`
	UserClientIp string                                    `position:"Query" name:"UserClientIp"`
	DomainName   string                                    `position:"Query" name:"DomainName"`
	Lang         string                                    `position:"Query" name:"Lang"`
	AliyunDns    string                                    `position:"Query" name:"AliyunDns"`
	DnsLists     *SaveTaskForModifyingDomainDnsDnsListList `position:"Query" type:"Repeated" name:"DnsList"`
}

func (r SaveTaskForModifyingDomainDnsRequest) Invoke(client *sdk.Client) (response *SaveTaskForModifyingDomainDnsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SaveTaskForModifyingDomainDnsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Domain", "2016-05-11", "SaveTaskForModifyingDomainDns", "", "")

	resp := struct {
		*responses.BaseResponse
		SaveTaskForModifyingDomainDnsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SaveTaskForModifyingDomainDnsResponse

	err = client.DoAction(&req, &resp)
	return
}

type SaveTaskForModifyingDomainDnsResponse struct {
	RequestId string
	TaskNo    string
}

type SaveTaskForModifyingDomainDnsDnsListList []string

func (list *SaveTaskForModifyingDomainDnsDnsListList) UnmarshalJSON(data []byte) error {
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
