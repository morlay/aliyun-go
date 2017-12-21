package ess

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateScalingConfigurationRequest struct {
	DataDisk3Size               int                                          `position:"Query" name:"DataDisk.3.Size"`
	ImageId                     string                                       `position:"Query" name:"ImageId"`
	DataDisk1SnapshotId         string                                       `position:"Query" name:"DataDisk.1.SnapshotId"`
	DataDisk3Category           string                                       `position:"Query" name:"DataDisk.3.Category"`
	DataDisk1Device             string                                       `position:"Query" name:"DataDisk.1.Device"`
	ScalingGroupId              string                                       `position:"Query" name:"ScalingGroupId"`
	DataDisk2Device             string                                       `position:"Query" name:"DataDisk.2.Device"`
	InstanceTypess              *CreateScalingConfigurationInstanceTypesList `position:"Query" type:"Repeated" name:"InstanceTypes"`
	IoOptimized                 string                                       `position:"Query" name:"IoOptimized"`
	SecurityGroupId             string                                       `position:"Query" name:"SecurityGroupId"`
	InternetMaxBandwidthOut     int                                          `position:"Query" name:"InternetMaxBandwidthOut"`
	SecurityEnhancementStrategy string                                       `position:"Query" name:"SecurityEnhancementStrategy"`
	KeyPairName                 string                                       `position:"Query" name:"KeyPairName"`
	SystemDiskCategory          string                                       `position:"Query" name:"SystemDiskCategory"`
	UserData                    string                                       `position:"Query" name:"UserData"`
	DataDisk4Category           string                                       `position:"Query" name:"DataDisk.4.Category"`
	DataDisk2SnapshotId         string                                       `position:"Query" name:"DataDisk.2.SnapshotId"`
	DataDisk4Size               int                                          `position:"Query" name:"DataDisk.4.Size"`
	InstanceType                string                                       `position:"Query" name:"InstanceType"`
	DataDisk2Category           string                                       `position:"Query" name:"DataDisk.2.Category"`
	DataDisk1Size               int                                          `position:"Query" name:"DataDisk.1.Size"`
	DataDisk3SnapshotId         string                                       `position:"Query" name:"DataDisk.3.SnapshotId"`
	ResourceOwnerAccount        string                                       `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                string                                       `position:"Query" name:"OwnerAccount"`
	DataDisk2Size               int                                          `position:"Query" name:"DataDisk.2.Size"`
	RamRoleName                 string                                       `position:"Query" name:"RamRoleName"`
	OwnerId                     int64                                        `position:"Query" name:"OwnerId"`
	ScalingConfigurationName    string                                       `position:"Query" name:"ScalingConfigurationName"`
	Tags                        string                                       `position:"Query" name:"Tags"`
	DataDisk2DeleteWithInstance string                                       `position:"Query" name:"DataDisk.2.DeleteWithInstance"`
	DataDisk1Category           string                                       `position:"Query" name:"DataDisk.1.Category"`
	DataDisk3DeleteWithInstance string                                       `position:"Query" name:"DataDisk.3.DeleteWithInstance"`
	LoadBalancerWeight          int                                          `position:"Query" name:"LoadBalancerWeight"`
	InstanceName                string                                       `position:"Query" name:"InstanceName"`
	SystemDiskSize              int                                          `position:"Query" name:"SystemDiskSize"`
	DataDisk4SnapshotId         string                                       `position:"Query" name:"DataDisk.4.SnapshotId"`
	DataDisk4Device             string                                       `position:"Query" name:"DataDisk.4.Device"`
	InternetChargeType          string                                       `position:"Query" name:"InternetChargeType"`
	DataDisk3Device             string                                       `position:"Query" name:"DataDisk.3.Device"`
	DataDisk4DeleteWithInstance string                                       `position:"Query" name:"DataDisk.4.DeleteWithInstance"`
	InternetMaxBandwidthIn      int                                          `position:"Query" name:"InternetMaxBandwidthIn"`
	DataDisk1DeleteWithInstance string                                       `position:"Query" name:"DataDisk.1.DeleteWithInstance"`
}

func (r CreateScalingConfigurationRequest) Invoke(client *sdk.Client) (response *CreateScalingConfigurationResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateScalingConfigurationRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ess", "2014-08-28", "CreateScalingConfiguration", "ess", "")

	resp := struct {
		*responses.BaseResponse
		CreateScalingConfigurationResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateScalingConfigurationResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateScalingConfigurationResponse struct {
	ScalingConfigurationId string
	RequestId              string
}

type CreateScalingConfigurationInstanceTypesList []string

func (list *CreateScalingConfigurationInstanceTypesList) UnmarshalJSON(data []byte) error {
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
