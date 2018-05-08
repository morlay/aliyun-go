package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDampPoliciesByCidRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDampPoliciesByCidRequest) Invoke(client *sdk.Client) (resp *DescribeDampPoliciesByCidResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeDampPoliciesByCid", "rds", "")
	resp = &DescribeDampPoliciesByCidResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDampPoliciesByCidResponse struct {
	responses.BaseResponse
	RequestId common.String
	Policies  DescribeDampPoliciesByCidPolicyList
}

type DescribeDampPoliciesByCidPolicy struct {
	PolicyName common.String
	Comment    common.String
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
