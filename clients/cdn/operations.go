package cdn

import (
	"encoding/json"

	"github.com/morlay/aliyun-go/core"
)

func (c *CdnClient) SetPageCompressConfig(req *SetPageCompressConfigArgs) (resp *SetPageCompressConfigResponse, err error) {
	resp = &SetPageCompressConfigResponse{}
	err = c.Request("SetPageCompressConfig", req, resp)
	return
}

type SetPageCompressConfigArgs struct {
	SecurityToken string
	Enable        string
	DomainName    string
	OwnerId       int64
}
type SetPageCompressConfigResponse struct {
	RequestId string
}

func (c *CdnClient) AddCdnDomain(req *AddCdnDomainArgs) (resp *AddCdnDomainResponse, err error) {
	resp = &AddCdnDomainResponse{}
	err = c.Request("AddCdnDomain", req, resp)
	return
}

type AddCdnDomainArgs struct {
	TopLevelDomain  string
	Sources         string
	OwnerAccount    string
	DomainName      string
	LiveType        string
	OwnerId         int64
	ResourceGroupId string
	SourcePort      int
	Priorities      string
	SecurityToken   string
	CdnType         string
	Scope           string
	SourceType      string
	CheckUrl        string
	Region          string
}
type AddCdnDomainResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeLiveStreamOnlineBps(req *DescribeLiveStreamOnlineBpsArgs) (resp *DescribeLiveStreamOnlineBpsResponse, err error) {
	resp = &DescribeLiveStreamOnlineBpsResponse{}
	err = c.Request("DescribeLiveStreamOnlineBps", req, resp)
	return
}

type DescribeLiveStreamOnlineBpsLiveStreamOnlineBpsInfo struct {
	StreamUrl string
	UpBps     float32
	DownBps   float32
	Time      string
}
type DescribeLiveStreamOnlineBpsArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
	StreamName    string
}
type DescribeLiveStreamOnlineBpsResponse struct {
	RequestId                string
	TotalUserNumber          int64
	FlvBps                   float32
	HlsBps                   float32
	LiveStreamOnlineBpsInfos DescribeLiveStreamOnlineBpsLiveStreamOnlineBpsInfoList
}

type DescribeLiveStreamOnlineBpsLiveStreamOnlineBpsInfoList []DescribeLiveStreamOnlineBpsLiveStreamOnlineBpsInfo

func (list *DescribeLiveStreamOnlineBpsLiveStreamOnlineBpsInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamOnlineBpsLiveStreamOnlineBpsInfo)
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

func (c *CdnClient) DescribeLiveStreamFrameLossRatio(req *DescribeLiveStreamFrameLossRatioArgs) (resp *DescribeLiveStreamFrameLossRatioResponse, err error) {
	resp = &DescribeLiveStreamFrameLossRatioResponse{}
	err = c.Request("DescribeLiveStreamFrameLossRatio", req, resp)
	return
}

type DescribeLiveStreamFrameLossRatioFrameLossRatioInfo struct {
	StreamUrl      string
	FrameLossRatio float32
	Time           string
}
type DescribeLiveStreamFrameLossRatioArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
	StreamName    string
}
type DescribeLiveStreamFrameLossRatioResponse struct {
	RequestId           string
	FrameLossRatioInfos DescribeLiveStreamFrameLossRatioFrameLossRatioInfoList
}

type DescribeLiveStreamFrameLossRatioFrameLossRatioInfoList []DescribeLiveStreamFrameLossRatioFrameLossRatioInfo

func (list *DescribeLiveStreamFrameLossRatioFrameLossRatioInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamFrameLossRatioFrameLossRatioInfo)
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

func (c *CdnClient) DescribeDomainBpsDataByTimeStamp(req *DescribeDomainBpsDataByTimeStampArgs) (resp *DescribeDomainBpsDataByTimeStampResponse, err error) {
	resp = &DescribeDomainBpsDataByTimeStampResponse{}
	err = c.Request("DescribeDomainBpsDataByTimeStamp", req, resp)
	return
}

type DescribeDomainBpsDataByTimeStampBpsDataModel struct {
	LocationName string
	IspName      string
	Bps          int64
}
type DescribeDomainBpsDataByTimeStampArgs struct {
	IspNames      string
	SecurityToken string
	LocationNames string
	DomainName    string
	OwnerId       int64
	TimePoint     string
}
type DescribeDomainBpsDataByTimeStampResponse struct {
	RequestId   string
	DomainName  string
	TimeStamp   string
	BpsDataList DescribeDomainBpsDataByTimeStampBpsDataModelList
}

type DescribeDomainBpsDataByTimeStampBpsDataModelList []DescribeDomainBpsDataByTimeStampBpsDataModel

func (list *DescribeDomainBpsDataByTimeStampBpsDataModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainBpsDataByTimeStampBpsDataModel)
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

func (c *CdnClient) SetLocationAccessRestriction(req *SetLocationAccessRestrictionArgs) (resp *SetLocationAccessRestrictionResponse, err error) {
	resp = &SetLocationAccessRestrictionResponse{}
	err = c.Request("SetLocationAccessRestriction", req, resp)
	return
}

type SetLocationAccessRestrictionArgs struct {
	SecurityToken string
	DomainName    string
	Location      string
	OwnerId       int64
	Type          string
}
type SetLocationAccessRestrictionResponse struct {
	RequestId string
}

func (c *CdnClient) CreateLiveStreamRecordIndexFiles(req *CreateLiveStreamRecordIndexFilesArgs) (resp *CreateLiveStreamRecordIndexFilesResponse, err error) {
	resp = &CreateLiveStreamRecordIndexFilesResponse{}
	err = c.Request("CreateLiveStreamRecordIndexFiles", req, resp)
	return
}

type CreateLiveStreamRecordIndexFilesRecordInfo struct {
	RecordId   string
	RecordUrl  string
	Duration   float32
	Height     int
	Width      int
	CreateTime string
}
type CreateLiveStreamRecordIndexFilesArgs struct {
	OssBucket     string
	AppName       string
	SecurityToken string
	DomainName    string
	OssEndpoint   string
	EndTime       string
	StartTime     string
	OwnerId       int64
	StreamName    string
	OssObject     string
}
type CreateLiveStreamRecordIndexFilesResponse struct {
	RequestId  string
	RecordInfo CreateLiveStreamRecordIndexFilesRecordInfo
}

func (c *CdnClient) DescribeCdnDomainLogs(req *DescribeCdnDomainLogsArgs) (resp *DescribeCdnDomainLogsResponse, err error) {
	resp = &DescribeCdnDomainLogsResponse{}
	err = c.Request("DescribeCdnDomainLogs", req, resp)
	return
}

type DescribeCdnDomainLogsDomainLogModel struct {
	DomainName       string
	DomainLogDetails DescribeCdnDomainLogsDomainLogDetailList
}

type DescribeCdnDomainLogsDomainLogDetail struct {
	LogName   string
	LogPath   string
	LogSize   int64
	StartTime string
	EndTime   string
}
type DescribeCdnDomainLogsArgs struct {
	SecurityToken string
	DomainName    string
	PageSize      int64
	EndTime       string
	StartTime     string
	OwnerId       int64
	PageNumber    int64
	LogDay        string
}
type DescribeCdnDomainLogsResponse struct {
	RequestId      string
	PageNumber     int64
	PageSize       int64
	TotalCount     int64
	DomainLogModel DescribeCdnDomainLogsDomainLogModel
}

type DescribeCdnDomainLogsDomainLogDetailList []DescribeCdnDomainLogsDomainLogDetail

func (list *DescribeCdnDomainLogsDomainLogDetailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCdnDomainLogsDomainLogDetail)
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

func (c *CdnClient) DescribeDomainQpsData(req *DescribeDomainQpsDataArgs) (resp *DescribeDomainQpsDataResponse, err error) {
	resp = &DescribeDomainQpsDataResponse{}
	err = c.Request("DescribeDomainQpsData", req, resp)
	return
}

type DescribeDomainQpsDataDataModule struct {
	TimeStamp            string
	Value                string
	DomesticValue        string
	OverseasValue        string
	AccValue             string
	AccDomesticValue     string
	AccOverseasValue     string
	DynamicValue         string
	DynamicDomesticValue string
	DynamicOverseasValue string
	StaticValue          string
	StaticDomesticValue  string
	StaticOverseasValue  string
}
type DescribeDomainQpsDataArgs struct {
	FixTimeGap     string
	TimeMerge      string
	DomainName     string
	EndTime        string
	LocationNameEn string
	StartTime      string
	IspNameEn      string
	OwnerId        int64
	DomainType     string
	SecurityToken  string
	Interval       string
}
type DescribeDomainQpsDataResponse struct {
	RequestId       string
	DomainName      string
	DataInterval    string
	StartTime       string
	EndTime         string
	QpsDataInterval DescribeDomainQpsDataDataModuleList
}

type DescribeDomainQpsDataDataModuleList []DescribeDomainQpsDataDataModule

func (list *DescribeDomainQpsDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainQpsDataDataModule)
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

func (c *CdnClient) DescribeDomainISPData(req *DescribeDomainISPDataArgs) (resp *DescribeDomainISPDataResponse, err error) {
	resp = &DescribeDomainISPDataResponse{}
	err = c.Request("DescribeDomainISPData", req, resp)
	return
}

type DescribeDomainISPDataISPProportionData struct {
	ISP             string
	Proportion      string
	IspEname        string
	AvgObjectSize   string
	AvgResponseTime string
	Bps             string
	ByteHitRate     string
	Qps             string
	ReqErrRate      string
	ReqHitRate      string
	AvgResponseRate string
	TotalBytes      string
	BytesProportion string
	TotalQuery      string
}
type DescribeDomainISPDataArgs struct {
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
}
type DescribeDomainISPDataResponse struct {
	RequestId    string
	DomainName   string
	DataInterval string
	StartTime    string
	EndTime      string
	Value        DescribeDomainISPDataISPProportionDataList
}

type DescribeDomainISPDataISPProportionDataList []DescribeDomainISPDataISPProportionData

func (list *DescribeDomainISPDataISPProportionDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainISPDataISPProportionData)
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

func (c *CdnClient) DescribeOneMinuteData(req *DescribeOneMinuteDataArgs) (resp *DescribeOneMinuteDataResponse, err error) {
	resp = &DescribeOneMinuteDataResponse{}
	err = c.Request("DescribeOneMinuteData", req, resp)
	return
}

type DescribeOneMinuteDataArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	DataTime             string
	DomainName           string
	OwnerId              int64
}
type DescribeOneMinuteDataResponse struct {
	RequestId string
	Bps       string
	Qps       string
}

func (c *CdnClient) SetWafConfig(req *SetWafConfigArgs) (resp *SetWafConfigResponse, err error) {
	resp = &SetWafConfigResponse{}
	err = c.Request("SetWafConfig", req, resp)
	return
}

type SetWafConfigArgs struct {
	SecurityToken string
	Enable        string
	DomainName    string
	OwnerId       int64
}
type SetWafConfigResponse struct {
	RequestId string
}

func (c *CdnClient) SetLiveStreamsNotifyUrlConfig(req *SetLiveStreamsNotifyUrlConfigArgs) (resp *SetLiveStreamsNotifyUrlConfigResponse, err error) {
	resp = &SetLiveStreamsNotifyUrlConfigResponse{}
	err = c.Request("SetLiveStreamsNotifyUrlConfig", req, resp)
	return
}

type SetLiveStreamsNotifyUrlConfigArgs struct {
	SecurityToken string
	DomainName    string
	NotifyUrl     string
	OwnerId       int64
}
type SetLiveStreamsNotifyUrlConfigResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeLiveStreamRecordIndexFiles(req *DescribeLiveStreamRecordIndexFilesArgs) (resp *DescribeLiveStreamRecordIndexFilesResponse, err error) {
	resp = &DescribeLiveStreamRecordIndexFilesResponse{}
	err = c.Request("DescribeLiveStreamRecordIndexFiles", req, resp)
	return
}

type DescribeLiveStreamRecordIndexFilesRecordIndexInfo struct {
	RecordId   string
	RecordUrl  string
	DomainName string
	AppName    string
	StreamName string
	OssObject  string
	StartTime  string
	EndTime    string
	Duration   float32
	Height     int
	Width      int
	CreateTime string
}
type DescribeLiveStreamRecordIndexFilesArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
	StreamName    string
}
type DescribeLiveStreamRecordIndexFilesResponse struct {
	RequestId           string
	RecordIndexInfoList DescribeLiveStreamRecordIndexFilesRecordIndexInfoList
}

type DescribeLiveStreamRecordIndexFilesRecordIndexInfoList []DescribeLiveStreamRecordIndexFilesRecordIndexInfo

func (list *DescribeLiveStreamRecordIndexFilesRecordIndexInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamRecordIndexFilesRecordIndexInfo)
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

func (c *CdnClient) DescribeDomainPvData(req *DescribeDomainPvDataArgs) (resp *DescribeDomainPvDataResponse, err error) {
	resp = &DescribeDomainPvDataResponse{}
	err = c.Request("DescribeDomainPvData", req, resp)
	return
}

type DescribeDomainPvDataUsageData struct {
	TimeStamp string
	Value     string
}
type DescribeDomainPvDataArgs struct {
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
}
type DescribeDomainPvDataResponse struct {
	RequestId      string
	DomainName     string
	DataInterval   string
	StartTime      string
	EndTime        string
	PvDataInterval DescribeDomainPvDataUsageDataList
}

type DescribeDomainPvDataUsageDataList []DescribeDomainPvDataUsageData

func (list *DescribeDomainPvDataUsageDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainPvDataUsageData)
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

func (c *CdnClient) DescribeDomainMax95BpsData(req *DescribeDomainMax95BpsDataArgs) (resp *DescribeDomainMax95BpsDataResponse, err error) {
	resp = &DescribeDomainMax95BpsDataResponse{}
	err = c.Request("DescribeDomainMax95BpsData", req, resp)
	return
}

type DescribeDomainMax95BpsDataArgs struct {
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
}
type DescribeDomainMax95BpsDataResponse struct {
	RequestId        string
	DomainName       string
	StartTime        string
	EndTime          string
	Max95Bps         string
	DomesticMax95Bps string
	OverseasMax95Bps string
}

func (c *CdnClient) DescribeDomainTopReferVisit(req *DescribeDomainTopReferVisitArgs) (resp *DescribeDomainTopReferVisitResponse, err error) {
	resp = &DescribeDomainTopReferVisitResponse{}
	err = c.Request("DescribeDomainTopReferVisit", req, resp)
	return
}

type DescribeDomainTopReferVisitReferList struct {
	ReferDetail     string
	VisitData       string
	VisitProportion float32
	Flow            string
	FlowProportion  float32
}
type DescribeDomainTopReferVisitArgs struct {
	SecurityToken string
	DomainName    string
	SortBy        string
	StartTime     string
	OwnerId       int64
}
type DescribeDomainTopReferVisitResponse struct {
	RequestId    string
	DomainName   string
	StartTime    string
	TopReferList DescribeDomainTopReferVisitReferListList
}

type DescribeDomainTopReferVisitReferListList []DescribeDomainTopReferVisitReferList

func (list *DescribeDomainTopReferVisitReferListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainTopReferVisitReferList)
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

func (c *CdnClient) DescribeDomainReqHitRateData(req *DescribeDomainReqHitRateDataArgs) (resp *DescribeDomainReqHitRateDataResponse, err error) {
	resp = &DescribeDomainReqHitRateDataResponse{}
	err = c.Request("DescribeDomainReqHitRateData", req, resp)
	return
}

type DescribeDomainReqHitRateDataDataModule struct {
	TimeStamp string
	Value     string
}
type DescribeDomainReqHitRateDataArgs struct {
	SecurityToken string
	DomainName    string
	EndTime       string
	Interval      string
	StartTime     string
	OwnerId       int64
}
type DescribeDomainReqHitRateDataResponse struct {
	RequestId          string
	DomainName         string
	DataInterval       string
	StartTime          string
	EndTime            string
	ReqHitRateInterval DescribeDomainReqHitRateDataDataModuleList
}

type DescribeDomainReqHitRateDataDataModuleList []DescribeDomainReqHitRateDataDataModule

func (list *DescribeDomainReqHitRateDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainReqHitRateDataDataModule)
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

func (c *CdnClient) DescribeLiveStreamNumberList(req *DescribeLiveStreamNumberListArgs) (resp *DescribeLiveStreamNumberListResponse, err error) {
	resp = &DescribeLiveStreamNumberListResponse{}
	err = c.Request("DescribeLiveStreamNumberList", req, resp)
	return
}

type DescribeLiveStreamNumberListStreamNumberInfo struct {
	Number int
	Time   string
}
type DescribeLiveStreamNumberListArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
}
type DescribeLiveStreamNumberListResponse struct {
	RequestId         string
	DomainName        string
	StreamNumberInfos DescribeLiveStreamNumberListStreamNumberInfoList
}

type DescribeLiveStreamNumberListStreamNumberInfoList []DescribeLiveStreamNumberListStreamNumberInfo

func (list *DescribeLiveStreamNumberListStreamNumberInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamNumberListStreamNumberInfo)
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

func (c *CdnClient) AddLiveStreamTranscode(req *AddLiveStreamTranscodeArgs) (resp *AddLiveStreamTranscodeResponse, err error) {
	resp = &AddLiveStreamTranscodeResponse{}
	err = c.Request("AddLiveStreamTranscode", req, resp)
	return
}

