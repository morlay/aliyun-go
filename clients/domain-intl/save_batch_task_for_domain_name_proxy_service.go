package domain_intl

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveBatchTaskForDomainNameProxyServiceRequest struct {
	requests.RpcRequest
	UserClientIp string                                                `position:"Query" name:"UserClientIp"`
	DomainNames  *SaveBatchTaskForDomainNameProxyServiceDomainNameList `position:"Query" type:"Repeated" name:"DomainName"`
	Lang         string                                                `position:"Query" name:"Lang"`
	Status       string                                                `position:"Query" name:"Status"`
}

func (req *SaveBatchTaskForDomainNameProxyServiceRequest) Invoke(client *sdk.Client) (resp *SaveBatchTaskForDomainNameProxyServiceResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveBatchTaskForDomainNameProxyService", "domain", "")
	resp = &SaveBatchTaskForDomainNameProxyServiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveBatchTaskForDomainNameProxyServiceResponse struct {
	responses.BaseResponse
	RequestId string
	TaskNo    string
}

type SaveBatchTaskForDomainNameProxyServiceDomainNameList []string

func (list *SaveBatchTaskForDomainNameProxyServiceDomainNameList) UnmarshalJSON(data []byte) error {
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
