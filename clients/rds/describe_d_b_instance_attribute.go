package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Items     DescribeDBInstanceAttributeDBInstanceAttributeList
}

type DescribeDBInstanceAttributeDBInstanceAttribute struct {
	DBInstanceDiskUsed                common.String
	GuardDBInstanceName               common.String
	CanTempUpgrade                    bool
	TempUpgradeTimeStart              common.String
	TempUpgradeTimeEnd                common.String
	TempUpgradeRecoveryTime           common.String
	TempUpgradeRecoveryClass          common.String
	TempUpgradeRecoveryCpu            common.Integer
	TempUpgradeRecoveryMemory         common.Integer
	TempUpgradeRecoveryMaxIOPS        common.String
	TempUpgradeRecoveryMaxConnections common.String
	InsId                             common.Integer
	DBInstanceId                      common.String
	PayType                           common.String
	DBInstanceClassType               common.String
	DBInstanceType                    common.String
	RegionId                          common.String
	ConnectionString                  common.String
	Port                              common.String
	Engine                            common.String
	EngineVersion                     common.String
	DBInstanceClass                   common.String
	DBInstanceMemory                  common.Long
	DBInstanceStorage                 common.Integer
	VpcCloudInstanceId                common.String
	DBInstanceNetType                 common.String
	DBInstanceStatus                  common.String
	DBInstanceDescription             common.String
	LockMode                          common.String
	LockReason                        common.String
	ReadDelayTime                     common.String
	DBMaxQuantity                     common.Integer
	AccountMaxQuantity                common.Integer
	CreationTime                      common.String
	ExpireTime                        common.String
	MaintainTime                      common.String
	AvailabilityValue                 common.String
	MaxIOPS                           common.Integer
	MaxConnections                    common.Integer
	MasterInstanceId                  common.String
	DBInstanceCPU                     common.String
	IncrementSourceDBInstanceId       common.String
	GuardDBInstanceId                 common.String
	ReplicateId                       common.String
	TempDBInstanceId                  common.String
	SecurityIPList                    common.String
	ZoneId                            common.String
	InstanceNetworkType               common.String
	AdvancedFeatures                  common.String
	Category                          common.String
	AccountType                       common.String
	SupportUpgradeAccountType         common.String
	VpcId                             common.String
	VSwitchId                         common.String
	ConnectionMode                    common.String
	ResourceGroupId                   common.String
	ReadOnlyDBInstanceIds             DescribeDBInstanceAttributeReadOnlyDBInstanceIdList
}

type DescribeDBInstanceAttributeReadOnlyDBInstanceId struct {
	DBInstanceId common.String
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
