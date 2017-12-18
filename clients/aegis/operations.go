package aegis

import "encoding/json"

func (c *AegisClient) RenewInstance(req *RenewInstanceArgs) (resp *RenewInstanceResponse, err error) {
	resp = &RenewInstanceResponse{}
	err = c.Request("RenewInstance", req, resp)
	return
}

type RenewInstanceArgs struct {
	Duration     int
	InstanceId   string
	ClientToken  string
	OwnerId      int64
	PricingCycle string
}
type RenewInstanceResponse struct {
	OrderId   string
	RequestId string
}

func (c *AegisClient) GetEntityList(req *GetEntityListArgs) (resp *GetEntityListResponse, err error) {
	resp = &GetEntityListResponse{}
	err = c.Request("GetEntityList", req, resp)
	return
}

type GetEntityListData struct {
	List     GetEntityListEntityList
	PageInfo GetEntityListPageInfo
}

type GetEntityListEntity struct {
	Uuid         string
	GroupId      int64
	Ip           string
	InstanceName string
	InstanceId   string
	Region       string
	Os           string
	Flag         string
	BuyVersion   string
	AegisOnline  bool
	AegisVersion string
}

type GetEntityListPageInfo struct {
	CurrentPage int
	PageSize    int
	TotalCount  int
	Count       int
}
type GetEntityListArgs struct {
	GroupId     int64
	PageSize    int
	Remark      string
	EventType   string
	CurrentPage int
	RegionNo    string
}
type GetEntityListResponse struct {
	RequestId string
	Code      string
	Success   bool
	Message   string
	Data      GetEntityListData
}

type GetEntityListEntityList []GetEntityListEntity

func (list *GetEntityListEntityList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetEntityListEntity)
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

func (c *AegisClient) GetAccountStatistics(req *GetAccountStatisticsArgs) (resp *GetAccountStatisticsResponse, err error) {
	resp = &GetAccountStatisticsResponse{}
	err = c.Request("GetAccountStatistics", req, resp)
	return
}

type GetAccountStatisticsData struct {
	RemoteLogin  int
	CrackSuccess int
}
type GetAccountStatisticsArgs struct {
	EndTime   string
	StartTime string
}
type GetAccountStatisticsResponse struct {
	RequestId string
	Code      string
	Success   bool
	Message   string
	Data      GetAccountStatisticsData
}

func (c *AegisClient) GetStatisticsByUuid(req *GetStatisticsByUuidArgs) (resp *GetStatisticsByUuidResponse, err error) {
	resp = &GetStatisticsByUuidResponse{}
	err = c.Request("GetStatisticsByUuid", req, resp)
	return
}

type GetStatisticsByUuidEntity struct {
	Uuid    string
	Account int
	Health  int
	Patch   int
	Trojan  int
	Online  bool
}
type GetStatisticsByUuidArgs struct {
	Uuid string
}
type GetStatisticsByUuidResponse struct {
	RequestId string
	Code      string
	Success   bool
	Message   string
	Data      GetStatisticsByUuidEntityList
}

type GetStatisticsByUuidEntityList []GetStatisticsByUuidEntity

func (list *GetStatisticsByUuidEntityList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetStatisticsByUuidEntity)
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

func (c *AegisClient) QueryLoginEvent(req *QueryLoginEventArgs) (resp *QueryLoginEventResponse, err error) {
	resp = &QueryLoginEventResponse{}
	err = c.Request("QueryLoginEvent", req, resp)
	return
}

type QueryLoginEventData struct {
	List     QueryLoginEventEntityList
	PageInfo QueryLoginEventPageInfo
}

type QueryLoginEventEntity struct {
	Uuid          string
	LoginTime     string
	LoginType     int
	LoginTypeName string
	BuyVersion    string
	LoginSourceIp string
	GroupId       int
	InstanceName  string
	InstanceId    string
	Ip            string
	Region        string
	Status        int
	StatusName    string
	Location      string
	UserName      string
}

type QueryLoginEventPageInfo struct {
	CurrentPage int
	PageSize    int
	TotalCount  int
	Count       int
}
type QueryLoginEventArgs struct {
	EndTime     string
	CurrentPage int
	StartTime   string
	Uuid        string
	Status      int
}
type QueryLoginEventResponse struct {
	RequestId string
	Code      string
	Success   bool
	Message   string
	Data      QueryLoginEventData
}

