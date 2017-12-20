package drds

import (
	"encoding/json"

	"github.com/morlay/aliyun-go/core"
)

func (c *DrdsClient) DeleteDrdsDB(req *DeleteDrdsDBArgs) (resp *DeleteDrdsDBResponse, err error) {
	resp = &DeleteDrdsDBResponse{}
	err = c.Request("DeleteDrdsDB", req, resp)
	return
}

type DeleteDrdsDBArgs struct {
	DbName         string
	DrdsInstanceId string
}
type DeleteDrdsDBResponse struct {
	RequestId string
	Success   core.Bool
}

func (c *DrdsClient) DescribeDrdsInstances(req *DescribeDrdsInstancesArgs) (resp *DescribeDrdsInstancesResponse, err error) {
	resp = &DescribeDrdsInstancesResponse{}
	err = c.Request("DescribeDrdsInstances", req, resp)
	return
}

type DescribeDrdsInstancesInstance struct {
	DrdsInstanceId string
	Type           string
	RegionId       string
	ZoneId         string
	Description    string
	NetworkType    string
	Status         string
	CreateTime     int64
	Version        int64
	Vips           DescribeDrdsInstancesVipList
}

type DescribeDrdsInstancesVip struct {
	IP   string
	Port string
	Type string
}
type DescribeDrdsInstancesArgs struct {
	Type string
}
type DescribeDrdsInstancesResponse struct {
	RequestId string
	Success   core.Bool
	Data      DescribeDrdsInstancesInstanceList
}

type DescribeDrdsInstancesVipList []DescribeDrdsInstancesVip

func (list *DescribeDrdsInstancesVipList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDrdsInstancesVip)
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

type DescribeDrdsInstancesInstanceList []DescribeDrdsInstancesInstance

func (list *DescribeDrdsInstancesInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDrdsInstancesInstance)
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

func (c *DrdsClient) DeleteFailedDrdsDB(req *DeleteFailedDrdsDBArgs) (resp *DeleteFailedDrdsDBResponse, err error) {
	resp = &DeleteFailedDrdsDBResponse{}
	err = c.Request("DeleteFailedDrdsDB", req, resp)
	return
}

type DeleteFailedDrdsDBArgs struct {
	DbName         string
	DrdsInstanceId string
}
type DeleteFailedDrdsDBResponse struct {
	RequestId string
	Success   core.Bool
}

func (c *DrdsClient) RemoveDrdsInstance(req *RemoveDrdsInstanceArgs) (resp *RemoveDrdsInstanceResponse, err error) {
	resp = &RemoveDrdsInstanceResponse{}
	err = c.Request("RemoveDrdsInstance", req, resp)
	return
}

type RemoveDrdsInstanceArgs struct {
	DrdsInstanceId string
}
type RemoveDrdsInstanceResponse struct {
	RequestId string
	Success   core.Bool
}

func (c *DrdsClient) DescribeDrdsInstance(req *DescribeDrdsInstanceArgs) (resp *DescribeDrdsInstanceResponse, err error) {
	resp = &DescribeDrdsInstanceResponse{}
	err = c.Request("DescribeDrdsInstance", req, resp)
	return
}

type DescribeDrdsInstanceData struct {
	DrdsInstanceId string
	Type           string
	RegionId       string
	ZoneId         string
	Description    string
	NetworkType    string
	Status         string
	CreateTime     int64
	Version        int64
	Specification  string
	Vips           DescribeDrdsInstanceVipList
}

type DescribeDrdsInstanceVip struct {
	IP        string
	Port      string
	Type      string
	VpcId     string
	VswitchId string
}
type DescribeDrdsInstanceArgs struct {
	DrdsInstanceId string
}
type DescribeDrdsInstanceResponse struct {
	RequestId string
	Success   core.Bool
	Data      DescribeDrdsInstanceData
}

type DescribeDrdsInstanceVipList []DescribeDrdsInstanceVip

func (list *DescribeDrdsInstanceVipList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDrdsInstanceVip)
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

