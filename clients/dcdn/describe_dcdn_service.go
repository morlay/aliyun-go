package dcdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDcdnServiceRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDcdnServiceRequest) Invoke(client *sdk.Client) (resp *DescribeDcdnServiceResponse, err error) {
	req.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnService", "dcdn", "")
	resp = &DescribeDcdnServiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDcdnServiceResponse struct {
	responses.BaseResponse
	RequestId          string
	InstanceId         string
	InternetChargeType string
	OpeningTime        string
	ChangingChargeType string
	ChangingAffectTime string
	OperationLocks     DescribeDcdnServiceLockReasonList
}

type DescribeDcdnServiceLockReason struct {
	LockReason string
}

type DescribeDcdnServiceLockReasonList []DescribeDcdnServiceLockReason

func (list *DescribeDcdnServiceLockReasonList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnServiceLockReason)
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
