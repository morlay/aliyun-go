package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeGlobalAccelerationInstancesRequest struct {
	requests.RpcRequest
	IpAddress                    string `position:"Query" name:"IpAddress"`
	ResourceOwnerId              int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthType                string `position:"Query" name:"BandwidthType"`
	ResourceOwnerAccount         string `position:"Query" name:"ResourceOwnerAccount"`
	ServiceLocation              string `position:"Query" name:"ServiceLocation"`
	OwnerAccount                 string `position:"Query" name:"OwnerAccount"`
	OwnerId                      int64  `position:"Query" name:"OwnerId"`
	GlobalAccelerationInstanceId string `position:"Query" name:"GlobalAccelerationInstanceId"`
	ServerId                     string `position:"Query" name:"ServerId"`
	PageNumber                   int    `position:"Query" name:"PageNumber"`
	Name                         string `position:"Query" name:"Name"`
	PageSize                     int    `position:"Query" name:"PageSize"`
	Status                       string `position:"Query" name:"Status"`
}

func (req *DescribeGlobalAccelerationInstancesRequest) Invoke(client *sdk.Client) (resp *DescribeGlobalAccelerationInstancesResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeGlobalAccelerationInstances", "vpc", "")
	resp = &DescribeGlobalAccelerationInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeGlobalAccelerationInstancesResponse struct {
	responses.BaseResponse
	RequestId                   common.String
	TotalCount                  common.Integer
	PageNumber                  common.Integer
	PageSize                    common.Integer
	GlobalAccelerationInstances DescribeGlobalAccelerationInstancesGlobalAccelerationInstanceList
}

type DescribeGlobalAccelerationInstancesGlobalAccelerationInstance struct {
	RegionId                     common.String
	GlobalAccelerationInstanceId common.String
	IpAddress                    common.String
	Status                       common.String
	Bandwidth                    common.String
	InternetChargeType           common.String
	ChargeType                   common.String
	BandwidthType                common.String
	AccelerationLocation         common.String
	ServiceLocation              common.String
	Name                         common.String
	Description                  common.String
	ExpiredTime                  common.String
	CreationTime                 common.String
	OperationLocks               DescribeGlobalAccelerationInstancesLockReasonList
	BackendServers               DescribeGlobalAccelerationInstancesBackendServerList
	PublicIpAddresses            DescribeGlobalAccelerationInstancesPublicIpAddressList
}

type DescribeGlobalAccelerationInstancesLockReason struct {
	LockReason common.String
}

type DescribeGlobalAccelerationInstancesBackendServer struct {
	RegionId        common.String
	ServerId        common.String
	ServerIpAddress common.String
	ServerType      common.String
}

type DescribeGlobalAccelerationInstancesPublicIpAddress struct {
	AllocationId common.String
	IpAddress    common.String
}

type DescribeGlobalAccelerationInstancesGlobalAccelerationInstanceList []DescribeGlobalAccelerationInstancesGlobalAccelerationInstance

func (list *DescribeGlobalAccelerationInstancesGlobalAccelerationInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeGlobalAccelerationInstancesGlobalAccelerationInstance)
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

type DescribeGlobalAccelerationInstancesLockReasonList []DescribeGlobalAccelerationInstancesLockReason

func (list *DescribeGlobalAccelerationInstancesLockReasonList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeGlobalAccelerationInstancesLockReason)
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

type DescribeGlobalAccelerationInstancesBackendServerList []DescribeGlobalAccelerationInstancesBackendServer

func (list *DescribeGlobalAccelerationInstancesBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeGlobalAccelerationInstancesBackendServer)
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

type DescribeGlobalAccelerationInstancesPublicIpAddressList []DescribeGlobalAccelerationInstancesPublicIpAddress

func (list *DescribeGlobalAccelerationInstancesPublicIpAddressList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeGlobalAccelerationInstancesPublicIpAddress)
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
