package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCdnServiceRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeCdnServiceRequest) Invoke(client *sdk.Client) (response *DescribeCdnServiceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeCdnServiceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeCdnService", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeCdnServiceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeCdnServiceResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeCdnServiceResponse struct {
	RequestId          string
	InstanceId         string
	InternetChargeType string
	OpeningTime        string
	ChangingChargeType string
	ChangingAffectTime string
	OperationLocks     DescribeCdnServiceLockReasonList
}

type DescribeCdnServiceLockReason struct {
	LockReason string
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
