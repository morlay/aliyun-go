package cloudwf

import "encoding/json"

func (c *CloudwfClient) ListApStaStatus(req *ListApStaStatusArgs) (resp *ListApStaStatusResponse, err error) {
	resp = &ListApStaStatusResponse{}
	err = c.Request("ListApStaStatus", req, resp)
	return
}

type ListApStaStatusArgs struct {
	OrderCol       string
	SearchProtocal string
	SearchSsid     string
	SearchIp       string
	Length         int
	SearchUsername string
	SearchMac      string
	PageIndex      int
	Id             int64
	OrderDir       string
}
type ListApStaStatusResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ListApPosition(req *ListApPositionArgs) (resp *ListApPositionResponse, err error) {
	resp = &ListApPositionResponse{}
	err = c.Request("ListApPosition", req, resp)
	return
}

type ListApPositionArgs struct {
	MapId int64
}
type ListApPositionResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) RepairGroupAp(req *RepairGroupApArgs) (resp *RepairGroupApResponse, err error) {
	resp = &RepairGroupApResponse{}
	err = c.Request("RepairGroupAp", req, resp)
	return
}

type RepairGroupApArgs struct {
	Id int64
}
type RepairGroupApResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetSubAccountStatus(req *GetSubAccountStatusArgs) (resp *GetSubAccountStatusResponse, err error) {
	resp = &GetSubAccountStatusResponse{}
	err = c.Request("GetSubAccountStatus", req, resp)
	return
}

type GetSubAccountStatusArgs struct {
}
type GetSubAccountStatusResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ListStaStatus(req *ListStaStatusArgs) (resp *ListStaStatusResponse, err error) {
	resp = &ListStaStatusResponse{}
	err = c.Request("ListStaStatus", req, resp)
	return
}

type ListStaStatusArgs struct {
	OrderCol          string
	SearchGroupName   string
	SearchStatus      int
	Length            int
	SearchUsername    string
	OrderDir          string
	SearchProtocal    string
	SearchSsid        string
	SearchApName      string
	SearchIp          string
	PageIndex         int
	SearchMac         string
	SearchDescription string
}
type ListStaStatusResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ListStaOnoffLog(req *ListStaOnoffLogArgs) (resp *ListStaOnoffLogResponse, err error) {
	resp = &ListStaOnoffLogResponse{}
	err = c.Request("ListStaOnoffLog", req, resp)
	return
}

type ListStaOnoffLogArgs struct {
	OrderCol       string
	SearchSsid     string
	SearchApName   string
	Length         int
	SearchUsername string
	PageIndex      int
	Id             int64
	OrderDir       string
}
type ListStaOnoffLogResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) AreaUpdate(req *AreaUpdateArgs) (resp *AreaUpdateResponse, err error) {
	resp = &AreaUpdateResponse{}
	err = c.Request("AreaUpdate", req, resp)
	return
}

type AreaUpdateArgs struct {
	Name string
	Dids string
	Aid  int64
	Sid  int64
}
type AreaUpdateResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ProfileTrade(req *ProfileTradeArgs) (resp *ProfileTradeResponse, err error) {
	resp = &ProfileTradeResponse{}
	err = c.Request("ProfileTrade", req, resp)
	return
}

type ProfileTradeArgs struct {
	Gsid int64
}
type ProfileTradeResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) DeviceBatchCreate(req *DeviceBatchCreateArgs) (resp *DeviceBatchCreateResponse, err error) {
	resp = &DeviceBatchCreateResponse{}
	err = c.Request("DeviceBatchCreate", req, resp)
	return
}

type DeviceBatchCreateArgs struct {
	Sn         string
	DeviceType int
}
type DeviceBatchCreateResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) SaveStaStatus(req *SaveStaStatusArgs) (resp *SaveStaStatusResponse, err error) {
	resp = &SaveStaStatusResponse{}
	err = c.Request("SaveStaStatus", req, resp)
	return
}

type SaveStaStatusArgs struct {
	Description string
	Id          int64
}
type SaveStaStatusResponse struct {
	RequestId string
	Success   bool
	Message   string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) PutOssFile(req *PutOssFileArgs) (resp *PutOssFileResponse, err error) {
	resp = &PutOssFileResponse{}
	err = c.Request("PutOssFile", req, resp)
	return
}

type PutOssFileArgs struct {
	JsonData string
}
type PutOssFileResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) SaveApRadioSsidConfig(req *SaveApRadioSsidConfigArgs) (resp *SaveApRadioSsidConfigResponse, err error) {
	resp = &SaveApRadioSsidConfigResponse{}
	err = c.Request("SaveApRadioSsidConfig", req, resp)
	return
}

type SaveApRadioSsidConfigArgs struct {
	Hidden             int
	DynamicVlan        int
	Ssid               string
	Cir                int
	Mac                string
	Ieee80211w         int
	Network            int
	Isolate            int
	ApAssetId          int64
	EncKey             string
	MulticastForward   int
	Encryption         string
	Wmm                int
	AuthCache          int
	Disabled           int
	Id                 int64
	RadioIndex         int
	IgnoreWeakProbe    int
	Maxassoc           int
	DisassocLowAck     int
	DisassocWeakRssi   int
	SsidLb             int
	MaxInactivity      int
	VlanDhcp           int
	InstantlyEffective int
	ShortPreamble      int
}
type SaveApRadioSsidConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) BusinessUpdate(req *BusinessUpdateArgs) (resp *BusinessUpdateResponse, err error) {
	resp = &BusinessUpdateResponse{}
	err = c.Request("BusinessUpdate", req, resp)
	return
}

type BusinessUpdateArgs struct {
	Warn             int
	BusinessCity     string
	WarnEmail        string
	BusinessAddress  string
	Bid              int64
	BusinessManager  string
	BusinessProvince string
}
type BusinessUpdateResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) UserAnalyse(req *UserAnalyseArgs) (resp *UserAnalyseResponse, err error) {
	resp = &UserAnalyseResponse{}
	err = c.Request("UserAnalyse", req, resp)
	return
}

type UserAnalyseArgs struct {
	Gsid int64
}
type UserAnalyseResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) BatchChangeGroupApName(req *BatchChangeGroupApNameArgs) (resp *BatchChangeGroupApNameResponse, err error) {
	resp = &BatchChangeGroupApNameResponse{}
	err = c.Request("BatchChangeGroupApName", req, resp)
	return
}

type BatchChangeGroupApNameArgs struct {
	JsonData string
}
type BatchChangeGroupApNameResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetOnlineStaTimeSer(req *GetOnlineStaTimeSerArgs) (resp *GetOnlineStaTimeSerResponse, err error) {
	resp = &GetOnlineStaTimeSerResponse{}
	err = c.Request("GetOnlineStaTimeSer", req, resp)
	return
}

type GetOnlineStaTimeSerArgs struct {
	ZoomStart int64
	CompanyId int64
	ApgroupId int64
	Start     int64
	ZoomEnd   int64
	End       int64
}
type GetOnlineStaTimeSerResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ShopSetfiltermac(req *ShopSetfiltermacArgs) (resp *ShopSetfiltermacResponse, err error) {
	resp = &ShopSetfiltermacResponse{}
	err = c.Request("ShopSetfiltermac", req, resp)
	return
}

type ShopSetfiltermacArgs struct {
	Mac string
	Sid int64
}
type ShopSetfiltermacResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetBindAp4Umeng(req *GetBindAp4UmengArgs) (resp *GetBindAp4UmengResponse, err error) {
	resp = &GetBindAp4UmengResponse{}
	err = c.Request("GetBindAp4Umeng", req, resp)
	return
}

type GetBindAp4UmengArgs struct {
}
type GetBindAp4UmengResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ProfileMedia(req *ProfileMediaArgs) (resp *ProfileMediaResponse, err error) {
	resp = &ProfileMediaResponse{}
	err = c.Request("ProfileMedia", req, resp)
	return
}