type AddLiveStreamTranscodeArgs struct {
	Template      string
	App           string
	SecurityToken string
	OwnerAccount  string
	Domain        string
	Record        string
	OwnerId       int64
	Snapshot      string
}
type AddLiveStreamTranscodeResponse struct {
	RequestId string
}

func (c *CdnClient) AddLiveAppSnapshotConfig(req *AddLiveAppSnapshotConfigArgs) (resp *AddLiveAppSnapshotConfigResponse, err error) {
	resp = &AddLiveAppSnapshotConfigResponse{}
	err = c.Request("AddLiveAppSnapshotConfig", req, resp)
	return
}

type AddLiveAppSnapshotConfigArgs struct {
	TimeInterval       int
	OssBucket          string
	AppName            string
	SecurityToken      string
	DomainName         string
	OssEndpoint        string
	SequenceOssObject  string
	OverwriteOssObject string
	OwnerId            int64
}
type AddLiveAppSnapshotConfigResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeUserConfigs(req *DescribeUserConfigsArgs) (resp *DescribeUserConfigsResponse, err error) {
	resp = &DescribeUserConfigsResponse{}
	err = c.Request("DescribeUserConfigs", req, resp)
	return
}

type DescribeUserConfigsConfigs struct {
	OssLogConfig       DescribeUserConfigsOssLogConfig
	GreenManagerConfig DescribeUserConfigsGreenManagerConfig
}

type DescribeUserConfigsOssLogConfig struct {
	Enable string
	Bucket string
	Prefix string
}

type DescribeUserConfigsGreenManagerConfig struct {
	Quota string
	Ratio string
}
type DescribeUserConfigsArgs struct {
	SecurityToken string
	OwnerId       int64
	Config        string
}
type DescribeUserConfigsResponse struct {
	RequestId string
	Configs   DescribeUserConfigsConfigs
}

func (c *CdnClient) SetOptimizeConfig(req *SetOptimizeConfigArgs) (resp *SetOptimizeConfigResponse, err error) {
	resp = &SetOptimizeConfigResponse{}
	err = c.Request("SetOptimizeConfig", req, resp)
	return
}

type SetOptimizeConfigArgs struct {
	SecurityToken string
	Enable        string
	DomainName    string
	OwnerId       int64
}
type SetOptimizeConfigResponse struct {
	RequestId string
}

func (c *CdnClient) SetVideoSeekConfig(req *SetVideoSeekConfigArgs) (resp *SetVideoSeekConfigResponse, err error) {
	resp = &SetVideoSeekConfigResponse{}
	err = c.Request("SetVideoSeekConfig", req, resp)
	return
}

type SetVideoSeekConfigArgs struct {
	SecurityToken string
	Enable        string
	DomainName    string
	OwnerId       int64
}
type SetVideoSeekConfigResponse struct {
	RequestId string
}

func (c *CdnClient) AddLivePullStreamInfo(req *AddLivePullStreamInfoArgs) (resp *AddLivePullStreamInfoResponse, err error) {
	resp = &AddLivePullStreamInfoResponse{}
	err = c.Request("AddLivePullStreamInfo", req, resp)
	return
}

type AddLivePullStreamInfoArgs struct {
	SourceUrl     string
	AppName       string
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
	StreamName    string
}
type AddLivePullStreamInfoResponse struct {
	RequestId string
}

func (c *CdnClient) SetReqHeaderConfig(req *SetReqHeaderConfigArgs) (resp *SetReqHeaderConfigResponse, err error) {
	resp = &SetReqHeaderConfigResponse{}
	err = c.Request("SetReqHeaderConfig", req, resp)
	return
}

type SetReqHeaderConfigArgs struct {
	SecurityToken string
	DomainName    string
	OwnerId       int64
	Value         string
	Key           string
}
type SetReqHeaderConfigResponse struct {
	RequestId string
}

func (c *CdnClient) ModifyCdnService(req *ModifyCdnServiceArgs) (resp *ModifyCdnServiceResponse, err error) {
	resp = &ModifyCdnServiceResponse{}
	err = c.Request("ModifyCdnService", req, resp)
	return
}

type ModifyCdnServiceArgs struct {
	SecurityToken      string
	InternetChargeType string
	OwnerId            int64
}
type ModifyCdnServiceResponse struct {
	RequestId string
}

func (c *CdnClient) SetErrorPageConfig(req *SetErrorPageConfigArgs) (resp *SetErrorPageConfigResponse, err error) {
	resp = &SetErrorPageConfigResponse{}
	err = c.Request("SetErrorPageConfig", req, resp)
	return
}

type SetErrorPageConfigArgs struct {
	PageType      string
	SecurityToken string
	DomainName    string
	CustomPageUrl string
	OwnerId       int64
}
type SetErrorPageConfigResponse struct {
	RequestId string
}

func (c *CdnClient) SetCcConfig(req *SetCcConfigArgs) (resp *SetCcConfigResponse, err error) {
	resp = &SetCcConfigResponse{}
	err = c.Request("SetCcConfig", req, resp)
	return
}

type SetCcConfigArgs struct {
	AllowIps      string
	SecurityToken string
	DomainName    string
	OwnerId       int64
	BlockIps      string
}
type SetCcConfigResponse struct {
	RequestId string
}

func (c *CdnClient) SetHttpHeaderConfig(req *SetHttpHeaderConfigArgs) (resp *SetHttpHeaderConfigResponse, err error) {
	resp = &SetHttpHeaderConfigResponse{}
	err = c.Request("SetHttpHeaderConfig", req, resp)
	return
}

type SetHttpHeaderConfigArgs struct {
	HeaderValue   string
	SecurityToken string
	DomainName    string
	HeaderKey     string
	OwnerId       int64
}
type SetHttpHeaderConfigResponse struct {
	RequestId string
}

func (c *CdnClient) ResumeLiveStream(req *ResumeLiveStreamArgs) (resp *ResumeLiveStreamResponse, err error) {
	resp = &ResumeLiveStreamResponse{}
	err = c.Request("ResumeLiveStream", req, resp)
	return
}

type ResumeLiveStreamArgs struct {
	AppName        string
	SecurityToken  string
	LiveStreamType string
	DomainName     string
	OwnerId        int64
	StreamName     string
}
type ResumeLiveStreamResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeLiveStreamTranscodeStreamNum(req *DescribeLiveStreamTranscodeStreamNumArgs) (resp *DescribeLiveStreamTranscodeStreamNumResponse, err error) {
	resp = &DescribeLiveStreamTranscodeStreamNumResponse{}
	err = c.Request("DescribeLiveStreamTranscodeStreamNum", req, resp)
	return
}

type DescribeLiveStreamTranscodeStreamNumArgs struct {
	PullDomain    string
	SecurityToken string
	PushDomain    string
	OwnerId       int64
}
type DescribeLiveStreamTranscodeStreamNumResponse struct {
	RequestId         string
	Total             int64
	TranscodedNumber  int64
	UntranscodeNumber int64
}

func (c *CdnClient) DescribeDomainsBySource(req *DescribeDomainsBySourceArgs) (resp *DescribeDomainsBySourceResponse, err error) {
	resp = &DescribeDomainsBySourceResponse{}
	err = c.Request("DescribeDomainsBySource", req, resp)
	return
}

type DescribeDomainsBySourceDomainsData struct {
	Source      string
	DomainInfos DescribeDomainsBySourceDomainInfoList
	Domains     DescribeDomainsBySourceDomainList
}

type DescribeDomainsBySourceDomainInfo struct {
	DomainName  string
	DomainCname string
	CreateTime  string
	UpdateTime  string
	Status      string
}
type DescribeDomainsBySourceArgs struct {
	Sources       string
	SecurityToken string
	OwnerId       int64
}
type DescribeDomainsBySourceResponse struct {
	RequestId   string
	Sources     string
	DomainsList DescribeDomainsBySourceDomainsDataList
}

type DescribeDomainsBySourceDomainInfoList []DescribeDomainsBySourceDomainInfo

func (list *DescribeDomainsBySourceDomainInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainsBySourceDomainInfo)
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

type DescribeDomainsBySourceDomainList []string

func (list *DescribeDomainsBySourceDomainList) UnmarshalJSON(data []byte) error {
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

type DescribeDomainsBySourceDomainsDataList []DescribeDomainsBySourceDomainsData

func (list *DescribeDomainsBySourceDomainsDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainsBySourceDomainsData)
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

func (c *CdnClient) DescribeLiveStreamRelayPushData(req *DescribeLiveStreamRelayPushDataArgs) (resp *DescribeLiveStreamRelayPushDataResponse, err error) {
	resp = &DescribeLiveStreamRelayPushDataResponse{}
	err = c.Request("DescribeLiveStreamRelayPushData", req, resp)
	return
}

type DescribeLiveStreamRelayPushDataRelayPushDetailModel struct {
	Time          string
	Stream        string
	FrameRate     float32
	BitRate       float32
	FrameLossRate float32
	ServerAddr    string
	ClientAddr    string
}
type DescribeLiveStreamRelayPushDataArgs struct {
	RelayDomain   string
	SecurityToken string
	EndTime       string
	StartTime     string
	OwnerId       int64
}
type DescribeLiveStreamRelayPushDataResponse struct {
	RequestId                string
	RelayPushDetailModelList DescribeLiveStreamRelayPushDataRelayPushDetailModelList
}

type DescribeLiveStreamRelayPushDataRelayPushDetailModelList []DescribeLiveStreamRelayPushDataRelayPushDetailModel

func (list *DescribeLiveStreamRelayPushDataRelayPushDetailModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamRelayPushDataRelayPushDetailModel)
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

func (c *CdnClient) DescribeRefreshQuota(req *DescribeRefreshQuotaArgs) (resp *DescribeRefreshQuotaResponse, err error) {
	resp = &DescribeRefreshQuotaResponse{}
	err = c.Request("DescribeRefreshQuota", req, resp)
	return
}

type DescribeRefreshQuotaArgs struct {
	SecurityToken string
	OwnerId       int64
}
type DescribeRefreshQuotaResponse struct {
	RequestId     string
	UrlQuota      string
	DirQuota      string
	UrlRemain     string
	DirRemain     string
	PreloadQuota  string
	BlockQuota    string
	PreloadRemain string
	BlockRemain   string
}

func (c *CdnClient) DescribeLiveStreamTranscodeInfo(req *DescribeLiveStreamTranscodeInfoArgs) (resp *DescribeLiveStreamTranscodeInfoResponse, err error) {
	resp = &DescribeLiveStreamTranscodeInfoResponse{}
	err = c.Request("DescribeLiveStreamTranscodeInfo", req, resp)
	return
}

type DescribeLiveStreamTranscodeInfoDomainTranscodeInfo struct {
	TranscodeApp      string
	TranscodeId       string
	TranscodeName     string
	TranscodeRecord   string
	TranscodeSnapshot string
	TranscodeTemplate string
}
type DescribeLiveStreamTranscodeInfoArgs struct {
	SecurityToken       string
	OwnerId             int64
	DomainTranscodeName string
}
type DescribeLiveStreamTranscodeInfoResponse struct {
	RequestId           string
	DomainTranscodeList DescribeLiveStreamTranscodeInfoDomainTranscodeInfoList
}

type DescribeLiveStreamTranscodeInfoDomainTranscodeInfoList []DescribeLiveStreamTranscodeInfoDomainTranscodeInfo

func (list *DescribeLiveStreamTranscodeInfoDomainTranscodeInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamTranscodeInfoDomainTranscodeInfo)
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

func (c *CdnClient) DescribeCdnDomainDetail(req *DescribeCdnDomainDetailArgs) (resp *DescribeCdnDomainDetailResponse, err error) {
	resp = &DescribeCdnDomainDetailResponse{}
	err = c.Request("DescribeCdnDomainDetail", req, resp)
	return
}

type DescribeCdnDomainDetailGetDomainDetailModel struct {
	GmtCreated              string
	GmtModified             string
	SourceType              string
	DomainStatus            string
	SourcePort              int
	CdnType                 string
	Cname                   string
	HttpsCname              string
	DomainName              string
	Description             string
	ServerCertificateStatus string
	ServerCertificate       string
	Region                  string
	Scope                   string
	CertificateName         string
	ResourceGroupId         string
	SourceModels            DescribeCdnDomainDetailSourceModelList
	Sources                 DescribeCdnDomainDetailSourceList
}

type DescribeCdnDomainDetailSourceModel struct {
	Content  string
	Type     string
	Port     int
	Enabled  string
	Priority string
}
type DescribeCdnDomainDetailArgs struct {
	SecurityToken string
	DomainName    string
	OwnerId       int64
}
type DescribeCdnDomainDetailResponse struct {
	RequestId            string
	GetDomainDetailModel DescribeCdnDomainDetailGetDomainDetailModel
}

type DescribeCdnDomainDetailSourceModelList []DescribeCdnDomainDetailSourceModel

func (list *DescribeCdnDomainDetailSourceModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCdnDomainDetailSourceModel)
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

type DescribeCdnDomainDetailSourceList []string

func (list *DescribeCdnDomainDetailSourceList) UnmarshalJSON(data []byte) error {
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

func (c *CdnClient) SetDynamicConfig(req *SetDynamicConfigArgs) (resp *SetDynamicConfigResponse, err error) {
	resp = &SetDynamicConfigResponse{}
	err = c.Request("SetDynamicConfig", req, resp)
	return
}

type SetDynamicConfigArgs struct {
	DynamicOrigin       string
	StaticType          string
	SecurityToken       string
	StaticUri           string
	DomainName          string
	StaticPath          string
	DynamicCacheControl string
	OwnerId             int64
}
type SetDynamicConfigResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeLiveStreamPushData(req *DescribeLiveStreamPushDataArgs) (resp *DescribeLiveStreamPushDataResponse, err error) {
	resp = &DescribeLiveStreamPushDataResponse{}
	err = c.Request("DescribeLiveStreamPushData", req, resp)
	return
}

type DescribeLiveStreamPushDataPushStreamModel struct {
	Time          string
	Stream        string
	FrameRate     float32
	BitRate       float32
	FrameLossRate float32
	ServerAddr    string
	ClientAddr    string
}
type DescribeLiveStreamPushDataArgs struct {
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
}
type DescribeLiveStreamPushDataResponse struct {
	RequestId           string
	PushStreamModelList DescribeLiveStreamPushDataPushStreamModelList
}

type DescribeLiveStreamPushDataPushStreamModelList []DescribeLiveStreamPushDataPushStreamModel

func (list *DescribeLiveStreamPushDataPushStreamModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamPushDataPushStreamModel)
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

func (c *CdnClient) DescribeLiveStreamRelayPushErrors(req *DescribeLiveStreamRelayPushErrorsArgs) (resp *DescribeLiveStreamRelayPushErrorsResponse, err error) {
	resp = &DescribeLiveStreamRelayPushErrorsResponse{}
	err = c.Request("DescribeLiveStreamRelayPushErrors", req, resp)
	return
}

type DescribeLiveStreamRelayPushErrorsRelayPushErrorsModel struct {
	ErrorCode string
}
type DescribeLiveStreamRelayPushErrorsArgs struct {
	RelayDomain   string
	SecurityToken string
	EndTime       string
	StartTime     string
	OwnerId       int64
}
type DescribeLiveStreamRelayPushErrorsResponse struct {
	RequestId                string
	RelayPushErrorsModelList DescribeLiveStreamRelayPushErrorsRelayPushErrorsModelList
}

type DescribeLiveStreamRelayPushErrorsRelayPushErrorsModelList []DescribeLiveStreamRelayPushErrorsRelayPushErrorsModel

func (list *DescribeLiveStreamRelayPushErrorsRelayPushErrorsModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamRelayPushErrorsRelayPushErrorsModel)
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

func (c *CdnClient) DescribeIpInfo(req *DescribeIpInfoArgs) (resp *DescribeIpInfoResponse, err error) {
	resp = &DescribeIpInfoResponse{}
	err = c.Request("DescribeIpInfo", req, resp)
	return
}

type DescribeIpInfoArgs struct {
	SecurityToken string
	IP            string
	OwnerId       int64
}
type DescribeIpInfoResponse struct {
	RequestId   string
	CdnIp       string
	ISP         string
	IspEname    string
	Region      string
	RegionEname string
}

func (c *CdnClient) PushObjectCache(req *PushObjectCacheArgs) (resp *PushObjectCacheResponse, err error) {
	resp = &PushObjectCacheResponse{}
	err = c.Request("PushObjectCache", req, resp)
	return
}

type PushObjectCacheArgs struct {
	SecurityToken string
	ObjectPath    string
	OwnerId       int64
}
type PushObjectCacheResponse struct {
	RequestId  string
	PushTaskId string
}

func (c *CdnClient) DescribeUserCustomerLabels(req *DescribeUserCustomerLabelsArgs) (resp *DescribeUserCustomerLabelsResponse, err error) {
	resp = &DescribeUserCustomerLabelsResponse{}
	err = c.Request("DescribeUserCustomerLabels", req, resp)
	return
}

type DescribeUserCustomerLabelsArgs struct {
	Uid           int64
	SecurityToken string
	OwnerId       int64
}
type DescribeUserCustomerLabelsResponse struct {
	RequestId   string
	IsInnerUser core.Bool
}

func (c *CdnClient) MigrateDomainToHttpsDelivery(req *MigrateDomainToHttpsDeliveryArgs) (resp *MigrateDomainToHttpsDeliveryResponse, err error) {
	resp = &MigrateDomainToHttpsDeliveryResponse{}
	err = c.Request("MigrateDomainToHttpsDelivery", req, resp)
	return
}

type MigrateDomainToHttpsDeliveryArgs struct {
	PrivateKey        string
	ServerCertificate string
	SecurityToken     string
	OwnerAccount      string
	DomainName        string
	OwnerId           int64
}
type MigrateDomainToHttpsDeliveryResponse struct {
	RequestId string
}

func (c *CdnClient) DeleteLiveAppRecordConfig(req *DeleteLiveAppRecordConfigArgs) (resp *DeleteLiveAppRecordConfigResponse, err error) {
	resp = &DeleteLiveAppRecordConfigResponse{}
	err = c.Request("DeleteLiveAppRecordConfig", req, resp)
	return
}

type DeleteLiveAppRecordConfigArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	OwnerId       int64
}
type DeleteLiveAppRecordConfigResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeDomainSrcBpsData(req *DescribeDomainSrcBpsDataArgs) (resp *DescribeDomainSrcBpsDataResponse, err error) {
	resp = &DescribeDomainSrcBpsDataResponse{}
	err = c.Request("DescribeDomainSrcBpsData", req, resp)
	return
}

type DescribeDomainSrcBpsDataDataModule struct {
	TimeStamp string
	Value     string
}
type DescribeDomainSrcBpsDataArgs struct {
	FixTimeGap    string
	SecurityToken string
	TimeMerge     string
	DomainName    string
	EndTime       string
	Interval      string
	StartTime     string
	OwnerId       int64
}
type DescribeDomainSrcBpsDataResponse struct {
	RequestId             string
	DomainName            string
	DataInterval          string
	StartTime             string
	EndTime               string
	SrcBpsDataPerInterval DescribeDomainSrcBpsDataDataModuleList
}

type DescribeDomainSrcBpsDataDataModuleList []DescribeDomainSrcBpsDataDataModule

func (list *DescribeDomainSrcBpsDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainSrcBpsDataDataModule)
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

func (c *CdnClient) DescribeDomainFlowData(req *DescribeDomainFlowDataArgs) (resp *DescribeDomainFlowDataResponse, err error) {
	resp = &DescribeDomainFlowDataResponse{}
	err = c.Request("DescribeDomainFlowData", req, resp)
	return
}

type DescribeDomainFlowDataDataModule struct {
	TimeStamp            string
	Value                string
	DomesticValue        string
	OverseasValue        string
	DynamicValue         string
	DynamicDomesticValue string
	DynamicOverseasValue string
	StaticValue          string
	StaticDomesticValue  string
	StaticOverseasValue  string
}
type DescribeDomainFlowDataArgs struct {
	FixTimeGap     string
	TimeMerge      string
	DomainName     string
	EndTime        string
	LocationNameEn string
	StartTime      string
	IspNameEn      string
	OwnerId        int64
	DomainType     string
	SecurityToken  string
	Interval       string
}
type DescribeDomainFlowDataResponse struct {
	RequestId           string
	DomainName          string
	DataInterval        string
	StartTime           string
	EndTime             string
	FlowDataPerInterval DescribeDomainFlowDataDataModuleList
}

type DescribeDomainFlowDataDataModuleList []DescribeDomainFlowDataDataModule

func (list *DescribeDomainFlowDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainFlowDataDataModule)
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

func (c *CdnClient) DescribeLiveStreamSnapshotInfo(req *DescribeLiveStreamSnapshotInfoArgs) (resp *DescribeLiveStreamSnapshotInfoResponse, err error) {
	resp = &DescribeLiveStreamSnapshotInfoResponse{}
	err = c.Request("DescribeLiveStreamSnapshotInfo", req, resp)
	return
}

type DescribeLiveStreamSnapshotInfoLiveStreamSnapshotInfo struct {
	OssEndpoint string
	OssBucket   string
	OssObject   string
	CreateTime  string
}
type DescribeLiveStreamSnapshotInfoArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	Limit         int
	EndTime       string
	StartTime     string
	OwnerId       int64
	StreamName    string
}
type DescribeLiveStreamSnapshotInfoResponse struct {
	RequestId                  string
	NextStartTime              string
	LiveStreamSnapshotInfoList DescribeLiveStreamSnapshotInfoLiveStreamSnapshotInfoList
}

