package domain_intl

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveBatchTaskForTransferProhibitionLockRequest struct {
	requests.RpcRequest
	UserClientIp string                                                 `position:"Query" name:"UserClientIp"`
	DomainNames  *SaveBatchTaskForTransferProhibitionLockDomainNameList `position:"Query" type:"Repeated" name:"DomainName"`
	Lang         string                                                 `position:"Query" name:"Lang"`
	Status       string                                                 `position:"Query" name:"Status"`
}

func (req *SaveBatchTaskForTransferProhibitionLockRequest) Invoke(client *sdk.Client) (resp *SaveBatchTaskForTransferProhibitionLockResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveBatchTaskForTransferProhibitionLock", "domain", "")
	resp = &SaveBatchTaskForTransferProhibitionLockResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveBatchTaskForTransferProhibitionLockResponse struct {
	responses.BaseResponse
	RequestId string
	TaskNo    string
}

type SaveBatchTaskForTransferProhibitionLockDomainNameList []string

func (list *SaveBatchTaskForTransferProhibitionLockDomainNameList) UnmarshalJSON(data []byte) error {
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