type ProfileMediaArgs struct {
	Gsid int64
}
type ProfileMediaResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetStaDetailedStatus(req *GetStaDetailedStatusArgs) (resp *GetStaDetailedStatusResponse, err error) {
	resp = &GetStaDetailedStatusResponse{}
	err = c.Request("GetStaDetailedStatus", req, resp)
	return
}

type GetStaDetailedStatusArgs struct {
	Id int64
}
type GetStaDetailedStatusResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ReportHour(req *ReportHourArgs) (resp *ReportHourResponse, err error) {
	resp = &ReportHourResponse{}
	err = c.Request("ReportHour", req, resp)
	return
}

type ReportHourArgs struct {
	BeginDate string
	EndDate   string
	Agsid     int64
}
type ReportHourResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ReportRealtime(req *ReportRealtimeArgs) (resp *ReportRealtimeResponse, err error) {
	resp = &ReportRealtimeResponse{}
	err = c.Request("ReportRealtime", req, resp)
	return
}

type ReportRealtimeArgs struct {
	Agsid int64
}
type ReportRealtimeResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ShopCamera(req *ShopCameraArgs) (resp *ShopCameraResponse, err error) {
	resp = &ShopCameraResponse{}
	err = c.Request("ShopCamera", req, resp)
	return
}

type ShopCameraArgs struct {
	Gsid int64
}
type ShopCameraResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) AddApgroupConfig(req *AddApgroupConfigArgs) (resp *AddApgroupConfigResponse, err error) {
	resp = &AddApgroupConfigResponse{}
	err = c.Request("AddApgroupConfig", req, resp)
	return
}

type AddApgroupConfigArgs struct {
	ParentApgroupId int64
	Name            string
	Description     string
}
type AddApgroupConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ShopActionReturning(req *ShopActionReturningArgs) (resp *ShopActionReturningResponse, err error) {
	resp = &ShopActionReturningResponse{}
	err = c.Request("ShopActionReturning", req, resp)
	return
}

type ShopActionReturningArgs struct {
	Gsid int64
}
type ShopActionReturningResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) DeleteApgroupConfig(req *DeleteApgroupConfigArgs) (resp *DeleteApgroupConfigResponse, err error) {
	resp = &DeleteApgroupConfigResponse{}
	err = c.Request("DeleteApgroupConfig", req, resp)
	return
}

type DeleteApgroupConfigArgs struct {
	Id int64
}
type DeleteApgroupConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GroupIntime(req *GroupIntimeArgs) (resp *GroupIntimeResponse, err error) {
	resp = &GroupIntimeResponse{}
	err = c.Request("GroupIntime", req, resp)
	return
}

type GroupIntimeArgs struct {
	Gsid int64
}
type GroupIntimeResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ListApPositionMap(req *ListApPositionMapArgs) (resp *ListApPositionMapResponse, err error) {
	resp = &ListApPositionMapResponse{}
	err = c.Request("ListApPositionMap", req, resp)
	return
}

type ListApPositionMapArgs struct {
	OrderCol          string
	SearchName        string
	TotalItem         int
	Length            int
	MapType           int
	PageIndex         int
	SearchApgroupName string
	OrderDir          string
}
type ListApPositionMapResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ShopShowList(req *ShopShowListArgs) (resp *ShopShowListResponse, err error) {
	resp = &ShopShowListResponse{}
	err = c.Request("ShopShowList", req, resp)
	return
}

type ShopShowListArgs struct {
	Gid        int64
	Address    string
	Name       string
	Dirc       string
	Page       int
	Bid        int64
	Per        int
	ShopStatus int
}
type ShopShowListResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) OnoffGroupApRadio(req *OnoffGroupApRadioArgs) (resp *OnoffGroupApRadioResponse, err error) {
	resp = &OnoffGroupApRadioResponse{}
	err = c.Request("OnoffGroupApRadio", req, resp)
	return
}

type OnoffGroupApRadioArgs struct {
	ApgroupId int64
	Disabled  int
}
type OnoffGroupApRadioResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) SaveProbeDataSubscriber(req *SaveProbeDataSubscriberArgs) (resp *SaveProbeDataSubscriberResponse, err error) {
	resp = &SaveProbeDataSubscriberResponse{}
	err = c.Request("SaveProbeDataSubscriber", req, resp)
	return
}

type SaveProbeDataSubscriberArgs struct {
	ApiUrl         string
	ParamGenScript string
	Name           string
	HttpMethod     string
	Description    string
	Id             int64
	Type           int
	ResourceIdss   SaveProbeDataSubscriberResourceIdsList
}
type SaveProbeDataSubscriberResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

type SaveProbeDataSubscriberResourceIdsList []int64

func (list *SaveProbeDataSubscriberResourceIdsList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

func (c *CloudwfClient) GetStaRunHistoryTimeSer(req *GetStaRunHistoryTimeSerArgs) (resp *GetStaRunHistoryTimeSerResponse, err error) {
	resp = &GetStaRunHistoryTimeSerResponse{}
	err = c.Request("GetStaRunHistoryTimeSer", req, resp)
	return
}

type GetStaRunHistoryTimeSerArgs struct {
	Id int64
}
type GetStaRunHistoryTimeSerResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ShopGroupInfo(req *ShopGroupInfoArgs) (resp *ShopGroupInfoResponse, err error) {
	resp = &ShopGroupInfoResponse{}
	err = c.Request("ShopGroupInfo", req, resp)
	return
}

type ShopGroupInfoArgs struct {
	Gid int64
}
type ShopGroupInfoResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) SaveApgroupScanConfig(req *SaveApgroupScanConfigArgs) (resp *SaveApgroupScanConfigResponse, err error) {
	resp = &SaveApgroupScanConfigResponse{}
	err = c.Request("SaveApgroupScanConfig", req, resp)
	return
}

type SaveApgroupScanConfigArgs struct {
	JsonData  string
	ApgroupId int64
}
type SaveApgroupScanConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetProbeDataSubscriberConfig(req *GetProbeDataSubscriberConfigArgs) (resp *GetProbeDataSubscriberConfigResponse, err error) {
	resp = &GetProbeDataSubscriberConfigResponse{}
	err = c.Request("GetProbeDataSubscriberConfig", req, resp)
	return
}

type GetProbeDataSubscriberConfigArgs struct {
	Id int64
}
type GetProbeDataSubscriberConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetApgroupSsidConfigProgress(req *GetApgroupSsidConfigProgressArgs) (resp *GetApgroupSsidConfigProgressResponse, err error) {
	resp = &GetApgroupSsidConfigProgressResponse{}
	err = c.Request("GetApgroupSsidConfigProgress", req, resp)
	return
}

type GetApgroupSsidConfigProgressArgs struct {
	Id int64
}
type GetApgroupSsidConfigProgressResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) InnerCheckOrder(req *InnerCheckOrderArgs) (resp *InnerCheckOrderResponse, err error) {
	resp = &InnerCheckOrderResponse{}
	err = c.Request("InnerCheckOrder", req, resp)
	return
}

type InnerCheckOrderArgs struct {
	Data string
}
type InnerCheckOrderResponse struct {
	RequestId string
	Success   bool
	Message   string
	Code      string
	Data      string
}

func (c *CloudwfClient) ReportZoneHour(req *ReportZoneHourArgs) (resp *ReportZoneHourResponse, err error) {
	resp = &ReportZoneHourResponse{}
	err = c.Request("ReportZoneHour", req, resp)
	return
}

