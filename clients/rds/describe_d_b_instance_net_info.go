package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDBInstanceNetInfoRequest struct {
	ResourceOwnerId          int64  `position:"Query" name:"ResourceOwnerId"`
	Flag                     string `position:"Query" name:"Flag"`
	DBInstanceNetRWSplitType string `position:"Query" name:"DBInstanceNetRWSplitType"`
	ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken              string `position:"Query" name:"ClientToken"`
	OwnerAccount             string `position:"Query" name:"OwnerAccount"`
	DBInstanceId             string `position:"Query" name:"DBInstanceId"`
	OwnerId                  int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeDBInstanceNetInfoRequest) Invoke(client *sdk.Client) (response *DescribeDBInstanceNetInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDBInstanceNetInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceNetInfo", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDBInstanceNetInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDBInstanceNetInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDBInstanceNetInfoResponse struct {
	RequestId           string
	InstanceNetworkType string
	DBInstanceNetInfos  DescribeDBInstanceNetInfoDBInstanceNetInfoList
}

type DescribeDBInstanceNetInfoDBInstanceNetInfo struct {
	Upgradeable          string
	ExpiredTime          string
	ConnectionString     string
	IPAddress            string
	IPType               string
	Port                 string
	VPCId                string
	VSwitchId            string
	ConnectionStringType string
	MaxDelayTime         string
	DistributionType     string
	SecurityIPGroups     DescribeDBInstanceNetInfoSecurityIPGroupList
	DBInstanceWeights    DescribeDBInstanceNetInfoDBInstanceWeightList
}

type DescribeDBInstanceNetInfoSecurityIPGroup struct {
	SecurityIPGroupName string
	SecurityIPs         string
}

type DescribeDBInstanceNetInfoDBInstanceWeight struct {
	DBInstanceId   string
	DBInstanceType string
	Availability   string
	Weight         string
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