func (c *DrdsClient) ModifyDrdsInstanceDescription(req *ModifyDrdsInstanceDescriptionArgs) (resp *ModifyDrdsInstanceDescriptionResponse, err error) {
	resp = &ModifyDrdsInstanceDescriptionResponse{}
	err = c.Request("ModifyDrdsInstanceDescription", req, resp)
	return
}

type ModifyDrdsInstanceDescriptionArgs struct {
	Description    string
	DrdsInstanceId string
}
type ModifyDrdsInstanceDescriptionResponse struct {
	RequestId string
	Success   core.Bool
}

func (c *DrdsClient) DescribeDrdsInstanceNetInfoForInner(req *DescribeDrdsInstanceNetInfoForInnerArgs) (resp *DescribeDrdsInstanceNetInfoForInnerResponse, err error) {
	resp = &DescribeDrdsInstanceNetInfoForInnerResponse{}
	err = c.Request("DescribeDrdsInstanceNetInfoForInner", req, resp)
	return
}

type DescribeDrdsInstanceNetInfoForInnerNetInfo struct {
	IP       string
	Port     string
	Type     string
	IsForVpc core.Bool
}
type DescribeDrdsInstanceNetInfoForInnerArgs struct {
	DrdsInstanceId string
}
type DescribeDrdsInstanceNetInfoForInnerResponse struct {
	RequestId      string
	Success        core.Bool
	DrdsInstanceId string
	NetworkType    string
	NetInfos       DescribeDrdsInstanceNetInfoForInnerNetInfoList
}

type DescribeDrdsInstanceNetInfoForInnerNetInfoList []DescribeDrdsInstanceNetInfoForInnerNetInfo

func (list *DescribeDrdsInstanceNetInfoForInnerNetInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDrdsInstanceNetInfoForInnerNetInfo)
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

func (c *DrdsClient) DescribeDrdsDBs(req *DescribeDrdsDBsArgs) (resp *DescribeDrdsDBsResponse, err error) {
	resp = &DescribeDrdsDBsResponse{}
	err = c.Request("DescribeDrdsDBs", req, resp)
	return
}

type DescribeDrdsDBsDb struct {
	DbName     string
	Status     int
	CreateTime string
	Msg        string
	Mode       string
}
type DescribeDrdsDBsArgs struct {
	DrdsInstanceId string
}
type DescribeDrdsDBsResponse struct {
	RequestId string
	Success   core.Bool
	Data      DescribeDrdsDBsDbList
}

type DescribeDrdsDBsDbList []DescribeDrdsDBsDb

func (list *DescribeDrdsDBsDbList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDrdsDBsDb)
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

func (c *DrdsClient) ModifyDrdsDBPasswd(req *ModifyDrdsDBPasswdArgs) (resp *ModifyDrdsDBPasswdResponse, err error) {
	resp = &ModifyDrdsDBPasswdResponse{}
	err = c.Request("ModifyDrdsDBPasswd", req, resp)
	return
}

type ModifyDrdsDBPasswdArgs struct {
	NewPasswd      string
	DbName         string
	DrdsInstanceId string
}
type ModifyDrdsDBPasswdResponse struct {
	RequestId string
	Success   core.Bool
}

func (c *DrdsClient) DescribeCreateDrdsInstanceStatus(req *DescribeCreateDrdsInstanceStatusArgs) (resp *DescribeCreateDrdsInstanceStatusResponse, err error) {
	resp = &DescribeCreateDrdsInstanceStatusResponse{}
	err = c.Request("DescribeCreateDrdsInstanceStatus", req, resp)
	return
}

type DescribeCreateDrdsInstanceStatusData struct {
	Status string
}
type DescribeCreateDrdsInstanceStatusArgs struct {
	DrdsInstanceId string
}
type DescribeCreateDrdsInstanceStatusResponse struct {
	RequestId string
	Success   core.Bool
	Data      DescribeCreateDrdsInstanceStatusData
}