type DescribeLiveStreamSnapshotInfoLiveStreamSnapshotInfoList []DescribeLiveStreamSnapshotInfoLiveStreamSnapshotInfo

func (list *DescribeLiveStreamSnapshotInfoLiveStreamSnapshotInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamSnapshotInfoLiveStreamSnapshotInfo)
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

func (c *CdnClient) DescribeUserDomains(req *DescribeUserDomainsArgs) (resp *DescribeUserDomainsResponse, err error) {
	resp = &DescribeUserDomainsResponse{}
	err = c.Request("DescribeUserDomains", req, resp)
	return
}

type DescribeUserDomainsPageData struct {
	DomainName      string
	Cname           string
	CdnType         string
	DomainStatus    string
	GmtCreated      string
	GmtModified     string
	Description     string
	SourceType      string
	SslProtocol     string
	ResourceGroupId string
	Sandbox         string
	Sources         DescribeUserDomainsSourceList
}
type DescribeUserDomainsArgs struct {
	FuncFilter       string
	Sources          string
	DomainName       string
	OwnerId          int64
	FuncId           string
	PageNumber       int
	DomainStatus     string
	DomainSearchType string
	CheckDomainShow  core.Bool
	ResourceGroupId  string
	SecurityToken    string
	CdnType          string
	PageSize         int
}
type DescribeUserDomainsResponse struct {
	RequestId  string
	PageNumber int64
	PageSize   int64
	TotalCount int64
	Domains    DescribeUserDomainsPageDataList
}

type DescribeUserDomainsSourceList []string

func (list *DescribeUserDomainsSourceList) UnmarshalJSON(data []byte) error {
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

type DescribeUserDomainsPageDataList []DescribeUserDomainsPageData

func (list *DescribeUserDomainsPageDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeUserDomainsPageData)
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

func (c *CdnClient) DeleteLiveDomainMapping(req *DeleteLiveDomainMappingArgs) (resp *DeleteLiveDomainMappingResponse, err error) {
	resp = &DeleteLiveDomainMappingResponse{}
	err = c.Request("DeleteLiveDomainMapping", req, resp)
	return
}

type DeleteLiveDomainMappingArgs struct {
	PullDomain    string
	SecurityToken string
	PushDomain    string
	OwnerId       int64
}
type DeleteLiveDomainMappingResponse struct {
	RequestId string
}

func (c *CdnClient) RefreshObjectCaches(req *RefreshObjectCachesArgs) (resp *RefreshObjectCachesResponse, err error) {
	resp = &RefreshObjectCachesResponse{}
	err = c.Request("RefreshObjectCaches", req, resp)
	return
}

type RefreshObjectCachesArgs struct {
	SecurityToken string
	ObjectPath    string
	OwnerId       int64
	ObjectType    string
}
type RefreshObjectCachesResponse struct {
	RequestId     string
	RefreshTaskId string
}

func (c *CdnClient) StopCdnDomain(req *StopCdnDomainArgs) (resp *StopCdnDomainResponse, err error) {
	resp = &StopCdnDomainResponse{}
	err = c.Request("StopCdnDomain", req, resp)
	return
}

type StopCdnDomainArgs struct {
	SecurityToken string
	DomainName    string
	OwnerId       int64
}
type StopCdnDomainResponse struct {
	RequestId string
}

func (c *CdnClient) BatchDescribeDomainBpsData(req *BatchDescribeDomainBpsDataArgs) (resp *BatchDescribeDomainBpsDataResponse, err error) {
	resp = &BatchDescribeDomainBpsDataResponse{}
	err = c.Request("BatchDescribeDomainBpsData", req, resp)
	return
}

type BatchDescribeDomainBpsDataDataModule struct {
	Timestamp  string
	L1Bps      float32
	DomainName string
}
type BatchDescribeDomainBpsDataArgs struct {
	StartTime     string
	PageNumber    int
	SecurityToken string
	PageSize      int
	DomainName    string
	EndTime       string
	OwnerId       int64
	Version       string
}
type BatchDescribeDomainBpsDataResponse struct {
	RequestId  string
	PageNumber int
	PageSize   int
	TotalCount int
	BpsDatas   BatchDescribeDomainBpsDataDataModuleList
}

type BatchDescribeDomainBpsDataDataModuleList []BatchDescribeDomainBpsDataDataModule

func (list *BatchDescribeDomainBpsDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]BatchDescribeDomainBpsDataDataModule)
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

func (c *CdnClient) SetRemoteReqAuthConfig(req *SetRemoteReqAuthConfigArgs) (resp *SetRemoteReqAuthConfigResponse, err error) {
	resp = &SetRemoteReqAuthConfigResponse{}
	err = c.Request("SetRemoteReqAuthConfig", req, resp)
	return
}

type SetRemoteReqAuthConfigArgs struct {
	AuthPath      string
	DomainName    string
	AuthEnable    string
	OwnerId       int64
	TimeOut       string
	AuthType      string
	AuthProvider  string
	SecurityToken string
	AuthCrash     string
	AuthAddr      string
	AuthFileType  string
}
type SetRemoteReqAuthConfigResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeDomainTopUrlVisit(req *DescribeDomainTopUrlVisitArgs) (resp *DescribeDomainTopUrlVisitResponse, err error) {
	resp = &DescribeDomainTopUrlVisitResponse{}
	err = c.Request("DescribeDomainTopUrlVisit", req, resp)
	return
}

type DescribeDomainTopUrlVisitUrlList struct {
	UrlDetail       string
	VisitData       string
	VisitProportion float32
	Flow            string
	FlowProportion  float32
}
type DescribeDomainTopUrlVisitArgs struct {
	SecurityToken string
	DomainName    string
	SortBy        string
	StartTime     string
	OwnerId       int64
}
type DescribeDomainTopUrlVisitResponse struct {
	RequestId  string
	DomainName string
	StartTime  string
	AllUrlList DescribeDomainTopUrlVisitUrlListList
	Url200List DescribeDomainTopUrlVisitUrlListList
	Url300List DescribeDomainTopUrlVisitUrlListList
	Url400List DescribeDomainTopUrlVisitUrlListList
	Url500List DescribeDomainTopUrlVisitUrlListList
}

type DescribeDomainTopUrlVisitUrlListList []DescribeDomainTopUrlVisitUrlList

func (list *DescribeDomainTopUrlVisitUrlListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainTopUrlVisitUrlList)
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

func (c *CdnClient) DescribeUserVipsByDomain(req *DescribeUserVipsByDomainArgs) (resp *DescribeUserVipsByDomainResponse, err error) {
	resp = &DescribeUserVipsByDomainResponse{}
	err = c.Request("DescribeUserVipsByDomain", req, resp)
	return
}

type DescribeUserVipsByDomainArgs struct {
	SecurityToken string
	DomainName    string
	Available     string
	OwnerId       int64
}
type DescribeUserVipsByDomainResponse struct {
	RequestId  string
	DomainName int64
	Vips       DescribeUserVipsByDomainVipList
}

type DescribeUserVipsByDomainVipList []string

func (list *DescribeUserVipsByDomainVipList) UnmarshalJSON(data []byte) error {
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

func (c *CdnClient) StartMixStreamsService(req *StartMixStreamsServiceArgs) (resp *StartMixStreamsServiceResponse, err error) {
	resp = &StartMixStreamsServiceResponse{}
	err = c.Request("StartMixStreamsService", req, resp)
	return
}

type StartMixStreamsServiceMixStreamsInfo struct {
	DomainName string
	AppName    string
	StreamName string
}
type StartMixStreamsServiceArgs struct {
	MixType        string
	SecurityToken  string
	MainDomainName string
	MixStreamName  string
	MixTemplate    string
	MixDomainName  string
	OwnerId        int64
	MainAppName    string
	MixAppName     string
	MainStreamName string
}
type StartMixStreamsServiceResponse struct {
	RequestId          string
	MixStreamsInfoList StartMixStreamsServiceMixStreamsInfoList
}

type StartMixStreamsServiceMixStreamsInfoList []StartMixStreamsServiceMixStreamsInfo

func (list *StartMixStreamsServiceMixStreamsInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]StartMixStreamsServiceMixStreamsInfo)
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

func (c *CdnClient) DescribeDomainUvData(req *DescribeDomainUvDataArgs) (resp *DescribeDomainUvDataResponse, err error) {
	resp = &DescribeDomainUvDataResponse{}
	err = c.Request("DescribeDomainUvData", req, resp)
	return
}

type DescribeDomainUvDataUsageData struct {
	TimeStamp string
	Value     string
}
type DescribeDomainUvDataArgs struct {
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
}
type DescribeDomainUvDataResponse struct {
	RequestId      string
	DomainName     string
	DataInterval   string
	StartTime      string
	EndTime        string
	UvDataInterval DescribeDomainUvDataUsageDataList
}

type DescribeDomainUvDataUsageDataList []DescribeDomainUvDataUsageData

func (list *DescribeDomainUvDataUsageDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainUvDataUsageData)
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

func (c *CdnClient) StopMixStreamsService(req *StopMixStreamsServiceArgs) (resp *StopMixStreamsServiceResponse, err error) {
	resp = &StopMixStreamsServiceResponse{}
	err = c.Request("StopMixStreamsService", req, resp)
	return
}

type StopMixStreamsServiceMixStreamsInfo struct {
	DomainName string
	AppName    string
	StreamName string
}
type StopMixStreamsServiceArgs struct {
	SecurityToken  string
	MainDomainName string
	MixStreamName  string
	MixDomainName  string
	OwnerId        int64
	MainAppName    string
	MixAppName     string
	MainStreamName string
}
type StopMixStreamsServiceResponse struct {
	RequestId          string
	MixStreamsInfoList StopMixStreamsServiceMixStreamsInfoList
}

type StopMixStreamsServiceMixStreamsInfoList []StopMixStreamsServiceMixStreamsInfo

func (list *StopMixStreamsServiceMixStreamsInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]StopMixStreamsServiceMixStreamsInfo)
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

func (c *CdnClient) SetForceRedirectConfig(req *SetForceRedirectConfigArgs) (resp *SetForceRedirectConfigResponse, err error) {
	resp = &SetForceRedirectConfigResponse{}
	err = c.Request("SetForceRedirectConfig", req, resp)
	return
}

type SetForceRedirectConfigArgs struct {
	SecurityToken string
	DomainName    string
	RedirectType  string
	OwnerId       int64
}
type SetForceRedirectConfigResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeDomainQoSRt(req *DescribeDomainQoSRtArgs) (resp *DescribeDomainQoSRtResponse, err error) {
	resp = &DescribeDomainQoSRtResponse{}
	err = c.Request("DescribeDomainQoSRt", req, resp)
	return
}

type DescribeDomainQoSRtData struct {
	More5s string
	Time   string
	More3s string
}
type DescribeDomainQoSRtArgs struct {
	Ip            string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
	Version       string
	SecurityToken string
}
type DescribeDomainQoSRtResponse struct {
	DomainName string
	StartTime  string
	EndTime    string
	Ip         string
	Content    DescribeDomainQoSRtDataList
}

type DescribeDomainQoSRtDataList []DescribeDomainQoSRtData

func (list *DescribeDomainQoSRtDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainQoSRtData)
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

func (c *CdnClient) ClearUserBlackList(req *ClearUserBlackListArgs) (resp *ClearUserBlackListResponse, err error) {
	resp = &ClearUserBlackListResponse{}
	err = c.Request("ClearUserBlackList", req, resp)
	return
}

type ClearUserBlackListArgs struct {
	SecurityToken string
	OwnerAccount  string
	OwnerId       int64
}
type ClearUserBlackListResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeLiveStreamStreamStatus(req *DescribeLiveStreamStreamStatusArgs) (resp *DescribeLiveStreamStreamStatusResponse, err error) {
	resp = &DescribeLiveStreamStreamStatusResponse{}
	err = c.Request("DescribeLiveStreamStreamStatus", req, resp)
	return
}

type DescribeLiveStreamStreamStatusArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	OwnerId       int64
	StreamName    string
}
type DescribeLiveStreamStreamStatusResponse struct {
	RequestId    string
	StreamStatus string
}

func (c *CdnClient) SetIpAllowListConfig(req *SetIpAllowListConfigArgs) (resp *SetIpAllowListConfigResponse, err error) {
	resp = &SetIpAllowListConfigResponse{}
	err = c.Request("SetIpAllowListConfig", req, resp)
	return
}

type SetIpAllowListConfigArgs struct {
	AllowIps      string
	SecurityToken string
	DomainName    string
	OwnerId       int64
}
type SetIpAllowListConfigResponse struct {
	RequestId string
}

func (c *CdnClient) SetPathForceTtlCodeConfig(req *SetPathForceTtlCodeConfigArgs) (resp *SetPathForceTtlCodeConfigResponse, err error) {
	resp = &SetPathForceTtlCodeConfigResponse{}
	err = c.Request("SetPathForceTtlCodeConfig", req, resp)
	return
}

