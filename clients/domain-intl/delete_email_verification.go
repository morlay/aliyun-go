package domain_intl

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteEmailVerificationRequest struct {
	requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	Email        string `position:"Query" name:"Email"`
}

func (req *DeleteEmailVerificationRequest) Invoke(client *sdk.Client) (resp *DeleteEmailVerificationResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "DeleteEmailVerification", "domain", "")
	resp = &DeleteEmailVerificationResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteEmailVerificationResponse struct {
	responses.BaseResponse
	RequestId   common.String
	SuccessList DeleteEmailVerificationSendResultList
	FailList    DeleteEmailVerificationSendResultList
}

type DeleteEmailVerificationSendResult struct {
	Email   common.String
	Code    common.String
	Message common.String
}

type DeleteEmailVerificationSendResultList []DeleteEmailVerificationSendResult

func (list *DeleteEmailVerificationSendResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DeleteEmailVerificationSendResult)
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
