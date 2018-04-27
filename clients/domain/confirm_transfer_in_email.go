package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ConfirmTransferInEmailRequest struct {
	requests.RpcRequest
	UserClientIp string                                `position:"Query" name:"UserClientIp"`
	DomainNames  *ConfirmTransferInEmailDomainNameList `position:"Query" type:"Repeated" name:"DomainName"`
	Lang         string                                `position:"Query" name:"Lang"`
	Email        string                                `position:"Query" name:"Email"`
}

func (req *ConfirmTransferInEmailRequest) Invoke(client *sdk.Client) (resp *ConfirmTransferInEmailResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "ConfirmTransferInEmail", "", "")
	resp = &ConfirmTransferInEmailResponse{}
	err = client.DoAction(req, resp)
	return
}

type ConfirmTransferInEmailResponse struct {
	responses.BaseResponse
	RequestId   string
	SuccessList ConfirmTransferInEmailSuccessListList
	FailList    ConfirmTransferInEmailFailListList
}

type ConfirmTransferInEmailDomainNameList []string

func (list *ConfirmTransferInEmailDomainNameList) UnmarshalJSON(data []byte) error {
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

type ConfirmTransferInEmailSuccessListList []string

func (list *ConfirmTransferInEmailSuccessListList) UnmarshalJSON(data []byte) error {
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

type ConfirmTransferInEmailFailListList []string

func (list *ConfirmTransferInEmailFailListList) UnmarshalJSON(data []byte) error {
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
