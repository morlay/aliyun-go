package polardb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDBInstancesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InstanceNetworkType  string `position:"Query" name:"InstanceNetworkType"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (r DescribeDBInstancesRequest) Invoke(client *sdk.Client) (response *DescribeDBInstancesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDBInstancesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("polardb", "2017-08-01", "DescribeDBInstances", "polardb", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDBInstancesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDBInstancesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDBInstancesResponse struct {
	RequestId        string
	PageNumber       int
	TotalRecordCount int
	PageRecordCount  int
	Items            DescribeDBInstancesDBInstanceList
}

type DescribeDBInstancesDBInstance struct {
	DBInstanceId          string
	DBInstanceDescription string
	PayType               string
	DBInstanceType        string
	DBInstanceClass       string
	InstanceNetworkType   string
	RegionId              string
	ZoneId                string
	DBClusterId           string
	ExpireTime            string
	DBInstanceStatus      string
	Engine                string
	DBType                string
	DBVersion             string
	DBInstanceType1       string
	LockMode              string
	LockReason            string
	CreateTime            string
	VpcId                 string
	VSwitchId             string
}

type DescribeDBInstancesDBInstanceList []DescribeDBInstancesDBInstance

func (list *DescribeDBInstancesDBInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancesDBInstance)
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
