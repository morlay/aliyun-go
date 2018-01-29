package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RunInstancesRequest struct {
	requests.RpcRequest
	ResourceOwnerId             int64                             `position:"Query" name:"ResourceOwnerId"`
	SecurityEnhancementStrategy string                            `position:"Query" name:"SecurityEnhancementStrategy"`
	KeyPairName                 string                            `position:"Query" name:"KeyPairName"`
	SpotPriceLimit              float32                           `position:"Query" name:"SpotPriceLimit"`
	HostName                    string                            `position:"Query" name:"HostName"`
	Password                    string                            `position:"Query" name:"Password"`
	Tags                        *RunInstancesTagList              `position:"Query" type:"Repeated" name:"Tag"`
	DryRun                      string                            `position:"Query" name:"DryRun"`
	OwnerId                     int64                             `position:"Query" name:"OwnerId"`
	VSwitchId                   string                            `position:"Query" name:"VSwitchId"`
	SpotStrategy                string                            `position:"Query" name:"SpotStrategy"`
	InstanceName                string                            `position:"Query" name:"InstanceName"`
	InternetChargeType          string                            `position:"Query" name:"InternetChargeType"`
	ZoneId                      string                            `position:"Query" name:"ZoneId"`
	InternetMaxBandwidthIn      int                               `position:"Query" name:"InternetMaxBandwidthIn"`
	ImageId                     string                            `position:"Query" name:"ImageId"`
	ClientToken                 string                            `position:"Query" name:"ClientToken"`
	IoOptimized                 string                            `position:"Query" name:"IoOptimized"`
	SecurityGroupId             string                            `position:"Query" name:"SecurityGroupId"`
	InternetMaxBandwidthOut     int                               `position:"Query" name:"InternetMaxBandwidthOut"`
	Description                 string                            `position:"Query" name:"Description"`
	SystemDiskCategory          string                            `position:"Query" name:"SystemDiskCategory"`
	UserData                    string                            `position:"Query" name:"UserData"`
	InstanceType                string                            `position:"Query" name:"InstanceType"`
	NetworkInterfaces           *RunInstancesNetworkInterfaceList `position:"Query" type:"Repeated" name:"NetworkInterface"`
	Amount                      int                               `position:"Query" name:"Amount"`
	ResourceOwnerAccount        string                            `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                string                            `position:"Query" name:"OwnerAccount"`
	SystemDiskDiskName          string                            `position:"Query" name:"SystemDiskDiskName"`
	RamRoleName                 string                            `position:"Query" name:"RamRoleName"`
	AutoReleaseTime             string                            `position:"Query" name:"AutoReleaseTime"`
	DataDisks                   *RunInstancesDataDiskList         `position:"Query" type:"Repeated" name:"DataDisk"`
	SystemDiskSize              string                            `position:"Query" name:"SystemDiskSize"`
	SystemDiskDescription       string                            `position:"Query" name:"SystemDiskDescription"`
}

func (req *RunInstancesRequest) Invoke(client *sdk.Client) (resp *RunInstancesResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "RunInstances", "ecs", "")
	resp = &RunInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type RunInstancesTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

type RunInstancesNetworkInterface struct {
	PrimaryIpAddress     string `name:"PrimaryIpAddress"`
	VSwitchId            string `name:"VSwitchId"`
	SecurityGroupId      string `name:"SecurityGroupId"`
	NetworkInterfaceName string `name:"NetworkInterfaceName"`
	Description          string `name:"Description"`
}

type RunInstancesDataDisk struct {
	Size               int    `name:"Size"`
	SnapshotId         string `name:"SnapshotId"`
	Category           string `name:"Category"`
	Encrypted          string `name:"Encrypted"`
	DiskName           string `name:"DiskName"`
	Description        string `name:"Description"`
	DeleteWithInstance string `name:"DeleteWithInstance"`
}

type RunInstancesResponse struct {
	responses.BaseResponse
	RequestId      string
	InstanceIdSets RunInstancesInstanceIdSetList
}

type RunInstancesTagList []RunInstancesTag

func (list *RunInstancesTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RunInstancesTag)
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

type RunInstancesNetworkInterfaceList []RunInstancesNetworkInterface

func (list *RunInstancesNetworkInterfaceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RunInstancesNetworkInterface)
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

type RunInstancesDataDiskList []RunInstancesDataDisk

func (list *RunInstancesDataDiskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RunInstancesDataDisk)
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

type RunInstancesInstanceIdSetList []string

func (list *RunInstancesInstanceIdSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
