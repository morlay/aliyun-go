package ocs

import "encoding/json"

func (c *OcsClient) ModifyInstanceAttribute(req *ModifyInstanceAttributeArgs) (resp *ModifyInstanceAttributeResponse, err error) {
	resp = &ModifyInstanceAttributeResponse{}
	err = c.Request("ModifyInstanceAttribute", req, resp)
	return
}

type ModifyInstanceAttributeArgs struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	OwnerAccount         string
	OcsInstanceId        string
	OcsInstanceName      string
	OldPassword          string
	NewPassword          string
}
type ModifyInstanceAttributeResponse struct {
	RequestId string
}

func (c *OcsClient) DeactivateInstance(req *DeactivateInstanceArgs) (resp *DeactivateInstanceResponse, err error) {
	resp = &DeactivateInstanceResponse{}
	err = c.Request("DeactivateInstance", req, resp)
	return
}

type DeactivateInstanceArgs struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	OwnerAccount         string
	OcsInstanceId        string
}
type DeactivateInstanceResponse struct {
	RequestId string
}

func (c *OcsClient) ModifyInstanceCapacity(req *ModifyInstanceCapacityArgs) (resp *ModifyInstanceCapacityResponse, err error) {
	resp = &ModifyInstanceCapacityResponse{}
	err = c.Request("ModifyInstanceCapacity", req, resp)
	return
}

type ModifyInstanceCapacityArgs struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	OwnerAccount         string
	OcsInstanceId        string
	Capacity             int64
}
type ModifyInstanceCapacityResponse struct {
	RequestId string
}

func (c *OcsClient) ModifySecurityIps(req *ModifySecurityIpsArgs) (resp *ModifySecurityIpsResponse, err error) {
	resp = &ModifySecurityIpsResponse{}
	err = c.Request("ModifySecurityIps", req, resp)
	return
}

type ModifySecurityIpsArgs struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	OwnerAccount         string
	OcsInstanceId        string
	SecurityIps          string
}
type ModifySecurityIpsResponse struct {
	RequestId string
}

func (c *OcsClient) DescribeSecurityIps(req *DescribeSecurityIpsArgs) (resp *DescribeSecurityIpsResponse, err error) {
	resp = &DescribeSecurityIpsResponse{}
	err = c.Request("DescribeSecurityIps", req, resp)
	return
}

type DescribeSecurityIpsDescribeOcsSecurityIpsResponseDTO struct {
	SecurityIps string
}
type DescribeSecurityIpsArgs struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	OwnerAccount         string
	OcsInstanceId        string
}
type DescribeSecurityIpsResponse struct {
	RequestId                         string
	DescribeOcsSecurityIpsResponseDTO DescribeSecurityIpsDescribeOcsSecurityIpsResponseDTO
}

func (c *OcsClient) DescribeMonitorItems(req *DescribeMonitorItemsArgs) (resp *DescribeMonitorItemsResponse, err error) {
	resp = &DescribeMonitorItemsResponse{}
	err = c.Request("DescribeMonitorItems", req, resp)
	return
}

type DescribeMonitorItemsGetOcsMonitorItemsResponseDTO struct {
	MonitorItems DescribeMonitorItemsMonitorItemList
}

type DescribeMonitorItemsMonitorItem struct {
	MonitorKey string
	Unit       string
}
type DescribeMonitorItemsArgs struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	OwnerAccount         string
	OcsInstanceId        string
}
type DescribeMonitorItemsResponse struct {
	RequestId                     string
	GetOcsMonitorItemsResponseDTO DescribeMonitorItemsGetOcsMonitorItemsResponseDTO
}

type DescribeMonitorItemsMonitorItemList []DescribeMonitorItemsMonitorItem

func (list *DescribeMonitorItemsMonitorItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMonitorItemsMonitorItem)
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

func (c *OcsClient) DescribeMonitorValues(req *DescribeMonitorValuesArgs) (resp *DescribeMonitorValuesResponse, err error) {
	resp = &DescribeMonitorValuesResponse{}
	err = c.Request("DescribeMonitorValues", req, resp)
	return
}

type DescribeMonitorValuesGetOcsMonitorValuesResponse struct {
	Date           string
	OcsInstanceIds DescribeMonitorValuesOcsMonitorDTOList
}

