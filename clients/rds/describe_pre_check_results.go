package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribePreCheckResultsRequest struct {
	requests.RpcRequest
	PreCheckId           string `position:"Query" name:"PreCheckId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribePreCheckResultsRequest) Invoke(client *sdk.Client) (resp *DescribePreCheckResultsResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribePreCheckResults", "rds", "")
	resp = &DescribePreCheckResultsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribePreCheckResultsResponse struct {
	responses.BaseResponse
	RequestId    string
	DBInstanceId string
	Items        DescribePreCheckResultsPreCheckResultList
}

type DescribePreCheckResultsPreCheckResult struct {
	PreCheckName   string
	PreCheckResult string
	FailReasion    string
	RepairMethod   string
}

type DescribePreCheckResultsPreCheckResultList []DescribePreCheckResultsPreCheckResult

func (list *DescribePreCheckResultsPreCheckResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePreCheckResultsPreCheckResult)
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
