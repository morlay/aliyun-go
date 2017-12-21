package polardb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId string
	Items     DescribeDBClusterAttributeDBClusterAttributeList
}

type DescribeDBClusterAttributeDBClusterAttribute struct {
	RegionId             string
	DBClusterNetworkType string
	VPCId                string
	VSwitchId            string
	PayType              string
	DBClusterId          string
	DBClusterStatus      string
	DBClusterDescription string
	Engine               string
	DBType               string
	DBVersion            string
	Storage              int64
	ConnectionString     int64
	Port                 int64
	DBClusterNetType     string
	LockMode             string
	LockReason           string
	CreationTime         string
	ExpireTime           string
	DbInstances          DescribeDBClusterAttributeDbInstanceList
}

type DescribeDBClusterAttributeDbInstance struct {
	DBInstanceId          string
	DBInstanceStatus      string
	DBInstanceDescription string
	Engine                string
	DBType                string
	DBVersion             string
	DBInstanceStorage     string
	LockMode              string
	LockReason            string
	MaintainStartTime     string
	MaintainEndTime       string
	CreationTime          string
	DBInstanceClass       string
	SecurityIPList        string
	DBInstanceType        string
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
