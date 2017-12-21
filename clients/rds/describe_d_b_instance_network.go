package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDBInstanceNetworkRequest struct {
	requests.RpcRequest
	EndTime      string `position:"Query" name:"EndTime"`
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
	StartTime    string `position:"Query" name:"StartTime"`
}

func (req *DescribeDBInstanceNetworkRequest) Invoke(client *sdk.Client) (resp *DescribeDBInstanceNetworkResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceNetwork", "rds", "")
	resp = &DescribeDBInstanceNetworkResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDBInstanceNetworkResponse struct {
	responses.BaseResponse
	RequestId    string
	DBInstanceId string
	StartTime    string
	EndTime      string
	Topology     DescribeDBInstanceNetworkTopologyItemList
}

type DescribeDBInstanceNetworkTopologyItem struct {
	StartPoint        string
	EndPoint          string
	NetworkTrafficIn  string
	NetworkTrafficOut string
	NetworkLatency    string
	BackendLatency    string
	NetworkErrors     string
}

type DescribeDBInstanceNetworkTopologyItemList []DescribeDBInstanceNetworkTopologyItem

func (list *DescribeDBInstanceNetworkTopologyItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceNetworkTopologyItem)
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
