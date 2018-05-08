package dds

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
	DBInstanceIds        string `position:"Query" name:"DBInstanceIds"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	DBInstanceType       string `position:"Query" name:"DBInstanceType"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	Engine               string `position:"Query" name:"Engine"`
	PageSize             int    `position:"Query" name:"PageSize"`
}

func (req *DescribeDBInstancesRequest) Invoke(client *sdk.Client) (resp *DescribeDBInstancesResponse, err error) {
	req.InitWithApiInfo("Dds", "2015-12-01", "DescribeDBInstances", "dds", "")
	resp = &DescribeDBInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDBInstancesResponse struct {
	responses.BaseResponse
	RequestId   common.String
	PageNumber  common.Integer
	PageSize    common.Integer
	TotalCount  common.Integer
	DBInstances DescribeDBInstancesDBInstanceList
}

type DescribeDBInstancesDBInstance struct {
	DBInstanceId          common.String
	DBInstanceDescription common.String
	RegionId              common.String
	ZoneId                common.String
	Engine                common.String
	EngineVersion         common.String
	DBInstanceClass       common.String
	DBInstanceStorage     common.Integer
	DBInstanceStatus      common.String
	LockMode              common.String
	ChargeType            common.String
	NetworkType           common.String
	CreationTime          common.String
	ExpireTime            common.String
	DBInstanceType        common.String
	LastDowngradeTime     common.Integer
	MongosList            DescribeDBInstancesMongosAttributeList
	ShardList             DescribeDBInstancesShardAttributeList
}

type DescribeDBInstancesMongosAttribute struct {
	NodeId          common.String
	NodeDescription common.String
	NodeClass       common.String
	ConnectSting    common.String
	Port            common.Integer
}

type DescribeDBInstancesShardAttribute struct {
	NodeId          common.String
	NodeDescription common.String
	NodeClass       common.String
	NodeStorage     common.Integer
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

type DescribeDBInstancesMongosAttributeList []DescribeDBInstancesMongosAttribute

func (list *DescribeDBInstancesMongosAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancesMongosAttribute)
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

type DescribeDBInstancesShardAttributeList []DescribeDBInstancesShardAttribute

func (list *DescribeDBInstancesShardAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancesShardAttribute)
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