type ReportZoneHourArgs struct {
	BeginDate string
	EndDate   string
	Agsid     int64
}
type ReportZoneHourResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetGroupApRadioOnoffProgress(req *GetGroupApRadioOnoffProgressArgs) (resp *GetGroupApRadioOnoffProgressResponse, err error) {
	resp = &GetGroupApRadioOnoffProgressResponse{}
	err = c.Request("GetGroupApRadioOnoffProgress", req, resp)
	return
}

type GetGroupApRadioOnoffProgressArgs struct {
	Id int64
}
type GetGroupApRadioOnoffProgressResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetAccountConfig(req *GetAccountConfigArgs) (resp *GetAccountConfigResponse, err error) {
	resp = &GetAccountConfigResponse{}
	err = c.Request("GetAccountConfig", req, resp)
	return
}

type GetAccountConfigArgs struct {
	Id int64
}
type GetAccountConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetGroupApChangeNameTemplate(req *GetGroupApChangeNameTemplateArgs) (resp *GetGroupApChangeNameTemplateResponse, err error) {
	resp = &GetGroupApChangeNameTemplateResponse{}
	err = c.Request("GetGroupApChangeNameTemplate", req, resp)
	return
}

type GetGroupApChangeNameTemplateArgs struct {
}
type GetGroupApChangeNameTemplateResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetScanMode(req *GetScanModeArgs) (resp *GetScanModeResponse, err error) {
	resp = &GetScanModeResponse{}
	err = c.Request("GetScanMode", req, resp)
	return
}

type GetScanModeArgs struct {
	MacLists GetScanModeMacListList
}
type GetScanModeResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

type GetScanModeMacListList []string

func (list *GetScanModeMacListList) UnmarshalJSON(data []byte) error {
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

func (c *CloudwfClient) FindAp(req *FindApArgs) (resp *FindApResponse, err error) {
	resp = &FindApResponse{}
	err = c.Request("FindAp", req, resp)
	return
}

type FindApArgs struct {
	Id int64
}
type FindApResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) HeadquartersTrend(req *HeadquartersTrendArgs) (resp *HeadquartersTrendResponse, err error) {
	resp = &HeadquartersTrendResponse{}
	err = c.Request("HeadquartersTrend", req, resp)
	return
}

type HeadquartersTrendArgs struct {
	Bid int64
}
type HeadquartersTrendResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) SaveApConfig(req *SaveApConfigArgs) (resp *SaveApConfigResponse, err error) {
	resp = &SaveApConfigResponse{}
	err = c.Request("SaveApConfig", req, resp)
	return
}

type SaveApConfigArgs struct {
	Country     string
	ApAssetId   int64
	LogLevel    int
	Name        string
	EchoInt     int
	Scan        int
	Description string
	Id          int64
	Dai         string
	LogIp       string
	Mac         string
}
type SaveApConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetGroupApRadioConfigTemplate(req *GetGroupApRadioConfigTemplateArgs) (resp *GetGroupApRadioConfigTemplateResponse, err error) {
	resp = &GetGroupApRadioConfigTemplateResponse{}
	err = c.Request("GetGroupApRadioConfigTemplate", req, resp)
	return
}

type GetGroupApRadioConfigTemplateArgs struct {
}
type GetGroupApRadioConfigTemplateResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) CheckUmengDataAnalysisPermission(req *CheckUmengDataAnalysisPermissionArgs) (resp *CheckUmengDataAnalysisPermissionResponse, err error) {
	resp = &CheckUmengDataAnalysisPermissionResponse{}
	err = c.Request("CheckUmengDataAnalysisPermission", req, resp)
	return
}

type CheckUmengDataAnalysisPermissionArgs struct {
}
type CheckUmengDataAnalysisPermissionResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ReportDay(req *ReportDayArgs) (resp *ReportDayResponse, err error) {
	resp = &ReportDayResponse{}
	err = c.Request("ReportDay", req, resp)
	return
}

type ReportDayArgs struct {
	BeginDate string
	EndDate   string
	Agsid     int64
}
type ReportDayResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) AreaShowList(req *AreaShowListArgs) (resp *AreaShowListResponse, err error) {
	resp = &AreaShowListResponse{}
	err = c.Request("AreaShowList", req, resp)
	return
}

type AreaShowListArgs struct {
	Page int
	Per  int
	Sid  int64
}
type AreaShowListResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ProfileHistory(req *ProfileHistoryArgs) (resp *ProfileHistoryResponse, err error) {
	resp = &ProfileHistoryResponse{}
	err = c.Request("ProfileHistory", req, resp)
	return
}

type ProfileHistoryArgs struct {
	Idtype     int64
	EndMonth   string
	BeginMonth string
	Agsid      int64
}
type ProfileHistoryResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetApgroupScanConfigSaveProgress(req *GetApgroupScanConfigSaveProgressArgs) (resp *GetApgroupScanConfigSaveProgressResponse, err error) {
	resp = &GetApgroupScanConfigSaveProgressResponse{}
	err = c.Request("GetApgroupScanConfigSaveProgress", req, resp)
	return
}

type GetApgroupScanConfigSaveProgressArgs struct {
	Id int64
}
type GetApgroupScanConfigSaveProgressResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ResetApConfig(req *ResetApConfigArgs) (resp *ResetApConfigResponse, err error) {
	resp = &ResetApConfigResponse{}
	err = c.Request("ResetApConfig", req, resp)
	return
}

type ResetApConfigArgs struct {
	Id int64
}
type ResetApConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ProfileBase(req *ProfileBaseArgs) (resp *ProfileBaseResponse, err error) {
	resp = &ProfileBaseResponse{}
	err = c.Request("ProfileBase", req, resp)
	return
}

type ProfileBaseArgs struct {
	Gsid int64
}
type ProfileBaseResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetGroupApRepairProgress(req *GetGroupApRepairProgressArgs) (resp *GetGroupApRepairProgressResponse, err error) {
	resp = &GetGroupApRepairProgressResponse{}
	err = c.Request("GetGroupApRepairProgress", req, resp)
	return
}

type GetGroupApRepairProgressArgs struct {
	Id int64
}
type GetGroupApRepairProgressResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) InnerProduceCloudWF(req *InnerProduceCloudWFArgs) (resp *InnerProduceCloudWFResponse, err error) {
	resp = &InnerProduceCloudWFResponse{}
	err = c.Request("InnerProduceCloudWF", req, resp)
	return
}

type InnerProduceCloudWFArgs struct {
	Data string
}
type InnerProduceCloudWFResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
}

func (c *CloudwfClient) UpLoadMap(req *UpLoadMapArgs) (resp *UpLoadMapResponse, err error) {
	resp = &UpLoadMapResponse{}
	err = c.Request("UpLoadMap", req, resp)
	return
}

type UpLoadMapArgs struct {
	FileName   string
	UploadId   string
	ObjectName string
	ChunkIndex int
	ChunkCnt   int
}
type UpLoadMapResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) DeleteApRadioSsidConfig(req *DeleteApRadioSsidConfigArgs) (resp *DeleteApRadioSsidConfigResponse, err error) {
	resp = &DeleteApRadioSsidConfigResponse{}
	err = c.Request("DeleteApRadioSsidConfig", req, resp)
	return
}

type DeleteApRadioSsidConfigArgs struct {
	InstantlyEffective int
	Id                 int64
}
type DeleteApRadioSsidConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ShopInfo(req *ShopInfoArgs) (resp *ShopInfoResponse, err error) {
	resp = &ShopInfoResponse{}
	err = c.Request("ShopInfo", req, resp)
	return
}

type ShopInfoArgs struct {
	Sid int64
}
type ShopInfoResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ProfileDistrict(req *ProfileDistrictArgs) (resp *ProfileDistrictResponse, err error) {
	resp = &ProfileDistrictResponse{}
	err = c.Request("ProfileDistrict", req, resp)
	return
}

