package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SubmitEmailVerificationRequest struct {
	requests.RpcRequest
	SendIfExist  string `position:"Query" name:"SendIfExist"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	Email        string `position:"Query" name:"Email"`
}

func (req *SubmitEmailVerificationRequest) Invoke(client *sdk.Client) (resp *SubmitEmailVerificationResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "SubmitEmailVerification", "", "")
	resp = &SubmitEmailVerificationResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitEmailVerificationResponse struct {
	responses.BaseResponse
	RequestId   common.String
	SuccessList SubmitEmailVerificationSendResultList
	FailList    SubmitEmailVerificationSendResultList
	ExistList   SubmitEmailVerificationSendResultList
}

type SubmitEmailVerificationSendResult struct {
	Email   common.String
	Code    common.String
	Message common.String
}

type SubmitEmailVerificationSendResultList []SubmitEmailVerificationSendResult

func (list *SubmitEmailVerificationSendResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitEmailVerificationSendResult)
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
