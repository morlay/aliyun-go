package domain_intl

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ResendEmailVerificationRequest struct {
	requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	Email        string `position:"Query" name:"Email"`
}

func (req *ResendEmailVerificationRequest) Invoke(client *sdk.Client) (resp *ResendEmailVerificationResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "ResendEmailVerification", "domain", "")
	resp = &ResendEmailVerificationResponse{}
	err = client.DoAction(req, resp)
	return
}

type ResendEmailVerificationResponse struct {
	responses.BaseResponse
	RequestId   common.String
	SuccessList ResendEmailVerificationSendResultList
	FailList    ResendEmailVerificationSendResultList
}

type ResendEmailVerificationSendResult struct {
	Email   common.String
	Code    common.String
	Message common.String
}

type ResendEmailVerificationSendResultList []ResendEmailVerificationSendResult

func (list *ResendEmailVerificationSendResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ResendEmailVerificationSendResult)
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