type ProfileDistrictArgs struct {
	Gsid int64
}
type ProfileDistrictResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ListApgroupConfig(req *ListApgroupConfigArgs) (resp *ListApgroupConfigResponse, err error) {
	resp = &ListApgroupConfigResponse{}
	err = c.Request("ListApgroupConfig", req, resp)
	return
}

type ListApgroupConfigArgs struct {
	OrderCol      string
	SearchName    string
	SearchCompany string
	Length        int
	PageIndex     int
	OrderDir      string
}
type ListApgroupConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ShopGetredress(req *ShopGetredressArgs) (resp *ShopGetredressResponse, err error) {
	resp = &ShopGetredressResponse{}
	err = c.Request("ShopGetredress", req, resp)
	return
}

type ShopGetredressArgs struct {
	Sid int64
}
type ShopGetredressResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) InnerRefund(req *InnerRefundArgs) (resp *InnerRefundResponse, err error) {
	resp = &InnerRefundResponse{}
	err = c.Request("InnerRefund", req, resp)
	return
}

type InnerRefundArgs struct {
	Data string
}
type InnerRefundResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
}

func (c *CloudwfClient) ListGroupApBriefConfig(req *ListGroupApBriefConfigArgs) (resp *ListGroupApBriefConfigResponse, err error) {
	resp = &ListGroupApBriefConfigResponse{}
	err = c.Request("ListGroupApBriefConfig", req, resp)
	return
}

type ListGroupApBriefConfigArgs struct {
	OrderCol   string
	SearchName string
	ApgroupId  int64
	ColCnt     int
	Length     int
	PageIndex  int
	SearchMac  string
	OrderDir   string
}
type ListGroupApBriefConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ListApPositionStatus(req *ListApPositionStatusArgs) (resp *ListApPositionStatusResponse, err error) {
	resp = &ListApPositionStatusResponse{}
	err = c.Request("ListApPositionStatus", req, resp)
	return
}

type ListApPositionStatusArgs struct {
	JsonData string
}
type ListApPositionStatusResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetOssServerSign(req *GetOssServerSignArgs) (resp *GetOssServerSignResponse, err error) {
	resp = &GetOssServerSignResponse{}
	err = c.Request("GetOssServerSign", req, resp)
	return
}

type GetOssServerSignArgs struct {
}
type GetOssServerSignResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) DelApPosition(req *DelApPositionArgs) (resp *DelApPositionResponse, err error) {
	resp = &DelApPositionResponse{}
	err = c.Request("DelApPosition", req, resp)
	return
}

type DelApPositionArgs struct {
	ApAssetId int64
	MapId     int64
}
type DelApPositionResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ShopActionCustome(req *ShopActionCustomeArgs) (resp *ShopActionCustomeResponse, err error) {
	resp = &ShopActionCustomeResponse{}
	err = c.Request("ShopActionCustome", req, resp)
	return
}

type ShopActionCustomeArgs struct {
	Gsid int64
}
type ShopActionCustomeResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ResetAp(req *ResetApArgs) (resp *ResetApResponse, err error) {
	resp = &ResetApResponse{}
	err = c.Request("ResetAp", req, resp)
	return
}

type ResetApArgs struct {
	Id int64
}
type ResetApResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) BusinessCreate(req *BusinessCreateArgs) (resp *BusinessCreateResponse, err error) {
	resp = &BusinessCreateResponse{}
	err = c.Request("BusinessCreate", req, resp)
	return
}

type BusinessCreateArgs struct {
	BusinessCity     string
	Combo            string
	WarnEmail        string
	BusinessManager  string
	BusinessType     int
	Warn             int
	BusinessName     string
	BusinessTopType  int
	BusinessAddress  string
	BusinessTel      string
	BusinessProvince string
	BusinessSubtype  int
}
type BusinessCreateResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) SaveApgroupConfig(req *SaveApgroupConfigArgs) (resp *SaveApgroupConfigResponse, err error) {
	resp = &SaveApgroupConfigResponse{}
	err = c.Request("SaveApgroupConfig", req, resp)
	return
}

type SaveApgroupConfigArgs struct {
	Country     string
	LogLevel    int
	Name        string
	EchoInt     int
	Scan        int
	Description string
	Id          int64
	Dai         string
	LogIp       string
}
type SaveApgroupConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetLatestStaStatistic(req *GetLatestStaStatisticArgs) (resp *GetLatestStaStatisticResponse, err error) {
	resp = &GetLatestStaStatisticResponse{}
	err = c.Request("GetLatestStaStatistic", req, resp)
	return
}

type GetLatestStaStatisticArgs struct {
	ApgroupId int64
}
type GetLatestStaStatisticResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) SaveGroupApRadioConfig(req *SaveGroupApRadioConfigArgs) (resp *SaveGroupApRadioConfigResponse, err error) {
	resp = &SaveGroupApRadioConfigResponse{}
	err = c.Request("SaveGroupApRadioConfig", req, resp)
	return
}

type SaveGroupApRadioConfigArgs struct {
	JsonData string
}
type SaveGroupApRadioConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ExcelToJson(req *ExcelToJsonArgs) (resp *ExcelToJsonResponse, err error) {
	resp = &ExcelToJsonResponse{}
	err = c.Request("ExcelToJson", req, resp)
	return
}

type ExcelToJsonArgs struct {
	UploadData string
}
type ExcelToJsonResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ShopDataAlarm(req *ShopDataAlarmArgs) (resp *ShopDataAlarmResponse, err error) {
	resp = &ShopDataAlarmResponse{}
	err = c.Request("ShopDataAlarm", req, resp)
	return
}

type ShopDataAlarmArgs struct {
	WarnPhone string
	Warn      int
	CloseWarn int
	WarnEmail string
	Sid       int64
}
type ShopDataAlarmResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) AreaCreate(req *AreaCreateArgs) (resp *AreaCreateResponse, err error) {
	resp = &AreaCreateResponse{}
	err = c.Request("AreaCreate", req, resp)
	return
}

type AreaCreateArgs struct {
	Name string
	Dids string
	Sid  int64
}
type AreaCreateResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetApgroupPortalConfigProgress(req *GetApgroupPortalConfigProgressArgs) (resp *GetApgroupPortalConfigProgressResponse, err error) {
	resp = &GetApgroupPortalConfigProgressResponse{}
	err = c.Request("GetApgroupPortalConfigProgress", req, resp)
	return
}

type GetApgroupPortalConfigProgressArgs struct {
	Id int64
}
type GetApgroupPortalConfigProgressResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) HeadquartersToolsContrast(req *HeadquartersToolsContrastArgs) (resp *HeadquartersToolsContrastResponse, err error) {
	resp = &HeadquartersToolsContrastResponse{}
	err = c.Request("HeadquartersToolsContrast", req, resp)
	return
}

type HeadquartersToolsContrastArgs struct {
	Bid int64
}
type HeadquartersToolsContrastResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetCrowdList(req *GetCrowdListArgs) (resp *GetCrowdListResponse, err error) {
	resp = &GetCrowdListResponse{}
	err = c.Request("GetCrowdList", req, resp)
	return
}

type GetCrowdListArgs struct {
	Gsid      int64
	ClassType int
	GsType    string
	EndTime   string
	Page      int
	StartTime string
	Per       int
}
type GetCrowdListResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) SaveApPortalConfig(req *SaveApPortalConfigArgs) (resp *SaveApPortalConfigResponse, err error) {
	resp = &SaveApPortalConfigResponse{}
	err = c.Request("SaveApPortalConfig", req, resp)
	return
}

type SaveApPortalConfigArgs struct {
	AuthKey      string
	PortalUrl    string
	PortalStatus bool
	Whitelist    string
	CheckUrl     string
	ApConfigId   int64
	AuthSecret   string
	WebAuthUrl   string
	Network      int
}
type SaveApPortalConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) SavePortalConfig(req *SavePortalConfigArgs) (resp *SavePortalConfigResponse, err error) {
	resp = &SavePortalConfigResponse{}
	err = c.Request("SavePortalConfig", req, resp)
	return
}

