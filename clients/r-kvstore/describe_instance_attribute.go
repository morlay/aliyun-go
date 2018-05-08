package r_kvstore

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeInstanceAttributeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeInstanceAttributeRequest) Invoke(client *sdk.Client) (resp *DescribeInstanceAttributeResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeInstanceAttribute", "redisa", "")
	resp = &DescribeInstanceAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeInstanceAttributeResponse struct {
	responses.BaseResponse
	RequestId common.String
	Instances DescribeInstanceAttributeDBInstanceAttributeList
}

type DescribeInstanceAttributeDBInstanceAttribute struct {
	InstanceId          common.String
	InstanceName        common.String
	ConnectionDomain    common.String
	Port                common.Long
	InstanceStatus      common.String
	RegionId            common.String
	Capacity            common.Long
	InstanceClass       common.String
	QPS                 common.Long
	Bandwidth           common.Long
	Connections         common.Long
	ZoneId              common.String
	Config              common.String
	ChargeType          common.String
	NodeType            common.String
	NetworkType         common.String
	VpcId               common.String
	VSwitchId           common.String
	PrivateIp           common.String
	CreateTime          common.String
	EndTime             common.String
	HasRenewChangeOrder common.String
	IsRds               bool
	Engine              common.String
	EngineVersion       common.String
	MaintainStartTime   common.String
	MaintainEndTime     common.String
	AvailabilityValue   common.String
	SecurityIPList      common.String
	InstanceType        common.String
	ArchitectureType    common.String
	NodeType1           common.String
	PackageType         common.String
}

type DescribeInstanceAttributeDBInstanceAttributeList []DescribeInstanceAttributeDBInstanceAttribute

func (list *DescribeInstanceAttributeDBInstanceAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceAttributeDBInstanceAttribute)
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
