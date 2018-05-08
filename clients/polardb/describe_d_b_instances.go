package polardb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDBInstancesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InstanceNetworkType  string `position:"Query" name:"InstanceNetworkType"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeDBInstancesRequest) Invoke(client *sdk.Client) (resp *DescribeDBInstancesResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "DescribeDBInstances", "polardb", "")
	resp = &DescribeDBInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDBInstancesResponse struct {
	responses.BaseResponse
	RequestId        common.String
	PageNumber       common.Integer
	TotalRecordCount common.Integer
	PageRecordCount  common.Integer
	Items            DescribeDBInstancesDBInstanceList
}

type DescribeDBInstancesDBInstance struct {
	DBInstanceId          common.String
	DBInstanceDescription common.String
	PayType               common.String
	DBInstanceType        common.String
	DBInstanceClass       common.String
	InstanceNetworkType   common.String
	RegionId              common.String
	ZoneId                common.String
	DBClusterId           common.String
	ExpireTime            common.String
	DBInstanceStatus      common.String
	Engine                common.String
	DBType                common.String
	DBVersion             common.String
	DBInstanceType1       common.String
	LockMode              common.String
	LockReason            common.String
	CreateTime            common.String
	VpcId                 common.String
	VSwitchId             common.String
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