func (c *DrdsClient) DescribeRdsList(req *DescribeRdsListArgs) (resp *DescribeRdsListResponse, err error) {
	resp = &DescribeRdsListResponse{}
	err = c.Request("DescribeRdsList", req, resp)
	return
}

type DescribeRdsListRdsInstance struct {
	InstanceId       int
	InstanceName     string
	ConnectUrl       string
	Port             int
	InstanceStatus   string
	DbType           string
	ReadWeight       int
	ReadOnlyChildren DescribeRdsListChildList
}

type DescribeRdsListChild struct {
	InstanceId     string
	InstanceName   string
	ConnectUrl     string
	Port           int
	InstanceStatus string
	DbType         string
	ReadWeight     int
}
type DescribeRdsListArgs struct {
	DbName         string
	DrdsInstanceId string
}
type DescribeRdsListResponse struct {
	RequestId string
	Success   core.Bool
	Data      DescribeRdsListRdsInstanceList
}

type DescribeRdsListChildList []DescribeRdsListChild

func (list *DescribeRdsListChildList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRdsListChild)
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

type DescribeRdsListRdsInstanceList []DescribeRdsListRdsInstance

func (list *DescribeRdsListRdsInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRdsListRdsInstance)
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

func (c *DrdsClient) CreateDrdsInstance(req *CreateDrdsInstanceArgs) (resp *CreateDrdsInstanceResponse, err error) {
	resp = &CreateDrdsInstanceResponse{}
	err = c.Request("CreateDrdsInstance", req, resp)
	return
}

type CreateDrdsInstanceData struct {
	OrderId            int64
	DrdsInstanceIdList CreateDrdsInstanceDrdsInstanceIdListList
}
type CreateDrdsInstanceArgs struct {
	VswitchId      string
	Quantity       int
	InstanceSeries string
	VpcId          string
	Description    string
	ZoneId         string
	Specification  string
	Type           string
	PayType        string
}
type CreateDrdsInstanceResponse struct {
	RequestId string
	Success   core.Bool
	Data      CreateDrdsInstanceData
}

type CreateDrdsInstanceDrdsInstanceIdListList []string

