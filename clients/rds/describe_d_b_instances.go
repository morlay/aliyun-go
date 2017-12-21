package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDBInstancesRequest struct {
	requests.RpcRequest
	ConnectionMode       string `position:"Query" name:"ConnectionMode"`
	Tag4value            string `position:"Query" name:"Tag.4.value"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Tag2key              string `position:"Query" name:"Tag.2.key"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	SearchKey            string `position:"Query" name:"SearchKey"`
	Tag3key              string `position:"Query" name:"Tag.3.key"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	Tag1value            string `position:"Query" name:"Tag.1.value"`
	Engine               string `position:"Query" name:"Engine"`
	PageSize             int    `position:"Query" name:"PageSize"`
	DBInstanceStatus     string `position:"Query" name:"DBInstanceStatus"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	Tag3value            string `position:"Query" name:"Tag.3.value"`
	ProxyId              string `position:"Query" name:"ProxyId"`
	Tag5key              string `position:"Query" name:"Tag.5.key"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tag5value            string `position:"Query" name:"Tag.5.value"`
	DBInstanceType       string `position:"Query" name:"DBInstanceType"`
	Tags                 string `position:"Query" name:"Tags"`
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	Tag1key              string `position:"Query" name:"Tag.1.key"`
	VpcId                string `position:"Query" name:"VpcId"`
	Tag2value            string `position:"Query" name:"Tag.2.value"`
	Tag4key              string `position:"Query" name:"Tag.4.key"`
	InstanceNetworkType  string `position:"Query" name:"InstanceNetworkType"`
}

func (req *DescribeDBInstancesRequest) Invoke(client *sdk.Client) (resp *DescribeDBInstancesResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstances", "rds", "")
	resp = &DescribeDBInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDBInstancesResponse struct {
	responses.BaseResponse
	RequestId        string
	PageNumber       int
	TotalRecordCount int
	PageRecordCount  int
	Items            DescribeDBInstancesDBInstanceList
}

type DescribeDBInstancesDBInstance struct {
	InsId                 int
	DBInstanceId          string
	DBInstanceDescription string
	PayType               string
	DBInstanceType        string
	RegionId              string
	ExpireTime            string
	DBInstanceStatus      string
	Engine                string
	DBInstanceNetType     string
	ConnectionMode        string
	LockMode              string
	DBInstanceClass       string
	InstanceNetworkType   string
	LockReason            string
	ZoneId                string
	MutriORsignle         bool
	CreateTime            string
	EngineVersion         string
	GuardDBInstanceId     string
	TempDBInstanceId      string
	MasterInstanceId      string
	VpcId                 string
	VSwitchId             string
	ResourceGroupId       string
	ReadOnlyDBInstanceIds DescribeDBInstancesReadOnlyDBInstanceIdList
}

type DescribeDBInstancesReadOnlyDBInstanceId struct {
	DBInstanceId string
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

type DescribeDBInstancesReadOnlyDBInstanceIdList []DescribeDBInstancesReadOnlyDBInstanceId

func (list *DescribeDBInstancesReadOnlyDBInstanceIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancesReadOnlyDBInstanceId)
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
