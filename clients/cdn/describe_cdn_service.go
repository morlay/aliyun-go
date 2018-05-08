package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCdnServiceRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeCdnServiceRequest) Invoke(client *sdk.Client) (resp *DescribeCdnServiceResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeCdnService", "", "")
	resp = &DescribeCdnServiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCdnServiceResponse struct {
	responses.BaseResponse
	RequestId          common.String
	InstanceId         common.String
	InternetChargeType common.String
	OpeningTime        common.String
	ChangingChargeType common.String
	ChangingAffectTime common.String
	OperationLocks     DescribeCdnServiceLockReasonList
}

type DescribeCdnServiceLockReason struct {
	LockReason common.String
}

type DescribeCdnServiceLockReasonList []DescribeCdnServiceLockReason

func (list *DescribeCdnServiceLockReasonList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCdnServiceLockReason)
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