type SetPathForceTtlCodeConfigArgs struct {
	DomainName    string
	OwnerId       int64
	Version       string
	CodeString    string
	Path          string
	SecurityToken string
}
type SetPathForceTtlCodeConfigResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeDomainHitRateData(req *DescribeDomainHitRateDataArgs) (resp *DescribeDomainHitRateDataResponse, err error) {
	resp = &DescribeDomainHitRateDataResponse{}
	err = c.Request("DescribeDomainHitRateData", req, resp)
	return
}

type DescribeDomainHitRateDataDataModule struct {
	TimeStamp string
	Value     string
}
type DescribeDomainHitRateDataArgs struct {
	SecurityToken  string
	TimeMerge      string
	DomainName     string
	EndTime        string
	LocationNameEn string
	Interval       string
	StartTime      string
	IspNameEn      string
	OwnerId        int64
}
type DescribeDomainHitRateDataResponse struct {
	RequestId       string
	DomainName      string
	DataInterval    string
	StartTime       string
	EndTime         string
	HitRateInterval DescribeDomainHitRateDataDataModuleList
}

type DescribeDomainHitRateDataDataModuleList []DescribeDomainHitRateDataDataModule

func (list *DescribeDomainHitRateDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainHitRateDataDataModule)
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

func (c *CdnClient) DescribeLiveSpecificDomainMapping(req *DescribeLiveSpecificDomainMappingArgs) (resp *DescribeLiveSpecificDomainMappingResponse, err error) {
	resp = &DescribeLiveSpecificDomainMappingResponse{}
	err = c.Request("DescribeLiveSpecificDomainMapping", req, resp)
	return
}

type DescribeLiveSpecificDomainMappingDomainMappingModel struct {
	PushDomain string
	PullDomain string
}
type DescribeLiveSpecificDomainMappingArgs struct {
	PullDomain    string
	SecurityToken string
	PushDomain    string
	OwnerId       int64
}
type DescribeLiveSpecificDomainMappingResponse struct {
	RequestId           string
	DomainMappingModels DescribeLiveSpecificDomainMappingDomainMappingModelList
}

type DescribeLiveSpecificDomainMappingDomainMappingModelList []DescribeLiveSpecificDomainMappingDomainMappingModel

func (list *DescribeLiveSpecificDomainMappingDomainMappingModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveSpecificDomainMappingDomainMappingModel)
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

func (c *CdnClient) SetUserAgentAcessRestriction(req *SetUserAgentAcessRestrictionArgs) (resp *SetUserAgentAcessRestrictionResponse, err error) {
	resp = &SetUserAgentAcessRestrictionResponse{}
	err = c.Request("SetUserAgentAcessRestriction", req, resp)
	return
}

type SetUserAgentAcessRestrictionArgs struct {
	SecurityToken string
	DomainName    string
	UserAgent     string
	OwnerId       int64
	Type          string
}
type SetUserAgentAcessRestrictionResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeDomainRegionData(req *DescribeDomainRegionDataArgs) (resp *DescribeDomainRegionDataResponse, err error) {
	resp = &DescribeDomainRegionDataResponse{}
	err = c.Request("DescribeDomainRegionData", req, resp)
	return
}

type DescribeDomainRegionDataRegionProportionData struct {
	Region          string
	Proportion      string
	RegionEname     string
	AvgObjectSize   string
	AvgResponseTime string
	Bps             string
	ByteHitRate     string
	Qps             string
	ReqErrRate      string
	ReqHitRate      string
	AvgResponseRate string
	TotalBytes      string
	BytesProportion string
	TotalQuery      string
}
type DescribeDomainRegionDataArgs struct {
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
}
type DescribeDomainRegionDataResponse struct {
	RequestId    string
	DomainName   string
	DataInterval string
	StartTime    string
	EndTime      string
	Value        DescribeDomainRegionDataRegionProportionDataList
}

type DescribeDomainRegionDataRegionProportionDataList []DescribeDomainRegionDataRegionProportionData

func (list *DescribeDomainRegionDataRegionProportionDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainRegionDataRegionProportionData)
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

func (c *CdnClient) DescribeLiveStreamRecordIndexFile(req *DescribeLiveStreamRecordIndexFileArgs) (resp *DescribeLiveStreamRecordIndexFileResponse, err error) {
	resp = &DescribeLiveStreamRecordIndexFileResponse{}
	err = c.Request("DescribeLiveStreamRecordIndexFile", req, resp)
	return
}

type DescribeLiveStreamRecordIndexFileRecordIndexInfo struct {
	RecordId   string
	RecordUrl  string
	DomainName string
	AppName    string
	StreamName string
	OssObject  string
	StartTime  string
	EndTime    string
	Duration   float32
	Height     int
	Width      int
	CreateTime string
}
type DescribeLiveStreamRecordIndexFileArgs struct {
	RecordId      string
	AppName       string
	SecurityToken string
	DomainName    string
	OwnerId       int64
	StreamName    string
}
type DescribeLiveStreamRecordIndexFileResponse struct {
	RequestId       string
	RecordIndexInfo DescribeLiveStreamRecordIndexFileRecordIndexInfo
}

func (c *CdnClient) DescribeCdnMonitorData(req *DescribeCdnMonitorDataArgs) (resp *DescribeCdnMonitorDataResponse, err error) {
	resp = &DescribeCdnMonitorDataResponse{}
	err = c.Request("DescribeCdnMonitorData", req, resp)
	return
}

type DescribeCdnMonitorDataCDNMonitorData struct {
	TimeStamp         string
	QueryPerSecond    string
	BytesPerSecond    string
	BytesHitRate      string
	RequestHitRate    string
	AverageObjectSize string
}
type DescribeCdnMonitorDataArgs struct {
	SecurityToken string
	DomainName    string
	EndTime       string
	Interval      string
	StartTime     string
	OwnerId       int64
}
type DescribeCdnMonitorDataResponse struct {
	RequestId       string
	DomainName      string
	MonitorInterval int64
	StartTime       string
	EndTime         string
	MonitorDatas    DescribeCdnMonitorDataCDNMonitorDataList
}

type DescribeCdnMonitorDataCDNMonitorDataList []DescribeCdnMonitorDataCDNMonitorData

func (list *DescribeCdnMonitorDataCDNMonitorDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCdnMonitorDataCDNMonitorData)
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

func (c *CdnClient) SetRemoveQueryStringConfig(req *SetRemoveQueryStringConfigArgs) (resp *SetRemoveQueryStringConfigResponse, err error) {
	resp = &SetRemoveQueryStringConfigResponse{}
	err = c.Request("SetRemoveQueryStringConfig", req, resp)
	return
}

type SetRemoveQueryStringConfigArgs struct {
	KeepOssArgs   string
	SecurityToken string
	DomainName    string
	AliRemoveArgs string
	OwnerId       int64
}
type SetRemoveQueryStringConfigResponse struct {
	RequestId string
}

func (c *CdnClient) SetRangeConfig(req *SetRangeConfigArgs) (resp *SetRangeConfigResponse, err error) {
	resp = &SetRangeConfigResponse{}
	err = c.Request("SetRangeConfig", req, resp)
	return
}

type SetRangeConfigArgs struct {
	SecurityToken string
	Enable        string
	DomainName    string
	OwnerId       int64
}
type SetRangeConfigResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeLiveStreamHlsOnlineUserNumByDomain(req *DescribeLiveStreamHlsOnlineUserNumByDomainArgs) (resp *DescribeLiveStreamHlsOnlineUserNumByDomainResponse, err error) {
	resp = &DescribeLiveStreamHlsOnlineUserNumByDomainResponse{}
	err = c.Request("DescribeLiveStreamHlsOnlineUserNumByDomain", req, resp)
	return
}

type DescribeLiveStreamHlsOnlineUserNumByDomainLiveStreamOnlineUserNumInfo struct {
	StreamUrl  string
	UserNumber int64
	Time       string
}
type DescribeLiveStreamHlsOnlineUserNumByDomainArgs struct {
	AppName       string
	SecurityToken string
	HlsSwitch     string
	DomainName    string
	PageSize      int64
	OwnerId       int64
	PageNumber    int64
}
type DescribeLiveStreamHlsOnlineUserNumByDomainResponse struct {
	RequestId       string
	TotalUserNumber int64
	Count           int64
	PageNumber      int64
	PageSize        int64
	OnlineUserInfo  DescribeLiveStreamHlsOnlineUserNumByDomainLiveStreamOnlineUserNumInfoList
}

type DescribeLiveStreamHlsOnlineUserNumByDomainLiveStreamOnlineUserNumInfoList []DescribeLiveStreamHlsOnlineUserNumByDomainLiveStreamOnlineUserNumInfo

func (list *DescribeLiveStreamHlsOnlineUserNumByDomainLiveStreamOnlineUserNumInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamHlsOnlineUserNumByDomainLiveStreamOnlineUserNumInfo)
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

func (c *CdnClient) DescribeLiveStreamFrameInfo(req *DescribeLiveStreamFrameInfoArgs) (resp *DescribeLiveStreamFrameInfoResponse, err error) {
	resp = &DescribeLiveStreamFrameInfoResponse{}
	err = c.Request("DescribeLiveStreamFrameInfo", req, resp)
	return
}

type DescribeLiveStreamFrameInfoFrameDataModel struct {
	Time       string
	Stream     string
	ClientAddr string
	Server     string
	AudioRate  float32
	AudioByte  float32
	FrameRate  float32
	FrameByte  float32
}
type DescribeLiveStreamFrameInfoArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
	StreamName    string
}
type DescribeLiveStreamFrameInfoResponse struct {
	RequestId      string
	FrameDataInfos DescribeLiveStreamFrameInfoFrameDataModelList
}

type DescribeLiveStreamFrameInfoFrameDataModelList []DescribeLiveStreamFrameInfoFrameDataModel

func (list *DescribeLiveStreamFrameInfoFrameDataModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamFrameInfoFrameDataModel)
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

func (c *CdnClient) DescribeDomainFileSizeProportionData(req *DescribeDomainFileSizeProportionDataArgs) (resp *DescribeDomainFileSizeProportionDataResponse, err error) {
	resp = &DescribeDomainFileSizeProportionDataResponse{}
	err = c.Request("DescribeDomainFileSizeProportionData", req, resp)
	return
}

type DescribeDomainFileSizeProportionDataUsageData struct {
	TimeStamp string
	Value     DescribeDomainFileSizeProportionDataFileSizeProportionDataList
}

type DescribeDomainFileSizeProportionDataFileSizeProportionData struct {
	FileSize   string
	Proportion string
}
type DescribeDomainFileSizeProportionDataArgs struct {
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
}
type DescribeDomainFileSizeProportionDataResponse struct {
	RequestId                      string
	DomainName                     string
	DataInterval                   string
	StartTime                      string
	EndTime                        string
	FileSizeProportionDataInterval DescribeDomainFileSizeProportionDataUsageDataList
}

type DescribeDomainFileSizeProportionDataFileSizeProportionDataList []DescribeDomainFileSizeProportionDataFileSizeProportionData

func (list *DescribeDomainFileSizeProportionDataFileSizeProportionDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainFileSizeProportionDataFileSizeProportionData)
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

type DescribeDomainFileSizeProportionDataUsageDataList []DescribeDomainFileSizeProportionDataUsageData

func (list *DescribeDomainFileSizeProportionDataUsageDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainFileSizeProportionDataUsageData)
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

func (c *CdnClient) SetHttpsOptionConfig(req *SetHttpsOptionConfigArgs) (resp *SetHttpsOptionConfigResponse, err error) {
	resp = &SetHttpsOptionConfigResponse{}
	err = c.Request("SetHttpsOptionConfig", req, resp)
	return
}

type SetHttpsOptionConfigArgs struct {
	SecurityToken string
	DomainName    string
	Http2         string
	OwnerId       int64
}
type SetHttpsOptionConfigResponse struct {
	RequestId string
}

func (c *CdnClient) UpdateLiveAppSnapshotConfig(req *UpdateLiveAppSnapshotConfigArgs) (resp *UpdateLiveAppSnapshotConfigResponse, err error) {
	resp = &UpdateLiveAppSnapshotConfigResponse{}
	err = c.Request("UpdateLiveAppSnapshotConfig", req, resp)
	return
}

type UpdateLiveAppSnapshotConfigArgs struct {
	TimeInterval       int
	OssBucket          string
	AppName            string
	SecurityToken      string
	DomainName         string
	OssEndpoint        string
	SequenceOssObject  string
	OverwriteOssObject string
	OwnerId            int64
}
type UpdateLiveAppSnapshotConfigResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeLiveAppRecordConfig(req *DescribeLiveAppRecordConfigArgs) (resp *DescribeLiveAppRecordConfigResponse, err error) {
	resp = &DescribeLiveAppRecordConfigResponse{}
	err = c.Request("DescribeLiveAppRecordConfig", req, resp)
	return
}

type DescribeLiveAppRecordConfigLiveAppRecord struct {
	DomainName      string
	AppName         string
	OssEndpoint     string
	OssBucket       string
	OssObjectPrefix string
	CreateTime      string
}
type DescribeLiveAppRecordConfigArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	OwnerId       int64
}
type DescribeLiveAppRecordConfigResponse struct {
	RequestId     string
	LiveAppRecord DescribeLiveAppRecordConfigLiveAppRecord
}

func (c *CdnClient) SetReqAuthConfig(req *SetReqAuthConfigArgs) (resp *SetReqAuthConfigResponse, err error) {
	resp = &SetReqAuthConfigResponse{}
	err = c.Request("SetReqAuthConfig", req, resp)
	return
}

type SetReqAuthConfigArgs struct {
	Key1          string
	Key2          string
	SecurityToken string
	DomainName    string
	OwnerId       int64
	TimeOut       string
	AuthType      string
}
type SetReqAuthConfigResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeLiveSnapshotConfig(req *DescribeLiveSnapshotConfigArgs) (resp *DescribeLiveSnapshotConfigResponse, err error) {
	resp = &DescribeLiveSnapshotConfigResponse{}
	err = c.Request("DescribeLiveSnapshotConfig", req, resp)
	return
}

type DescribeLiveSnapshotConfigLiveStreamSnapshotConfig struct {
	DomainName         string
	AppName            string
	TimeInterval       int
	OssEndpoint        string
	OssBucket          string
	OverwriteOssObject string
	SequenceOssObject  string
	CreateTime         string
}
type DescribeLiveSnapshotConfigArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	PageSize      int
	OwnerId       int64
	PageNum       int
	StreamName    string
	Order         string
}
type DescribeLiveSnapshotConfigResponse struct {
	RequestId                    string
	PageNum                      int
	PageSize                     int
	Order                        string
	TotalNum                     int
	TotalPage                    int
	LiveStreamSnapshotConfigList DescribeLiveSnapshotConfigLiveStreamSnapshotConfigList
}

type DescribeLiveSnapshotConfigLiveStreamSnapshotConfigList []DescribeLiveSnapshotConfigLiveStreamSnapshotConfig

func (list *DescribeLiveSnapshotConfigLiveStreamSnapshotConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveSnapshotConfigLiveStreamSnapshotConfig)
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

func (c *CdnClient) DescribeDomainRealTimeData(req *DescribeDomainRealTimeDataArgs) (resp *DescribeDomainRealTimeDataResponse, err error) {
	resp = &DescribeDomainRealTimeDataResponse{}
	err = c.Request("DescribeDomainRealTimeData", req, resp)
	return
}

type DescribeDomainRealTimeDataDataModule struct {
	TimeStamp string
	Value     string
}
type DescribeDomainRealTimeDataArgs struct {
	Field         string
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
}
type DescribeDomainRealTimeDataResponse struct {
	RequestId       string
	DomainName      string
	Field           string
	StartTime       string
	EndTime         string
	DataPerInterval DescribeDomainRealTimeDataDataModuleList
}

type DescribeDomainRealTimeDataDataModuleList []DescribeDomainRealTimeDataDataModule

func (list *DescribeDomainRealTimeDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainRealTimeDataDataModule)
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

func (c *CdnClient) SetForwardSchemeConfig(req *SetForwardSchemeConfigArgs) (resp *SetForwardSchemeConfigResponse, err error) {
	resp = &SetForwardSchemeConfigResponse{}
	err = c.Request("SetForwardSchemeConfig", req, resp)
	return
}

type SetForwardSchemeConfigArgs struct {
	SchemeOrigin     string
	SchemeOriginPort string
	SecurityToken    string
	Enable           string
	DomainName       string
	OwnerId          int64
}
type SetForwardSchemeConfigResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeDomainHttpsData(req *DescribeDomainHttpsDataArgs) (resp *DescribeDomainHttpsDataResponse, err error) {
	resp = &DescribeDomainHttpsDataResponse{}
	err = c.Request("DescribeDomainHttpsData", req, resp)
	return
}