type DescribeMonitorValuesOcsMonitorDTO struct {
	OcsInstanceId string
	MonitorKeys   DescribeMonitorValuesOcsMonitorKeyDTOList
}

type DescribeMonitorValuesOcsMonitorKeyDTO struct {
	MonitorKey string
	Value      string
	Unit       string
	Date       string
}
type DescribeMonitorValuesArgs struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	OwnerAccount         string
	OcsInstanceId        string
	MonitorKeys          string
}
type DescribeMonitorValuesResponse struct {
	RequestId                   string
	GetOcsMonitorValuesResponse DescribeMonitorValuesGetOcsMonitorValuesResponse
}

type DescribeMonitorValuesOcsMonitorDTOList []DescribeMonitorValuesOcsMonitorDTO

func (list *DescribeMonitorValuesOcsMonitorDTOList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMonitorValuesOcsMonitorDTO)
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

type DescribeMonitorValuesOcsMonitorKeyDTOList []DescribeMonitorValuesOcsMonitorKeyDTO

func (list *DescribeMonitorValuesOcsMonitorKeyDTOList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMonitorValuesOcsMonitorKeyDTO)
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

func (c *OcsClient) DescribeInstances(req *DescribeInstancesArgs) (resp *DescribeInstancesResponse, err error) {
	resp = &DescribeInstancesResponse{}
	err = c.Request("DescribeInstances", req, resp)
	return
}

type DescribeInstancesGetOcsInstancesResponse struct {
	Total        int
	PageNo       int
	PageSize     int
	OcsInstances DescribeInstancesOcsInstanceList
}

type DescribeInstancesOcsInstance struct {
	OcsInstanceId     string
	OcsInstanceName   string
	Capacity          int64
	Qps               int64
	Bandwidth         int64
	Connections       int64
	ConnectionDomain  string
	Port              int
	UserName          string
	RegionId          string
	OcsInstanceStatus string
	GmtCreated        string
	EndTime           string
	ChargeType        string
	IzId              string
	NetworkType       string
	VpcId             string
	VSwitchId         string
	PrivateIp         string
}
type DescribeInstancesArgs struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	OwnerAccount         string
	OcsInstanceId        string
	OcsInstanceStatus    string
	PageNo               int
	PageSize             int
	NetworkType          string
	VpcId                string
	VSwitchId            string
	PrivateIps           string
}
type DescribeInstancesResponse struct {
	RequestId               string
	GetOcsInstancesResponse DescribeInstancesGetOcsInstancesResponse
}

type DescribeInstancesOcsInstanceList []DescribeInstancesOcsInstance

func (list *DescribeInstancesOcsInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesOcsInstance)
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

func (c *OcsClient) ReplaceAuthenticIP(req *ReplaceAuthenticIPArgs) (resp *ReplaceAuthenticIPResponse, err error) {
	resp = &ReplaceAuthenticIPResponse{}
	err = c.Request("ReplaceAuthenticIP", req, resp)
	return
}

type ReplaceAuthenticIPArgs struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	OwnerAccount         string
	OcsInstanceId        string
	OldAuthenticIP       string
	NewAuthenticIP       string
}
type ReplaceAuthenticIPResponse struct {
	RequestId string
}

func (c *OcsClient) RemoveAuthenticIP(req *RemoveAuthenticIPArgs) (resp *RemoveAuthenticIPResponse, err error) {
	resp = &RemoveAuthenticIPResponse{}
	err = c.Request("RemoveAuthenticIP", req, resp)
	return
}

type RemoveAuthenticIPArgs struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	OwnerAccount         string
	OcsInstanceId        string
	AuthenticIP          string
}
type RemoveAuthenticIPResponse struct {
	RequestId string
}

func (c *OcsClient) FlushInstance(req *FlushInstanceArgs) (resp *FlushInstanceResponse, err error) {
	resp = &FlushInstanceResponse{}
	err = c.Request("FlushInstance", req, resp)
	return
}

type FlushInstanceArgs struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	OwnerAccount         string
	OcsInstanceId        string
}
type FlushInstanceResponse struct {
	RequestId string
}

func (c *OcsClient) CreateInstance(req *CreateInstanceArgs) (resp *CreateInstanceResponse, err error) {
	resp = &CreateInstanceResponse{}
	err = c.Request("CreateInstance", req, resp)
	return
}

