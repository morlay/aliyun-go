package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryFailReasonForDomainRealNameVerificationRequest struct {
	requests.RpcRequest
	RealNameVerificationAction string `position:"Query" name:"RealNameVerificationAction"`
	UserClientIp               string `position:"Query" name:"UserClientIp"`
	DomainName                 string `position:"Query" name:"DomainName"`
	Lang                       string `position:"Query" name:"Lang"`
}

func (req *QueryFailReasonForDomainRealNameVerificationRequest) Invoke(client *sdk.Client) (resp *QueryFailReasonForDomainRealNameVerificationResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "QueryFailReasonForDomainRealNameVerification", "", "")
	resp = &QueryFailReasonForDomainRealNameVerificationResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryFailReasonForDomainRealNameVerificationResponse struct {
	responses.BaseResponse
	RequestId string
	Data      QueryFailReasonForDomainRealNameVerificationFailRecordList
}

type QueryFailReasonForDomainRealNameVerificationFailRecord struct {
	Date       string
	FailReason string
}

type QueryFailReasonForDomainRealNameVerificationFailRecordList []QueryFailReasonForDomainRealNameVerificationFailRecord

func (list *QueryFailReasonForDomainRealNameVerificationFailRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryFailReasonForDomainRealNameVerificationFailRecord)
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
