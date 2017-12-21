package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDBInstanceNetworkDetailRequest struct {
	requests.RpcRequest
	EndPoint     string `position:"Query" name:"EndPoint"`
	StartPoint   string `position:"Query" name:"StartPoint"`
	EndTime      string `position:"Query" name:"EndTime"`
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
	StartTime    string `position:"Query" name:"StartTime"`
}

func (req *DescribeDBInstanceNetworkDetailRequest) Invoke(client *sdk.Client) (resp *DescribeDBInstanceNetworkDetailResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceNetworkDetail", "rds", "")
	resp = &DescribeDBInstanceNetworkDetailResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDBInstanceNetworkDetailResponse struct {
	responses.BaseResponse
	RequestId         string
	DBInstanceId      string
	StartTime         string
	EndTime           string
	NewConnection     string
	ActiveConnection  string
	AbortedConnection string
	NetworkRequest    string
	NetworkTrafficIn  string
	NetworkTrafficOut string
	NetworkLatency    string
	BackendLatency    string
	NetworkErrors     string
	NetworkDetail     DescribeDBInstanceNetworkDetailNetworkKeyList
}

type DescribeDBInstanceNetworkDetailNetworkKey struct {
	Key    string
	Unit   string
	Values DescribeDBInstanceNetworkDetailNetworkValueList
}

type DescribeDBInstanceNetworkDetailNetworkValue struct {
	Value    string
	DateTime string
}

type DescribeDBInstanceNetworkDetailNetworkKeyList []DescribeDBInstanceNetworkDetailNetworkKey

func (list *DescribeDBInstanceNetworkDetailNetworkKeyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceNetworkDetailNetworkKey)
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

type DescribeDBInstanceNetworkDetailNetworkValueList []DescribeDBInstanceNetworkDetailNetworkValue

func (list *DescribeDBInstanceNetworkDetailNetworkValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceNetworkDetailNetworkValue)
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
