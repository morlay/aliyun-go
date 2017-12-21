package dds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDBInstancesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBInstanceIds        string `position:"Query" name:"DBInstanceIds"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	DBInstanceType       string `position:"Query" name:"DBInstanceType"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	Engine               string `position:"Query" name:"Engine"`
	PageSize             int    `position:"Query" name:"PageSize"`
}

func (r DescribeDBInstancesRequest) Invoke(client *sdk.Client) (response *DescribeDBInstancesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDBInstancesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Dds", "2015-12-01", "DescribeDBInstances", "dds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDBInstancesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDBInstancesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDBInstancesResponse struct {
	RequestId   string
	PageNumber  int
	PageSize    int
	TotalCount  int
	DBInstances DescribeDBInstancesDBInstanceList
}

type DescribeDBInstancesDBInstance struct {
	DBInstanceId          string
	DBInstanceDescription string
	RegionId              string
	ZoneId                string
	Engine                string
	EngineVersion         string
	DBInstanceClass       string
	DBInstanceStorage     int
	DBInstanceStatus      string
	LockMode              string
	ChargeType            string
	NetworkType           string
	CreationTime          string
	ExpireTime            string
	DBInstanceType        string
	LastDowngradeTime     int
	MongosList            DescribeDBInstancesMongosAttributeList
	ShardList             DescribeDBInstancesShardAttributeList
}

type DescribeDBInstancesMongosAttribute struct {
	NodeId          string
	NodeDescription string
	NodeClass       string
	ConnectSting    string
	Port            int
}

type DescribeDBInstancesShardAttribute struct {
	NodeId          string
	NodeDescription string
	NodeClass       string
	NodeStorage     int
}

type DescribeDBInstancesDBInstanceList []DescribeDBInstancesDBInstance

func (list *DescribeDBInstancesDBInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancesDBInstance)
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

type DescribeDBInstancesMongosAttributeList []DescribeDBInstancesMongosAttribute

func (list *DescribeDBInstancesMongosAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancesMongosAttribute)
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

type DescribeDBInstancesShardAttributeList []DescribeDBInstancesShardAttribute

func (list *DescribeDBInstancesShardAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancesShardAttribute)
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