type SavePortalConfigArgs struct {
	JsonData string
}
type SavePortalConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ShopSetredress(req *ShopSetredressArgs) (resp *ShopSetredressResponse, err error) {
	resp = &ShopSetredressResponse{}
	err = c.Request("ShopSetredress", req, resp)
	return
}

type ShopSetredressArgs struct {
	Workday     string
	Filterclose int
	Minstoptime int
	Holiday     string
	Hnum        string
	Sid         int64
	Clerk       int
	Filterstate int
	Wnum        string
	State       int
	Crowdfixed  int
	Maxstoptime int
}
type ShopSetredressResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) AreaDelete(req *AreaDeleteArgs) (resp *AreaDeleteResponse, err error) {
	resp = &AreaDeleteResponse{}
	err = c.Request("AreaDelete", req, resp)
	return
}

type AreaDeleteArgs struct {
	Aid int64
	Sid int64
}
type AreaDeleteResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ListApAssetCanBeAdded(req *ListApAssetCanBeAddedArgs) (resp *ListApAssetCanBeAddedResponse, err error) {
	resp = &ListApAssetCanBeAddedResponse{}
	err = c.Request("ListApAssetCanBeAdded", req, resp)
	return
}

type ListApAssetCanBeAddedArgs struct {
	SearchName  string
	ApgroupId   int64
	Length      int
	PageIndex   int
	SearchMac   string
	SearchModel string
}
type ListApAssetCanBeAddedResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetBid(req *GetBidArgs) (resp *GetBidResponse, err error) {
	resp = &GetBidResponse{}
	err = c.Request("GetBid", req, resp)
	return
}

type GetBidArgs struct {
}
type GetBidResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GroupOverview(req *GroupOverviewArgs) (resp *GroupOverviewResponse, err error) {
	resp = &GroupOverviewResponse{}
	err = c.Request("GroupOverview", req, resp)
	return
}

type GroupOverviewArgs struct {
	Gsid int64
}
type GroupOverviewResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ProfileConsume(req *ProfileConsumeArgs) (resp *ProfileConsumeResponse, err error) {
	resp = &ProfileConsumeResponse{}
	err = c.Request("ProfileConsume", req, resp)
	return
}

type ProfileConsumeArgs struct {
	Gsid int64
}
type ProfileConsumeResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ShopUpdate(req *ShopUpdateArgs) (resp *ShopUpdateResponse, err error) {
	resp = &ShopUpdateResponse{}
	err = c.Request("ShopUpdate", req, resp)
	return
}

type ShopUpdateArgs struct {
	ShopCoordinate    string
	ShopProvince      string
	ShopTopType       int
	ShopAddress       string
	ShopType          int
	WarnEmail         string
	Sid               int64
	ShopTel           string
	WarnpHone         string
	Warn              int
	ShopArea          int
	ShopRemarks       string
	ShopCity          string
	ShopSubtype       int
	ShopBrand         string
	ShopName          string
	ShopCloseWarn     int
	ShopManager       string
	ShopBusinessHours string
}
type ShopUpdateResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) HeadquartersRanking(req *HeadquartersRankingArgs) (resp *HeadquartersRankingResponse, err error) {
	resp = &HeadquartersRankingResponse{}
	err = c.Request("HeadquartersRanking", req, resp)
	return
}

type HeadquartersRankingArgs struct {
	Bid int64
}
type HeadquartersRankingResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ShopMarketingList(req *ShopMarketingListArgs) (resp *ShopMarketingListResponse, err error) {
	resp = &ShopMarketingListResponse{}
	err = c.Request("ShopMarketingList", req, resp)
	return
}

type ShopMarketingListArgs struct {
	Name string
	Page int
	Per  int
	Sid  int64
}
type ShopMarketingListResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) UserDataDelete(req *UserDataDeleteArgs) (resp *UserDataDeleteResponse, err error) {
	resp = &UserDataDeleteResponse{}
	err = c.Request("UserDataDelete", req, resp)
	return
}

type UserDataDeleteArgs struct {
	Iid int64
	Bid int64
}
type UserDataDeleteResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetApRunHistoryTimeSer(req *GetApRunHistoryTimeSerArgs) (resp *GetApRunHistoryTimeSerResponse, err error) {
	resp = &GetApRunHistoryTimeSerResponse{}
	err = c.Request("GetApRunHistoryTimeSer", req, resp)
	return
}

type GetApRunHistoryTimeSerArgs struct {
	Start int64
	End   int64
	Id    int64
}
type GetApRunHistoryTimeSerResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) DeletePositionMap(req *DeletePositionMapArgs) (resp *DeletePositionMapResponse, err error) {
	resp = &DeletePositionMapResponse{}
	err = c.Request("DeletePositionMap", req, resp)
	return
}

type DeletePositionMapArgs struct {
	MapId int64
}
type DeletePositionMapResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) BusinessShowList(req *BusinessShowListArgs) (resp *BusinessShowListResponse, err error) {
	resp = &BusinessShowListResponse{}
	err = c.Request("BusinessShowList", req, resp)
	return
}

type BusinessShowListArgs struct {
	Page int
	Per  int
}
type BusinessShowListResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) DeviceDelete(req *DeviceDeleteArgs) (resp *DeviceDeleteResponse, err error) {
	resp = &DeviceDeleteResponse{}
	err = c.Request("DeviceDelete", req, resp)
	return
}

type DeviceDeleteArgs struct {
	Did int64
	Mac string
}
type DeviceDeleteResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetBatchSaveApAssetProgress(req *GetBatchSaveApAssetProgressArgs) (resp *GetBatchSaveApAssetProgressResponse, err error) {
	resp = &GetBatchSaveApAssetProgressResponse{}
	err = c.Request("GetBatchSaveApAssetProgress", req, resp)
	return
}

type GetBatchSaveApAssetProgressArgs struct {
}
type GetBatchSaveApAssetProgressResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ListBriefConfigByAction(req *ListBriefConfigByActionArgs) (resp *ListBriefConfigByActionResponse, err error) {
	resp = &ListBriefConfigByActionResponse{}
	err = c.Request("ListBriefConfigByAction", req, resp)
	return
}

type ListBriefConfigByActionArgs struct {
	AncestorApgroupId int64
	Limit             int
	FuzzySearch       string
}
type ListBriefConfigByActionResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) AliyunRegisterApAsset(req *AliyunRegisterApAssetArgs) (resp *AliyunRegisterApAssetResponse, err error) {
	resp = &AliyunRegisterApAssetResponse{}
	err = c.Request("AliyunRegisterApAsset", req, resp)
	return
}

type AliyunRegisterApAssetArgs struct {
	ApgroupId int64
	Mac       string
	SerialNo  string
}
type AliyunRegisterApAssetResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) UserDataCreate(req *UserDataCreateArgs) (resp *UserDataCreateResponse, err error) {
	resp = &UserDataCreateResponse{}
	err = c.Request("UserDataCreate", req, resp)
	return
}

type UserDataCreateArgs struct {
	UploadFile string
	Name       string
	Bid        int64
	Type       string
}
type UserDataCreateResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ShopDeletemarketing(req *ShopDeletemarketingArgs) (resp *ShopDeletemarketingResponse, err error) {
	resp = &ShopDeletemarketingResponse{}
	err = c.Request("ShopDeletemarketing", req, resp)
	return
}

