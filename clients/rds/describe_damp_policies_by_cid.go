package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDampPoliciesByCidRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeDampPoliciesByCidRequest) Invoke(client *sdk.Client) (response *DescribeDampPoliciesByCidResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDampPoliciesByCidRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeDampPoliciesByCid", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDampPoliciesByCidResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDampPoliciesByCidResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDampPoliciesByCidResponse struct {
	RequestId string
	Policies  DescribeDampPoliciesByCidPolicyList
}

type DescribeDampPoliciesByCidPolicy struct {
	PolicyName string
	Comment    string
}

type DescribeDampPoliciesByCidPolicyList []DescribeDampPoliciesByCidPolicy

func (list *DescribeDampPoliciesByCidPolicyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDampPoliciesByCidPolicy)
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