type DescribeDomainHttpsDataHttpsStatisticsInfo struct {
	Time               string
	L1HttpsBps         float32
	L1HttpsInnerBps    float32
	L1HttpsOutBps      float32
	L1HttpsQps         int64
	L1HttpsInnerQps    int64
	L1HttpsOutQps      int64
	L1HttpsTtraf       int64
	HttpsSrcBps        int64
	HttpsSrcTraf       int64
	L1HttpsInnerTraf   int64
	L1HttpsOutTraf     int64
	HttpsByteHitRate   float32
	HttpsReqHitRate    float32
	L1HttpsHitRate     float32
	L1HttpsInner_acc   float32
	L1HttpsOut_acc     float32
	L1HttpsTacc        float32
	L1DyHttpsBps       float32
	L1DyHttpsInnerBps  float32
	L1DyHttpsOutBps    float32
	L1StHttpsBps       float32
	L1StHttpsInnerBps  float32
	L1StHttpsOutBps    float32
	L1DyHttpsTraf      float32
	L1DyHttpsInnerTraf float32
	L1DyHttpsOutTraf   float32
	L1StHttpsTraf      float32
	L1StHttpsInnerTraf float32
	L1StHttpsOutTraf   float32
	L1DyHttpsQps       float32
	L1DyHttpsInnerQps  float32
	L1DyHttpsOutQps    float32
	L1StHttpsQps       float32
	L1StHttpsInnerQps  float32
	L1StHttpsOutQps    float32
	L1DyHttpsAcc       float32
	L1DyHttpsInnerAcc  float32
	L1DyHttpsOutAcc    float32
	L1StHttpsAcc       float32
	L1StHttpsInnerAcc  float32
	L1StHttpsOutAcc    float32
}
type DescribeDomainHttpsDataArgs struct {
	DomainType    string
	FixTimeGap    string
	SecurityToken string
	TimeMerge     string
	DomainNames   string
	EndTime       string
	Interval      string
	StartTime     string
	Cls           string
	OwnerId       int64
}
type DescribeDomainHttpsDataResponse struct {
	RequestId            string
	DomainNames          string
	DataInterval         string
	HttpsStatisticsInfos DescribeDomainHttpsDataHttpsStatisticsInfoList
}

type DescribeDomainHttpsDataHttpsStatisticsInfoList []DescribeDomainHttpsDataHttpsStatisticsInfo

func (list *DescribeDomainHttpsDataHttpsStatisticsInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainHttpsDataHttpsStatisticsInfo)
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

func (c *CdnClient) DescribeLiveStreamFrameAndBitRateByDomain(req *DescribeLiveStreamFrameAndBitRateByDomainArgs) (resp *DescribeLiveStreamFrameAndBitRateByDomainResponse, err error) {
	resp = &DescribeLiveStreamFrameAndBitRateByDomainResponse{}
	err = c.Request("DescribeLiveStreamFrameAndBitRateByDomain", req, resp)
	return
}

type DescribeLiveStreamFrameAndBitRateByDomainFrameRateAndBitRateInfo struct {
	StreamUrl      string
	VideoFrameRate float32
	AudioFrameRate float32
	BitRate        float32
	Time           string
}
type DescribeLiveStreamFrameAndBitRateByDomainArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	PageSize      int64
	OwnerId       int64
	PageNumber    int64
}
type DescribeLiveStreamFrameAndBitRateByDomainResponse struct {
	RequestId                string
	Count                    int64
	PageNumber               int64
	PageSize                 int64
	FrameRateAndBitRateInfos DescribeLiveStreamFrameAndBitRateByDomainFrameRateAndBitRateInfoList
}

type DescribeLiveStreamFrameAndBitRateByDomainFrameRateAndBitRateInfoList []DescribeLiveStreamFrameAndBitRateByDomainFrameRateAndBitRateInfo

func (list *DescribeLiveStreamFrameAndBitRateByDomainFrameRateAndBitRateInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamFrameAndBitRateByDomainFrameRateAndBitRateInfo)
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

func (c *CdnClient) DescribeLiveStreamsControlHistory(req *DescribeLiveStreamsControlHistoryArgs) (resp *DescribeLiveStreamsControlHistoryResponse, err error) {
	resp = &DescribeLiveStreamsControlHistoryResponse{}
	err = c.Request("DescribeLiveStreamsControlHistory", req, resp)
	return
}

type DescribeLiveStreamsControlHistoryLiveStreamControlInfo struct {
	StreamName string
	ClientIP   string
	Action     string
	TimeStamp  string
}
type DescribeLiveStreamsControlHistoryArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
}
type DescribeLiveStreamsControlHistoryResponse struct {
	RequestId   string
	ControlInfo DescribeLiveStreamsControlHistoryLiveStreamControlInfoList
}

type DescribeLiveStreamsControlHistoryLiveStreamControlInfoList []DescribeLiveStreamsControlHistoryLiveStreamControlInfo

func (list *DescribeLiveStreamsControlHistoryLiveStreamControlInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamsControlHistoryLiveStreamControlInfo)
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

func (c *CdnClient) AddLiveAppRecordConfig(req *AddLiveAppRecordConfigArgs) (resp *AddLiveAppRecordConfigResponse, err error) {
	resp = &AddLiveAppRecordConfigResponse{}
	err = c.Request("AddLiveAppRecordConfig", req, resp)
	return
}

type AddLiveAppRecordConfigArgs struct {
	OssBucket       string
	AppName         string
	SecurityToken   string
	DomainName      string
	OssEndpoint     string
	OssObjectPrefix string
	OwnerId         int64
}
type AddLiveAppRecordConfigResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeLiveStreamRoomUserNumber(req *DescribeLiveStreamRoomUserNumberArgs) (resp *DescribeLiveStreamRoomUserNumberResponse, err error) {
	resp = &DescribeLiveStreamRoomUserNumberResponse{}
	err = c.Request("DescribeLiveStreamRoomUserNumber", req, resp)
	return
}

type DescribeLiveStreamRoomUserNumberLiveStreamOnlineUserNumInfo struct {
	StreamUrl  string
	UserNumber int64
	Time       string
}
type DescribeLiveStreamRoomUserNumberArgs struct {
	AppName       string
	SecurityToken string
	HlsSwitch     string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
	StreamName    string
}
type DescribeLiveStreamRoomUserNumberResponse struct {
	RequestId       string
	TotalUserNumber int64
	OnlineUserInfo  DescribeLiveStreamRoomUserNumberLiveStreamOnlineUserNumInfoList
}

type DescribeLiveStreamRoomUserNumberLiveStreamOnlineUserNumInfoList []DescribeLiveStreamRoomUserNumberLiveStreamOnlineUserNumInfo

func (list *DescribeLiveStreamRoomUserNumberLiveStreamOnlineUserNumInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamRoomUserNumberLiveStreamOnlineUserNumInfo)
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

func (c *CdnClient) DeleteCdnDomain(req *DeleteCdnDomainArgs) (resp *DeleteCdnDomainResponse, err error) {
	resp = &DeleteCdnDomainResponse{}
	err = c.Request("DeleteCdnDomain", req, resp)
	return
}

type DeleteCdnDomainArgs struct {
	ResourceGroupId string
	SecurityToken   string
	DomainName      string
	OwnerId         int64
}
type DeleteCdnDomainResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeDomainDownstreamBpsOfEdge(req *DescribeDomainDownstreamBpsOfEdgeArgs) (resp *DescribeDomainDownstreamBpsOfEdgeResponse, err error) {
	resp = &DescribeDomainDownstreamBpsOfEdgeResponse{}
	err = c.Request("DescribeDomainDownstreamBpsOfEdge", req, resp)
	return
}

type DescribeDomainDownstreamBpsOfEdgeDomainBpsModel struct {
	Time string
	Bps  float32
}
type DescribeDomainDownstreamBpsOfEdgeArgs struct {
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
}
type DescribeDomainDownstreamBpsOfEdgeResponse struct {
	RequestId string
	BpsDatas  DescribeDomainDownstreamBpsOfEdgeDomainBpsModelList
}

type DescribeDomainDownstreamBpsOfEdgeDomainBpsModelList []DescribeDomainDownstreamBpsOfEdgeDomainBpsModel

func (list *DescribeDomainDownstreamBpsOfEdgeDomainBpsModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainDownstreamBpsOfEdgeDomainBpsModel)
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

func (c *CdnClient) DescribeLiveStreamLimitInfo(req *DescribeLiveStreamLimitInfoArgs) (resp *DescribeLiveStreamLimitInfoResponse, err error) {
	resp = &DescribeLiveStreamLimitInfoResponse{}
	err = c.Request("DescribeLiveStreamLimitInfo", req, resp)
	return
}

type DescribeLiveStreamLimitInfoUserLimitMode struct {
	LimitDomain       string
	LimitNum          string
	LimitTranscodeNum string
}
type DescribeLiveStreamLimitInfoArgs struct {
	SecurityToken string
	OwnerId       int64
	LimitDomain   string
}
type DescribeLiveStreamLimitInfoResponse struct {
	RequestId      string
	UserLimitLists DescribeLiveStreamLimitInfoUserLimitModeList
}

type DescribeLiveStreamLimitInfoUserLimitModeList []DescribeLiveStreamLimitInfoUserLimitMode

func (list *DescribeLiveStreamLimitInfoUserLimitModeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamLimitInfoUserLimitMode)
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

func (c *CdnClient) DescribeLiveStreamPushErrors(req *DescribeLiveStreamPushErrorsArgs) (resp *DescribeLiveStreamPushErrorsResponse, err error) {
	resp = &DescribeLiveStreamPushErrorsResponse{}
	err = c.Request("DescribeLiveStreamPushErrors", req, resp)
	return
}

type DescribeLiveStreamPushErrorsPushErrorsModel struct {
	ErrorCode string
}
type DescribeLiveStreamPushErrorsArgs struct {
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
}
type DescribeLiveStreamPushErrorsResponse struct {
	RequestId           string
	PushErrorsModelList DescribeLiveStreamPushErrorsPushErrorsModelList
}

type DescribeLiveStreamPushErrorsPushErrorsModelList []DescribeLiveStreamPushErrorsPushErrorsModel

func (list *DescribeLiveStreamPushErrorsPushErrorsModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamPushErrorsPushErrorsModel)
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

func (c *CdnClient) DescribeLiveStreamRecordContent(req *DescribeLiveStreamRecordContentArgs) (resp *DescribeLiveStreamRecordContentResponse, err error) {
	resp = &DescribeLiveStreamRecordContentResponse{}
	err = c.Request("DescribeLiveStreamRecordContent", req, resp)
	return
}

type DescribeLiveStreamRecordContentRecordContentInfo struct {
	OssEndpoint     string
	OssBucket       string
	OssObjectPrefix string
	StartTime       string
	EndTime         string
	Duration        float32
}
type DescribeLiveStreamRecordContentArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
	StreamName    string
}
type DescribeLiveStreamRecordContentResponse struct {
	RequestId             string
	RecordContentInfoList DescribeLiveStreamRecordContentRecordContentInfoList
}

type DescribeLiveStreamRecordContentRecordContentInfoList []DescribeLiveStreamRecordContentRecordContentInfo

func (list *DescribeLiveStreamRecordContentRecordContentInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamRecordContentRecordContentInfo)
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

func (c *CdnClient) AddLiveDomainMapping(req *AddLiveDomainMappingArgs) (resp *AddLiveDomainMappingResponse, err error) {
	resp = &AddLiveDomainMappingResponse{}
	err = c.Request("AddLiveDomainMapping", req, resp)
	return
}

type AddLiveDomainMappingArgs struct {
	PullDomain    string
	SecurityToken string
	PushDomain    string
	OwnerId       int64
}
type AddLiveDomainMappingResponse struct {
	RequestId string
}

func (c *CdnClient) ModifyHttpHeaderConfig(req *ModifyHttpHeaderConfigArgs) (resp *ModifyHttpHeaderConfigResponse, err error) {
	resp = &ModifyHttpHeaderConfigResponse{}
	err = c.Request("ModifyHttpHeaderConfig", req, resp)
	return
}

type ModifyHttpHeaderConfigArgs struct {
	HeaderValue   string
	SecurityToken string
	ConfigID      string
	DomainName    string
	HeaderKey     string
	OwnerId       int64
}
type ModifyHttpHeaderConfigResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeLiveStreamRelayPushBitRate(req *DescribeLiveStreamRelayPushBitRateArgs) (resp *DescribeLiveStreamRelayPushBitRateResponse, err error) {
	resp = &DescribeLiveStreamRelayPushBitRateResponse{}
	err = c.Request("DescribeLiveStreamRelayPushBitRate", req, resp)
	return
}

type DescribeLiveStreamRelayPushBitRateRelayPushBitRateModel struct {
	Time          string
	VedioFrame    string
	VedioTimstamp string
	AudioFrame    string
	AudioTimstamp string
	RelayDomain   string
}
type DescribeLiveStreamRelayPushBitRateArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	OwnerId       int64
	StreamName    string
}
type DescribeLiveStreamRelayPushBitRateResponse struct {
	RequestId                 string
	RelayPushBitRateModelList DescribeLiveStreamRelayPushBitRateRelayPushBitRateModelList
}

type DescribeLiveStreamRelayPushBitRateRelayPushBitRateModelList []DescribeLiveStreamRelayPushBitRateRelayPushBitRateModel

func (list *DescribeLiveStreamRelayPushBitRateRelayPushBitRateModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamRelayPushBitRateRelayPushBitRateModel)
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

func (c *CdnClient) DescribeTopDomainsByFlow(req *DescribeTopDomainsByFlowArgs) (resp *DescribeTopDomainsByFlowResponse, err error) {
	resp = &DescribeTopDomainsByFlowResponse{}
	err = c.Request("DescribeTopDomainsByFlow", req, resp)
	return
}

type DescribeTopDomainsByFlowTopDomain struct {
	DomainName     string
	Rank           int64
	TotalTraffic   string
	TrafficPercent string
	MaxBps         int64
	MaxBpsTime     string
	TotalAccess    int64
}
type DescribeTopDomainsByFlowArgs struct {
	SecurityToken string
	Limit         int64
	EndTime       string
	StartTime     string
	OwnerId       int64
}
type DescribeTopDomainsByFlowResponse struct {
	RequestId         string
	StartTime         string
	EndTime           string
	DomainCount       int64
	DomainOnlineCount int64
	TopDomains        DescribeTopDomainsByFlowTopDomainList
}

type DescribeTopDomainsByFlowTopDomainList []DescribeTopDomainsByFlowTopDomain

func (list *DescribeTopDomainsByFlowTopDomainList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTopDomainsByFlowTopDomain)
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

func (c *CdnClient) DescribeLiveStreamOnlineUserNumByDomain(req *DescribeLiveStreamOnlineUserNumByDomainArgs) (resp *DescribeLiveStreamOnlineUserNumByDomainResponse, err error) {
	resp = &DescribeLiveStreamOnlineUserNumByDomainResponse{}
	err = c.Request("DescribeLiveStreamOnlineUserNumByDomain", req, resp)
	return
}

type DescribeLiveStreamOnlineUserNumByDomainLiveStreamOnlineUserNumInfo struct {
	StreamUrl  string
	UserNumber int64
	Time       string
}
type DescribeLiveStreamOnlineUserNumByDomainArgs struct {
	AppName       string
	SecurityToken string
	HlsSwitch     string
	DomainName    string
	PageSize      int64
	OwnerId       int64
	PageNumber    int64
}
type DescribeLiveStreamOnlineUserNumByDomainResponse struct {
	RequestId       string
	TotalUserNumber int64
	Count           int64
	PageNumber      int64
	PageSize        int64
	OnlineUserInfo  DescribeLiveStreamOnlineUserNumByDomainLiveStreamOnlineUserNumInfoList
}

type DescribeLiveStreamOnlineUserNumByDomainLiveStreamOnlineUserNumInfoList []DescribeLiveStreamOnlineUserNumByDomainLiveStreamOnlineUserNumInfo

func (list *DescribeLiveStreamOnlineUserNumByDomainLiveStreamOnlineUserNumInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamOnlineUserNumByDomainLiveStreamOnlineUserNumInfo)
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

func (c *CdnClient) SetIpBlackListConfig(req *SetIpBlackListConfigArgs) (resp *SetIpBlackListConfigResponse, err error) {
	resp = &SetIpBlackListConfigResponse{}
	err = c.Request("SetIpBlackListConfig", req, resp)
	return
}

type SetIpBlackListConfigArgs struct {
	SecurityToken string
	DomainName    string
	OwnerId       int64
	BlockIps      string
}
type SetIpBlackListConfigResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeDomainSrcFlowData(req *DescribeDomainSrcFlowDataArgs) (resp *DescribeDomainSrcFlowDataResponse, err error) {
	resp = &DescribeDomainSrcFlowDataResponse{}
	err = c.Request("DescribeDomainSrcFlowData", req, resp)
	return
}

type DescribeDomainSrcFlowDataDataModule struct {
	TimeStamp string
	Value     string
}
type DescribeDomainSrcFlowDataArgs struct {
	FixTimeGap    string
	SecurityToken string
	TimeMerge     string
	DomainName    string
	EndTime       string
	Interval      string
	StartTime     string
	OwnerId       int64
}
type DescribeDomainSrcFlowDataResponse struct {
	RequestId              string
	DomainName             string
	DataInterval           string
	StartTime              string
	EndTime                string
	SrcFlowDataPerInterval DescribeDomainSrcFlowDataDataModuleList
}

type DescribeDomainSrcFlowDataDataModuleList []DescribeDomainSrcFlowDataDataModule

func (list *DescribeDomainSrcFlowDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainSrcFlowDataDataModule)
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

func (c *CdnClient) DescribeExtensiveDomainData(req *DescribeExtensiveDomainDataArgs) (resp *DescribeExtensiveDomainDataResponse, err error) {
	resp = &DescribeExtensiveDomainDataResponse{}
	err = c.Request("DescribeExtensiveDomainData", req, resp)
	return
}

