package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListAvailableConfigRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
}

func (req *ListAvailableConfigRequest) Invoke(client *sdk.Client) (resp *ListAvailableConfigResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListAvailableConfig", "", "")
	resp = &ListAvailableConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListAvailableConfigResponse struct {
	responses.BaseResponse
	RequestId                    common.String
	SecurityGroupTypes           ListAvailableConfigSecurityGroupTypeList
	InstanceTypes                ListAvailableConfigInstanceTypeList
	EmrVerTypes                  ListAvailableConfigEmrVerTypeList
	ZoneTypes                    ListAvailableConfigZoneTypeList
	Vpcs                         ListAvailableConfigVpcList
	EmrSupportedInstanceTypeList ListAvailableConfigEmrSupportInstanceTypeList
}

type ListAvailableConfigSecurityGroupType struct {
	Name  common.String
	State common.String
	Id    common.String
}

type ListAvailableConfigInstanceType struct {
	Classify               common.String
	Type                   common.String
	CpuCore                common.Integer
	MemSize                common.Integer
	HasCloudDisk           bool
	HasEfficiencyCloudDisk bool
	HasSSDCloudDisk        bool
}

type ListAvailableConfigEmrVerType struct {
	Name       common.String
	EcmStack   bool
	SubModules ListAvailableConfigSubModuleList
}

type ListAvailableConfigSubModule struct {
	Name         common.String
	RequiredList ListAvailableConfigRequiredList
	OptionalList ListAvailableConfigOptionalList
}

type ListAvailableConfigRequired struct {
	DisplayName common.String
	Name        common.String
	OnlyDisplay bool
	StartTpe    common.Integer
	Version     common.String
}

type ListAvailableConfigOptional struct {
	DisplayName common.String
	Name        common.String
	OnlyDisplay bool
	StartTpe    common.Integer
	Version     common.String
}

type ListAvailableConfigZoneType struct {
	Name                          common.String
	Id                            common.String
	AvailableResources            ListAvailableConfigAvailableResourceList
	AvailableResourceCreationList ListAvailableConfigAvailableResourceCreationListList
	AvailableDiskCategoryList     ListAvailableConfigAvailableDiskCategoryListList
}

type ListAvailableConfigAvailableResource struct {
	IoOptimized            bool
	SystemDiskCategories   ListAvailableConfigSystemDiskCategoryList
	DataDiskCategories     ListAvailableConfigDataDiskCategoryList
	NetworkTypes           ListAvailableConfigNetworkTypeList
	SupportedInstanceTypes ListAvailableConfigSupportedInstanceTypeList
	InstanceTypeFamilies   ListAvailableConfigInstanceTypeFamilyList
	InstanceGenerations    ListAvailableConfigInstanceGenerationList
}

type ListAvailableConfigVpc struct {
	Id             common.String
	VpcName        common.String
	CidrBlock      common.String
	VSwitchs       ListAvailableConfigVSwitchList
	SecurityGroups ListAvailableConfigSecurityGroupList
}

type ListAvailableConfigVSwitch struct {
	Id          common.String
	VswitchName common.String
	CidrBlock   common.String
	ZoneId      common.String
}

type ListAvailableConfigSecurityGroup struct {
	Name common.String
	Id   common.String
}

type ListAvailableConfigEmrSupportInstanceType struct {
	ClusterType             common.String
	NodeTypeSupportInfoList ListAvailableConfigClusterNodeTypeSupportInfoList
}

type ListAvailableConfigClusterNodeTypeSupportInfo struct {
	ClusterNodeType         common.String
	SupportInstanceTypeList ListAvailableConfigSupportInstanceTypeListList
}

type ListAvailableConfigSecurityGroupTypeList []ListAvailableConfigSecurityGroupType

func (list *ListAvailableConfigSecurityGroupTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigSecurityGroupType)
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

type ListAvailableConfigInstanceTypeList []ListAvailableConfigInstanceType

func (list *ListAvailableConfigInstanceTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigInstanceType)
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

type ListAvailableConfigEmrVerTypeList []ListAvailableConfigEmrVerType

func (list *ListAvailableConfigEmrVerTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigEmrVerType)
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

type ListAvailableConfigZoneTypeList []ListAvailableConfigZoneType

func (list *ListAvailableConfigZoneTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigZoneType)
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

type ListAvailableConfigVpcList []ListAvailableConfigVpc

func (list *ListAvailableConfigVpcList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigVpc)
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

type ListAvailableConfigEmrSupportInstanceTypeList []ListAvailableConfigEmrSupportInstanceType

func (list *ListAvailableConfigEmrSupportInstanceTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigEmrSupportInstanceType)
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

type ListAvailableConfigSubModuleList []ListAvailableConfigSubModule

func (list *ListAvailableConfigSubModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigSubModule)
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

type ListAvailableConfigRequiredList []ListAvailableConfigRequired

func (list *ListAvailableConfigRequiredList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigRequired)
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

type ListAvailableConfigOptionalList []ListAvailableConfigOptional

func (list *ListAvailableConfigOptionalList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigOptional)
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

type ListAvailableConfigAvailableResourceList []ListAvailableConfigAvailableResource

func (list *ListAvailableConfigAvailableResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigAvailableResource)
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

type ListAvailableConfigAvailableResourceCreationListList []common.String

func (list *ListAvailableConfigAvailableResourceCreationListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type ListAvailableConfigAvailableDiskCategoryListList []common.String

func (list *ListAvailableConfigAvailableDiskCategoryListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type ListAvailableConfigSystemDiskCategoryList []common.String

func (list *ListAvailableConfigSystemDiskCategoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type ListAvailableConfigDataDiskCategoryList []common.String

func (list *ListAvailableConfigDataDiskCategoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type ListAvailableConfigNetworkTypeList []common.String

func (list *ListAvailableConfigNetworkTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type ListAvailableConfigSupportedInstanceTypeList []common.String

func (list *ListAvailableConfigSupportedInstanceTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type ListAvailableConfigInstanceTypeFamilyList []common.String

func (list *ListAvailableConfigInstanceTypeFamilyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type ListAvailableConfigInstanceGenerationList []common.String

func (list *ListAvailableConfigInstanceGenerationList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type ListAvailableConfigVSwitchList []ListAvailableConfigVSwitch

func (list *ListAvailableConfigVSwitchList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigVSwitch)
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

type ListAvailableConfigSecurityGroupList []ListAvailableConfigSecurityGroup

func (list *ListAvailableConfigSecurityGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigSecurityGroup)
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

type ListAvailableConfigClusterNodeTypeSupportInfoList []ListAvailableConfigClusterNodeTypeSupportInfo

func (list *ListAvailableConfigClusterNodeTypeSupportInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigClusterNodeTypeSupportInfo)
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

type ListAvailableConfigSupportInstanceTypeListList []common.String

func (list *ListAvailableConfigSupportInstanceTypeListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
