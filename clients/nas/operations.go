package nas

import "encoding/json"

func (c *NasClient) CreateFileSystem(req *CreateFileSystemArgs) (resp *CreateFileSystemResponse, err error) {
	resp = &CreateFileSystemResponse{}
	err = c.Request("CreateFileSystem", req, resp)
	return
}

type CreateFileSystemArgs struct {
	Description  string
	ProtocolType string
	StorageType  string
}
type CreateFileSystemResponse struct {
	RequestId    string
	FileSystemId string
}

func (c *NasClient) CreateMountTarget(req *CreateMountTargetArgs) (resp *CreateMountTargetResponse, err error) {
	resp = &CreateMountTargetResponse{}
	err = c.Request("CreateMountTarget", req, resp)
	return
}

type CreateMountTargetArgs struct {
	VSwitchId       string
	VpcId           string
	NetworkType     string
	AccessGroupName string
	FileSystemId    string
}
type CreateMountTargetResponse struct {
	RequestId         string
	MountTargetDomain string
}

func (c *NasClient) DescribeMountTargets(req *DescribeMountTargetsArgs) (resp *DescribeMountTargetsResponse, err error) {
	resp = &DescribeMountTargetsResponse{}
	err = c.Request("DescribeMountTargets", req, resp)
	return
}

type DescribeMountTargetsMountTarget struct {
	MountTargetDomain string
	NetworkType       string
	VpcId             string
	VswId             string
	AccessGroup       string
	Status            string
}
type DescribeMountTargetsArgs struct {
	MountTargetDomain string
	PageSize          int
	PageNumber        int
	FileSystemId      string
}
type DescribeMountTargetsResponse struct {
	RequestId    string
	TotalCount   int
	PageSize     int
	PageNumber   int
	MountTargets DescribeMountTargetsMountTargetList
}

type DescribeMountTargetsMountTargetList []DescribeMountTargetsMountTarget

func (list *DescribeMountTargetsMountTargetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMountTargetsMountTarget)
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

func (c *NasClient) CreateAccessRule(req *CreateAccessRuleArgs) (resp *CreateAccessRuleResponse, err error) {
	resp = &CreateAccessRuleResponse{}
	err = c.Request("CreateAccessRule", req, resp)
	return
}

type CreateAccessRuleArgs struct {
	RWAccessType    string
	SourceCidrIp    string
	UserAccessType  string
	Priority        int
	AccessGroupName string
}
type CreateAccessRuleResponse struct {
	RequestId    string
	AccessRuleId string
}

func (c *NasClient) DescribeFileSystems(req *DescribeFileSystemsArgs) (resp *DescribeFileSystemsResponse, err error) {
	resp = &DescribeFileSystemsResponse{}
	err = c.Request("DescribeFileSystems", req, resp)
	return
}

type DescribeFileSystemsFileSystem struct {
	FileSystemId string
	Destription  string
	CreateTime   string
	RegionId     string
	ProtocolType string
	StorageType  string
	MeteredSize  int64
	MountTargets DescribeFileSystemsMountTargetList
	Packages     DescribeFileSystems_PackageList
}

type DescribeFileSystemsMountTarget struct {
	MountTargetDomain string
}

type DescribeFileSystems_Package struct {
	PackageId string
}
type DescribeFileSystemsArgs struct {
	PageSize     int
	PageNumber   int
	FileSystemId string
}
type DescribeFileSystemsResponse struct {
	RequestId   string
	TotalCount  int
	PageSize    int
	PageNumber  int
	FileSystems DescribeFileSystemsFileSystemList
}

type DescribeFileSystemsMountTargetList []DescribeFileSystemsMountTarget

func (list *DescribeFileSystemsMountTargetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeFileSystemsMountTarget)
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

type DescribeFileSystems_PackageList []DescribeFileSystems_Package

func (list *DescribeFileSystems_PackageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeFileSystems_Package)
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

type DescribeFileSystemsFileSystemList []DescribeFileSystemsFileSystem

func (list *DescribeFileSystemsFileSystemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeFileSystemsFileSystem)
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

func (c *NasClient) ModifyFileSystem(req *ModifyFileSystemArgs) (resp *ModifyFileSystemResponse, err error) {
	resp = &ModifyFileSystemResponse{}
	err = c.Request("ModifyFileSystem", req, resp)
	return
}

type ModifyFileSystemArgs struct {
	Description  string
	FileSystemId string
}
type ModifyFileSystemResponse struct {
	RequestId string
}

func (c *NasClient) ModifyAccessRule(req *ModifyAccessRuleArgs) (resp *ModifyAccessRuleResponse, err error) {
	resp = &ModifyAccessRuleResponse{}
	err = c.Request("ModifyAccessRule", req, resp)
	return
}

type ModifyAccessRuleArgs struct {
	RWAccessType    string
	SourceCidrIp    string
	UserAccessType  string
	Priority        int
	AccessGroupName string
	AccessRuleId    string
}
type ModifyAccessRuleResponse struct {
	RequestId string
}

func (c *NasClient) DeleteMountTarget(req *DeleteMountTargetArgs) (resp *DeleteMountTargetResponse, err error) {
	resp = &DeleteMountTargetResponse{}
	err = c.Request("DeleteMountTarget", req, resp)
	return
}

type DeleteMountTargetArgs struct {
	MountTargetDomain string
	FileSystemId      string
}
type DeleteMountTargetResponse struct {
	RequestId string
}