type DescribeExtensiveDomainDataUsageData struct {
	ExactDomain string
	TimeStamp   string
	Acc         string
	Flow        string
}
type DescribeExtensiveDomainDataArgs struct {
	SecurityToken   string
	ExtensiveDomain string
	PageSize        int
	EndTime         string
	StartTime       string
	OwnerId         int64
	PageNumber      int
}
type DescribeExtensiveDomainDataResponse struct {
	RequestId       string
	ExtensiveDomain string
	DataInterval    string
	StartTime       string
	EndTime         string
	PageNumber      string
	TotalCount      string
	PageSize        string
	DataPerInterval DescribeExtensiveDomainDataUsageDataList
}

type DescribeExtensiveDomainDataUsageDataList []DescribeExtensiveDomainDataUsageData

func (list *DescribeExtensiveDomainDataUsageDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeExtensiveDomainDataUsageData)
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

func (c *CdnClient) SetFileTypeForceTtlCodeConfig(req *SetFileTypeForceTtlCodeConfigArgs) (resp *SetFileTypeForceTtlCodeConfigResponse, err error) {
	resp = &SetFileTypeForceTtlCodeConfigResponse{}
	err = c.Request("SetFileTypeForceTtlCodeConfig", req, resp)
	return
}

type SetFileTypeForceTtlCodeConfigArgs struct {
	FileType      string
	DomainName    string
	OwnerId       int64
	Version       string
	CodeString    string
	SecurityToken string
}
type SetFileTypeForceTtlCodeConfigResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeDomainUpstreamBpsOfEdge(req *DescribeDomainUpstreamBpsOfEdgeArgs) (resp *DescribeDomainUpstreamBpsOfEdgeResponse, err error) {
	resp = &DescribeDomainUpstreamBpsOfEdgeResponse{}
	err = c.Request("DescribeDomainUpstreamBpsOfEdge", req, resp)
	return
}

type DescribeDomainUpstreamBpsOfEdgeDomainBpsModel struct {
	Time string
	Bps  float32
}
type DescribeDomainUpstreamBpsOfEdgeArgs struct {
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
}
type DescribeDomainUpstreamBpsOfEdgeResponse struct {
	RequestId string
	BpsDatas  DescribeDomainUpstreamBpsOfEdgeDomainBpsModelList
}

type DescribeDomainUpstreamBpsOfEdgeDomainBpsModelList []DescribeDomainUpstreamBpsOfEdgeDomainBpsModel

func (list *DescribeDomainUpstreamBpsOfEdgeDomainBpsModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainUpstreamBpsOfEdgeDomainBpsModel)
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

func (c *CdnClient) DescribeLiveStreamRoomBitRate(req *DescribeLiveStreamRoomBitRateArgs) (resp *DescribeLiveStreamRoomBitRateResponse, err error) {
	resp = &DescribeLiveStreamRoomBitRateResponse{}
	err = c.Request("DescribeLiveStreamRoomBitRate", req, resp)
	return
}

type DescribeLiveStreamRoomBitRateFrameRateAndBitRateInfo struct {
	StreamUrl      string
	VideoFrameRate float32
	AudioFrameRate float32
	BitRate        float32
	Time           string
}
type DescribeLiveStreamRoomBitRateArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
	StreamName    string
}
type DescribeLiveStreamRoomBitRateResponse struct {
	RequestId                string
	FrameRateAndBitRateInfos DescribeLiveStreamRoomBitRateFrameRateAndBitRateInfoList
}

type DescribeLiveStreamRoomBitRateFrameRateAndBitRateInfoList []DescribeLiveStreamRoomBitRateFrameRateAndBitRateInfo

func (list *DescribeLiveStreamRoomBitRateFrameRateAndBitRateInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamRoomBitRateFrameRateAndBitRateInfo)
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

func (c *CdnClient) DescribeCdnRegionAndIsp(req *DescribeCdnRegionAndIspArgs) (resp *DescribeCdnRegionAndIspResponse, err error) {
	resp = &DescribeCdnRegionAndIspResponse{}
	err = c.Request("DescribeCdnRegionAndIsp", req, resp)
	return
}

type DescribeCdnRegionAndIspRegion struct {
	NameZh string
	NameEn string
}

type DescribeCdnRegionAndIspIsp struct {
	NameZh string
	NameEn string
}
type DescribeCdnRegionAndIspArgs struct {
	SecurityToken string
	OwnerId       int64
}
type DescribeCdnRegionAndIspResponse struct {
	RequestId string
	Regions   DescribeCdnRegionAndIspRegionList
	Isps      DescribeCdnRegionAndIspIspList
}

type DescribeCdnRegionAndIspRegionList []DescribeCdnRegionAndIspRegion

func (list *DescribeCdnRegionAndIspRegionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCdnRegionAndIspRegion)
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

type DescribeCdnRegionAndIspIspList []DescribeCdnRegionAndIspIsp

func (list *DescribeCdnRegionAndIspIspList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCdnRegionAndIspIsp)
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

func (c *CdnClient) DeleteHttpHeaderConfig(req *DeleteHttpHeaderConfigArgs) (resp *DeleteHttpHeaderConfigResponse, err error) {
	resp = &DeleteHttpHeaderConfigResponse{}
	err = c.Request("DeleteHttpHeaderConfig", req, resp)
	return
}

type DeleteHttpHeaderConfigArgs struct {
	SecurityToken string
	ConfigID      string
	DomainName    string
	OwnerId       int64
}
type DeleteHttpHeaderConfigResponse struct {
	RequestId string
}

func (c *CdnClient) StartCdnDomain(req *StartCdnDomainArgs) (resp *StartCdnDomainResponse, err error) {
	resp = &StartCdnDomainResponse{}
	err = c.Request("StartCdnDomain", req, resp)
	return
}

type StartCdnDomainArgs struct {
	SecurityToken string
	DomainName    string
	OwnerId       int64
}
type StartCdnDomainResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeLiveStreamOnlineUserNum(req *DescribeLiveStreamOnlineUserNumArgs) (resp *DescribeLiveStreamOnlineUserNumResponse, err error) {
	resp = &DescribeLiveStreamOnlineUserNumResponse{}
	err = c.Request("DescribeLiveStreamOnlineUserNum", req, resp)
	return
}

type DescribeLiveStreamOnlineUserNumLiveStreamOnlineUserNumInfo struct {
	StreamUrl  string
	UserNumber int64
	Time       string
}
type DescribeLiveStreamOnlineUserNumArgs struct {
	AppName       string
	SecurityToken string
	HlsSwitch     string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
	StreamName    string
}
type DescribeLiveStreamOnlineUserNumResponse struct {
	RequestId       string
	TotalUserNumber int64
	OnlineUserInfo  DescribeLiveStreamOnlineUserNumLiveStreamOnlineUserNumInfoList
}

type DescribeLiveStreamOnlineUserNumLiveStreamOnlineUserNumInfoList []DescribeLiveStreamOnlineUserNumLiveStreamOnlineUserNumInfo

func (list *DescribeLiveStreamOnlineUserNumLiveStreamOnlineUserNumInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamOnlineUserNumLiveStreamOnlineUserNumInfo)
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

func (c *CdnClient) DeleteLiveAppSnapshotConfig(req *DeleteLiveAppSnapshotConfigArgs) (resp *DeleteLiveAppSnapshotConfigResponse, err error) {
	resp = &DeleteLiveAppSnapshotConfigResponse{}
	err = c.Request("DeleteLiveAppSnapshotConfig", req, resp)
	return
}

type DeleteLiveAppSnapshotConfigArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	OwnerId       int64
}
type DeleteLiveAppSnapshotConfigResponse struct {
	RequestId string
}

func (c *CdnClient) GetUserDomainBlackList(req *GetUserDomainBlackListArgs) (resp *GetUserDomainBlackListResponse, err error) {
	resp = &GetUserDomainBlackListResponse{}
	err = c.Request("GetUserDomainBlackList", req, resp)
	return
}

type GetUserDomainBlackListArgs struct {
	SecurityToken string
	OwnerAccount  string
	DomainName    string
	OwnerId       int64
}
type GetUserDomainBlackListResponse struct {
	RequestId string
	Bind      string
}

func (c *CdnClient) DescribeCdnService(req *DescribeCdnServiceArgs) (resp *DescribeCdnServiceResponse, err error) {
	resp = &DescribeCdnServiceResponse{}
	err = c.Request("DescribeCdnService", req, resp)
	return
}

type DescribeCdnServiceLockReason struct {
	LockReason string
}
type DescribeCdnServiceArgs struct {
	SecurityToken string
	OwnerId       int64
}
type DescribeCdnServiceResponse struct {
	RequestId          string
	InstanceId         string
	InternetChargeType string
	OpeningTime        string
	ChangingChargeType string
	ChangingAffectTime string
	OperationLocks     DescribeCdnServiceLockReasonList
}

type DescribeCdnServiceLockReasonList []DescribeCdnServiceLockReason

func (list *DescribeCdnServiceLockReasonList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCdnServiceLockReason)
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

func (c *CdnClient) DescribeDomainUpstreamOfCenter(req *DescribeDomainUpstreamOfCenterArgs) (resp *DescribeDomainUpstreamOfCenterResponse, err error) {
	resp = &DescribeDomainUpstreamOfCenterResponse{}
	err = c.Request("DescribeDomainUpstreamOfCenter", req, resp)
	return
}

type DescribeDomainUpstreamOfCenterDomainBpsModel struct {
	Time string
	Bps  float32
}
type DescribeDomainUpstreamOfCenterArgs struct {
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
}
type DescribeDomainUpstreamOfCenterResponse struct {
	RequestId string
	BpsDatas  DescribeDomainUpstreamOfCenterDomainBpsModelList
}

type DescribeDomainUpstreamOfCenterDomainBpsModelList []DescribeDomainUpstreamOfCenterDomainBpsModel

func (list *DescribeDomainUpstreamOfCenterDomainBpsModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainUpstreamOfCenterDomainBpsModel)
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

func (c *CdnClient) DescribeLiveStreamBitRateData(req *DescribeLiveStreamBitRateDataArgs) (resp *DescribeLiveStreamBitRateDataResponse, err error) {
	resp = &DescribeLiveStreamBitRateDataResponse{}
	err = c.Request("DescribeLiveStreamBitRateData", req, resp)
	return
}

type DescribeLiveStreamBitRateDataFrameRateAndBitRateInfo struct {
	StreamUrl      string
	VideoFrameRate float32
	AudioFrameRate float32
	BitRate        float32
	Time           string
}
type DescribeLiveStreamBitRateDataArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
	StreamName    string
}
type DescribeLiveStreamBitRateDataResponse struct {
	RequestId                string
	FrameRateAndBitRateInfos DescribeLiveStreamBitRateDataFrameRateAndBitRateInfoList
}

type DescribeLiveStreamBitRateDataFrameRateAndBitRateInfoList []DescribeLiveStreamBitRateDataFrameRateAndBitRateInfo

func (list *DescribeLiveStreamBitRateDataFrameRateAndBitRateInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamBitRateDataFrameRateAndBitRateInfo)
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

func (c *CdnClient) ModifyCdnDomain(req *ModifyCdnDomainArgs) (resp *ModifyCdnDomainResponse, err error) {
	resp = &ModifyCdnDomainResponse{}
	err = c.Request("ModifyCdnDomain", req, resp)
	return
}

type ModifyCdnDomainArgs struct {
	TopLevelDomain  string
	SourcePort      int
	ResourceGroupId string
	Priorities      string
	Sources         string
	SecurityToken   string
	DomainName      string
	SourceType      string
	OwnerId         int64
}
type ModifyCdnDomainResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeLiveStreamDomainAppInfo(req *DescribeLiveStreamDomainAppInfoArgs) (resp *DescribeLiveStreamDomainAppInfoResponse, err error) {
	resp = &DescribeLiveStreamDomainAppInfoResponse{}
	err = c.Request("DescribeLiveStreamDomainAppInfo", req, resp)
	return
}

type DescribeLiveStreamDomainAppInfoDomainAppInfo struct {
	AppDomain    string
	AppId        string
	AppKey       string
	AppOssBucket string
	AppOssHost   string
	AppOwnerId   string
	AppSecret    string
	UpdateTime   string
}
type DescribeLiveStreamDomainAppInfoArgs struct {
	SecurityToken string
	AppDomain     string
	OwnerId       int64
}
type DescribeLiveStreamDomainAppInfoResponse struct {
	RequestId     string
	DomainAppList DescribeLiveStreamDomainAppInfoDomainAppInfoList
}

type DescribeLiveStreamDomainAppInfoDomainAppInfoList []DescribeLiveStreamDomainAppInfoDomainAppInfo

func (list *DescribeLiveStreamDomainAppInfoDomainAppInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamDomainAppInfoDomainAppInfo)
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

func (c *CdnClient) DeleteLiveStreamTranscode(req *DeleteLiveStreamTranscodeArgs) (resp *DeleteLiveStreamTranscodeResponse, err error) {
	resp = &DeleteLiveStreamTranscodeResponse{}
	err = c.Request("DeleteLiveStreamTranscode", req, resp)
	return
}

type DeleteLiveStreamTranscodeArgs struct {
	Template      string
	App           string
	SecurityToken string
	OwnerAccount  string
	Domain        string
	OwnerId       int64
}
type DeleteLiveStreamTranscodeResponse struct {
	RequestId string
}

func (c *CdnClient) ClearUserDomainBlackList(req *ClearUserDomainBlackListArgs) (resp *ClearUserDomainBlackListResponse, err error) {
	resp = &ClearUserDomainBlackListResponse{}
	err = c.Request("ClearUserDomainBlackList", req, resp)
	return
}

type ClearUserDomainBlackListArgs struct {
	SecurityToken string
	OwnerAccount  string
	DomainName    string
	OwnerId       int64
}
type ClearUserDomainBlackListResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeDomainBpsData(req *DescribeDomainBpsDataArgs) (resp *DescribeDomainBpsDataResponse, err error) {
	resp = &DescribeDomainBpsDataResponse{}
	err = c.Request("DescribeDomainBpsData", req, resp)
	return
}

type DescribeDomainBpsDataDataModule struct {
	TimeStamp            string
	Value                string
	DomesticValue        string
	OverseasValue        string
	DynamicValue         string
	DynamicDomesticValue string
	DynamicOverseasValue string
	StaticValue          string
	StaticDomesticValue  string
	StaticOverseasValue  string
	L2Value              string
	DomesticL2Value      string
	OverseasL2Value      string
}

type DescribeDomainBpsDataDataModule1 struct {
	TimeStamp string
	Value     string
}
type DescribeDomainBpsDataArgs struct {
	FixTimeGap     string
	TimeMerge      string
	DomainName     string
	EndTime        string
	LocationNameEn string
	StartTime      string
	IspNameEn      string
	OwnerId        int64
	DomainType     string
	SecurityToken  string
	Interval       string
}
type DescribeDomainBpsDataResponse struct {
	RequestId          string
	DomainName         string
	DataInterval       string
	StartTime          string
	EndTime            string
	LocationNameEn     string
	IspNameEn          string
	LocationName       string
	IspName            string
	BpsDataPerInterval DescribeDomainBpsDataDataModuleList
	SupplyBpsDatas     DescribeDomainBpsDataDataModule1List
}

type DescribeDomainBpsDataDataModuleList []DescribeDomainBpsDataDataModule

func (list *DescribeDomainBpsDataDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainBpsDataDataModule)
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

type DescribeDomainBpsDataDataModule1List []DescribeDomainBpsDataDataModule1

func (list *DescribeDomainBpsDataDataModule1List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainBpsDataDataModule1)
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

func (c *CdnClient) SetSourceHostConfig(req *SetSourceHostConfigArgs) (resp *SetSourceHostConfigResponse, err error) {
	resp = &SetSourceHostConfigResponse{}
	err = c.Request("SetSourceHostConfig", req, resp)
	return
}

type SetSourceHostConfigArgs struct {
	SecurityToken string
	Enable        string
	DomainName    string
	OwnerId       int64
	BackSrcDomain string
}
type SetSourceHostConfigResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeLiveStreamBpsData(req *DescribeLiveStreamBpsDataArgs) (resp *DescribeLiveStreamBpsDataResponse, err error) {
	resp = &DescribeLiveStreamBpsDataResponse{}
	err = c.Request("DescribeLiveStreamBpsData", req, resp)
	return
}

type DescribeLiveStreamBpsDataDomainBpsModel struct {
	Time string
	Bps  float32
}
type DescribeLiveStreamBpsDataArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
	StreamName    string
}
type DescribeLiveStreamBpsDataResponse struct {
	RequestId string
	BpsDatas  DescribeLiveStreamBpsDataDomainBpsModelList
}

type DescribeLiveStreamBpsDataDomainBpsModelList []DescribeLiveStreamBpsDataDomainBpsModel

func (list *DescribeLiveStreamBpsDataDomainBpsModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamBpsDataDomainBpsModel)
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

func (c *CdnClient) DescribeLivePullStreamConfig(req *DescribeLivePullStreamConfigArgs) (resp *DescribeLivePullStreamConfigResponse, err error) {
	resp = &DescribeLivePullStreamConfigResponse{}
	err = c.Request("DescribeLivePullStreamConfig", req, resp)
	return
}

