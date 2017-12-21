package polardb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDBInstanceAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeDBInstanceAttributeRequest) Invoke(client *sdk.Client) (response *DescribeDBInstanceAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDBInstanceAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("polardb", "2017-08-01", "DescribeDBInstanceAttribute", "polardb", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDBInstanceAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDBInstanceAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDBInstanceAttributeResponse struct {
	RequestId string
	Items     DescribeDBInstanceAttributeDBInstanceAttributeList
}

type DescribeDBInstanceAttributeDBInstanceAttribute struct {
	DBInstanceId          string
	DBClusterDescription  string
	DBClusterId           string
	PayType               string
	DBInstanceType        string
	RegionId              string
	ZoneId                string
	Engine                string
	DBType                string
	DBVersion             string
	DBInstanceClass       string
	DBInstanceStorage     int64
	DBInstanceStatus      string
	DBInstanceDescription string
	ConnectionString      int64
	Port                  int64
	DBInstanceNetType     string
	LockMode              string
	LockReason            string
	CreationTime          string
	ExpireTime            string
	MaintainStartTime     string
	MaintainEndTime       string
	MaxConnections        int
	MaxIOPS               int
	SecurityIPList        string
	InstanceNetworkType   string
	VpcId                 string
	VSwitchId             string
	DBInstanceType1       string
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