type ShopDeletemarketingArgs struct {
	Id  int64
	Sid int64
}
type ShopDeletemarketingResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetOnlineApTimeSer(req *GetOnlineApTimeSerArgs) (resp *GetOnlineApTimeSerResponse, err error) {
	resp = &GetOnlineApTimeSerResponse{}
	err = c.Request("GetOnlineApTimeSer", req, resp)
	return
}

type GetOnlineApTimeSerArgs struct {
	ZoomStart int64
	CompanyId int64
	ApgroupId int64
	Start     int64
	ZoomEnd   int64
	End       int64
}
type GetOnlineApTimeSerResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetApgroupConfigProgress(req *GetApgroupConfigProgressArgs) (resp *GetApgroupConfigProgressResponse, err error) {
	resp = &GetApgroupConfigProgressResponse{}
	err = c.Request("GetApgroupConfigProgress", req, resp)
	return
}

type GetApgroupConfigProgressArgs struct {
	Id int64
}
type GetApgroupConfigProgressResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) BatchSaveApPosition(req *BatchSaveApPositionArgs) (resp *BatchSaveApPositionResponse, err error) {
	resp = &BatchSaveApPositionResponse{}
	err = c.Request("BatchSaveApPosition", req, resp)
	return
}

type BatchSaveApPositionArgs struct {
	JsonData string
}
type BatchSaveApPositionResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) FrequencyAnalyse(req *FrequencyAnalyseArgs) (resp *FrequencyAnalyseResponse, err error) {
	resp = &FrequencyAnalyseResponse{}
	err = c.Request("FrequencyAnalyse", req, resp)
	return
}

type FrequencyAnalyseArgs struct {
	Gsid int64
}
type FrequencyAnalyseResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) DeviceShowList(req *DeviceShowListArgs) (resp *DeviceShowListResponse, err error) {
	resp = &DeviceShowListResponse{}
	err = c.Request("DeviceShowList", req, resp)
	return
}

type DeviceShowListArgs struct {
	Dirc       string
	Page       int
	Per        int
	DeviceType int
	Sid        int64
}
type DeviceShowListResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) DeviceCreate(req *DeviceCreateArgs) (resp *DeviceCreateResponse, err error) {
	resp = &DeviceCreateResponse{}
	err = c.Request("DeviceCreate", req, resp)
	return
}

type DeviceCreateArgs struct {
	DeviceNum      string
	DevicePosition string
	DeviceName     string
	DeviceType     int
	Sid            int64
}
type DeviceCreateResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ListApStatus(req *ListApStatusArgs) (resp *ListApStatusResponse, err error) {
	resp = &ListApStatusResponse{}
	err = c.Request("ListApStatus", req, resp)
	return
}

type ListApStatusArgs struct {
	OrderCol          string
	SearchName        string
	SearchGroupName   string
	SearchStatus      int
	SearchWanIp       string
	SearchApModelName string
	Length            int
	OrderDir          string
	SearchBssEquals   int
	SearchSwVersion   int64
	SearchCompanyName string
	SearchMac         string
	PageIndex         int
}
type ListApStatusResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) KickSta(req *KickStaArgs) (resp *KickStaResponse, err error) {
	resp = &KickStaResponse{}
	err = c.Request("KickSta", req, resp)
	return
}

type KickStaArgs struct {
	Id int64
}
type KickStaResponse struct {
	RequestId string
	Success   bool
	Message   string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ApgroupBatchDeleteAp(req *ApgroupBatchDeleteApArgs) (resp *ApgroupBatchDeleteApResponse, err error) {
	resp = &ApgroupBatchDeleteApResponse{}
	err = c.Request("ApgroupBatchDeleteAp", req, resp)
	return
}

type ApgroupBatchDeleteApArgs struct {
	ApAssetIds string
}
type ApgroupBatchDeleteApResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) RepairAp(req *RepairApArgs) (resp *RepairApResponse, err error) {
	resp = &RepairApResponse{}
	err = c.Request("RepairAp", req, resp)
	return
}

type RepairApArgs struct {
	Id int64
}
type RepairApResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) SaveApRadioConfig(req *SaveApRadioConfigArgs) (resp *SaveApRadioConfigResponse, err error) {
	resp = &SaveApRadioConfigResponse{}
	err = c.Request("SaveApRadioConfig", req, resp)
	return
}

type SaveApRadioConfigArgs struct {
	RequireMode        string
	Htmode             string
	Frag               int
	Minrate            int
	Probereq           int
	Channel            int
	Shortgi            int
	Hwmode             string
	Uapsd              int
	BeaconInt          int
	Mac                string
	Rts                int
	Txpower            int
	Noscan             int
	BcastRate          int
	Disabled           int
	InstantlyEffective int
	Id                 int64
	RadioIndex         int
}
type SaveApRadioConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetRadioRunHistoryTimeSer(req *GetRadioRunHistoryTimeSerArgs) (resp *GetRadioRunHistoryTimeSerResponse, err error) {
	resp = &GetRadioRunHistoryTimeSerResponse{}
	err = c.Request("GetRadioRunHistoryTimeSer", req, resp)
	return
}

type GetRadioRunHistoryTimeSerArgs struct {
	Id int64
}
type GetRadioRunHistoryTimeSerResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) DeleteApgroupSsidConfig(req *DeleteApgroupSsidConfigArgs) (resp *DeleteApgroupSsidConfigResponse, err error) {
	resp = &DeleteApgroupSsidConfigResponse{}
	err = c.Request("DeleteApgroupSsidConfig", req, resp)
	return
}

type DeleteApgroupSsidConfigArgs struct {
	ApgroupId int64
	Id        int64
}
type DeleteApgroupSsidConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetLatestApStatistic(req *GetLatestApStatisticArgs) (resp *GetLatestApStatisticResponse, err error) {
	resp = &GetLatestApStatisticResponse{}
	err = c.Request("GetLatestApStatistic", req, resp)
	return
}

type GetLatestApStatisticArgs struct {
	ApgroupId int64
}
type GetLatestApStatisticResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetApDetailedStatus(req *GetApDetailedStatusArgs) (resp *GetApDetailedStatusResponse, err error) {
	resp = &GetApDetailedStatusResponse{}
	err = c.Request("GetApDetailedStatus", req, resp)
	return
}

type GetApDetailedStatusArgs struct {
	Id int64
}
type GetApDetailedStatusResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ShopGetfiltermac(req *ShopGetfiltermacArgs) (resp *ShopGetfiltermacResponse, err error) {
	resp = &ShopGetfiltermacResponse{}
	err = c.Request("ShopGetfiltermac", req, resp)
	return
}

type ShopGetfiltermacArgs struct {
	Sid int64
}
type ShopGetfiltermacResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ShopGroupUpdate(req *ShopGroupUpdateArgs) (resp *ShopGroupUpdateResponse, err error) {
	resp = &ShopGroupUpdateResponse{}
	err = c.Request("ShopGroupUpdate", req, resp)
	return
}

type ShopGroupUpdateArgs struct {
	Gid         int64
	ShopIds     string
	Name        string
	Description string
}
type ShopGroupUpdateResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ReportZoneDay(req *ReportZoneDayArgs) (resp *ReportZoneDayResponse, err error) {
	resp = &ReportZoneDayResponse{}
	err = c.Request("ReportZoneDay", req, resp)
	return
}

type ReportZoneDayArgs struct {
	BeginDate string
	EndDate   string
	Agsid     int64
}
type ReportZoneDayResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ReportZoneMinute(req *ReportZoneMinuteArgs) (resp *ReportZoneMinuteResponse, err error) {
	resp = &ReportZoneMinuteResponse{}
	err = c.Request("ReportZoneMinute", req, resp)
	return
}

type ReportZoneMinuteArgs struct {
	BeginDate string
	EndDate   string
	Agsid     int64
}
type ReportZoneMinuteResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ApgroupBatchAddAp(req *ApgroupBatchAddApArgs) (resp *ApgroupBatchAddApResponse, err error) {
	resp = &ApgroupBatchAddApResponse{}
	err = c.Request("ApgroupBatchAddAp", req, resp)
	return
}

