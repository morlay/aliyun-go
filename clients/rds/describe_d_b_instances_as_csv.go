package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId string
	Items     DescribeDBInstancesAsCsvDBInstanceAttributeList
}

type DescribeDBInstancesAsCsvDBInstanceAttribute struct {
	InsId                       int
	DBInstanceId                string
	DBInstanceName              string
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
	Tags                        string
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