type DescribeLivePullStreamConfigLiveAppRecord struct {
	DomainName string
	StreamName string
	SourceUrl  string
	StartTime  string
	EndTime    string
}
type DescribeLivePullStreamConfigArgs struct {
	SecurityToken string
	DomainName    string
	OwnerId       int64
}
type DescribeLivePullStreamConfigResponse struct {
	RequestId         string
	LiveAppRecordList DescribeLivePullStreamConfigLiveAppRecordList
}

type DescribeLivePullStreamConfigLiveAppRecordList []DescribeLivePullStreamConfigLiveAppRecord

func (list *DescribeLivePullStreamConfigLiveAppRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLivePullStreamConfigLiveAppRecord)
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

func (c *CdnClient) SetUserDomainBlackList(req *SetUserDomainBlackListArgs) (resp *SetUserDomainBlackListResponse, err error) {
	resp = &SetUserDomainBlackListResponse{}
	err = c.Request("SetUserDomainBlackList", req, resp)
	return
}

type SetUserDomainBlackListArgs struct {
	SecurityToken string
	OwnerAccount  string
	DomainName    string
	OwnerId       int64
}
type SetUserDomainBlackListResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeL2VipsByDomain(req *DescribeL2VipsByDomainArgs) (resp *DescribeL2VipsByDomainResponse, err error) {
	resp = &DescribeL2VipsByDomainResponse{}
	err = c.Request("DescribeL2VipsByDomain", req, resp)
	return
}

type DescribeL2VipsByDomainArgs struct {
	SecurityToken string
	DomainName    string
	OwnerId       int64
}
type DescribeL2VipsByDomainResponse struct {
	RequestId  string
	DomainName string
	Vips       DescribeL2VipsByDomainVipList
}

type DescribeL2VipsByDomainVipList []string

func (list *DescribeL2VipsByDomainVipList) UnmarshalJSON(data []byte) error {
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

func (c *CdnClient) SetIgnoreQueryStringConfig(req *SetIgnoreQueryStringConfigArgs) (resp *SetIgnoreQueryStringConfigResponse, err error) {
	resp = &SetIgnoreQueryStringConfigResponse{}
	err = c.Request("SetIgnoreQueryStringConfig", req, resp)
	return
}

type SetIgnoreQueryStringConfigArgs struct {
	KeepOssArgs   string
	HashKeyArgs   string
	SecurityToken string
	Enable        string
	DomainName    string
	OwnerId       int64
}
type SetIgnoreQueryStringConfigResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeDomainsUsageByDay(req *DescribeDomainsUsageByDayArgs) (resp *DescribeDomainsUsageByDayResponse, err error) {
	resp = &DescribeDomainsUsageByDayResponse{}
	err = c.Request("DescribeDomainsUsageByDay", req, resp)
	return
}

type DescribeDomainsUsageByDayUsageByDay struct {
	TimeStamp      string
	Qps            string
	BytesHitRate   string
	RequestHitRate string
	MaxBps         string
	MaxBpsTime     string
	MaxSrcBps      string
	MaxSrcBpsTime  string
	TotalAccess    string
	TotalTraffic   string
}

type DescribeDomainsUsageByDayUsageTotal struct {
	BytesHitRate   string
	RequestHitRate string
	MaxBps         string
	MaxBpsTime     string
	MaxSrcBps      string
	MaxSrcBpsTime  string
	TotalAccess    string
	TotalTraffic   string
}
type DescribeDomainsUsageByDayArgs struct {
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
}
type DescribeDomainsUsageByDayResponse struct {
	RequestId    string
	DomainName   string
	DataInterval string
	StartTime    string
	EndTime      string
	UsageByDays  DescribeDomainsUsageByDayUsageByDayList
	UsageTotal   DescribeDomainsUsageByDayUsageTotal
}

type DescribeDomainsUsageByDayUsageByDayList []DescribeDomainsUsageByDayUsageByDay

func (list *DescribeDomainsUsageByDayUsageByDayList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainsUsageByDayUsageByDay)
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

func (c *CdnClient) DescribeDomainOnlineUserNumber(req *DescribeDomainOnlineUserNumberArgs) (resp *DescribeDomainOnlineUserNumberResponse, err error) {
	resp = &DescribeDomainOnlineUserNumberResponse{}
	err = c.Request("DescribeDomainOnlineUserNumber", req, resp)
	return
}

type DescribeDomainOnlineUserNumberLiveStreamOnlineUserNumInfo struct {
	Time       string
	UserNumber int64
}
type DescribeDomainOnlineUserNumberArgs struct {
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
}
type DescribeDomainOnlineUserNumberResponse struct {
	RequestId                    string
	LiveStreamOnlineUserNumInfos DescribeDomainOnlineUserNumberLiveStreamOnlineUserNumInfoList
}

type DescribeDomainOnlineUserNumberLiveStreamOnlineUserNumInfoList []DescribeDomainOnlineUserNumberLiveStreamOnlineUserNumInfo

func (list *DescribeDomainOnlineUserNumberLiveStreamOnlineUserNumInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainOnlineUserNumberLiveStreamOnlineUserNumInfo)
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

func (c *CdnClient) SetL2OssKeyConfig(req *SetL2OssKeyConfigArgs) (resp *SetL2OssKeyConfigResponse, err error) {
	resp = &SetL2OssKeyConfigResponse{}
	err = c.Request("SetL2OssKeyConfig", req, resp)
	return
}

type SetL2OssKeyConfigArgs struct {
	SecurityToken  string
	DomainName     string
	OwnerId        int64
	PrivateOssAuth string
}
type SetL2OssKeyConfigResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeLiveStreamsBlockList(req *DescribeLiveStreamsBlockListArgs) (resp *DescribeLiveStreamsBlockListResponse, err error) {
	resp = &DescribeLiveStreamsBlockListResponse{}
	err = c.Request("DescribeLiveStreamsBlockList", req, resp)
	return
}

type DescribeLiveStreamsBlockListArgs struct {
	SecurityToken string
	DomainName    string
	OwnerId       int64
}
type DescribeLiveStreamsBlockListResponse struct {
	RequestId  string
	DomainName string
	StreamUrls DescribeLiveStreamsBlockListStreamUrlList
}

type DescribeLiveStreamsBlockListStreamUrlList []string

func (list *DescribeLiveStreamsBlockListStreamUrlList) UnmarshalJSON(data []byte) error {
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

func (c *CdnClient) SetRefererConfig(req *SetRefererConfigArgs) (resp *SetRefererConfigResponse, err error) {
	resp = &SetRefererConfigResponse{}
	err = c.Request("SetRefererConfig", req, resp)
	return
}

type SetRefererConfigArgs struct {
	ReferList     string
	SecurityToken string
	DomainName    string
	ReferType     string
	DisableAst    string
	OwnerId       int64
	AllowEmpty    string
}
type SetRefererConfigResponse struct {
	RequestId string
}

func (c *CdnClient) OpenCdnService(req *OpenCdnServiceArgs) (resp *OpenCdnServiceResponse, err error) {
	resp = &OpenCdnServiceResponse{}
	err = c.Request("OpenCdnService", req, resp)
	return
}

type OpenCdnServiceArgs struct {
	SecurityToken      string
	InternetChargeType string
	OwnerId            int64
}
type OpenCdnServiceResponse struct {
	RequestId string
}

func (c *CdnClient) SetHttpErrorPageConfig(req *SetHttpErrorPageConfigArgs) (resp *SetHttpErrorPageConfigResponse, err error) {
	resp = &SetHttpErrorPageConfigResponse{}
	err = c.Request("SetHttpErrorPageConfig", req, resp)
	return
}

type SetHttpErrorPageConfigArgs struct {
	SecurityToken string
	DomainName    string
	PageUrl       string
	OwnerId       int64
	ErrorCode     string
}
type SetHttpErrorPageConfigResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeLiveStreamsOnlineList(req *DescribeLiveStreamsOnlineListArgs) (resp *DescribeLiveStreamsOnlineListResponse, err error) {
	resp = &DescribeLiveStreamsOnlineListResponse{}
	err = c.Request("DescribeLiveStreamsOnlineList", req, resp)
	return
}

type DescribeLiveStreamsOnlineListLiveStreamOnlineInfo struct {
	DomainName  string
	AppName     string
	StreamName  string
	PublishTime string
	PublishUrl  string
}
type DescribeLiveStreamsOnlineListArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	OwnerId       int64
}
type DescribeLiveStreamsOnlineListResponse struct {
	RequestId  string
	OnlineInfo DescribeLiveStreamsOnlineListLiveStreamOnlineInfoList
}

type DescribeLiveStreamsOnlineListLiveStreamOnlineInfoList []DescribeLiveStreamsOnlineListLiveStreamOnlineInfo

func (list *DescribeLiveStreamsOnlineListLiveStreamOnlineInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamsOnlineListLiveStreamOnlineInfo)
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

func (c *CdnClient) DescribeDomainConfigs(req *DescribeDomainConfigsArgs) (resp *DescribeDomainConfigsResponse, err error) {
	resp = &DescribeDomainConfigsResponse{}
	err = c.Request("DescribeDomainConfigs", req, resp)
	return
}

type DescribeDomainConfigsDomainConfigs struct {
	CacheExpiredConfigs     DescribeDomainConfigsCacheExpiredConfigList
	HttpErrorPageConfigs    DescribeDomainConfigsHttpErrorPageConfigList
	HttpHeaderConfigs       DescribeDomainConfigsHttpHeaderConfigList
	DynamicConfigs          DescribeDomainConfigsDynamicConfigList
	ReqHeaderConfigs        DescribeDomainConfigsReqHeaderConfigList
	SetVarsConfigs          DescribeDomainConfigsSetVarsConfigList
	CcConfig                DescribeDomainConfigsCcConfig
	ErrorPageConfig         DescribeDomainConfigsErrorPageConfig
	OptimizeConfig          DescribeDomainConfigsOptimizeConfig
	PageCompressConfig      DescribeDomainConfigsPageCompressConfig
	IgnoreQueryStringConfig DescribeDomainConfigsIgnoreQueryStringConfig
	RangeConfig             DescribeDomainConfigsRangeConfig
	RefererConfig           DescribeDomainConfigsRefererConfig
	ReqAuthConfig           DescribeDomainConfigsReqAuthConfig
	SrcHostConfig           DescribeDomainConfigsSrcHostConfig
	VideoSeekConfig         DescribeDomainConfigsVideoSeekConfig
	WafConfig               DescribeDomainConfigsWafConfig
	NotifyUrlConfig         DescribeDomainConfigsNotifyUrlConfig
	RedirectTypeConfig      DescribeDomainConfigsRedirectTypeConfig
	ForwardSchemeConfig     DescribeDomainConfigsForwardSchemeConfig
	RemoveQueryStringConfig DescribeDomainConfigsRemoveQueryStringConfig
	L2OssKeyConfig          DescribeDomainConfigsL2OssKeyConfig
	MacServiceConfig        DescribeDomainConfigsMacServiceConfig
	GreenManagerConfig      DescribeDomainConfigsGreenManagerConfig
	HttpsOptionConfig       DescribeDomainConfigsHttpsOptionConfig
	AliBusinessConfig       DescribeDomainConfigsAliBusinessConfig
	IpAllowListConfig       DescribeDomainConfigsIpAllowListConfig
}

type DescribeDomainConfigsCacheExpiredConfig struct {
	ConfigId     string
	CacheType    string
	CacheContent string
	TTL          string
	Weight       string
	Status       string
}

type DescribeDomainConfigsHttpErrorPageConfig struct {
	ConfigId  string
	ErrorCode string
	PageUrl   string
	Status    string
}

type DescribeDomainConfigsHttpHeaderConfig struct {
	ConfigId    string
	HeaderKey   string
	HeaderValue string
	Status      string
}

type DescribeDomainConfigsDynamicConfig struct {
	ConfigId            string
	DynamicOrigin       string
	StaticType          string
	StaticUri           string
	StaticPath          string
	DynamicCacheControl string
	Status              string
}

type DescribeDomainConfigsReqHeaderConfig struct {
	ConfigId string
	Key      string
	Value    string
	Status   string
}

type DescribeDomainConfigsSetVarsConfig struct {
	ConfigId string
	VarName  string
	VarValue string
	Status   string
}

type DescribeDomainConfigsCcConfig struct {
	ConfigId string
	Enable   string
	AllowIps string
	BlockIps string
	Status   string
}

type DescribeDomainConfigsErrorPageConfig struct {
	ConfigId      string
	ErrorCode     string
	PageType      string
	CustomPageUrl string
	Status        string
}

type DescribeDomainConfigsOptimizeConfig struct {
	ConfigId string
	Enable   string
	Status   string
}

type DescribeDomainConfigsPageCompressConfig struct {
	ConfigId string
	Enable   string
	Status   string
}

type DescribeDomainConfigsIgnoreQueryStringConfig struct {
	ConfigId    string
	HashKeyArgs string
	Enable      string
	Status      string
}

type DescribeDomainConfigsRangeConfig struct {
	ConfigId string
	Enable   string
	Status   string
}

type DescribeDomainConfigsRefererConfig struct {
	ConfigId   string
	ReferType  string
	ReferList  string
	AllowEmpty string
	DisableAst string
	Status     string
}

type DescribeDomainConfigsReqAuthConfig struct {
	ConfigId         string
	AuthType         string
	Key1             string
	Key2             string
	Status           string
	AliAuthWhiteList string
	AuthM3u8         string
	AuthAddr         string
	AuthRemoteDesc   string
	TimeOut          string
}

type DescribeDomainConfigsSrcHostConfig struct {
	ConfigId   string
	DomainName string
	Status     string
}

type DescribeDomainConfigsVideoSeekConfig struct {
	ConfigId string
	Enable   string
	Status   string
}

type DescribeDomainConfigsWafConfig struct {
	ConfigId string
	Enable   string
	Status   string
}

type DescribeDomainConfigsNotifyUrlConfig struct {
	Enable    string
	NotifyUrl string
}

type DescribeDomainConfigsRedirectTypeConfig struct {
	RedirectType string
}

type DescribeDomainConfigsForwardSchemeConfig struct {
	ConfigId         string
	Enable           string
	SchemeOrigin     string
	SchemeOriginPort string
	Status           string
}

type DescribeDomainConfigsRemoveQueryStringConfig struct {
	AliRemoveArgs string
	ConfigId      string
	Status        string
}

type DescribeDomainConfigsL2OssKeyConfig struct {
	PrivateOssAuth string
	ConfigId       string
	Status         string
}

type DescribeDomainConfigsMacServiceConfig struct {
	AppList       string
	Enabled       string
	ProcessResult string
	ConfigId      string
	Status        string
}

type DescribeDomainConfigsGreenManagerConfig struct {
	Enabled  string
	ConfigId string
	Status   string
}

type DescribeDomainConfigsHttpsOptionConfig struct {
	Http2    string
	ConfigId string
	Status   string
}

type DescribeDomainConfigsAliBusinessConfig struct {
	AliBusinessTable string
	AliBusinessType  string
	ConfigId         string
	Status           string
}

type DescribeDomainConfigsIpAllowListConfig struct {
	ConfigId  string
	IpList    string
	IpAclXfwd string
	Status    string
}
type DescribeDomainConfigsArgs struct {
	SecurityToken string
	DomainName    string
	ConfigList    string
	OwnerId       int64
}
type DescribeDomainConfigsResponse struct {
	RequestId     string
	DomainConfigs DescribeDomainConfigsDomainConfigs
}

type DescribeDomainConfigsCacheExpiredConfigList []DescribeDomainConfigsCacheExpiredConfig

func (list *DescribeDomainConfigsCacheExpiredConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainConfigsCacheExpiredConfig)
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

type DescribeDomainConfigsHttpErrorPageConfigList []DescribeDomainConfigsHttpErrorPageConfig

func (list *DescribeDomainConfigsHttpErrorPageConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainConfigsHttpErrorPageConfig)
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

type DescribeDomainConfigsHttpHeaderConfigList []DescribeDomainConfigsHttpHeaderConfig

func (list *DescribeDomainConfigsHttpHeaderConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainConfigsHttpHeaderConfig)
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

type DescribeDomainConfigsDynamicConfigList []DescribeDomainConfigsDynamicConfig

func (list *DescribeDomainConfigsDynamicConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainConfigsDynamicConfig)
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

type DescribeDomainConfigsReqHeaderConfigList []DescribeDomainConfigsReqHeaderConfig

func (list *DescribeDomainConfigsReqHeaderConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainConfigsReqHeaderConfig)
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

type DescribeDomainConfigsSetVarsConfigList []DescribeDomainConfigsSetVarsConfig

func (list *DescribeDomainConfigsSetVarsConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainConfigsSetVarsConfig)
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

func (c *CdnClient) DeleteLivePullStreamInfo(req *DeleteLivePullStreamInfoArgs) (resp *DeleteLivePullStreamInfoResponse, err error) {
	resp = &DeleteLivePullStreamInfoResponse{}
	err = c.Request("DeleteLivePullStreamInfo", req, resp)
	return
}

type DeleteLivePullStreamInfoArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	OwnerId       int64
	StreamName    string
}
type DeleteLivePullStreamInfoResponse struct {
	RequestId string
}

func (c *CdnClient) SetUserBlackList(req *SetUserBlackListArgs) (resp *SetUserBlackListResponse, err error) {
	resp = &SetUserBlackListResponse{}
	err = c.Request("SetUserBlackList", req, resp)
	return
}

