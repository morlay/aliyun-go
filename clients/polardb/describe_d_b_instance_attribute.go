package polardb

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
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDBInstanceAttributeRequest) Invoke(client *sdk.Client) (resp *DescribeDBInstanceAttributeResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "DescribeDBInstanceAttribute", "polardb", "")
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
	DBInstanceId          common.String
	DBClusterDescription  common.String
	DBClusterId           common.String
	PayType               common.String
	DBInstanceType        common.String
	RegionId              common.String
	ZoneId                common.String
	Engine                common.String
	DBType                common.String
	DBVersion             common.String
	DBInstanceClass       common.String
	DBInstanceStorage     common.Long
	DBInstanceStatus      common.String
	DBInstanceDescription common.String
	ConnectionString      common.Long
	Port                  common.Long
	DBInstanceNetType     common.String
	LockMode              common.String
	LockReason            common.String
	CreationTime          common.String
	ExpireTime            common.String
	MaintainStartTime     common.String
	MaintainEndTime       common.String
	MaxConnections        common.Integer
	MaxIOPS               common.Integer
	SecurityIPList        common.String
	InstanceNetworkType   common.String
	VpcId                 common.String
	VSwitchId             common.String
	DBInstanceType1       common.String
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