func (list *CreateDrdsInstanceDrdsInstanceIdListList) UnmarshalJSON(data []byte) error {
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

func (c *DrdsClient) ModifyDrdsIpWhiteList(req *ModifyDrdsIpWhiteListArgs) (resp *ModifyDrdsIpWhiteListResponse, err error) {
	resp = &ModifyDrdsIpWhiteListResponse{}
	err = c.Request("ModifyDrdsIpWhiteList", req, resp)
	return
}

type ModifyDrdsIpWhiteListArgs struct {
	Mode           core.Bool
	DbName         string
	GroupAttribute string
	IpWhiteList    string
	DrdsInstanceId string
	GroupName      string
}
type ModifyDrdsIpWhiteListResponse struct {
	RequestId string
	Success   core.Bool
}

func (c *DrdsClient) DescribeDrdsDBIpWhiteList(req *DescribeDrdsDBIpWhiteListArgs) (resp *DescribeDrdsDBIpWhiteListResponse, err error) {
	resp = &DescribeDrdsDBIpWhiteListResponse{}
	err = c.Request("DescribeDrdsDBIpWhiteList", req, resp)
	return
}

type DescribeDrdsDBIpWhiteListData struct {
	IpWhiteList DescribeDrdsDBIpWhiteListIpWhiteListList
}
type DescribeDrdsDBIpWhiteListArgs struct {
	DbName         string
	DrdsInstanceId string
	GroupName      string
}
type DescribeDrdsDBIpWhiteListResponse struct {
	RequestId string
	Success   core.Bool
	Data      DescribeDrdsDBIpWhiteListData
}

type DescribeDrdsDBIpWhiteListIpWhiteListList []string

func (list *DescribeDrdsDBIpWhiteListIpWhiteListList) UnmarshalJSON(data []byte) error {
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

func (c *DrdsClient) CreateReadOnlyAccount(req *CreateReadOnlyAccountArgs) (resp *CreateReadOnlyAccountResponse, err error) {
	resp = &CreateReadOnlyAccountResponse{}
	err = c.Request("CreateReadOnlyAccount", req, resp)
	return
}

type CreateReadOnlyAccountData struct {
	DbName         string
	DrdsInstanceId string
	AccountName    string
}
type CreateReadOnlyAccountArgs struct {
	Password       string
	DbName         string
	DrdsInstanceId string
}
type CreateReadOnlyAccountResponse struct {
	RequestId string
	Success   core.Bool
	Data      CreateReadOnlyAccountData
}

func (c *DrdsClient) CreateDrdsDB(req *CreateDrdsDBArgs) (resp *CreateDrdsDBResponse, err error) {
	resp = &CreateDrdsDBResponse{}
	err = c.Request("CreateDrdsDB", req, resp)
	return
}

type CreateDrdsDBArgs struct {
	Encode         string
	Password       string
	DbName         string
	RdsInstances   string
	DrdsInstanceId string
}
type CreateDrdsDBResponse struct {
	RequestId string
	Success   core.Bool
}

func (c *DrdsClient) DescribeShardDbConnectionInfo(req *DescribeShardDbConnectionInfoArgs) (resp *DescribeShardDbConnectionInfoResponse, err error) {
	resp = &DescribeShardDbConnectionInfoResponse{}
	err = c.Request("DescribeShardDbConnectionInfo", req, resp)
	return
}

type DescribeShardDbConnectionInfoConnectionInfo struct {
	InstanceName               string
	InstanceUrl                string
	SubDbName                  string
	DbStatus                   string
	DbType                     string
	MinPoolSize                int
	MaxPoolSize                int
	IdleTimeOut                int
	BlockingTimeout            int
	ConnectionProperties       string
	PreparedStatementCacheSize int
	UserName                   string
}
type DescribeShardDbConnectionInfoArgs struct {
	DbName         string
	DrdsInstanceId string
	SubDbName      string
}
type DescribeShardDbConnectionInfoResponse struct {
	RequestId      string
	Success        core.Bool
	ConnectionInfo DescribeShardDbConnectionInfoConnectionInfo
}

func (c *DrdsClient) ModifyFullTableScan(req *ModifyFullTableScanArgs) (resp *ModifyFullTableScanResponse, err error) {
	resp = &ModifyFullTableScanResponse{}
	err = c.Request("ModifyFullTableScan", req, resp)
	return
}

type ModifyFullTableScanArgs struct {
	DbName         string
	TableNames     string
	DrdsInstanceId string
	FullTableScan  core.Bool
}
type ModifyFullTableScanResponse struct {
	RequestId string
	Success   core.Bool
}

func (c *DrdsClient) DescribeShardDBs(req *DescribeShardDBsArgs) (resp *DescribeShardDBsResponse, err error) {
	resp = &DescribeShardDBsResponse{}
	err = c.Request("DescribeShardDBs", req, resp)
	return
}

type DescribeShardDBsDbIntancePair struct {
	SubDbName    string
	InstanceName string
}
type DescribeShardDBsArgs struct {
	DbName         string
	DrdsInstanceId string
}
type DescribeShardDBsResponse struct {
	RequestId string
	Success   core.Bool
	Data      DescribeShardDBsDbIntancePairList
}

type DescribeShardDBsDbIntancePairList []DescribeShardDBsDbIntancePair

func (list *DescribeShardDBsDbIntancePairList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeShardDBsDbIntancePair)
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

func (c *DrdsClient) DescribeReadOnlyAccount(req *DescribeReadOnlyAccountArgs) (resp *DescribeReadOnlyAccountResponse, err error) {
	resp = &DescribeReadOnlyAccountResponse{}
	err = c.Request("DescribeReadOnlyAccount", req, resp)
	return
}

type DescribeReadOnlyAccountData struct {
	DbName         string
	DrdsInstanceId string
	AccountName    string
}
type DescribeReadOnlyAccountArgs struct {
	DbName         string
	DrdsInstanceId string
}
type DescribeReadOnlyAccountResponse struct {
	RequestId string
	Success   core.Bool
	Data      DescribeReadOnlyAccountData
}

func (c *DrdsClient) DescribeRegions(req *DescribeRegionsArgs) (resp *DescribeRegionsResponse, err error) {
	resp = &DescribeRegionsResponse{}
	err = c.Request("DescribeRegions", req, resp)
	return
}

type DescribeRegionsDrdsRegion struct {
	RegionId           string
	RegionName         string
	ZoneId             string
	ZoneName           string
	InstanceSeriesList DescribeRegionsInstanceSeriesList
}

type DescribeRegionsInstanceSeries struct {
	SeriesName string
	SpecList   DescribeRegionsSpecListList
}
type DescribeRegionsArgs struct {
}
type DescribeRegionsResponse struct {
	RequestId   string
	Success     core.Bool
	DrdsRegions DescribeRegionsDrdsRegionList
}

type DescribeRegionsInstanceSeriesList []DescribeRegionsInstanceSeries

func (list *DescribeRegionsInstanceSeriesList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRegionsInstanceSeries)
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

type DescribeRegionsSpecListList []string

func (list *DescribeRegionsSpecListList) UnmarshalJSON(data []byte) error {
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

type DescribeRegionsDrdsRegionList []DescribeRegionsDrdsRegion

func (list *DescribeRegionsDrdsRegionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRegionsDrdsRegion)
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

func (c *DrdsClient) DescribeDrdsDB(req *DescribeDrdsDBArgs) (resp *DescribeDrdsDBResponse, err error) {
	resp = &DescribeDrdsDBResponse{}
	err = c.Request("DescribeDrdsDB", req, resp)
	return
}

type DescribeDrdsDBData struct {
	DbName     string
	Status     int
	CreateTime string
	Msg        string
	Mode       string
}
type DescribeDrdsDBArgs struct {
	DbName         string
	DrdsInstanceId string
}
type DescribeDrdsDBResponse struct {
	RequestId string
	Success   core.Bool
	Data      DescribeDrdsDBData
}

func (c *DrdsClient) RemoveReadOnlyAccount(req *RemoveReadOnlyAccountArgs) (resp *RemoveReadOnlyAccountResponse, err error) {
	resp = &RemoveReadOnlyAccountResponse{}
	err = c.Request("RemoveReadOnlyAccount", req, resp)
	return
}

type RemoveReadOnlyAccountArgs struct {
	DbName         string
	AccountName    string
	DrdsInstanceId string
}
type RemoveReadOnlyAccountResponse struct {
	RequestId string
	Success   core.Bool
}

func (c *DrdsClient) ModifyReadOnlyAccountPassword(req *ModifyReadOnlyAccountPasswordArgs) (resp *ModifyReadOnlyAccountPasswordResponse, err error) {
	resp = &ModifyReadOnlyAccountPasswordResponse{}
	err = c.Request("ModifyReadOnlyAccountPassword", req, resp)
	return
}

type ModifyReadOnlyAccountPasswordArgs struct {
	NewPasswd      string
	DbName         string
	AccountName    string
	OriginPassword string
	DrdsInstanceId string
}
type ModifyReadOnlyAccountPasswordResponse struct {
	RequestId string
	Success   core.Bool
}

func (c *DrdsClient) ModifyRdsReadWeight(req *ModifyRdsReadWeightArgs) (resp *ModifyRdsReadWeightResponse, err error) {
	resp = &ModifyRdsReadWeightResponse{}
	err = c.Request("ModifyRdsReadWeight", req, resp)
	return
}

type ModifyRdsReadWeightArgs struct {
	InstanceNames  string
	DbName         string
	Weights        string
	DrdsInstanceId string
}
type ModifyRdsReadWeightResponse struct {
	RequestId string
	Success   core.Bool
}
