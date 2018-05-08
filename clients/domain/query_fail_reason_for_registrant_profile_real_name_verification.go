package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryFailReasonForRegistrantProfileRealNameVerificationRequest struct {
	requests.RpcRequest
	UserClientIp        string `position:"Query" name:"UserClientIp"`
	RegistrantProfileID int64  `position:"Query" name:"RegistrantProfileID"`
	Lang                string `position:"Query" name:"Lang"`
}

func (req *QueryFailReasonForRegistrantProfileRealNameVerificationRequest) Invoke(client *sdk.Client) (resp *QueryFailReasonForRegistrantProfileRealNameVerificationResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "QueryFailReasonForRegistrantProfileRealNameVerification", "", "")
	resp = &QueryFailReasonForRegistrantProfileRealNameVerificationResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryFailReasonForRegistrantProfileRealNameVerificationResponse struct {
	responses.BaseResponse
	RequestId common.String
	Data      QueryFailReasonForRegistrantProfileRealNameVerificationFailRecordList
}

type QueryFailReasonForRegistrantProfileRealNameVerificationFailRecord struct {
	Date       common.String
	FailReason common.String
}

type QueryFailReasonForRegistrantProfileRealNameVerificationFailRecordList []QueryFailReasonForRegistrantProfileRealNameVerificationFailRecord

func (list *QueryFailReasonForRegistrantProfileRealNameVerificationFailRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryFailReasonForRegistrantProfileRealNameVerificationFailRecord)
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
