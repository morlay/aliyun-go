package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SaveBatchTaskForDomainNameProxyServiceRequest struct {
	requests.RpcRequest
	UserClientIp string                                                `position:"Query" name:"UserClientIp"`
	DomainNames  *SaveBatchTaskForDomainNameProxyServiceDomainNameList `position:"Query" type:"Repeated" name:"DomainName"`
	Lang         string                                                `position:"Query" name:"Lang"`
	Status       string                                                `position:"Query" name:"Status"`
}

func (req *SaveBatchTaskForDomainNameProxyServiceRequest) Invoke(client *sdk.Client) (resp *SaveBatchTaskForDomainNameProxyServiceResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "SaveBatchTaskForDomainNameProxyService", "", "")
	resp = &SaveBatchTaskForDomainNameProxyServiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveBatchTaskForDomainNameProxyServiceResponse struct {
	responses.BaseResponse
	RequestId common.String
	TaskNo    common.String
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
