package vod

import (
	"encoding/json"

	"github.com/morlay/aliyun-go/core"
)

func (c *VodClient) RefreshUploadVideo(req *RefreshUploadVideoArgs) (resp *RefreshUploadVideoResponse, err error) {
	resp = &RefreshUploadVideoResponse{}
	err = c.Request("RefreshUploadVideo", req, resp)
	return
}

type RefreshUploadVideoArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	VideoId              string
	OwnerId              int64
}
type RefreshUploadVideoResponse struct {
	RequestId     string
	UploadAuth    string
	UploadAddress string
}

func (c *VodClient) SetMessageCallback(req *SetMessageCallbackArgs) (resp *SetMessageCallbackResponse, err error) {
	resp = &SetMessageCallbackResponse{}
	err = c.Request("SetMessageCallback", req, resp)
	return
}

type SetMessageCallbackArgs struct {
	CallbackType         string
	ResourceOwnerId      string
	CallbackSwitch       string
	ResourceOwnerAccount string
	OwnerAccount         string
	EventTypeList        string
	CallbackURL          string
	OwnerId              string
}
type SetMessageCallbackResponse struct {
	RequestId string
}

func (c *VodClient) DescribeCdnDomainLogs(req *DescribeCdnDomainLogsArgs) (resp *DescribeCdnDomainLogsResponse, err error) {
	resp = &DescribeCdnDomainLogsResponse{}
	err = c.Request("DescribeCdnDomainLogs", req, resp)
	return
}

type DescribeCdnDomainLogsDomainLogModel struct {
	DomainName       string
	DomainLogDetails DescribeCdnDomainLogsDomainLogDetailList
}