func (c *NasClient) ModifyMountTarget(req *ModifyMountTargetArgs) (resp *ModifyMountTargetResponse, err error) {
	resp = &ModifyMountTargetResponse{}
	err = c.Request("ModifyMountTarget", req, resp)
	return
}

type ModifyMountTargetArgs struct {
	MountTargetDomain string
	AccessGroupName   string
	FileSystemId      string
	Status            string
}
type ModifyMountTargetResponse struct {
	RequestId string
}

func (c *NasClient) DescribeAccessRules(req *DescribeAccessRulesArgs) (resp *DescribeAccessRulesResponse, err error) {
	resp = &DescribeAccessRulesResponse{}
	err = c.Request("DescribeAccessRules", req, resp)
	return
}

type DescribeAccessRulesAccessRule struct {
	SourceCidrIp string
	Priority     int
	AccessRuleId string
	RWAccess     string
	UserAccess   string
}
type DescribeAccessRulesArgs struct {
	PageSize        int
	AccessGroupName string
	AccessRuleId    string
	PageNumber      int
}
type DescribeAccessRulesResponse struct {
	RequestId   string
	TotalCount  int
	PageSize    int
	PageNumber  int
	AccessRules DescribeAccessRulesAccessRuleList
}

type DescribeAccessRulesAccessRuleList []DescribeAccessRulesAccessRule

func (list *DescribeAccessRulesAccessRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccessRulesAccessRule)
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

func (c *NasClient) DescribeAccessGroups(req *DescribeAccessGroupsArgs) (resp *DescribeAccessGroupsResponse, err error) {
	resp = &DescribeAccessGroupsResponse{}
	err = c.Request("DescribeAccessGroups", req, resp)
	return
}

type DescribeAccessGroupsAccessGroup struct {
	AccessGroupName  string
	AccessGroupType  string
	RuleCount        int
	MountTargetCount int
	Description      string
}
type DescribeAccessGroupsArgs struct {
	PageSize        int
	AccessGroupName string
	PageNumber      int
}
type DescribeAccessGroupsResponse struct {
	RequestId    string
	TotalCount   int
	PageSize     int
	PageNumber   int
	AccessGroups DescribeAccessGroupsAccessGroupList
}

type DescribeAccessGroupsAccessGroupList []DescribeAccessGroupsAccessGroup

func (list *DescribeAccessGroupsAccessGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccessGroupsAccessGroup)
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

func (c *NasClient) CreateAccessGroup(req *CreateAccessGroupArgs) (resp *CreateAccessGroupResponse, err error) {
	resp = &CreateAccessGroupResponse{}
	err = c.Request("CreateAccessGroup", req, resp)
	return
}

type CreateAccessGroupArgs struct {
	Description     string
	AccessGroupType string
	AccessGroupName string
}
type CreateAccessGroupResponse struct {
	RequestId       string
	AccessGroupName string
}

func (c *NasClient) DescribeRegions(req *DescribeRegionsArgs) (resp *DescribeRegionsResponse, err error) {
	resp = &DescribeRegionsResponse{}
	err = c.Request("DescribeRegions", req, resp)
	return
}

type DescribeRegionsRegion struct {
	RegionId  string
	LocalName string
}
type DescribeRegionsArgs struct {
	PageSize   int
	PageNumber int
}
type DescribeRegionsResponse struct {
	RequestId  string
	TotalCount int
	PageSize   int
	PageNumber int
	Regions    DescribeRegionsRegionList
}

type DescribeRegionsRegionList []DescribeRegionsRegion

func (list *DescribeRegionsRegionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRegionsRegion)
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

func (c *NasClient) DeleteAccessRule(req *DeleteAccessRuleArgs) (resp *DeleteAccessRuleResponse, err error) {
	resp = &DeleteAccessRuleResponse{}
	err = c.Request("DeleteAccessRule", req, resp)
	return
}

type DeleteAccessRuleArgs struct {
	AccessGroupName string
	AccessRuleId    string
}
type DeleteAccessRuleResponse struct {
	RequestId string
}

func (c *NasClient) DeleteFileSystem(req *DeleteFileSystemArgs) (resp *DeleteFileSystemResponse, err error) {
	resp = &DeleteFileSystemResponse{}
	err = c.Request("DeleteFileSystem", req, resp)
	return
}

type DeleteFileSystemArgs struct {
	FileSystemId string
}
type DeleteFileSystemResponse struct {
	RequestId string
}

func (c *NasClient) DeleteAccessGroup(req *DeleteAccessGroupArgs) (resp *DeleteAccessGroupResponse, err error) {
	resp = &DeleteAccessGroupResponse{}
	err = c.Request("DeleteAccessGroup", req, resp)
	return
}

type DeleteAccessGroupArgs struct {
	AccessGroupName string
}
type DeleteAccessGroupResponse struct {
	RequestId string
}

func (c *NasClient) ModifyAccessGroup(req *ModifyAccessGroupArgs) (resp *ModifyAccessGroupResponse, err error) {
	resp = &ModifyAccessGroupResponse{}
	err = c.Request("ModifyAccessGroup", req, resp)
	return
}

type ModifyAccessGroupArgs struct {
	Description     string
	AccessGroupName string
}
type ModifyAccessGroupResponse struct {
	RequestId string
}
