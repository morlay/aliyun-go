package dds

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
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ReplicaSetRole       string `position:"Query" name:"ReplicaSetRole"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	NodeId               string `position:"Query" name:"NodeId"`
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
	req.InitWithApiInfo("Dds", "2015-12-01", "DescribeDBInstancePerformance", "dds", "")

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
	StartTime       string
	EndTime         string
	PerformanceKeys DescribeDBInstancePerformancePerformanceKeyList
}

type DescribeDBInstancePerformancePerformanceKey struct {
	Key               string
	Unit              string
	ValueFormat       string
	PerformanceValues DescribeDBInstancePerformancePerformanceValueList
}

type DescribeDBInstancePerformancePerformanceValue struct {
	Value string
	Date  string
}

type DescribeDBInstancePerformancePerformanceKeyList []DescribeDBInstancePerformancePerformanceKey

func (list *DescribeDBInstancePerformancePerformanceKeyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancePerformancePerformanceKey)
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

type DescribeDBInstancePerformancePerformanceValueList []DescribeDBInstancePerformancePerformanceValue

func (list *DescribeDBInstancePerformancePerformanceValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancePerformancePerformanceValue)
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
