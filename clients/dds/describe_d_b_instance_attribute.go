package dds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDBInstanceAttributeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Engine               string `position:"Query" name:"Engine"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDBInstanceAttributeRequest) Invoke(client *sdk.Client) (resp *DescribeDBInstanceAttributeResponse, err error) {
	req.InitWithApiInfo("Dds", "2015-12-01", "DescribeDBInstanceAttribute", "dds", "")
	resp = &DescribeDBInstanceAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDBInstanceAttributeResponse struct {
	responses.BaseResponse
	RequestId   common.String
	DBInstances DescribeDBInstanceAttributeDBInstanceList
}

type DescribeDBInstanceAttributeDBInstance struct {
	DBInstanceId          common.String
	DBInstanceDescription common.String
	RegionId              common.String
	ZoneId                common.String
	Engine                common.String
	EngineVersion         common.String
	StorageEngine         common.String
	DBInstanceClass       common.String
	DBInstanceStorage     common.Integer
	DBInstanceStatus      common.String
	LockMode              common.String
	ChargeType            common.String
	CreationTime          common.String
	ReplicaSetName        common.String
	NetworkType           common.String
	ExpireTime            common.String
	MaintainStartTime     common.String
	MaintainEndTime       common.String
	DBInstanceType        common.String
	LastDowngradeTime     common.Integer
	MongosList            DescribeDBInstanceAttributeMongosAttributeList
	ShardList             DescribeDBInstanceAttributeShardAttributeList
}

type DescribeDBInstanceAttributeMongosAttribute struct {
	NodeId          common.String
	NodeDescription common.String
	NodeClass       common.String
	ConnectSting    common.String
	Port            common.Integer
}

type DescribeDBInstanceAttributeShardAttribute struct {
	NodeId          common.String
	NodeDescription common.String
	NodeClass       common.String
	NodeStorage     common.Integer
}

type DescribeDBInstanceAttributeDBInstanceList []DescribeDBInstanceAttributeDBInstance

func (list *DescribeDBInstanceAttributeDBInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceAttributeDBInstance)
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

type DescribeDBInstanceAttributeMongosAttributeList []DescribeDBInstanceAttributeMongosAttribute

func (list *DescribeDBInstanceAttributeMongosAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceAttributeMongosAttribute)
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

type DescribeDBInstanceAttributeShardAttributeList []DescribeDBInstanceAttributeShardAttribute

func (list *DescribeDBInstanceAttributeShardAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceAttributeShardAttribute)
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