type DescribeCdnDomainLogsDomainLogDetail struct {
	LogPath   string
	StartTime string
	EndTime   string
	LogSize   int64
	LogName   string
}
type DescribeCdnDomainLogsArgs struct {
	ResourceOwnerId      string
	ResourceOwnerAccount string
	PageNo               int64
	OwnerAccount         string
	DomainName           string
	PageSize             int64
	EndTime              string
	StartTime            string
	OwnerId              string
	LogDay               string
}
type DescribeCdnDomainLogsResponse struct {
	RequestId      string
	PageNo         int64
	PageSize       int64
	Total          int64
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

func (c *VodClient) ListAIVideoSummaryJob(req *ListAIVideoSummaryJobArgs) (resp *ListAIVideoSummaryJobResponse, err error) {
	resp = &ListAIVideoSummaryJobResponse{}
	err = c.Request("ListAIVideoSummaryJob", req, resp)
	return
}

type ListAIVideoSummaryJobAIVideoSummaryJob struct {
	JobId        string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
}
type ListAIVideoSummaryJobArgs struct {
	ResourceOwnerId      string
	ResourceOwnerAccount string
	AIVideoSummaryJobIds string
	OwnerAccount         string
	OwnerId              string
}
type ListAIVideoSummaryJobResponse struct {
	RequestId                    string
	AIVideoSummaryJobList        ListAIVideoSummaryJobAIVideoSummaryJobList
	NonExistAIVideoSummaryJobIds ListAIVideoSummaryJobNonExistAIVideoSummaryJobIdList
}

type ListAIVideoSummaryJobAIVideoSummaryJobList []ListAIVideoSummaryJobAIVideoSummaryJob

func (list *ListAIVideoSummaryJobAIVideoSummaryJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAIVideoSummaryJobAIVideoSummaryJob)
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

type ListAIVideoSummaryJobNonExistAIVideoSummaryJobIdList []string

func (list *ListAIVideoSummaryJobNonExistAIVideoSummaryJobIdList) UnmarshalJSON(data []byte) error {
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

func (c *VodClient) GetVideoPlayInfo(req *GetVideoPlayInfoArgs) (resp *GetVideoPlayInfoResponse, err error) {
	resp = &GetVideoPlayInfoResponse{}
	err = c.Request("GetVideoPlayInfo", req, resp)
	return
}

type GetVideoPlayInfoPlayInfo struct {
	AccessKeyId     string
	AccessKeySecret string
	AuthInfo        string
	SecurityToken   string
	Region          string
	PlayDomain      string
}

type GetVideoPlayInfoVideoInfo struct {
	CoverURL   string
	CustomerId int64
	Duration   float32
	Status     string
	Title      string
	VideoId    string
}
type GetVideoPlayInfoArgs struct {
	SignVersion          string
	ResourceOwnerId      int64
	ClientVersion        string
	ResourceOwnerAccount string
	Channel              string
	PlaySign             string
	VideoId              string
	OwnerId              int64
	ClientTS             int64
}
type GetVideoPlayInfoResponse struct {
	RequestId string
	PlayInfo  GetVideoPlayInfoPlayInfo
	VideoInfo GetVideoPlayInfoVideoInfo
}

func (c *VodClient) GetCDNStatisSum(req *GetCDNStatisSumArgs) (resp *GetCDNStatisSumResponse, err error) {
	resp = &GetCDNStatisSumResponse{}
	err = c.Request("GetCDNStatisSum", req, resp)
	return
}

type GetCDNStatisSumCDNMetric struct {
	StatTime              string
	FlowDataDomesticValue int64
	FlowDataOverseasValue int64
	FlowDataValue         int64
	BpsDataDomesticValue  int64
	BpsDataOverseasValue  int64
	BpsDataValue          int64
}
type GetCDNStatisSumArgs struct {
	ResourceOwnerId      string
	StartStatisTime      string
	ResourceOwnerAccount string
	Level                string
	OwnerAccount         string
	OwnerId              string
	EndStatisTime        string
}
type GetCDNStatisSumResponse struct {
	RequestId        string
	SumFlowDataValue int64
	MaxBpsDataValue  int64
	CDNStatisList    GetCDNStatisSumCDNMetricList
}

type GetCDNStatisSumCDNMetricList []GetCDNStatisSumCDNMetric

func (list *GetCDNStatisSumCDNMetricList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetCDNStatisSumCDNMetric)
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

func (c *VodClient) SubmitAIVideoCoverJob(req *SubmitAIVideoCoverJobArgs) (resp *SubmitAIVideoCoverJobResponse, err error) {
	resp = &SubmitAIVideoCoverJobResponse{}
	err = c.Request("SubmitAIVideoCoverJob", req, resp)
	return
}

type SubmitAIVideoCoverJobAIVideoCoverJob struct {
	JobId        string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
}
type SubmitAIVideoCoverJobArgs struct {
	UserData             string
	ResourceOwnerId      string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              string
	MediaId              string
	AIVideoCoverConfig   string
}
type SubmitAIVideoCoverJobResponse struct {
	RequestId       string
	AIVideoCoverJob SubmitAIVideoCoverJobAIVideoCoverJob
}

func (c *VodClient) ListAIVideoCensorJob(req *ListAIVideoCensorJobArgs) (resp *ListAIVideoCensorJobResponse, err error) {
	resp = &ListAIVideoCensorJobResponse{}
	err = c.Request("ListAIVideoCensorJob", req, resp)
	return
}

type ListAIVideoCensorJobAIVideoCensorJob struct {
	JobId        string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
}
type ListAIVideoCensorJobArgs struct {
	AIVideoCensorJobIds  string
	ResourceOwnerId      string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              string
}
type ListAIVideoCensorJobResponse struct {
	RequestId                   string
	AIVideoCensorJobList        ListAIVideoCensorJobAIVideoCensorJobList
	NonExistAIVideoCensorJobIds ListAIVideoCensorJobNonExistAIVideoCensorJobIdList
}

type ListAIVideoCensorJobAIVideoCensorJobList []ListAIVideoCensorJobAIVideoCensorJob

func (list *ListAIVideoCensorJobAIVideoCensorJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAIVideoCensorJobAIVideoCensorJob)
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

type ListAIVideoCensorJobNonExistAIVideoCensorJobIdList []string

func (list *ListAIVideoCensorJobNonExistAIVideoCensorJobIdList) UnmarshalJSON(data []byte) error {
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

func (c *VodClient) GetVideoList(req *GetVideoListArgs) (resp *GetVideoListResponse, err error) {
	resp = &GetVideoListResponse{}
	err = c.Request("GetVideoList", req, resp)
	return
}

type GetVideoListVideo struct {
	VideoId      string
	Title        string
	Tags         string
	Status       string
	Size         int64
	Duration     float32
	Description  string
	CreateTime   string
	CreationTime string
	ModifyTime   string
	CoverURL     string
	CateId       int
	CateName     string
	Snapshots    GetVideoListSnapshotList
}
type GetVideoListArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	CateId               int
	PageNo               int
	PageSize             int
	EndTime              string
	SortBy               string
	StartTime            string
	OwnerId              int64
	Status               string
}
type GetVideoListResponse struct {
	RequestId string
	Total     int
	VideoList GetVideoListVideoList
}

type GetVideoListSnapshotList []string

func (list *GetVideoListSnapshotList) UnmarshalJSON(data []byte) error {
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

type GetVideoListVideoList []GetVideoListVideo

func (list *GetVideoListVideoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetVideoListVideo)
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

func (c *VodClient) ListAIASRJob(req *ListAIASRJobArgs) (resp *ListAIASRJobResponse, err error) {
	resp = &ListAIASRJobResponse{}
	err = c.Request("ListAIASRJob", req, resp)
	return
}

type ListAIASRJobAIASRJob struct {
	JobId        string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
}
type ListAIASRJobArgs struct {
	ResourceOwnerId      string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              string
	AIASRJobIds          string
}
type ListAIASRJobResponse struct {
	RequestId           string
	AIASRJobList        ListAIASRJobAIASRJobList
	NonExistAIASRJobIds ListAIASRJobNonExistAIASRJobIdList
}

type ListAIASRJobAIASRJobList []ListAIASRJobAIASRJob

func (list *ListAIASRJobAIASRJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAIASRJobAIASRJob)
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

type ListAIASRJobNonExistAIASRJobIdList []string

func (list *ListAIASRJobNonExistAIASRJobIdList) UnmarshalJSON(data []byte) error {
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

func (c *VodClient) SetEditingProjectMaterials(req *SetEditingProjectMaterialsArgs) (resp *SetEditingProjectMaterialsResponse, err error) {
	resp = &SetEditingProjectMaterialsResponse{}
	err = c.Request("SetEditingProjectMaterials", req, resp)
	return
}

type SetEditingProjectMaterialsArgs struct {
	ResourceOwnerId      string
	ResourceOwnerAccount string
	OwnerAccount         string
	MaterialIds          string
	OwnerId              string
	ProjectId            string
}
type SetEditingProjectMaterialsResponse struct {
	RequestId string
}

func (c *VodClient) DescribeDomainFlowData(req *DescribeDomainFlowDataArgs) (resp *DescribeDomainFlowDataResponse, err error) {
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
	ResourceOwnerId      string
	ResourceOwnerAccount string
	TimeMerge            string
	OwnerAccount         string
	DomainName           string
	EndTime              string
	LocationNameEn       string
	StartTime            string
	IspNameEn            string
	OwnerId              string
	Interval             string
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

func (c *VodClient) GetImageInfo(req *GetImageInfoArgs) (resp *GetImageInfoResponse, err error) {
	resp = &GetImageInfoResponse{}
	err = c.Request("GetImageInfo", req, resp)
	return
}

type GetImageInfoImageInfo struct {
	ImageId      string
	Title        string
	CreationTime string
	ImageType    string
	Tags         string
	URL          string
	Mezzanine    GetImageInfoMezzanine
}

type GetImageInfoMezzanine struct {
	OriginalFileName string
}
type GetImageInfoArgs struct {
	ResourceOwnerId      int64
	ImageId              string
	ResourceOwnerAccount string
	OwnerId              int64
	AuthTimeout          int64
}
type GetImageInfoResponse struct {
	RequestId string
	ImageInfo GetImageInfoImageInfo
}

func (c *VodClient) CreateUploadVideo(req *CreateUploadVideoArgs) (resp *CreateUploadVideoResponse, err error) {
	resp = &CreateUploadVideoResponse{}
	err = c.Request("CreateUploadVideo", req, resp)
	return
}

type CreateUploadVideoArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	TranscodeMode        string
	IP                   string
	Description          string
	FileSize             int64
	OwnerId              int64
	Title                string
	Tags                 string
	CoverURL             string
	UserData             string
	FileName             string
	TemplateGroupId      string
	CateId               int
}
type CreateUploadVideoResponse struct {
	RequestId     string
	VideoId       string
	UploadAddress string
	UploadAuth    string
}

func (c *VodClient) ListAIVideoCategoryJob(req *ListAIVideoCategoryJobArgs) (resp *ListAIVideoCategoryJobResponse, err error) {
	resp = &ListAIVideoCategoryJobResponse{}
	err = c.Request("ListAIVideoCategoryJob", req, resp)
	return
}

type ListAIVideoCategoryJobAIVideoCategoryJob struct {
	JobId        string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
}
type ListAIVideoCategoryJobArgs struct {
	ResourceOwnerId       string
	AIVideoCategoryJobIds string
	ResourceOwnerAccount  string
	OwnerAccount          string
	OwnerId               string
}
type ListAIVideoCategoryJobResponse struct {
	RequestId                     string
	AIVideoCategoryJobList        ListAIVideoCategoryJobAIVideoCategoryJobList
	NonExistAIVideoCategoryJobIds ListAIVideoCategoryJobNonExistAIVideoCategoryJobIdList
}

type ListAIVideoCategoryJobAIVideoCategoryJobList []ListAIVideoCategoryJobAIVideoCategoryJob

func (list *ListAIVideoCategoryJobAIVideoCategoryJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAIVideoCategoryJobAIVideoCategoryJob)
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

type ListAIVideoCategoryJobNonExistAIVideoCategoryJobIdList []string

func (list *ListAIVideoCategoryJobNonExistAIVideoCategoryJobIdList) UnmarshalJSON(data []byte) error {
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

func (c *VodClient) ListAIJob(req *ListAIJobArgs) (resp *ListAIJobResponse, err error) {
	resp = &ListAIJobResponse{}
	err = c.Request("ListAIJob", req, resp)
	return
}

type ListAIJobAIJob struct {
	JobId        string
	MediaId      string
	Type         string
	Status       string
	Code         string
	Message      string
	CreationTime string
	CompleteTime string
	Data         string
}
type ListAIJobArgs struct {
	ResourceOwnerId      string
	ResourceOwnerAccount string
	OwnerAccount         string
	JobIds               string
	OwnerId              string
}
type ListAIJobResponse struct {
	RequestId        string
	AIJobList        ListAIJobAIJobList
	NonExistAIJobIds ListAIJobNonExistAIJobIdList
}

type ListAIJobAIJobList []ListAIJobAIJob

func (list *ListAIJobAIJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAIJobAIJob)
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

type ListAIJobNonExistAIJobIdList []string

func (list *ListAIJobNonExistAIJobIdList) UnmarshalJSON(data []byte) error {
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

func (c *VodClient) AddEditingProject(req *AddEditingProjectArgs) (resp *AddEditingProjectResponse, err error) {
	resp = &AddEditingProjectResponse{}
	err = c.Request("AddEditingProject", req, resp)
	return
}

type AddEditingProjectProject struct {
	ProjectId    string
	CreationTime string
	ModifiedTime string
	Status       string
	Description  string
	Title        string
}
type AddEditingProjectArgs struct {
	CoverURL             string
	ResourceOwnerId      string
	ResourceOwnerAccount string
	OwnerAccount         string
	Description          string
	Timeline             string
	OwnerId              string
	Title                string
}
type AddEditingProjectResponse struct {
	RequestId string
	Project   AddEditingProjectProject
}

func (c *VodClient) GetMessageCallback(req *GetMessageCallbackArgs) (resp *GetMessageCallbackResponse, err error) {
	resp = &GetMessageCallbackResponse{}
	err = c.Request("GetMessageCallback", req, resp)
	return
}

type GetMessageCallbackMessageCallback struct {
	CallbackSwitch string
	CallbackURL    string
	EventTypeList  string
}
type GetMessageCallbackArgs struct {
	ResourceOwnerId      string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              string
}
type GetMessageCallbackResponse struct {
	RequestId       string
	MessageCallback GetMessageCallbackMessageCallback
}

func (c *VodClient) OpenVodService(req *OpenVodServiceArgs) (resp *OpenVodServiceResponse, err error) {
	resp = &OpenVodServiceResponse{}
	err = c.Request("OpenVodService", req, resp)
	return
}

type OpenVodServiceArgs struct {
	ResourceOwnerId      string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              string
}
type OpenVodServiceResponse struct {
	RequestId string
	Success   core.Bool
	Code      string
	Message   string
}

func (c *VodClient) GetPlayInfo(req *GetPlayInfoArgs) (resp *GetPlayInfoResponse, err error) {
	resp = &GetPlayInfoResponse{}
	err = c.Request("GetPlayInfo", req, resp)
	return
}

type GetPlayInfoPlayInfo struct {
	Width      int64
	Height     int64
	Size       int64
	PlayURL    string
	Bitrate    string
	Definition string
	Duration   string
	Format     string
	Fps        string
	Encrypt    int64
	Plaintext  string
	Complexity string
	StreamType string
	Rand       string
	JobId      string
}

type GetPlayInfoVideoBase struct {
	CoverURL     string
	Duration     string
	Status       string
	Title        string
	VideoId      string
	MediaType    string
	CreationTime string
}
type GetPlayInfoArgs struct {
	ResourceOwnerId      int64
	StreamType           string
	Formats              string
	ResourceOwnerAccount string
	Channel              string
	VideoId              string
	PlayerVersion        string
	OwnerId              int64
	Rand                 string
	ReAuthInfo           string
	Definition           string
	AuthTimeout          int64
	AuthInfo             string
}
type GetPlayInfoResponse struct {
	RequestId    string
	PlayInfoList GetPlayInfoPlayInfoList
	VideoBase    GetPlayInfoVideoBase
}

type GetPlayInfoPlayInfoList []GetPlayInfoPlayInfo

func (list *GetPlayInfoPlayInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetPlayInfoPlayInfo)
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

func (c *VodClient) GetEditingProject(req *GetEditingProjectArgs) (resp *GetEditingProjectResponse, err error) {
	resp = &GetEditingProjectResponse{}
	err = c.Request("GetEditingProject", req, resp)
	return
}

type GetEditingProjectProject struct {
	ProjectId    string
	CreationTime string
	ModifiedTime string
	Status       string
	Description  string
	Title        string
	Timeline     string
	CoverURL     string
}
type GetEditingProjectArgs struct {
	ResourceOwnerId      string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              string
	ProjectId            string
}
type GetEditingProjectResponse struct {
	RequestId string
	Project   GetEditingProjectProject
}

func (c *VodClient) ListAIVideoTerrorismRecogJob(req *ListAIVideoTerrorismRecogJobArgs) (resp *ListAIVideoTerrorismRecogJobResponse, err error) {
	resp = &ListAIVideoTerrorismRecogJobResponse{}
	err = c.Request("ListAIVideoTerrorismRecogJob", req, resp)
	return
}

type ListAIVideoTerrorismRecogJobAIVideoTerrorismRecogJob struct {
	JobId        string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
}
type ListAIVideoTerrorismRecogJobArgs struct {
	ResourceOwnerId             string
	AIVideoTerrorismRecogJobIds string
	ResourceOwnerAccount        string
	OwnerAccount                string
	OwnerId                     string
}
type ListAIVideoTerrorismRecogJobResponse struct {
	RequestId                    string
	AIVideoTerrorismRecogJobList ListAIVideoTerrorismRecogJobAIVideoTerrorismRecogJobList
	NonExistTerrorismRecogJobIds ListAIVideoTerrorismRecogJobNonExistTerrorismRecogJobIdList
}

type ListAIVideoTerrorismRecogJobAIVideoTerrorismRecogJobList []ListAIVideoTerrorismRecogJobAIVideoTerrorismRecogJob

func (list *ListAIVideoTerrorismRecogJobAIVideoTerrorismRecogJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAIVideoTerrorismRecogJobAIVideoTerrorismRecogJob)
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

type ListAIVideoTerrorismRecogJobNonExistTerrorismRecogJobIdList []string

func (list *ListAIVideoTerrorismRecogJobNonExistTerrorismRecogJobIdList) UnmarshalJSON(data []byte) error {
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

func (c *VodClient) CreateUploadImage(req *CreateUploadImageArgs) (resp *CreateUploadImageResponse, err error) {
	resp = &CreateUploadImageResponse{}
	err = c.Request("CreateUploadImage", req, resp)
	return
}

type CreateUploadImageArgs struct {
	ResourceOwnerId      int64
	ImageType            string
	OriginalFileName     string
	ResourceOwnerAccount string
	ImageExt             string
	OwnerId              int64
	Title                string
	Tags                 string
}
type CreateUploadImageResponse struct {
	RequestId     string
	ImageId       string
	ImageURL      string
	UploadAddress string
	UploadAuth    string
}

func (c *VodClient) GetVideoInfo(req *GetVideoInfoArgs) (resp *GetVideoInfoResponse, err error) {
	resp = &GetVideoInfoResponse{}
	err = c.Request("GetVideoInfo", req, resp)
	return
}

type GetVideoInfoVideo struct {
	VideoId      string
	Title        string
	Tags         string
	Status       string
	Size         int64
	Duration     float32
	Description  string
	CreateTime   string
	CreationTime string
	ModifyTime   string
	CoverURL     string
	CateId       int
	CateName     string
	Snapshots    GetVideoInfoSnapshotList
}
type GetVideoInfoArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	VideoId              string
	ResultTypes          string
	OwnerId              int64
}
type GetVideoInfoResponse struct {
	RequestId string
	AI        string
	Video     GetVideoInfoVideo
}

type GetVideoInfoSnapshotList []string

func (list *GetVideoInfoSnapshotList) UnmarshalJSON(data []byte) error {
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

func (c *VodClient) DeleteStream(req *DeleteStreamArgs) (resp *DeleteStreamResponse, err error) {
	resp = &DeleteStreamResponse{}
	err = c.Request("DeleteStream", req, resp)
	return
}

type DeleteStreamArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	JobIds               string
	VideoId              string
	OwnerId              int64
}
type DeleteStreamResponse struct {
	RequestId string
}

func (c *VodClient) ListAIVideoPornRecogJob(req *ListAIVideoPornRecogJobArgs) (resp *ListAIVideoPornRecogJobResponse, err error) {
	resp = &ListAIVideoPornRecogJobResponse{}
	err = c.Request("ListAIVideoPornRecogJob", req, resp)
	return
}

type ListAIVideoPornRecogJobAIVideoPornRecogJob struct {
	JobId        string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
}
type ListAIVideoPornRecogJobArgs struct {
	ResourceOwnerId        string
	ResourceOwnerAccount   string
	OwnerAccount           string
	AIVideoPornRecogJobIds string
	OwnerId                string
}
type ListAIVideoPornRecogJobResponse struct {
	RequestId               string
	AIVideoPornRecogJobList ListAIVideoPornRecogJobAIVideoPornRecogJobList
	NonExistPornRecogJobIds ListAIVideoPornRecogJobNonExistPornRecogJobIdList
}

type ListAIVideoPornRecogJobAIVideoPornRecogJobList []ListAIVideoPornRecogJobAIVideoPornRecogJob

func (list *ListAIVideoPornRecogJobAIVideoPornRecogJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAIVideoPornRecogJobAIVideoPornRecogJob)
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

type ListAIVideoPornRecogJobNonExistPornRecogJobIdList []string

func (list *ListAIVideoPornRecogJobNonExistPornRecogJobIdList) UnmarshalJSON(data []byte) error {
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

func (c *VodClient) DeleteVideo(req *DeleteVideoArgs) (resp *DeleteVideoResponse, err error) {
	resp = &DeleteVideoResponse{}
	err = c.Request("DeleteVideo", req, resp)
	return
}

type DeleteVideoArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerId              int64
	VideoIds             string
}
type DeleteVideoResponse struct {
	RequestId string
}

func (c *VodClient) AddCategory(req *AddCategoryArgs) (resp *AddCategoryResponse, err error) {
	resp = &AddCategoryResponse{}
	err = c.Request("AddCategory", req, resp)
	return
}

type AddCategoryCategory struct {
	CateId   int64
	CateName string
	ParentId int64
	Level    int64
}
type AddCategoryArgs struct {
	ResourceOwnerId      string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              string
	ParentId             int64
	CateName             string
}
type AddCategoryResponse struct {
	RequestId string
	Category  AddCategoryCategory
}

func (c *VodClient) UpdateCategory(req *UpdateCategoryArgs) (resp *UpdateCategoryResponse, err error) {
	resp = &UpdateCategoryResponse{}
	err = c.Request("UpdateCategory", req, resp)
	return
}

type UpdateCategoryArgs struct {
	ResourceOwnerId      string
	ResourceOwnerAccount string
	CateId               int64
	OwnerAccount         string
	OwnerId              string
	CateName             string
}
type UpdateCategoryResponse struct {
	RequestId string
}

func (c *VodClient) ProduceEditingProjectVideo(req *ProduceEditingProjectVideoArgs) (resp *ProduceEditingProjectVideoResponse, err error) {
	resp = &ProduceEditingProjectVideoResponse{}
	err = c.Request("ProduceEditingProjectVideo", req, resp)
	return
}

type ProduceEditingProjectVideoArgs struct {
	CoverURL             string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	Timeline             string
	Description          string
	OwnerId              int64
	Title                string
	ProjectId            string
}
type ProduceEditingProjectVideoResponse struct {
	RequestId string
	MediaId   string
}

func (c *VodClient) GetOSSStatis(req *GetOSSStatisArgs) (resp *GetOSSStatisResponse, err error) {
	resp = &GetOSSStatisResponse{}
	err = c.Request("GetOSSStatis", req, resp)
	return
}

type GetOSSStatisOSSMetric struct {
	StatTime           string
	StorageUtilization int64
}
type GetOSSStatisArgs struct {
	ResourceOwnerId      string
	StartStatisTime      string
	ResourceOwnerAccount string
	Level                string
	OwnerAccount         string
	OwnerId              string
	EndStatisTime        string
}
type GetOSSStatisResponse struct {
	RequestId             string
	MaxStorageUtilization int64
	OssStatisList         GetOSSStatisOSSMetricList
}

type GetOSSStatisOSSMetricList []GetOSSStatisOSSMetric

func (list *GetOSSStatisOSSMetricList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetOSSStatisOSSMetric)
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

func (c *VodClient) GetVideoConfig(req *GetVideoConfigArgs) (resp *GetVideoConfigResponse, err error) {
	resp = &GetVideoConfigResponse{}
	err = c.Request("GetVideoConfig", req, resp)
	return
}

type GetVideoConfigArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	VideoId              string
	OwnerId              int64
	AuthInfo             string
}
type GetVideoConfigResponse struct {
	RequestId      string
	DownloadSwitch string
}

func (c *VodClient) SubmitAIASRJob(req *SubmitAIASRJobArgs) (resp *SubmitAIASRJobResponse, err error) {
	resp = &SubmitAIASRJobResponse{}
	err = c.Request("SubmitAIASRJob", req, resp)
	return
}

type SubmitAIASRJobAIASRJob struct {
	JobId        string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
}
type SubmitAIASRJobArgs struct {
	UserData             string
	ResourceOwnerId      string
	ResourceOwnerAccount string
	OwnerAccount         string
	AIASRConfig          string
	OwnerId              string
	MediaId              string
}
type SubmitAIASRJobResponse struct {
	RequestId string
	AIASRJob  SubmitAIASRJobAIASRJob
}

func (c *VodClient) SearchEditingProject(req *SearchEditingProjectArgs) (resp *SearchEditingProjectResponse, err error) {
	resp = &SearchEditingProjectResponse{}
	err = c.Request("SearchEditingProject", req, resp)
	return
}

type SearchEditingProjectProject struct {
	ProjectId    string
	CreationTime string
	ModifiedTime string
	Status       string
	Description  string
	Title        string
	CoverURL     string
}
type SearchEditingProjectArgs struct {
	ResourceOwnerId      string
	ResourceOwnerAccount string
	OwnerAccount         string
	EndTime              string
	StartTime            string
	OwnerId              string
	Title                string
	PageNo               int
	PageSize             int
	SortBy               string
	Status               string
}
type SearchEditingProjectResponse struct {
	RequestId   string
	Total       int
	ProjectList SearchEditingProjectProjectList
}

type SearchEditingProjectProjectList []SearchEditingProjectProject

func (list *SearchEditingProjectProjectList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SearchEditingProjectProject)
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

func (c *VodClient) ListAIVideoCoverJob(req *ListAIVideoCoverJobArgs) (resp *ListAIVideoCoverJobResponse, err error) {
	resp = &ListAIVideoCoverJobResponse{}
	err = c.Request("ListAIVideoCoverJob", req, resp)
	return
}

type ListAIVideoCoverJobAIVideoCoverJob struct {
	JobId        string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
}
type ListAIVideoCoverJobArgs struct {
	ResourceOwnerId      string
	ResourceOwnerAccount string
	OwnerAccount         string
	AIVideoCoverJobIds   string
	OwnerId              string
}
type ListAIVideoCoverJobResponse struct {
	RequestId                  string
	AIVideoCoverJobList        ListAIVideoCoverJobAIVideoCoverJobList
	NonExistAIVideoCoverJobIds ListAIVideoCoverJobNonExistAIVideoCoverJobIdList
}

type ListAIVideoCoverJobAIVideoCoverJobList []ListAIVideoCoverJobAIVideoCoverJob

func (list *ListAIVideoCoverJobAIVideoCoverJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAIVideoCoverJobAIVideoCoverJob)
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

type ListAIVideoCoverJobNonExistAIVideoCoverJobIdList []string

func (list *ListAIVideoCoverJobNonExistAIVideoCoverJobIdList) UnmarshalJSON(data []byte) error {
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

func (c *VodClient) DescribeDomainBpsData(req *DescribeDomainBpsDataArgs) (resp *DescribeDomainBpsDataResponse, err error) {
	resp = &DescribeDomainBpsDataResponse{}
	err = c.Request("DescribeDomainBpsData", req, resp)
	return
}

type DescribeDomainBpsDataDataModule struct {
	TimeStamp            string
	Value                string
	DomesticValue        string
	OverseasValue        string
	L2Value              string
	DomesticL2Value      string
	OverseasL2Value      string
	DynamicValue         int64
	DynamicDomesticValue string
	DynamicOverseasValue string
	StaticValue          string
	StaticDomesticValue  string
	StaticOverseasValue  string
}

type DescribeDomainBpsDataDataModule1 struct {
	TimeStamp string
	Value     string
}
type DescribeDomainBpsDataArgs struct {
	ResourceOwnerId      string
	ResourceOwnerAccount string
	TimeMerge            string
	OwnerAccount         string
	DomainName           string
	EndTime              string
	LocationNameEn       string
	StartTime            string
	IspNameEn            string
	OwnerId              string
	Interval             string
}
type DescribeDomainBpsDataResponse struct {
	RequestId          string
	DomainName         string
	DataInterval       string
	StartTime          string
	EndTime            string
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

func (c *VodClient) DeleteCategory(req *DeleteCategoryArgs) (resp *DeleteCategoryResponse, err error) {
	resp = &DeleteCategoryResponse{}
	err = c.Request("DeleteCategory", req, resp)
	return
}

type DeleteCategoryArgs struct {
	ResourceOwnerId      string
	ResourceOwnerAccount string
	CateId               int64
	OwnerAccount         string
	OwnerId              string
}
type DeleteCategoryResponse struct {
	RequestId string
}

func (c *VodClient) SubmitAIJob(req *SubmitAIJobArgs) (resp *SubmitAIJobResponse, err error) {
	resp = &SubmitAIJobResponse{}
	err = c.Request("SubmitAIJob", req, resp)
	return
}

type SubmitAIJobAIJob struct {
	JobId        string
	Type         string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
}
type SubmitAIJobArgs struct {
	UserData             string
	ResourceOwnerId      string
	Types                string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              string
	MediaId              string
	Config               string
}
type SubmitAIJobResponse struct {
	RequestId string
	AIJobList SubmitAIJobAIJobList
}

type SubmitAIJobAIJobList []SubmitAIJobAIJob

func (list *SubmitAIJobAIJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitAIJobAIJob)
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

func (c *VodClient) SubmitAIVideoCategoryJob(req *SubmitAIVideoCategoryJobArgs) (resp *SubmitAIVideoCategoryJobResponse, err error) {
	resp = &SubmitAIVideoCategoryJobResponse{}
	err = c.Request("SubmitAIVideoCategoryJob", req, resp)
	return
}

type SubmitAIVideoCategoryJobAIVideoCategoryJob struct {
	JobId        string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
}
type SubmitAIVideoCategoryJobArgs struct {
	AIVideoCategoryConfig string
	UserData              string
	ResourceOwnerId       string
	ResourceOwnerAccount  string
	OwnerAccount          string
	OwnerId               string
	MediaId               string
}
type SubmitAIVideoCategoryJobResponse struct {
	RequestId          string
	AIVideoCategoryJob SubmitAIVideoCategoryJobAIVideoCategoryJob
}

func (c *VodClient) SubmitAIVideoTerrorismRecogJob(req *SubmitAIVideoTerrorismRecogJobArgs) (resp *SubmitAIVideoTerrorismRecogJobResponse, err error) {
	resp = &SubmitAIVideoTerrorismRecogJobResponse{}
	err = c.Request("SubmitAIVideoTerrorismRecogJob", req, resp)
	return
}

type SubmitAIVideoTerrorismRecogJobAIVideoTerrorismRecogJob struct {
	JobId        string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
}
type SubmitAIVideoTerrorismRecogJobArgs struct {
	UserData                    string
	ResourceOwnerId             string
	ResourceOwnerAccount        string
	OwnerAccount                string
	AIVideoTerrorismRecogConfig string
	OwnerId                     string
	MediaId                     string
}
type SubmitAIVideoTerrorismRecogJobResponse struct {
	RequestId                string
	AIVideoTerrorismRecogJob SubmitAIVideoTerrorismRecogJobAIVideoTerrorismRecogJob
}

func (c *VodClient) GetEditingProjectMaterials(req *GetEditingProjectMaterialsArgs) (resp *GetEditingProjectMaterialsResponse, err error) {
	resp = &GetEditingProjectMaterialsResponse{}
	err = c.Request("GetEditingProjectMaterials", req, resp)
	return
}

type GetEditingProjectMaterialsMaterial struct {
	MaterialId   string
	Title        string
	Tags         string
	Status       string
	Size         int64
	Duration     float32
	Description  string
	CreationTime string
	ModifiedTime string
	CoverURL     string
	CateId       int
	CateName     string
	Source       string
	Snapshots    GetEditingProjectMaterialsSnapshotList
	Sprites      GetEditingProjectMaterialsSpriteList
}
type GetEditingProjectMaterialsArgs struct {
	ResourceOwnerId      string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              string
	Type                 string
	ProjectId            string
}
type GetEditingProjectMaterialsResponse struct {
	RequestId    string
	MaterialList GetEditingProjectMaterialsMaterialList
}

type GetEditingProjectMaterialsSnapshotList []string

func (list *GetEditingProjectMaterialsSnapshotList) UnmarshalJSON(data []byte) error {
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

type GetEditingProjectMaterialsSpriteList []string

func (list *GetEditingProjectMaterialsSpriteList) UnmarshalJSON(data []byte) error {
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

type GetEditingProjectMaterialsMaterialList []GetEditingProjectMaterialsMaterial

func (list *GetEditingProjectMaterialsMaterialList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetEditingProjectMaterialsMaterial)
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

func (c *VodClient) SubmitAIVideoSummaryJob(req *SubmitAIVideoSummaryJobArgs) (resp *SubmitAIVideoSummaryJobResponse, err error) {
	resp = &SubmitAIVideoSummaryJobResponse{}
	err = c.Request("SubmitAIVideoSummaryJob", req, resp)
	return
}

type SubmitAIVideoSummaryJobAIVideoSummaryJob struct {
	JobId        string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
}
type SubmitAIVideoSummaryJobArgs struct {
	UserData             string
	ResourceOwnerId      string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              string
	MediaId              string
	AIVideoSummaryConfig string
}
type SubmitAIVideoSummaryJobResponse struct {
	RequestId         string
	AIVideoSummaryJob SubmitAIVideoSummaryJobAIVideoSummaryJob
}

func (c *VodClient) DeleteEditingProject(req *DeleteEditingProjectArgs) (resp *DeleteEditingProjectResponse, err error) {
	resp = &DeleteEditingProjectResponse{}
	err = c.Request("DeleteEditingProject", req, resp)
	return
}

type DeleteEditingProjectArgs struct {
	ResourceOwnerId      string
	ResourceOwnerAccount string
	OwnerAccount         string
	ProjectIds           string
	OwnerId              string
}
type DeleteEditingProjectResponse struct {
	RequestId string
}

func (c *VodClient) SubmitAIVideoPornRecogJob(req *SubmitAIVideoPornRecogJobArgs) (resp *SubmitAIVideoPornRecogJobResponse, err error) {
	resp = &SubmitAIVideoPornRecogJobResponse{}
	err = c.Request("SubmitAIVideoPornRecogJob", req, resp)
	return
}

type SubmitAIVideoPornRecogJobAIVideoPornRecogJob struct {
	JobId        string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
}
type SubmitAIVideoPornRecogJobArgs struct {
	UserData               string
	ResourceOwnerId        string
	ResourceOwnerAccount   string
	OwnerAccount           string
	AIVideoPornRecogConfig string
	OwnerId                string
	MediaId                string
}
type SubmitAIVideoPornRecogJobResponse struct {
	RequestId           string
	AIVideoPornRecogJob SubmitAIVideoPornRecogJobAIVideoPornRecogJob
}

func (c *VodClient) SubmitAIVideoCensorJob(req *SubmitAIVideoCensorJobArgs) (resp *SubmitAIVideoCensorJobResponse, err error) {
	resp = &SubmitAIVideoCensorJobResponse{}
	err = c.Request("SubmitAIVideoCensorJob", req, resp)
	return
}

type SubmitAIVideoCensorJobAIVideoCensorJob struct {
	JobId        string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
}
type SubmitAIVideoCensorJobArgs struct {
	UserData             string
	ResourceOwnerId      string
	AIVideoCensorConfig  string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              string
	MediaId              string
}
type SubmitAIVideoCensorJobResponse struct {
	RequestId        string
	AIVideoCensorJob SubmitAIVideoCensorJobAIVideoCensorJob
}

func (c *VodClient) GetMezzanineInfo(req *GetMezzanineInfoArgs) (resp *GetMezzanineInfoResponse, err error) {
	resp = &GetMezzanineInfoResponse{}
	err = c.Request("GetMezzanineInfo", req, resp)
	return
}

type GetMezzanineInfoMezzanine struct {
	VideoId      string
	Bitrate      string
	CreationTime string
	Duration     string
	Fps          string
	Height       int64
	Width        int64
	Size         int64
	Status       string
	FileURL      string
	FileName     string
	CRC64        string
}
type GetMezzanineInfoArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	VideoId              string
	OwnerId              int64
	AuthTimeout          int64
}
type GetMezzanineInfoResponse struct {
	RequestId string
	Mezzanine GetMezzanineInfoMezzanine
}

func (c *VodClient) UpdateVideoInfo(req *UpdateVideoInfoArgs) (resp *UpdateVideoInfoResponse, err error) {
	resp = &UpdateVideoInfoResponse{}
	err = c.Request("UpdateVideoInfo", req, resp)
	return
}

type UpdateVideoInfoArgs struct {
	CoverURL             string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	CateId               int
	Description          string
	VideoId              string
	OwnerId              int64
	Title                string
	Tags                 string
}
type UpdateVideoInfoResponse struct {
	RequestId string
}

func (c *VodClient) GetVideoPlayAuth(req *GetVideoPlayAuthArgs) (resp *GetVideoPlayAuthResponse, err error) {
	resp = &GetVideoPlayAuthResponse{}
	err = c.Request("GetVideoPlayAuth", req, resp)
	return
}

type GetVideoPlayAuthVideoMeta struct {
	CoverURL string
	Duration float32
	Status   string
	Title    string
	VideoId  string
}
type GetVideoPlayAuthArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ReAuthInfo           string
	AuthInfoTimeout      int64
	VideoId              string
	OwnerId              int64
}
type GetVideoPlayAuthResponse struct {
	RequestId string
	PlayAuth  string
	VideoMeta GetVideoPlayAuthVideoMeta
}

func (c *VodClient) UpdateEditingProject(req *UpdateEditingProjectArgs) (resp *UpdateEditingProjectResponse, err error) {
	resp = &UpdateEditingProjectResponse{}
	err = c.Request("UpdateEditingProject", req, resp)
	return
}

type UpdateEditingProjectArgs struct {
	CoverURL             string
	ResourceOwnerId      string
	ResourceOwnerAccount string
	OwnerAccount         string
	Timeline             string
	Description          string
	OwnerId              string
	Title                string
	ProjectId            string
}
type UpdateEditingProjectResponse struct {
	RequestId string
}

func (c *VodClient) GetCategories(req *GetCategoriesArgs) (resp *GetCategoriesResponse, err error) {
	resp = &GetCategoriesResponse{}
	err = c.Request("GetCategories", req, resp)
	return
}

type GetCategoriesCategory struct {
	CateId   int64
	CateName string
	Level    int64
	ParentId int64
}
type GetCategoriesArgs struct {
	ResourceOwnerId      string
	ResourceOwnerAccount string
	CateId               int64
	PageNo               int64
	OwnerAccount         string
	PageSize             int64
	OwnerId              string
}
type GetCategoriesResponse struct {
	RequestId     string
	SubTotal      int64
	SubCategories GetCategoriesCategoryList
	Category      GetCategoriesCategory
}

type GetCategoriesCategoryList []GetCategoriesCategory

func (list *GetCategoriesCategoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetCategoriesCategory)
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

func (c *VodClient) SubmitSnapshotJob(req *SubmitSnapshotJobArgs) (resp *SubmitSnapshotJobResponse, err error) {
	resp = &SubmitSnapshotJobResponse{}
	err = c.Request("SubmitSnapshotJob", req, resp)
	return
}

type SubmitSnapshotJobSnapshotJob struct {
	JobId string
}
type SubmitSnapshotJobArgs struct {
	ResourceOwnerId      int64
	SpecifiedOffsetTime  int64
	ResourceOwnerAccount string
	Width                string
	Count                int64
	VideoId              string
	Interval             int64
	OwnerId              int64
	SpriteSnapshotConfig string
	Height               string
}
type SubmitSnapshotJobResponse struct {
	RequestId   string
	SnapshotJob SubmitSnapshotJobSnapshotJob
}

func (c *VodClient) CreateOrder(req *CreateOrderArgs) (resp *CreateOrderResponse, err error) {
	resp = &CreateOrderResponse{}
	err = c.Request("CreateOrder", req, resp)
	return
}

type CreateOrderArgs struct {
	ResourceOwnerId      string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              string
}
type CreateOrderResponse struct {
	RequestId string
	Success   core.Bool
	Code      string
	Message   string
}
