package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDBInstanceAttributeRequest struct {
	requests.RpcRequest
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDBInstanceAttributeRequest) Invoke(client *sdk.Client) (resp *DescribeDBInstanceAttributeResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceAttribute", "rds", "")
	resp = &DescribeDBInstanceAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDBInstanceAttributeResponse struct {
	responses.BaseResponse
	RequestId string
	Items     DescribeDBInstanceAttributeDBInstanceAttributeList
}

type DescribeDBInstanceAttributeDBInstanceAttribute struct {
	InsId                       int
	DBInstanceId                string
	PayType                     string
	DBInstanceClassType         string
	DBInstanceType              string
	RegionId                    string
	ConnectionString            string
	Port                        string
	Engine                      string
	EngineVersion               string
	DBInstanceClass             string
	DBInstanceMemory            int64
	DBInstanceStorage           int
	DBInstanceNetType           string
	DBInstanceStatus            string
	DBInstanceDescription       string
	LockMode                    string
	LockReason                  string
	ReadDelayTime               string
	DBMaxQuantity               int
	AccountMaxQuantity          int
	CreationTime                string
	ExpireTime                  string
	MaintainTime                string
	AvailabilityValue           string
	MaxIOPS                     int
	MaxConnections              int
	MasterInstanceId            string
	DBInstanceCPU               string
	IncrementSourceDBInstanceId string
	GuardDBInstanceId           string
	TempDBInstanceId            string
	SecurityIPList              string
	ZoneId                      string
	InstanceNetworkType         string
	Category                    string
	AccountType                 string
	SupportUpgradeAccountType   string
	VpcId                       string
	VSwitchId                   string
	ConnectionMode              string
	ResourceGroupId             string
	ReadOnlyDBInstanceIds       DescribeDBInstanceAttributeReadOnlyDBInstanceIdList
}

type DescribeDBInstanceAttributeReadOnlyDBInstanceId struct {
	DBInstanceId string
}

type DescribeDBInstanceAttributeDBInstanceAttributeList []DescribeDBInstanceAttributeDBInstanceAttribute

func (list *DescribeDBInstanceAttributeDBInstanceAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceAttributeDBInstanceAttribute)
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

type DescribeDBInstanceAttributeReadOnlyDBInstanceIdList []DescribeDBInstanceAttributeReadOnlyDBInstanceId

func (list *DescribeDBInstanceAttributeReadOnlyDBInstanceIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceAttributeReadOnlyDBInstanceId)
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
