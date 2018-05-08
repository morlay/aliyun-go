package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDBInstanceNetInfoRequest struct {
	requests.RpcRequest
	ResourceOwnerId          int64  `position:"Query" name:"ResourceOwnerId"`
	Flag                     string `position:"Query" name:"Flag"`
	DBInstanceNetRWSplitType string `position:"Query" name:"DBInstanceNetRWSplitType"`
	ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken              string `position:"Query" name:"ClientToken"`
	OwnerAccount             string `position:"Query" name:"OwnerAccount"`
	DBInstanceId             string `position:"Query" name:"DBInstanceId"`
	OwnerId                  int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDBInstanceNetInfoRequest) Invoke(client *sdk.Client) (resp *DescribeDBInstanceNetInfoResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceNetInfo", "rds", "")
	resp = &DescribeDBInstanceNetInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDBInstanceNetInfoResponse struct {
	responses.BaseResponse
	RequestId           common.String
	InstanceNetworkType common.String
	DBInstanceNetInfos  DescribeDBInstanceNetInfoDBInstanceNetInfoList
}

type DescribeDBInstanceNetInfoDBInstanceNetInfo struct {
	Upgradeable          common.String
	ExpiredTime          common.String
	ConnectionString     common.String
	IPAddress            common.String
	IPType               common.String
	Port                 common.String
	VPCId                common.String
	VSwitchId            common.String
	ConnectionStringType common.String
	MaxDelayTime         common.String
	DistributionType     common.String
	SecurityIPGroups     DescribeDBInstanceNetInfoSecurityIPGroupList
	DBInstanceWeights    DescribeDBInstanceNetInfoDBInstanceWeightList
}

type DescribeDBInstanceNetInfoSecurityIPGroup struct {
	SecurityIPGroupName common.String
	SecurityIPs         common.String
}

type DescribeDBInstanceNetInfoDBInstanceWeight struct {
	DBInstanceId   common.String
	DBInstanceType common.String
	Availability   common.String
	Weight         common.String
}

type DescribeDBInstanceNetInfoDBInstanceNetInfoList []DescribeDBInstanceNetInfoDBInstanceNetInfo

func (list *DescribeDBInstanceNetInfoDBInstanceNetInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceNetInfoDBInstanceNetInfo)
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

type DescribeDBInstanceNetInfoSecurityIPGroupList []DescribeDBInstanceNetInfoSecurityIPGroup

func (list *DescribeDBInstanceNetInfoSecurityIPGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceNetInfoSecurityIPGroup)
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

type DescribeDBInstanceNetInfoDBInstanceWeightList []DescribeDBInstanceNetInfoDBInstanceWeight

func (list *DescribeDBInstanceNetInfoDBInstanceWeightList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceNetInfoDBInstanceWeight)
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
