package polardb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDBClusterAttributeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDBClusterAttributeRequest) Invoke(client *sdk.Client) (resp *DescribeDBClusterAttributeResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "DescribeDBClusterAttribute", "polardb", "")
	resp = &DescribeDBClusterAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDBClusterAttributeResponse struct {
	responses.BaseResponse
	RequestId common.String
	Items     DescribeDBClusterAttributeDBClusterAttributeList
}

type DescribeDBClusterAttributeDBClusterAttribute struct {
	RegionId             common.String
	DBClusterNetworkType common.String
	VPCId                common.String
	VSwitchId            common.String
	PayType              common.String
	DBClusterId          common.String
	DBClusterStatus      common.String
	DBClusterDescription common.String
	Engine               common.String
	DBType               common.String
	DBVersion            common.String
	Storage              common.Long
	ConnectionString     common.Long
	Port                 common.Long
	DBClusterNetType     common.String
	LockMode             common.String
	LockReason           common.String
	CreationTime         common.String
	ExpireTime           common.String
	DbInstances          DescribeDBClusterAttributeDbInstanceList
}

type DescribeDBClusterAttributeDbInstance struct {
	DBInstanceId          common.String
	DBInstanceStatus      common.String
	DBInstanceDescription common.String
	Engine                common.String
	DBType                common.String
	DBVersion             common.String
	DBInstanceStorage     common.String
	LockMode              common.String
	LockReason            common.String
	MaintainStartTime     common.String
	MaintainEndTime       common.String
	CreationTime          common.String
	DBInstanceClass       common.String
	SecurityIPList        common.String
	DBInstanceType        common.String
}

type DescribeDBClusterAttributeDBClusterAttributeList []DescribeDBClusterAttributeDBClusterAttribute

func (list *DescribeDBClusterAttributeDBClusterAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBClusterAttributeDBClusterAttribute)
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

type DescribeDBClusterAttributeDbInstanceList []DescribeDBClusterAttributeDbInstance

func (list *DescribeDBClusterAttributeDbInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBClusterAttributeDbInstance)
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
