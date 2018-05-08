package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeOptimizeAdviceOnMissIndexRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeOptimizeAdviceOnMissIndexRequest) Invoke(client *sdk.Client) (resp *DescribeOptimizeAdviceOnMissIndexResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeOptimizeAdviceOnMissIndex", "rds", "")
	resp = &DescribeOptimizeAdviceOnMissIndexResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeOptimizeAdviceOnMissIndexResponse struct {
	responses.BaseResponse
	RequestId         common.String
	DBInstanceId      common.String
	TotalRecordsCount common.Integer
	PageNumber        common.Integer
	PageRecordCount   common.Integer
	Items             DescribeOptimizeAdviceOnMissIndexAdviceOnMissIndexList
}

type DescribeOptimizeAdviceOnMissIndexAdviceOnMissIndex struct {
	DBName      common.String
	TableName   common.String
	QueryColumn common.String
	SQLText     common.String
}

type DescribeOptimizeAdviceOnMissIndexAdviceOnMissIndexList []DescribeOptimizeAdviceOnMissIndexAdviceOnMissIndex

func (list *DescribeOptimizeAdviceOnMissIndexAdviceOnMissIndexList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeOptimizeAdviceOnMissIndexAdviceOnMissIndex)
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
