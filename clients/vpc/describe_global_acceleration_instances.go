package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId                   string
	TotalCount                  int
	PageNumber                  int
	PageSize                    int
	GlobalAccelerationInstances DescribeGlobalAccelerationInstancesGlobalAccelerationInstanceList
}

type DescribeGlobalAccelerationInstancesGlobalAccelerationInstance struct {
	RegionId                     string
	GlobalAccelerationInstanceId string
	IpAddress                    string
	Status                       string
	Bandwidth                    string
	InternetChargeType           string
	ChargeType                   string
	BandwidthType                string
	AccelerationLocation         string
	ServiceLocation              string
	Name                         string
	Description                  string
	ExpiredTime                  string
	CreationTime                 string
	OperationLocks               DescribeGlobalAccelerationInstancesLockReasonList
	BackendServers               DescribeGlobalAccelerationInstancesBackendServerList
	PublicIpAddresses            DescribeGlobalAccelerationInstancesPublicIpAddressList
}

type DescribeGlobalAccelerationInstancesLockReason struct {
	LockReason string
}

type DescribeGlobalAccelerationInstancesBackendServer struct {
	RegionId        string
	ServerId        string
	ServerIpAddress string
	ServerType      string
}

type DescribeGlobalAccelerationInstancesPublicIpAddress struct {
	AllocationId string
	IpAddress    string
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
