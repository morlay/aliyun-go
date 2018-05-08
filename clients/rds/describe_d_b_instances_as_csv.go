package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDBInstancesAsCsvRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDBInstancesAsCsvRequest) Invoke(client *sdk.Client) (resp *DescribeDBInstancesAsCsvResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstancesAsCsv", "rds", "")
	resp = &DescribeDBInstancesAsCsvResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDBInstancesAsCsvResponse struct {
	responses.BaseResponse
	RequestId common.String
	Items     DescribeDBInstancesAsCsvDBInstanceAttributeList
}

type DescribeDBInstancesAsCsvDBInstanceAttribute struct {
	DBInstanceId                common.String
	PayType                     common.String
	DBInstanceClassType         common.String
	DBInstanceType              common.String
	RegionId                    common.String
	ConnectionString            common.String
	Port                        common.String
	Engine                      common.String
	EngineVersion               common.String
	DBInstanceClass             common.String
	DBInstanceMemory            common.Long
	DBInstanceStorage           common.Integer
	DBInstanceNetType           common.String
	DBInstanceStatus            common.String
	DBInstanceDescription       common.String
	LockMode                    common.String
	LockReason                  common.String
	ReadDelayTime               common.String
	DBMaxQuantity               common.Integer
	AccountMaxQuantity          common.Integer
	CreationTime                common.String
	ExpireTime                  common.String
	MaintainTime                common.String
	AvailabilityValue           common.String
	MaxIOPS                     common.Integer
	MaxConnections              common.Integer
	MasterInstanceId            common.String
	DBInstanceCPU               common.String
	IncrementSourceDBInstanceId common.String
	GuardDBInstanceId           common.String
	TempDBInstanceId            common.String
	SecurityIPList              common.String
	ZoneId                      common.String
	InstanceNetworkType         common.String
	Category                    common.String
	AccountType                 common.String
	SupportUpgradeAccountType   common.String
	VpcId                       common.String
	VSwitchId                   common.String
	ConnectionMode              common.String
	Tags                        common.String
}

type DescribeDBInstancesAsCsvDBInstanceAttributeList []DescribeDBInstancesAsCsvDBInstanceAttribute

func (list *DescribeDBInstancesAsCsvDBInstanceAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancesAsCsvDBInstanceAttribute)
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