type ApgroupBatchAddApArgs struct {
	ApAssetIds string
	ApgroupId  int64
}
type ApgroupBatchAddApResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetApgroupDetailedConfig(req *GetApgroupDetailedConfigArgs) (resp *GetApgroupDetailedConfigResponse, err error) {
	resp = &GetApgroupDetailedConfigResponse{}
	err = c.Request("GetApgroupDetailedConfig", req, resp)
	return
}

type GetApgroupDetailedConfigArgs struct {
	Id int64
}
type GetApgroupDetailedConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ListBriefApConfig(req *ListBriefApConfigArgs) (resp *ListBriefApConfigResponse, err error) {
	resp = &ListBriefApConfigResponse{}
	err = c.Request("ListBriefApConfig", req, resp)
	return
}

type ListBriefApConfigArgs struct {
	SearchScan  int
	OrderCol    string
	SearchName  string
	Length      int
	SearchMac   string
	PageIndex   int
	OrderDir    string
	SearchModel string
}
type ListBriefApConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) HeadquartersOverview(req *HeadquartersOverviewArgs) (resp *HeadquartersOverviewResponse, err error) {
	resp = &HeadquartersOverviewResponse{}
	err = c.Request("HeadquartersOverview", req, resp)
	return
}

type HeadquartersOverviewArgs struct {
	Bid int64
}
type HeadquartersOverviewResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GroupDetails(req *GroupDetailsArgs) (resp *GroupDetailsResponse, err error) {
	resp = &GroupDetailsResponse{}
	err = c.Request("GroupDetails", req, resp)
	return
}

type GroupDetailsArgs struct {
	Gsid int64
}
type GroupDetailsResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) HeadquartersToolsCoincide(req *HeadquartersToolsCoincideArgs) (resp *HeadquartersToolsCoincideResponse, err error) {
	resp = &HeadquartersToolsCoincideResponse{}
	err = c.Request("HeadquartersToolsCoincide", req, resp)
	return
}

type HeadquartersToolsCoincideArgs struct {
	Bid int64
}
type HeadquartersToolsCoincideResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetApDetailedConfig(req *GetApDetailedConfigArgs) (resp *GetApDetailedConfigResponse, err error) {
	resp = &GetApDetailedConfigResponse{}
	err = c.Request("GetApDetailedConfig", req, resp)
	return
}

type GetApDetailedConfigArgs struct {
	Id int64
}
type GetApDetailedConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) UserDataShowList(req *UserDataShowListArgs) (resp *UserDataShowListResponse, err error) {
	resp = &UserDataShowListResponse{}
	err = c.Request("UserDataShowList", req, resp)
	return
}

type UserDataShowListArgs struct {
	Iid  int64
	Name string
	Page int
	Bid  int64
	Per  int
}
type UserDataShowListResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) SaveAccountConfig(req *SaveAccountConfigArgs) (resp *SaveAccountConfigResponse, err error) {
	resp = &SaveAccountConfigResponse{}
	err = c.Request("SaveAccountConfig", req, resp)
	return
}

type SaveAccountConfigArgs struct {
	JsonData string
}
type SaveAccountConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) BatchRegisterApAsset(req *BatchRegisterApAssetArgs) (resp *BatchRegisterApAssetResponse, err error) {
	resp = &BatchRegisterApAssetResponse{}
	err = c.Request("BatchRegisterApAsset", req, resp)
	return
}

type BatchRegisterApAssetArgs struct {
	JsonData string
}
type BatchRegisterApAssetResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ReportZoneRealtime(req *ReportZoneRealtimeArgs) (resp *ReportZoneRealtimeResponse, err error) {
	resp = &ReportZoneRealtimeResponse{}
	err = c.Request("ReportZoneRealtime", req, resp)
	return
}

type ReportZoneRealtimeArgs struct {
	Agsid int64
}
type ReportZoneRealtimeResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ShopGroupDelete(req *ShopGroupDeleteArgs) (resp *ShopGroupDeleteResponse, err error) {
	resp = &ShopGroupDeleteResponse{}
	err = c.Request("ShopGroupDelete", req, resp)
	return
}

type ShopGroupDeleteArgs struct {
	Gid int64
}
type ShopGroupDeleteResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ProfileHistoryList(req *ProfileHistoryListArgs) (resp *ProfileHistoryListResponse, err error) {
	resp = &ProfileHistoryListResponse{}
	err = c.Request("ProfileHistoryList", req, resp)
	return
}

type ProfileHistoryListArgs struct {
	Idtype int64
	Page   int
	Per    int
	Agsid  int64
}
type ProfileHistoryListResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GroupTrend(req *GroupTrendArgs) (resp *GroupTrendResponse, err error) {
	resp = &GroupTrendResponse{}
	err = c.Request("GroupTrend", req, resp)
	return
}

type GroupTrendArgs struct {
	Gsid int64
}
type GroupTrendResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) SaveApMapInfo(req *SaveApMapInfoArgs) (resp *SaveApMapInfoResponse, err error) {
	resp = &SaveApMapInfoResponse{}
	err = c.Request("SaveApMapInfo", req, resp)
	return
}

type SaveApMapInfoArgs struct {
	JsonData string
}
type SaveApMapInfoResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ListConfigByAction(req *ListConfigByActionArgs) (resp *ListConfigByActionResponse, err error) {
	resp = &ListConfigByActionResponse{}
	err = c.Request("ListConfigByAction", req, resp)
	return
}

type ListConfigByActionArgs struct {
	SearchName string
	Limit      int
	ActionName string
}
type ListConfigByActionResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) SaveApScanConfig(req *SaveApScanConfigArgs) (resp *SaveApScanConfigResponse, err error) {
	resp = &SaveApScanConfigResponse{}
	err = c.Request("SaveApScanConfig", req, resp)
	return
}

type SaveApScanConfigArgs struct {
	JsonData   string
	ApConfigId int64
}
type SaveApScanConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) PeripheryAnalyse(req *PeripheryAnalyseArgs) (resp *PeripheryAnalyseResponse, err error) {
	resp = &PeripheryAnalyseResponse{}
	err = c.Request("PeripheryAnalyse", req, resp)
	return
}

type PeripheryAnalyseArgs struct {
	Gsid int64
}
type PeripheryAnalyseResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ShopGroupCreate(req *ShopGroupCreateArgs) (resp *ShopGroupCreateResponse, err error) {
	resp = &ShopGroupCreateResponse{}
	err = c.Request("ShopGroupCreate", req, resp)
	return
}

type ShopGroupCreateArgs struct {
	ShopIds     string
	Name        string
	Description string
	Bid         int64
}
type ShopGroupCreateResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetGroupApRadioConfigProgress(req *GetGroupApRadioConfigProgressArgs) (resp *GetGroupApRadioConfigProgressResponse, err error) {
	resp = &GetGroupApRadioConfigProgressResponse{}
	err = c.Request("GetGroupApRadioConfigProgress", req, resp)
	return
}

type GetGroupApRadioConfigProgressArgs struct {
	Id int64
}
type GetGroupApRadioConfigProgressResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) AreaInfo(req *AreaInfoArgs) (resp *AreaInfoResponse, err error) {
	resp = &AreaInfoResponse{}
	err = c.Request("AreaInfo", req, resp)
	return
}

type AreaInfoArgs struct {
	Aid int64
	Sid int64
}
type AreaInfoResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) KickAndClearPMKcache(req *KickAndClearPMKcacheArgs) (resp *KickAndClearPMKcacheResponse, err error) {
	resp = &KickAndClearPMKcacheResponse{}
	err = c.Request("KickAndClearPMKcache", req, resp)
	return
}

