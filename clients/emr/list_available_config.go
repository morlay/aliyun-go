package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId          string
	SecurityGroupTypes ListAvailableConfigSecurityGroupTypeList
	InstanceTypes      ListAvailableConfigInstanceTypeList
	EmrVerTypes        ListAvailableConfigEmrVerTypeList
	ZoneTypes          ListAvailableConfigZoneTypeList
	Vpcs               ListAvailableConfigVpcList
}

type ListAvailableConfigSecurityGroupType struct {
	Name  string
	State string
	Id    string
}

type ListAvailableConfigInstanceType struct {
	Classify               string
	Type                   string
	CpuCore                int
	MemSize                int
	HasCloudDisk           bool
	HasEfficiencyCloudDisk bool
	HasSSDCloudDisk        bool
}

type ListAvailableConfigEmrVerType struct {
	Name       string
	SubModules ListAvailableConfigSubModuleList
}

type ListAvailableConfigSubModule struct {
	Name         string
	RequiredList ListAvailableConfigRequiredList
	OptionalList ListAvailableConfigOptionalList
}

type ListAvailableConfigRequired struct {
	DisplayName string
	Name        string
	OnlyDisplay bool
	StartTpe    int
	Version     string
}

type ListAvailableConfigOptional struct {
	DisplayName string
	Name        string
	OnlyDisplay bool
	StartTpe    int
	Version     string
}

type ListAvailableConfigZoneType struct {
	Name                          string
	Id                            string
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
	Id             string
	CidrBlock      string
	VSwitchs       ListAvailableConfigVSwitchList
	SecurityGroups ListAvailableConfigSecurityGroupList
}

type ListAvailableConfigVSwitch struct {
	Id        string
	CidrBlock string
	ZoneId    string
}

type ListAvailableConfigSecurityGroup struct {
	Name string
	Id   string
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

type ListAvailableConfigAvailableResourceCreationListList []string

func (list *ListAvailableConfigAvailableResourceCreationListList) UnmarshalJSON(data []byte) error {
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

type ListAvailableConfigAvailableDiskCategoryListList []string

func (list *ListAvailableConfigAvailableDiskCategoryListList) UnmarshalJSON(data []byte) error {
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

type ListAvailableConfigSystemDiskCategoryList []string

func (list *ListAvailableConfigSystemDiskCategoryList) UnmarshalJSON(data []byte) error {
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

type ListAvailableConfigDataDiskCategoryList []string

func (list *ListAvailableConfigDataDiskCategoryList) UnmarshalJSON(data []byte) error {
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

type ListAvailableConfigNetworkTypeList []string

func (list *ListAvailableConfigNetworkTypeList) UnmarshalJSON(data []byte) error {
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

type ListAvailableConfigSupportedInstanceTypeList []string

func (list *ListAvailableConfigSupportedInstanceTypeList) UnmarshalJSON(data []byte) error {
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

type ListAvailableConfigInstanceTypeFamilyList []string

func (list *ListAvailableConfigInstanceTypeFamilyList) UnmarshalJSON(data []byte) error {
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

type ListAvailableConfigInstanceGenerationList []string

func (list *ListAvailableConfigInstanceGenerationList) UnmarshalJSON(data []byte) error {
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
