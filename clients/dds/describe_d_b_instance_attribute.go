package dds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId   string
	DBInstances DescribeDBInstanceAttributeDBInstanceList
}

type DescribeDBInstanceAttributeDBInstance struct {
	DBInstanceId          string
	DBInstanceDescription string
	RegionId              string
	ZoneId                string
	Engine                string
	EngineVersion         string
	StorageEngine         string
	DBInstanceClass       string
	DBInstanceStorage     int
	DBInstanceStatus      string
	LockMode              string
	ChargeType            string
	CreationTime          string
	ReplicaSetName        string
	NetworkType           string
	ExpireTime            string
	MaintainStartTime     string
	MaintainEndTime       string
	DBInstanceType        string
	LastDowngradeTime     int
	MongosList            DescribeDBInstanceAttributeMongosAttributeList
	ShardList             DescribeDBInstanceAttributeShardAttributeList
}

type DescribeDBInstanceAttributeMongosAttribute struct {
	NodeId          string
	NodeDescription string
	NodeClass       string
	ConnectSting    string
	Port            int
}

type DescribeDBInstanceAttributeShardAttribute struct {
	NodeId          string
	NodeDescription string
	NodeClass       string
	NodeStorage     int
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