type KickAndClearPMKcacheArgs struct {
	Id int64
}
type KickAndClearPMKcacheResponse struct {
	RequestId string
	Success   bool
	Message   string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) HeadquartersToolsO2O(req *HeadquartersToolsO2OArgs) (resp *HeadquartersToolsO2OResponse, err error) {
	resp = &HeadquartersToolsO2OResponse{}
	err = c.Request("HeadquartersToolsO2O", req, resp)
	return
}

type HeadquartersToolsO2OArgs struct {
	Bid int64
}
type HeadquartersToolsO2OResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ListApDetailInfo(req *ListApDetailInfoArgs) (resp *ListApDetailInfoResponse, err error) {
	resp = &ListApDetailInfoResponse{}
	err = c.Request("ListApDetailInfo", req, resp)
	return
}

type ListApDetailInfoArgs struct {
	ApAssetId int64
}
type ListApDetailInfoResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ListApAsset(req *ListApAssetArgs) (resp *ListApAssetResponse, err error) {
	resp = &ListApAssetResponse{}
	err = c.Request("ListApAsset", req, resp)
	return
}

type ListApAssetArgs struct {
	OrderCol       string
	SearchName     string
	SearchSerialNo string
	Length         int
	PageIndex      int
	SearchMac      string
	OrderDir       string
	SearchModel    string
}
type ListApAssetResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ProfileTag(req *ProfileTagArgs) (resp *ProfileTagResponse, err error) {
	resp = &ProfileTagResponse{}
	err = c.Request("ProfileTag", req, resp)
	return
}

type ProfileTagArgs struct {
	Idtype     int64
	BeginDate  string
	EndDate    string
	AppType    int
	Tag        string
	Agsid      int64
	AreaNumber int
}
type ProfileTagResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetAddApsProgress(req *GetAddApsProgressArgs) (resp *GetAddApsProgressResponse, err error) {
	resp = &GetAddApsProgressResponse{}
	err = c.Request("GetAddApsProgress", req, resp)
	return
}

type GetAddApsProgressArgs struct {
	Id int64
}
type GetAddApsProgressResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ShopCreate(req *ShopCreateArgs) (resp *ShopCreateResponse, err error) {
	resp = &ShopCreateResponse{}
	err = c.Request("ShopCreate", req, resp)
	return
}

type ShopCreateArgs struct {
	ShopCoordinate    string
	ShopProvince      string
	ShopTopType       int
	ShopAddress       string
	ShopType          int
	WarnEmail         string
	ShopTel           string
	WarnpHone         string
	Warn              int
	ShopArea          int
	ShopRemarks       string
	ShopCity          string
	ShopSubtype       int
	ShopBrand         string
	ShopName          string
	ShopCloseWarn     int
	Bid               int64
	ShopManager       string
	ShopBusinessHours string
}
type ShopCreateResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ShopGroupShowList(req *ShopGroupShowListArgs) (resp *ShopGroupShowListResponse, err error) {
	resp = &ShopGroupShowListResponse{}
	err = c.Request("ShopGroupShowList", req, resp)
	return
}

type ShopGroupShowListArgs struct {
	Page int
	Bid  int64
	Per  int
}
type ShopGroupShowListResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ListApRadioStatus(req *ListApRadioStatusArgs) (resp *ListApRadioStatusResponse, err error) {
	resp = &ListApRadioStatusResponse{}
	err = c.Request("ListApRadioStatus", req, resp)
	return
}

type ListApRadioStatusArgs struct {
	SearchDisabled      int
	OrderCol            string
	SearchName          string
	SearchChannelEquals int
	Length              int
	SearchMac           string
	SearchApgroupName   string
	PageIndex           int
	OrderDir            string
	SearchApStatus      int
}
type ListApRadioStatusResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) BusinessInfo(req *BusinessInfoArgs) (resp *BusinessInfoResponse, err error) {
	resp = &BusinessInfoResponse{}
	err = c.Request("BusinessInfo", req, resp)
	return
}

type BusinessInfoArgs struct {
	Bid int64
}
type BusinessInfoResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) GetMapUrl(req *GetMapUrlArgs) (resp *GetMapUrlResponse, err error) {
	resp = &GetMapUrlResponse{}
	err = c.Request("GetMapUrl", req, resp)
	return
}

type GetMapUrlArgs struct {
	MapId int64
}
type GetMapUrlResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) DeviceUpdate(req *DeviceUpdateArgs) (resp *DeviceUpdateResponse, err error) {
	resp = &DeviceUpdateResponse{}
	err = c.Request("DeviceUpdate", req, resp)
	return
}

type DeviceUpdateArgs struct {
	DevicePosition string
	DeviceName     string
	Did            int64
}
type DeviceUpdateResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ShopCreatemarketing(req *ShopCreatemarketingArgs) (resp *ShopCreatemarketingResponse, err error) {
	resp = &ShopCreatemarketingResponse{}
	err = c.Request("ShopCreatemarketing", req, resp)
	return
}

type ShopCreatemarketingArgs struct {
	Etime string
	Name  string
	Stime string
	Sid   int64
}
type ShopCreatemarketingResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) SaveApgroupSsidConfig(req *SaveApgroupSsidConfigArgs) (resp *SaveApgroupSsidConfigResponse, err error) {
	resp = &SaveApgroupSsidConfigResponse{}
	err = c.Request("SaveApgroupSsidConfig", req, resp)
	return
}

type SaveApgroupSsidConfigArgs struct {
	JsonData string
}
type SaveApgroupSsidConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ListAccountConfig(req *ListAccountConfigArgs) (resp *ListAccountConfigResponse, err error) {
	resp = &ListAccountConfigResponse{}
	err = c.Request("ListAccountConfig", req, resp)
	return
}

type ListAccountConfigArgs struct {
	OrderCol    string
	Length      int
	SearchEmail string
	PageIndex   int
	OrderDir    string
}
type ListAccountConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ShopOverview(req *ShopOverviewArgs) (resp *ShopOverviewResponse, err error) {
	resp = &ShopOverviewResponse{}
	err = c.Request("ShopOverview", req, resp)
	return
}

type ShopOverviewArgs struct {
	Gsid int64
}
type ShopOverviewResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) SetScanMode(req *SetScanModeArgs) (resp *SetScanModeResponse, err error) {
	resp = &SetScanModeResponse{}
	err = c.Request("SetScanMode", req, resp)
	return
}

type SetScanModeArgs struct {
	Operation int
	MacLists  SetScanModeMacListList
}
type SetScanModeResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

type SetScanModeMacListList []string

func (list *SetScanModeMacListList) UnmarshalJSON(data []byte) error {
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

func (c *CloudwfClient) ReportMinute(req *ReportMinuteArgs) (resp *ReportMinuteResponse, err error) {
	resp = &ReportMinuteResponse{}
	err = c.Request("ReportMinute", req, resp)
	return
}

type ReportMinuteArgs struct {
	BeginDate string
	EndDate   string
	Agsid     int64
}
type ReportMinuteResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) UserDataUpdate(req *UserDataUpdateArgs) (resp *UserDataUpdateResponse, err error) {
	resp = &UserDataUpdateResponse{}
	err = c.Request("UserDataUpdate", req, resp)
	return
}

type UserDataUpdateArgs struct {
	Iid        int64
	UploadFile string
	Name       string
	Bid        int64
	Type       string
}
type UserDataUpdateResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}

func (c *CloudwfClient) ShopDelete(req *ShopDeleteArgs) (resp *ShopDeleteResponse, err error) {
	resp = &ShopDeleteResponse{}
	err = c.Request("ShopDelete", req, resp)
	return
}

type ShopDeleteArgs struct {
	Sid int64
}
type ShopDeleteResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