type CreateInstanceOcsInstance struct {
	OcsInstanceId     string
	OcsInstanceName   string
	Capacity          int64
	Qps               int64
	Bandwidth         int64
	Connections       int64
	ConnectionDomain  string
	Port              int
	UserName          string
	RegionId          string
	OcsInstanceStatus string
	GmtCreated        string
	EndTime           string
	ChargeType        string
	IzId              string
	NetworkType       string
	VpcId             string
	VSwitchId         string
	PrivateIp         string
}
type CreateInstanceArgs struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	OwnerAccount         string
	OcsInstanceName      string
	Password             string
	Capacity             int64
	Region               string
	IzNo                 string
	NetworkType          string
	VpcId                string
	VSwitchId            string
	PrivateIp            string
}
type CreateInstanceResponse struct {
	RequestId   string
	OcsInstance CreateInstanceOcsInstance
}

func (c *OcsClient) AddAuthenticIP(req *AddAuthenticIPArgs) (resp *AddAuthenticIPResponse, err error) {
	resp = &AddAuthenticIPResponse{}
	err = c.Request("AddAuthenticIP", req, resp)
	return
}

type AddAuthenticIPArgs struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	OwnerAccount         string
	OcsInstanceId        string
	AuthenticIP          string
}
type AddAuthenticIPResponse struct {
	RequestId string
}

func (c *OcsClient) DescribeAuthenticIP(req *DescribeAuthenticIPArgs) (resp *DescribeAuthenticIPResponse, err error) {
	resp = &DescribeAuthenticIPResponse{}
	err = c.Request("DescribeAuthenticIP", req, resp)
	return
}

type DescribeAuthenticIPGetAuthenticIpResponse struct {
	AuthenticIPs string
}
type DescribeAuthenticIPArgs struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	OwnerAccount         string
	OcsInstanceId        string
}
type DescribeAuthenticIPResponse struct {
	RequestId              string
	GetAuthenticIpResponse DescribeAuthenticIPGetAuthenticIpResponse
}

func (c *OcsClient) DeleteInstance(req *DeleteInstanceArgs) (resp *DeleteInstanceResponse, err error) {
	resp = &DeleteInstanceResponse{}
	err = c.Request("DeleteInstance", req, resp)
	return
}

type DeleteInstanceArgs struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	OwnerAccount         string
	OcsInstanceId        string
}
type DeleteInstanceResponse struct {
	RequestId string
}

func (c *OcsClient) DescribeHistoryMonitorValues(req *DescribeHistoryMonitorValuesArgs) (resp *DescribeHistoryMonitorValuesResponse, err error) {
	resp = &DescribeHistoryMonitorValuesResponse{}
	err = c.Request("DescribeHistoryMonitorValues", req, resp)
	return
}

type DescribeHistoryMonitorValuesArgs struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	OwnerAccount         string
	OcsInstanceId        string
	StartTime            string
	EndTime              string
	IntervalForHistory   string
}
type DescribeHistoryMonitorValuesResponse struct {
	RequestId      string
	MonitorHistory string
}

func (c *OcsClient) ActivateInstance(req *ActivateInstanceArgs) (resp *ActivateInstanceResponse, err error) {
	resp = &ActivateInstanceResponse{}
	err = c.Request("ActivateInstance", req, resp)
	return
}

type ActivateInstanceArgs struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	OwnerAccount         string
	OcsInstanceId        string
}
type ActivateInstanceResponse struct {
	RequestId string
}

func (c *OcsClient) DescribeRegions(req *DescribeRegionsArgs) (resp *DescribeRegionsResponse, err error) {
	resp = &DescribeRegionsResponse{}
	err = c.Request("DescribeRegions", req, resp)
	return
}

type DescribeRegionsGetDescribeRegionsResponse struct {
	Total    int
	PageNo   int
	PageSize int
	Regions  DescribeRegionsRegionList
}

type DescribeRegionsRegion struct {
	RegionId  string
	IzNumber  string
	LocalName string
}
type DescribeRegionsArgs struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	OwnerAccount         string
	OcsInstanceId        string
}
type DescribeRegionsResponse struct {
	RequestId                  string
	GetDescribeRegionsResponse DescribeRegionsGetDescribeRegionsResponse
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
