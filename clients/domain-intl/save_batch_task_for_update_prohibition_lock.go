package domain_intl

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveBatchTaskForUpdateProhibitionLockRequest struct {
	requests.RpcRequest
	UserClientIp string                                               `position:"Query" name:"UserClientIp"`
	DomainNames  *SaveBatchTaskForUpdateProhibitionLockDomainNameList `position:"Query" type:"Repeated" name:"DomainName"`
	Lang         string                                               `position:"Query" name:"Lang"`
	Status       string                                               `position:"Query" name:"Status"`
}

func (req *SaveBatchTaskForUpdateProhibitionLockRequest) Invoke(client *sdk.Client) (resp *SaveBatchTaskForUpdateProhibitionLockResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveBatchTaskForUpdateProhibitionLock", "domain", "")
	resp = &SaveBatchTaskForUpdateProhibitionLockResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveBatchTaskForUpdateProhibitionLockResponse struct {
	responses.BaseResponse
	RequestId string
	TaskNo    string
}

type SaveBatchTaskForUpdateProhibitionLockDomainNameList []string

func (list *SaveBatchTaskForUpdateProhibitionLockDomainNameList) UnmarshalJSON(data []byte) error {
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
