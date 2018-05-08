package polardb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDBInstancePerformanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Key                  string `position:"Query" name:"Key"`
}

func (req *DescribeDBInstancePerformanceRequest) Invoke(client *sdk.Client) (resp *DescribeDBInstancePerformanceResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "DescribeDBInstancePerformance", "polardb", "")
	resp = &DescribeDBInstancePerformanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDBInstancePerformanceResponse struct {
	responses.BaseResponse
	RequestId       common.String
	DBInstanceId    common.String
	Engine          common.String
	DBType          common.String
	DBVersion       common.String
	StartTime       common.String
	EndTime         common.String
	PerformanceKeys DescribeDBInstancePerformancePerformanceItemList
}

type DescribeDBInstancePerformancePerformanceItem struct {
	MetricName  common.String
	Measurement common.String
	Points      DescribeDBInstancePerformancePerformanceItemValueList
}

type DescribeDBInstancePerformancePerformanceItemValue struct {
	Value     common.Float
	Timestamp common.Long
}

type DescribeDBInstancePerformancePerformanceItemList []DescribeDBInstancePerformancePerformanceItem

func (list *DescribeDBInstancePerformancePerformanceItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancePerformancePerformanceItem)
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

type DescribeDBInstancePerformancePerformanceItemValueList []DescribeDBInstancePerformancePerformanceItemValue

func (list *DescribeDBInstancePerformancePerformanceItemValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancePerformancePerformanceItemValue)
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
