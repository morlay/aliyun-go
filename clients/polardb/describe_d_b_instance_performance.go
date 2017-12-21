package polardb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDBInstancePerformanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Key                  string `position:"Query" name:"Key"`
}

func (r DescribeDBInstancePerformanceRequest) Invoke(client *sdk.Client) (response *DescribeDBInstancePerformanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDBInstancePerformanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("polardb", "2017-08-01", "DescribeDBInstancePerformance", "polardb", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDBInstancePerformanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDBInstancePerformanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDBInstancePerformanceResponse struct {
	RequestId       string
	DBInstanceId    string
	Engine          string
	DBType          string
	DBVersion       string
	StartTime       string
	EndTime         string
	PerformanceKeys DescribeDBInstancePerformancePerformanceItemList
}

type DescribeDBInstancePerformancePerformanceItem struct {
	MetricName  string
	Measurement string
	Points      DescribeDBInstancePerformancePerformanceItemValueList
}

type DescribeDBInstancePerformancePerformanceItemValue struct {
	Value     float32
	Timestamp int64
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
