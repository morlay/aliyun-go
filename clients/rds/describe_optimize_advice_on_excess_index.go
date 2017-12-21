package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeOptimizeAdviceOnExcessIndexRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeOptimizeAdviceOnExcessIndexRequest) Invoke(client *sdk.Client) (resp *DescribeOptimizeAdviceOnExcessIndexResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeOptimizeAdviceOnExcessIndex", "rds", "")
	resp = &DescribeOptimizeAdviceOnExcessIndexResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeOptimizeAdviceOnExcessIndexResponse struct {
	responses.BaseResponse
	RequestId         string
	TotalRecordsCount int
	PageNumber        int
	PageRecordCount   int
	Items             DescribeOptimizeAdviceOnExcessIndexAdviceOnExcessIndexList
}

type DescribeOptimizeAdviceOnExcessIndexAdviceOnExcessIndex struct {
	DBName     string
	TableName  string
	IndexCount int64
}

type DescribeOptimizeAdviceOnExcessIndexAdviceOnExcessIndexList []DescribeOptimizeAdviceOnExcessIndexAdviceOnExcessIndex

func (list *DescribeOptimizeAdviceOnExcessIndexAdviceOnExcessIndexList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeOptimizeAdviceOnExcessIndexAdviceOnExcessIndex)
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
