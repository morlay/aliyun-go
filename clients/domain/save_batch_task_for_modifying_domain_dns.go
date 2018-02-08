package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveBatchTaskForModifyingDomainDnsRequest struct {
	requests.RpcRequest
	UserClientIp      string                                                  `position:"Query" name:"UserClientIp"`
	DomainNames       *SaveBatchTaskForModifyingDomainDnsDomainNameList       `position:"Query" type:"Repeated" name:"DomainName"`
	DomainNameServers *SaveBatchTaskForModifyingDomainDnsDomainNameServerList `position:"Query" type:"Repeated" name:"DomainNameServer"`
	Lang              string                                                  `position:"Query" name:"Lang"`
	AliyunDns         string                                                  `position:"Query" name:"AliyunDns"`
}

func (req *SaveBatchTaskForModifyingDomainDnsRequest) Invoke(client *sdk.Client) (resp *SaveBatchTaskForModifyingDomainDnsResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "SaveBatchTaskForModifyingDomainDns", "", "")
	resp = &SaveBatchTaskForModifyingDomainDnsResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveBatchTaskForModifyingDomainDnsResponse struct {
	responses.BaseResponse
	RequestId string
	TaskNo    string
}

type SaveBatchTaskForModifyingDomainDnsDomainNameList []string

func (list *SaveBatchTaskForModifyingDomainDnsDomainNameList) UnmarshalJSON(data []byte) error {
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

type SaveBatchTaskForModifyingDomainDnsDomainNameServerList []string

func (list *SaveBatchTaskForModifyingDomainDnsDomainNameServerList) UnmarshalJSON(data []byte) error {
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
