package domain_intl

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type TransferInCheckMailTokenRequest struct {
	requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	Token        string `position:"Query" name:"Token"`
}

func (req *TransferInCheckMailTokenRequest) Invoke(client *sdk.Client) (resp *TransferInCheckMailTokenResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "TransferInCheckMailToken", "domain", "")
	resp = &TransferInCheckMailTokenResponse{}
	err = client.DoAction(req, resp)
	return
}

type TransferInCheckMailTokenResponse struct {
	responses.BaseResponse
	RequestId   string
	SuccessList TransferInCheckMailTokenSuccessListList
	FailList    TransferInCheckMailTokenFailListList
}

type TransferInCheckMailTokenSuccessListList []string

func (list *TransferInCheckMailTokenSuccessListList) UnmarshalJSON(data []byte) error {
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

type TransferInCheckMailTokenFailListList []string

func (list *TransferInCheckMailTokenFailListList) UnmarshalJSON(data []byte) error {
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