type QueryLoginEventEntityList []QueryLoginEventEntity

func (list *QueryLoginEventEntityList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryLoginEventEntity)
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

func (c *AegisClient) CreateInstance(req *CreateInstanceArgs) (resp *CreateInstanceResponse, err error) {
	resp = &CreateInstanceResponse{}
	err = c.Request("CreateInstance", req, resp)
	return
}

type CreateInstanceArgs struct {
	Duration          int
	IsAutoRenew       bool
	ClientToken       string
	VmNumber          int
	OwnerId           int64
	VersionCode       int
	PricingCycle      string
	AutoRenewDuration int
}
type CreateInstanceResponse struct {
	OrderId    string
	InstanceId string
	RequestId  string
}

func (c *AegisClient) QueryCrackEvent(req *QueryCrackEventArgs) (resp *QueryCrackEventResponse, err error) {
	resp = &QueryCrackEventResponse{}
	err = c.Request("QueryCrackEvent", req, resp)
	return
}

type QueryCrackEventData struct {
	List     QueryCrackEventEntityList
	PageInfo QueryCrackEventPageInfo
}

type QueryCrackEventEntity struct {
	Uuid           string
	AttackTime     string
	AttackType     int
	AttackTypeName string
	BuyVersion     string
	CrackSourceIp  string
	CrackTimes     int
	GroupId        int
	InstanceName   string
	InstanceId     string
	Ip             string
	Region         string
	Status         int
	StatusName     string
	Location       string
	InWhite        int
	UserName       string
}

type QueryCrackEventPageInfo struct {
	CurrentPage int
	PageSize    int
	TotalCount  int
	Count       int
}
type QueryCrackEventArgs struct {
	EndTime     string
	CurrentPage int
	StartTime   string
	Uuid        string
	Status      int
}
type QueryCrackEventResponse struct {
	RequestId string
	Code      string
	Success   bool
	Message   string
	Data      QueryCrackEventData
}

type QueryCrackEventEntityList []QueryCrackEventEntity

func (list *QueryCrackEventEntityList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCrackEventEntity)
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

func (c *AegisClient) GetStatistics(req *GetStatisticsArgs) (resp *GetStatisticsResponse, err error) {
	resp = &GetStatisticsResponse{}
	err = c.Request("GetStatistics", req, resp)
	return
}

type GetStatisticsData struct {
	Account int
	Health  int
	Patch   int
	Trojan  int
}
type GetStatisticsArgs struct {
	EndTime   string
	StartTime string
}
type GetStatisticsResponse struct {
	RequestId string
	Code      string
	Success   bool
	Message   string
	Data      GetStatisticsData
}

func (c *AegisClient) ReleaseInstance(req *ReleaseInstanceArgs) (resp *ReleaseInstanceResponse, err error) {
	resp = &ReleaseInstanceResponse{}
	err = c.Request("ReleaseInstance", req, resp)
	return
}

type ReleaseInstanceArgs struct {
	InstanceId string
	OwnerId    int64
}
type ReleaseInstanceResponse struct {
	RequestId string
}

func (c *AegisClient) UpgradeInstance(req *UpgradeInstanceArgs) (resp *UpgradeInstanceResponse, err error) {
	resp = &UpgradeInstanceResponse{}
	err = c.Request("UpgradeInstance", req, resp)
	return
}

type UpgradeInstanceArgs struct {
	InstanceId  string
	ClientToken string
	VmNumber    int
	OwnerId     int64
	VersionCode int
}
type UpgradeInstanceResponse struct {
	OrderId   string
	RequestId string
}

func (c *AegisClient) GetCrackStatistics(req *GetCrackStatisticsArgs) (resp *GetCrackStatisticsResponse, err error) {
	resp = &GetCrackStatisticsResponse{}
	err = c.Request("GetCrackStatistics", req, resp)
	return
}

type GetCrackStatisticsData struct {
	Intercepted int
}
type GetCrackStatisticsArgs struct {
	EndTime   string
	StartTime string
}
type GetCrackStatisticsResponse struct {
	RequestId string
	Code      string
	Success   bool
	Message   string
	Data      GetCrackStatisticsData
}
