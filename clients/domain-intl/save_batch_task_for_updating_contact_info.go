package domain_intl

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveBatchTaskForUpdatingContactInfoRequest struct {
	requests.RpcRequest
	ContactType         string                                             `position:"Query" name:"ContactType"`
	UserClientIp        string                                             `position:"Query" name:"UserClientIp"`
	RegistrantProfileId int64                                              `position:"Query" name:"RegistrantProfileId"`
	DomainNames         *SaveBatchTaskForUpdatingContactInfoDomainNameList `position:"Query" type:"Repeated" name:"DomainName"`
	AddTransferLock     string                                             `position:"Query" name:"AddTransferLock"`
	Lang                string                                             `position:"Query" name:"Lang"`
}

func (req *SaveBatchTaskForUpdatingContactInfoRequest) Invoke(client *sdk.Client) (resp *SaveBatchTaskForUpdatingContactInfoResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveBatchTaskForUpdatingContactInfo", "domain", "")
	resp = &SaveBatchTaskForUpdatingContactInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveBatchTaskForUpdatingContactInfoResponse struct {
	responses.BaseResponse
	RequestId string
	TaskNo    string
}

type SaveBatchTaskForUpdatingContactInfoDomainNameList []string

func (list *SaveBatchTaskForUpdatingContactInfoDomainNameList) UnmarshalJSON(data []byte) error {
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