type SetUserBlackListArgs struct {
	ConfigUrl     string
	SecurityToken string
	OwnerAccount  string
	OwnerId       int64
}
type SetUserBlackListResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeDomainMonthBillingBpsData(req *DescribeDomainMonthBillingBpsDataArgs) (resp *DescribeDomainMonthBillingBpsDataResponse, err error) {
	resp = &DescribeDomainMonthBillingBpsDataResponse{}
	err = c.Request("DescribeDomainMonthBillingBpsData", req, resp)
	return
}

type DescribeDomainMonthBillingBpsDataArgs struct {
	SecurityToken      string
	InternetChargeType string
	DomainName         string
	EndTime            string
	StartTime          string
	OwnerId            int64
}
type DescribeDomainMonthBillingBpsDataResponse struct {
	RequestId              string
	DomainName             string
	StartTime              string
	EndTime                string
	Bandwidth95Bps         float32
	DomesticBandwidth95Bps float32
	OverseasBandwidth95Bps float32
	MonthavgBps            float32
	DomesticMonthavgBps    float32
	OverseasMonthavgBps    float32
	Month4thBps            float32
	DomesticMonth4thBps    float32
	OverseasMonth4thBps    float32
}

func (c *CdnClient) DescribeLiveRecordConfig(req *DescribeLiveRecordConfigArgs) (resp *DescribeLiveRecordConfigResponse, err error) {
	resp = &DescribeLiveRecordConfigResponse{}
	err = c.Request("DescribeLiveRecordConfig", req, resp)
	return
}

type DescribeLiveRecordConfigLiveAppRecord struct {
	DomainName      string
	AppName         string
	OssEndpoint     string
	OssBucket       string
	OssObjectPrefix string
	CreateTime      string
}
type DescribeLiveRecordConfigArgs struct {
	SecurityToken string
	DomainName    string
	OwnerId       int64
}
type DescribeLiveRecordConfigResponse struct {
	RequestId         string
	LiveAppRecordList DescribeLiveRecordConfigLiveAppRecordList
}

type DescribeLiveRecordConfigLiveAppRecordList []DescribeLiveRecordConfigLiveAppRecord

func (list *DescribeLiveRecordConfigLiveAppRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveRecordConfigLiveAppRecord)
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

func (c *CdnClient) DescribeRangeDataByLocateAndIspService(req *DescribeRangeDataByLocateAndIspServiceArgs) (resp *DescribeRangeDataByLocateAndIspServiceResponse, err error) {
	resp = &DescribeRangeDataByLocateAndIspServiceResponse{}
	err = c.Request("DescribeRangeDataByLocateAndIspService", req, resp)
	return
}

type DescribeRangeDataByLocateAndIspServiceArgs struct {
	IspNames      string
	SecurityToken string
	DomainNames   string
	LocationNames string
	EndTime       string
	StartTime     string
	OwnerId       int64
}
type DescribeRangeDataByLocateAndIspServiceResponse struct {
	RequestId  string
	JsonResult string
}

func (c *CdnClient) DescribeDomainPathData(req *DescribeDomainPathDataArgs) (resp *DescribeDomainPathDataResponse, err error) {
	resp = &DescribeDomainPathDataResponse{}
	err = c.Request("DescribeDomainPathData", req, resp)
	return
}

type DescribeDomainPathDataUsageData struct {
	Traffic int
	Acc     int
	Path    string
	Time    string
}
type DescribeDomainPathDataArgs struct {
	StartTime     string
	PageNumber    int
	Path          string
	SecurityToken string
	PageSize      int
	DomainName    string
	EndTime       string
	OwnerId       int64
	Version       string
}
type DescribeDomainPathDataResponse struct {
	DomainName          string
	StartTime           string
	EndTime             string
	PageSize            int
	PageNumber          int
	DataInterval        string
	TotalCount          int
	PathDataPerInterval DescribeDomainPathDataUsageDataList
}

type DescribeDomainPathDataUsageDataList []DescribeDomainPathDataUsageData

func (list *DescribeDomainPathDataUsageDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainPathDataUsageData)
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

func (c *CdnClient) DescribeLiveStreamsPublishList(req *DescribeLiveStreamsPublishListArgs) (resp *DescribeLiveStreamsPublishListResponse, err error) {
	resp = &DescribeLiveStreamsPublishListResponse{}
	err = c.Request("DescribeLiveStreamsPublishList", req, resp)
	return
}

type DescribeLiveStreamsPublishListLiveStreamPublishInfo struct {
	DomainName   string
	AppName      string
	StreamName   string
	StreamUrl    string
	PublishTime  string
	StopTime     string
	PublishUrl   string
	ClientAddr   string
	EdgeNodeAddr string
}
type DescribeLiveStreamsPublishListArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	PageSize      int64
	EndTime       string
	StartTime     string
	OwnerId       int64
	StreamName    string
	PageNumber    int64
}
type DescribeLiveStreamsPublishListResponse struct {
	RequestId   string
	PublishInfo DescribeLiveStreamsPublishListLiveStreamPublishInfoList
}

type DescribeLiveStreamsPublishListLiveStreamPublishInfoList []DescribeLiveStreamsPublishListLiveStreamPublishInfo

func (list *DescribeLiveStreamsPublishListLiveStreamPublishInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamsPublishListLiveStreamPublishInfo)
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

func (c *CdnClient) DescribeDomainAverageResponseTime(req *DescribeDomainAverageResponseTimeArgs) (resp *DescribeDomainAverageResponseTimeResponse, err error) {
	resp = &DescribeDomainAverageResponseTimeResponse{}
	err = c.Request("DescribeDomainAverageResponseTime", req, resp)
	return
}

type DescribeDomainAverageResponseTimeDataModule struct {
	TimeStamp string
	Value     string
}
type DescribeDomainAverageResponseTimeArgs struct {
	SecurityToken string
	TimeMerge     string
	DomainName    string
	EndTime       string
	Interval      string
	StartTime     string
	OwnerId       int64
}
type DescribeDomainAverageResponseTimeResponse struct {
	RequestId        string
	DomainName       string
	DataInterval     string
	StartTime        string
	EndTime          string
	AvgRTPerInterval DescribeDomainAverageResponseTimeDataModuleList
}

type DescribeDomainAverageResponseTimeDataModuleList []DescribeDomainAverageResponseTimeDataModule

func (list *DescribeDomainAverageResponseTimeDataModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainAverageResponseTimeDataModule)
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

func (c *CdnClient) DescribeCdnDomainBaseDetail(req *DescribeCdnDomainBaseDetailArgs) (resp *DescribeCdnDomainBaseDetailResponse, err error) {
	resp = &DescribeCdnDomainBaseDetailResponse{}
	err = c.Request("DescribeCdnDomainBaseDetail", req, resp)
	return
}

type DescribeCdnDomainBaseDetailDomainBaseDetailModel struct {
	Cname        string
	CdnType      string
	DomainStatus string
	SourceType   string
	Region       string
	DomainName   string
	Remark       string
	GmtModified  string
	GmtCreated   string
	Sources      DescribeCdnDomainBaseDetailSourceList
}
type DescribeCdnDomainBaseDetailArgs struct {
	SecurityToken string
	DomainName    string
	OwnerId       int64
}
type DescribeCdnDomainBaseDetailResponse struct {
	RequestId             string
	DomainBaseDetailModel DescribeCdnDomainBaseDetailDomainBaseDetailModel
}

type DescribeCdnDomainBaseDetailSourceList []string

func (list *DescribeCdnDomainBaseDetailSourceList) UnmarshalJSON(data []byte) error {
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

func (c *CdnClient) DescribeDomainSlowRatio(req *DescribeDomainSlowRatioArgs) (resp *DescribeDomainSlowRatioResponse, err error) {
	resp = &DescribeDomainSlowRatioResponse{}
	err = c.Request("DescribeDomainSlowRatio", req, resp)
	return
}

type DescribeDomainSlowRatioSlowRatioData struct {
	TotalUsers   int
	SlowUsers    int
	SlowRatio    float32
	RegionNameZh string
	RegionNameEn string
	IspNameZh    string
	IspNameEn    string
	Time         string
}
type DescribeDomainSlowRatioArgs struct {
	StartTime     string
	PageNumber    int
	SecurityToken string
	PageSize      int
	DomainName    string
	EndTime       string
	OwnerId       int64
	Version       string
}
type DescribeDomainSlowRatioResponse struct {
	EndTime                  string
	DataInterval             int
	PageNumber               int
	PageSize                 int
	TotalCount               int
	StartTime                string
	SlowRatioDataPerInterval DescribeDomainSlowRatioSlowRatioDataList
}

type DescribeDomainSlowRatioSlowRatioDataList []DescribeDomainSlowRatioSlowRatioData

func (list *DescribeDomainSlowRatioSlowRatioDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainSlowRatioSlowRatioData)
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

func (c *CdnClient) SetDomainServerCertificate(req *SetDomainServerCertificateArgs) (resp *SetDomainServerCertificateResponse, err error) {
	resp = &SetDomainServerCertificateResponse{}
	err = c.Request("SetDomainServerCertificate", req, resp)
	return
}

type SetDomainServerCertificateArgs struct {
	PrivateKey              string
	ServerCertificateStatus string
	ServerCertificate       string
	SecurityToken           string
	CertName                string
	DomainName              string
	OwnerId                 int64
	Region                  string
}
type SetDomainServerCertificateResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeLiveStreamsFrameRateAndBitRateData(req *DescribeLiveStreamsFrameRateAndBitRateDataArgs) (resp *DescribeLiveStreamsFrameRateAndBitRateDataResponse, err error) {
	resp = &DescribeLiveStreamsFrameRateAndBitRateDataResponse{}
	err = c.Request("DescribeLiveStreamsFrameRateAndBitRateData", req, resp)
	return
}

type DescribeLiveStreamsFrameRateAndBitRateDataFrameRateAndBitRateInfo struct {
	StreamUrl      string
	VideoFrameRate float32
	AudioFrameRate float32
	BitRate        float32
	Time           string
}
type DescribeLiveStreamsFrameRateAndBitRateDataArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
	StreamName    string
}
type DescribeLiveStreamsFrameRateAndBitRateDataResponse struct {
	RequestId                string
	FrameRateAndBitRateInfos DescribeLiveStreamsFrameRateAndBitRateDataFrameRateAndBitRateInfoList
}

type DescribeLiveStreamsFrameRateAndBitRateDataFrameRateAndBitRateInfoList []DescribeLiveStreamsFrameRateAndBitRateDataFrameRateAndBitRateInfo

func (list *DescribeLiveStreamsFrameRateAndBitRateDataFrameRateAndBitRateInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamsFrameRateAndBitRateDataFrameRateAndBitRateInfo)
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

func (c *CdnClient) DescribeDomainHttpCodeData(req *DescribeDomainHttpCodeDataArgs) (resp *DescribeDomainHttpCodeDataResponse, err error) {
	resp = &DescribeDomainHttpCodeDataResponse{}
	err = c.Request("DescribeDomainHttpCodeData", req, resp)
	return
}

type DescribeDomainHttpCodeDataUsageData struct {
	TimeStamp string
	Value     DescribeDomainHttpCodeDataCodeProportionDataList
}

type DescribeDomainHttpCodeDataCodeProportionData struct {
	Code       string
	Proportion string
	Count      string
}
type DescribeDomainHttpCodeDataArgs struct {
	SecurityToken  string
	TimeMerge      string
	DomainName     string
	EndTime        string
	LocationNameEn string
	Interval       string
	StartTime      string
	IspNameEn      string
	OwnerId        int64
}
type DescribeDomainHttpCodeDataResponse struct {
	RequestId    string
	DomainName   string
	DataInterval string
	StartTime    string
	EndTime      string
	HttpCodeData DescribeDomainHttpCodeDataUsageDataList
}

type DescribeDomainHttpCodeDataCodeProportionDataList []DescribeDomainHttpCodeDataCodeProportionData

func (list *DescribeDomainHttpCodeDataCodeProportionDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainHttpCodeDataCodeProportionData)
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

type DescribeDomainHttpCodeDataUsageDataList []DescribeDomainHttpCodeDataUsageData

func (list *DescribeDomainHttpCodeDataUsageDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainHttpCodeDataUsageData)
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

func (c *CdnClient) DescribeRefreshTasks(req *DescribeRefreshTasksArgs) (resp *DescribeRefreshTasksResponse, err error) {
	resp = &DescribeRefreshTasksResponse{}
	err = c.Request("DescribeRefreshTasks", req, resp)
	return
}

type DescribeRefreshTasksCDNTask struct {
	TaskId       string
	ObjectPath   string
	Process      string
	Status       string
	CreationTime string
	Description  string
	ObjectType   string
}
type DescribeRefreshTasksArgs struct {
	ObjectPath      string
	DomainName      string
	EndTime         string
	StartTime       string
	OwnerId         int64
	PageNumber      int
	ResourceGroupId string
	SecurityToken   string
	PageSize        int
	ObjectType      string
	TaskId          string
	Status          string
}
type DescribeRefreshTasksResponse struct {
	RequestId  string
	PageNumber int64
	PageSize   int64
	TotalCount int64
	Tasks      DescribeRefreshTasksCDNTaskList
}

type DescribeRefreshTasksCDNTaskList []DescribeRefreshTasksCDNTask

func (list *DescribeRefreshTasksCDNTaskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRefreshTasksCDNTask)
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

func (c *CdnClient) DescribeDomainCCData(req *DescribeDomainCCDataArgs) (resp *DescribeDomainCCDataResponse, err error) {
	resp = &DescribeDomainCCDataResponse{}
	err = c.Request("DescribeDomainCCData", req, resp)
	return
}

type DescribeDomainCCDataCCDatas struct {
	TimeStamp string
	Count     string
}
type DescribeDomainCCDataArgs struct {
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
}
type DescribeDomainCCDataResponse struct {
	RequestId    string
	DomainName   string
	DataInterval string
	StartTime    string
	EndTime      string
	CCDataList   DescribeDomainCCDataCCDatasList
}

type DescribeDomainCCDataCCDatasList []DescribeDomainCCDataCCDatas

func (list *DescribeDomainCCDataCCDatasList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainCCDataCCDatas)
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

func (c *CdnClient) ForbidLiveStream(req *ForbidLiveStreamArgs) (resp *ForbidLiveStreamResponse, err error) {
	resp = &ForbidLiveStreamResponse{}
	err = c.Request("ForbidLiveStream", req, resp)
	return
}

type ForbidLiveStreamArgs struct {
	ResumeTime     string
	AppName        string
	SecurityToken  string
	LiveStreamType string
	DomainName     string
	OwnerId        int64
	StreamName     string
}
type ForbidLiveStreamResponse struct {
	RequestId string
}

func (c *CdnClient) DescribeCdnTypes(req *DescribeCdnTypesArgs) (resp *DescribeCdnTypesResponse, err error) {
	resp = &DescribeCdnTypesResponse{}
	err = c.Request("DescribeCdnTypes", req, resp)
	return
}

type DescribeCdnTypesCdnType struct {
	Type string
	Desc string
}
type DescribeCdnTypesArgs struct {
	SecurityToken string
	OwnerAccount  string
	OwnerId       int64
}
type DescribeCdnTypesResponse struct {
	RequestId string
	CdnTypes  DescribeCdnTypesCdnTypeList
}

type DescribeCdnTypesCdnTypeList []DescribeCdnTypesCdnType

func (list *DescribeCdnTypesCdnTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCdnTypesCdnType)
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

func (c *CdnClient) DescribeDomainCCAttackInfo(req *DescribeDomainCCAttackInfoArgs) (resp *DescribeDomainCCAttackInfoResponse, err error) {
	resp = &DescribeDomainCCAttackInfoResponse{}
	err = c.Request("DescribeDomainCCAttackInfo", req, resp)
	return
}

type DescribeDomainCCAttackInfoAttackIpDatas struct {
	Ip          string
	AttackCount string
	Result      string
}

type DescribeDomainCCAttackInfoAttackedUrlDatas struct {
	Url         string
	AttackCount string
	Result      string
}
type DescribeDomainCCAttackInfoArgs struct {
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
}
type DescribeDomainCCAttackInfoResponse struct {
	RequestId           string
	DomainName          string
	StartTime           string
	EndTime             string
	AttackIpDataList    DescribeDomainCCAttackInfoAttackIpDatasList
	AttackedUrlDataList DescribeDomainCCAttackInfoAttackedUrlDatasList
}

type DescribeDomainCCAttackInfoAttackIpDatasList []DescribeDomainCCAttackInfoAttackIpDatas

func (list *DescribeDomainCCAttackInfoAttackIpDatasList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainCCAttackInfoAttackIpDatas)
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

type DescribeDomainCCAttackInfoAttackedUrlDatasList []DescribeDomainCCAttackInfoAttackedUrlDatas

func (list *DescribeDomainCCAttackInfoAttackedUrlDatasList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainCCAttackInfoAttackedUrlDatas)
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

func (c *CdnClient) SetOssLogConfig(req *SetOssLogConfigArgs) (resp *SetOssLogConfigResponse, err error) {
	resp = &SetOssLogConfigResponse{}
	err = c.Request("SetOssLogConfig", req, resp)
	return
}

type SetOssLogConfigArgs struct {
	Bucket        string
	SecurityToken string
	Enable        string
	Prefix        string
	OwnerId       int64
}
type SetOssLogConfigResponse struct {
	RequestId string
}
