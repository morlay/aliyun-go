package live

import "encoding/json"

func (c *LiveClient) StartMultipleStreamMixService(req *StartMultipleStreamMixServiceArgs) (resp *StartMultipleStreamMixServiceResponse, err error) {
	resp = &StartMultipleStreamMixServiceResponse{}
	err = c.Request("StartMultipleStreamMixService", req, resp)
	return
}

type StartMultipleStreamMixServiceArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	MixTemplate   string
	OwnerId       int64
	StreamName    string
}
type StartMultipleStreamMixServiceResponse struct {
	RequestId string
}

func (c *LiveClient) StartCaster(req *StartCasterArgs) (resp *StartCasterResponse, err error) {
	resp = &StartCasterResponse{}
	err = c.Request("StartCaster", req, resp)
	return
}

type StartCasterSceneInfo struct {
	SceneId   string
	StreamUrl string
}

type StartCasterSceneInfo1 struct {
	SceneId     string
	StreamUrl   string
	StreamInfos StartCasterStreamInfoList
}

type StartCasterStreamInfo struct {
	TranscodeConfig string
	VideoFormat     string
	OutputStreamUrl string
}
type StartCasterArgs struct {
	SecurityToken string
	CasterId      string
	OwnerId       int64
	Version       string
}
type StartCasterResponse struct {
	RequestId     string
	PvwSceneInfos StartCasterSceneInfoList
	PgmSceneInfos StartCasterSceneInfo1List
}

type StartCasterStreamInfoList []StartCasterStreamInfo

func (list *StartCasterStreamInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]StartCasterStreamInfo)
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

type StartCasterSceneInfoList []StartCasterSceneInfo

func (list *StartCasterSceneInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]StartCasterSceneInfo)
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

type StartCasterSceneInfo1List []StartCasterSceneInfo1

func (list *StartCasterSceneInfo1List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]StartCasterSceneInfo1)
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

func (c *LiveClient) EffectCasterUrgent(req *EffectCasterUrgentArgs) (resp *EffectCasterUrgentResponse, err error) {
	resp = &EffectCasterUrgentResponse{}
	err = c.Request("EffectCasterUrgent", req, resp)
	return
}

type EffectCasterUrgentArgs struct {
	SecurityToken string
	CasterId      string
	SceneId       string
	OwnerId       int64
	Version       string
}
type EffectCasterUrgentResponse struct {
	RequestId string
}

func (c *LiveClient) DeleteCasterLayout(req *DeleteCasterLayoutArgs) (resp *DeleteCasterLayoutResponse, err error) {
	resp = &DeleteCasterLayoutResponse{}
	err = c.Request("DeleteCasterLayout", req, resp)
	return
}

type DeleteCasterLayoutArgs struct {
	SecurityToken string
	CasterId      string
	OwnerId       int64
	Version       string
	LayoutId      string
}
type DeleteCasterLayoutResponse struct {
	RequestId string
	CasterId  string
	LayoutId  string
}

func (c *LiveClient) CreateLiveStreamRecordIndexFiles(req *CreateLiveStreamRecordIndexFilesArgs) (resp *CreateLiveStreamRecordIndexFilesResponse, err error) {
	resp = &CreateLiveStreamRecordIndexFilesResponse{}
	err = c.Request("CreateLiveStreamRecordIndexFiles", req, resp)
	return
}

type CreateLiveStreamRecordIndexFilesRecordInfo struct {
	RecordId    string
	RecordUrl   string
	DomainName  string
	AppName     string
	StreamName  string
	OssBucket   string
	OssEndpoint string
	OssObject   string
	StartTime   string
	EndTime     string
	Duration    float32
	Height      int
	Width       int
	CreateTime  string
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

func (c *LiveClient) DeleteLiveSnapshotDetectPornConfig(req *DeleteLiveSnapshotDetectPornConfigArgs) (resp *DeleteLiveSnapshotDetectPornConfigResponse, err error) {
	resp = &DeleteLiveSnapshotDetectPornConfigResponse{}
	err = c.Request("DeleteLiveSnapshotDetectPornConfig", req, resp)
	return
}

type DeleteLiveSnapshotDetectPornConfigArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	OwnerId       int64
}
type DeleteLiveSnapshotDetectPornConfigResponse struct {
	RequestId string
}

func (c *LiveClient) SetLiveStreamsNotifyUrlConfig(req *SetLiveStreamsNotifyUrlConfigArgs) (resp *SetLiveStreamsNotifyUrlConfigResponse, err error) {
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

func (c *LiveClient) DescribeLiveStreamRecordIndexFiles(req *DescribeLiveStreamRecordIndexFilesArgs) (resp *DescribeLiveStreamRecordIndexFilesResponse, err error) {
	resp = &DescribeLiveStreamRecordIndexFilesResponse{}
	err = c.Request("DescribeLiveStreamRecordIndexFiles", req, resp)
	return
}

type DescribeLiveStreamRecordIndexFilesRecordIndexInfo struct {
	RecordId    string
	RecordUrl   string
	DomainName  string
	AppName     string
	StreamName  string
	OssBucket   string
	OssEndpoint string
	OssObject   string
	StartTime   string
	EndTime     string
	Duration    float32
	Height      int
	Width       int
	CreateTime  string
}
type DescribeLiveStreamRecordIndexFilesArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	PageSize      int
	EndTime       string
	StartTime     string
	OwnerId       int64
	PageNum       int
	StreamName    string
	Order         string
}
type DescribeLiveStreamRecordIndexFilesResponse struct {
	RequestId           string
	PageNum             int
	PageSize            int
	Order               string
	TotalNum            int
	TotalPage           int
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

func (c *LiveClient) AddLiveStreamTranscode(req *AddLiveStreamTranscodeArgs) (resp *AddLiveStreamTranscodeResponse, err error) {
	resp = &AddLiveStreamTranscodeResponse{}
	err = c.Request("AddLiveStreamTranscode", req, resp)
	return
}

type AddLiveStreamTranscodeArgs struct {
	App           string
	Template      string
	SecurityToken string
	Domain        string
	OwnerId       int64
}
type AddLiveStreamTranscodeResponse struct {
	RequestId string
}

func (c *LiveClient) AddLiveAppSnapshotConfig(req *AddLiveAppSnapshotConfigArgs) (resp *AddLiveAppSnapshotConfigResponse, err error) {
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

func (c *LiveClient) AddLiveMixConfig(req *AddLiveMixConfigArgs) (resp *AddLiveMixConfigResponse, err error) {
	resp = &AddLiveMixConfigResponse{}
	err = c.Request("AddLiveMixConfig", req, resp)
	return
}

type AddLiveMixConfigArgs struct {
	Template      string
	AppName       string
	SecurityToken string
	DomainName    string
	OwnerId       int64
}
type AddLiveMixConfigResponse struct {
	RequestId string
}

func (c *LiveClient) ResumeLiveStream(req *ResumeLiveStreamArgs) (resp *ResumeLiveStreamResponse, err error) {
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

func (c *LiveClient) DeleteLiveMixConfig(req *DeleteLiveMixConfigArgs) (resp *DeleteLiveMixConfigResponse, err error) {
	resp = &DeleteLiveMixConfigResponse{}
	err = c.Request("DeleteLiveMixConfig", req, resp)
	return
}

type DeleteLiveMixConfigArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	OwnerId       int64
}
type DeleteLiveMixConfigResponse struct {
	RequestId string
}

func (c *LiveClient) ImagePornDetection(req *ImagePornDetectionArgs) (resp *ImagePornDetectionResponse, err error) {
	resp = &ImagePornDetectionResponse{}
	err = c.Request("ImagePornDetection", req, resp)
	return
}

type ImagePornDetectionArgs struct {
	SecurityToken string
	ImageUrl      string
	OwnerId       int64
}
type ImagePornDetectionResponse struct {
	RequestId string
	Label     string
	Rate      float32
}

func (c *LiveClient) DescribeLiveStreamsNotifyUrlConfig(req *DescribeLiveStreamsNotifyUrlConfigArgs) (resp *DescribeLiveStreamsNotifyUrlConfigResponse, err error) {
	resp = &DescribeLiveStreamsNotifyUrlConfigResponse{}
	err = c.Request("DescribeLiveStreamsNotifyUrlConfig", req, resp)
	return
}

type DescribeLiveStreamsNotifyUrlConfigLiveStreamsNotifyConfig struct {
	DomainName string
	NotifyUrl  string
}
type DescribeLiveStreamsNotifyUrlConfigArgs struct {
	SecurityToken string
	DomainName    string
	OwnerId       int64
}
type DescribeLiveStreamsNotifyUrlConfigResponse struct {
	RequestId               string
	LiveStreamsNotifyConfig DescribeLiveStreamsNotifyUrlConfigLiveStreamsNotifyConfig
}

func (c *LiveClient) UpdateLiveRecordNotifyConfig(req *UpdateLiveRecordNotifyConfigArgs) (resp *UpdateLiveRecordNotifyConfigResponse, err error) {
	resp = &UpdateLiveRecordNotifyConfigResponse{}
	err = c.Request("UpdateLiveRecordNotifyConfig", req, resp)
	return
}

type UpdateLiveRecordNotifyConfigArgs struct {
	SecurityToken    string
	DomainName       string
	NotifyUrl        string
	OwnerId          int64
	NeedStatusNotify bool
}
type UpdateLiveRecordNotifyConfigResponse struct {
	RequestId string
}

func (c *LiveClient) UpdateLiveDetectNotifyConfig(req *UpdateLiveDetectNotifyConfigArgs) (resp *UpdateLiveDetectNotifyConfigResponse, err error) {
	resp = &UpdateLiveDetectNotifyConfigResponse{}
	err = c.Request("UpdateLiveDetectNotifyConfig", req, resp)
	return
}

type UpdateLiveDetectNotifyConfigArgs struct {
	SecurityToken string
	DomainName    string
	NotifyUrl     string
	OwnerId       int64
}
type UpdateLiveDetectNotifyConfigResponse struct {
	RequestId string
}

func (c *LiveClient) DescribeLiveStreamTranscodeInfo(req *DescribeLiveStreamTranscodeInfoArgs) (resp *DescribeLiveStreamTranscodeInfoResponse, err error) {
	resp = &DescribeLiveStreamTranscodeInfoResponse{}
	err = c.Request("DescribeLiveStreamTranscodeInfo", req, resp)
	return
}

type DescribeLiveStreamTranscodeInfoDomainTranscodeInfo struct {
	TranscodeApp      string
	TranscodeName     string
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

func (c *LiveClient) AddCasterLayout(req *AddCasterLayoutArgs) (resp *AddCasterLayoutResponse, err error) {
	resp = &AddCasterLayoutResponse{}
	err = c.Request("AddCasterLayout", req, resp)
	return
}

type AddCasterLayoutAudioLayer struct {
	VolumeRate   float32
	ValidChannel string
}

type AddCasterLayoutVideoLayer struct {
	HeightNormalized    float32
	WidthNormalized     float32
	PositionRefer       string
	PositionNormalizeds AddCasterLayoutPositionNormalizedList
}
type AddCasterLayoutArgs struct {
	BlendLists    AddCasterLayoutBlendListList
	AudioLayers   AddCasterLayoutAudioLayerList
	SecurityToken string
	VideoLayers   AddCasterLayoutVideoLayerList
	CasterId      string
	MixLists      AddCasterLayoutMixListList
	OwnerId       int64
	Version       string
}
type AddCasterLayoutResponse struct {
	RequestId string
	LayoutId  string
}

type AddCasterLayoutPositionNormalizedList []float32

func (list *AddCasterLayoutPositionNormalizedList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]float32)
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

type AddCasterLayoutBlendListList []string

func (list *AddCasterLayoutBlendListList) UnmarshalJSON(data []byte) error {
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

type AddCasterLayoutAudioLayerList []AddCasterLayoutAudioLayer

func (list *AddCasterLayoutAudioLayerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddCasterLayoutAudioLayer)
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

type AddCasterLayoutVideoLayerList []AddCasterLayoutVideoLayer

func (list *AddCasterLayoutVideoLayerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddCasterLayoutVideoLayer)
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

type AddCasterLayoutMixListList []string

func (list *AddCasterLayoutMixListList) UnmarshalJSON(data []byte) error {
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

func (c *LiveClient) DescribeCasterLayouts(req *DescribeCasterLayoutsArgs) (resp *DescribeCasterLayoutsResponse, err error) {
	resp = &DescribeCasterLayoutsResponse{}
	err = c.Request("DescribeCasterLayouts", req, resp)
	return
}

type DescribeCasterLayoutsLayout struct {
	LayoutId    string
	VideoLayers DescribeCasterLayoutsVideoLayerList
	AudioLayers DescribeCasterLayoutsAudioLayerList
	BlendList   DescribeCasterLayoutsBlendListList
	MixList     DescribeCasterLayoutsMixListList
}

type DescribeCasterLayoutsVideoLayer struct {
	HeightNormalized    float32
	WidthNormalized     float32
	PositionRefer       string
	PositionNormalizeds DescribeCasterLayoutsPositionNormalizedList
}

type DescribeCasterLayoutsAudioLayer struct {
	VolumeRate   float32
	ValidChannel string
}
type DescribeCasterLayoutsArgs struct {
	SecurityToken string
	CasterId      string
	OwnerId       int64
	Version       string
}
type DescribeCasterLayoutsResponse struct {
	RequestId string
	Total     int
	Layouts   DescribeCasterLayoutsLayoutList
}

type DescribeCasterLayoutsVideoLayerList []DescribeCasterLayoutsVideoLayer

func (list *DescribeCasterLayoutsVideoLayerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterLayoutsVideoLayer)
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

type DescribeCasterLayoutsAudioLayerList []DescribeCasterLayoutsAudioLayer

func (list *DescribeCasterLayoutsAudioLayerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterLayoutsAudioLayer)
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

type DescribeCasterLayoutsBlendListList []string

func (list *DescribeCasterLayoutsBlendListList) UnmarshalJSON(data []byte) error {
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

type DescribeCasterLayoutsMixListList []string

func (list *DescribeCasterLayoutsMixListList) UnmarshalJSON(data []byte) error {
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

type DescribeCasterLayoutsPositionNormalizedList []string

func (list *DescribeCasterLayoutsPositionNormalizedList) UnmarshalJSON(data []byte) error {
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

type DescribeCasterLayoutsLayoutList []DescribeCasterLayoutsLayout

func (list *DescribeCasterLayoutsLayoutList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterLayoutsLayout)
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

func (c *LiveClient) DeleteLivePullStreamInfoConfig(req *DeleteLivePullStreamInfoConfigArgs) (resp *DeleteLivePullStreamInfoConfigResponse, err error) {
	resp = &DeleteLivePullStreamInfoConfigResponse{}
	err = c.Request("DeleteLivePullStreamInfoConfig", req, resp)
	return
}

type DeleteLivePullStreamInfoConfigArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	OwnerId       int64
	StreamName    string
}
type DeleteLivePullStreamInfoConfigResponse struct {
	RequestId string
}

func (c *LiveClient) DeleteLiveAppRecordConfig(req *DeleteLiveAppRecordConfigArgs) (resp *DeleteLiveAppRecordConfigResponse, err error) {
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

func (c *LiveClient) AddCasterVideoResource(req *AddCasterVideoResourceArgs) (resp *AddCasterVideoResourceResponse, err error) {
	resp = &AddCasterVideoResourceResponse{}
	err = c.Request("AddCasterVideoResource", req, resp)
	return
}

type AddCasterVideoResourceArgs struct {
	LiveStreamUrl string
	SecurityToken string
	LocationId    string
	CasterId      string
	ResourceName  string
	RepeatNum     int
	OwnerId       int64
	MaterialId    string
	Version       string
}
type AddCasterVideoResourceResponse struct {
	RequestId  string
	ResourceId string
}

func (c *LiveClient) DescribeLiveDetectNotifyConfig(req *DescribeLiveDetectNotifyConfigArgs) (resp *DescribeLiveDetectNotifyConfigResponse, err error) {
	resp = &DescribeLiveDetectNotifyConfigResponse{}
	err = c.Request("DescribeLiveDetectNotifyConfig", req, resp)
	return
}

type DescribeLiveDetectNotifyConfigLiveDetectNotifyConfig struct {
	DomainName string
	NotifyUrl  string
}
type DescribeLiveDetectNotifyConfigArgs struct {
	SecurityToken string
	DomainName    string
	OwnerId       int64
}
type DescribeLiveDetectNotifyConfigResponse struct {
	RequestId              string
	LiveDetectNotifyConfig DescribeLiveDetectNotifyConfigLiveDetectNotifyConfig
}

func (c *LiveClient) DescribeLiveStreamSnapshotInfo(req *DescribeLiveStreamSnapshotInfoArgs) (resp *DescribeLiveStreamSnapshotInfoResponse, err error) {
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

func (c *LiveClient) DescribeLiveRecordNotifyConfig(req *DescribeLiveRecordNotifyConfigArgs) (resp *DescribeLiveRecordNotifyConfigResponse, err error) {
	resp = &DescribeLiveRecordNotifyConfigResponse{}
	err = c.Request("DescribeLiveRecordNotifyConfig", req, resp)
	return
}

type DescribeLiveRecordNotifyConfigLiveRecordNotifyConfig struct {
	DomainName       string
	NotifyUrl        string
	NeedStatusNotify bool
}
type DescribeLiveRecordNotifyConfigArgs struct {
	SecurityToken string
	DomainName    string
	OwnerId       int64
}
type DescribeLiveRecordNotifyConfigResponse struct {
	RequestId              string
	LiveRecordNotifyConfig DescribeLiveRecordNotifyConfigLiveRecordNotifyConfig
}

func (c *LiveClient) DeleteCaster(req *DeleteCasterArgs) (resp *DeleteCasterResponse, err error) {
	resp = &DeleteCasterResponse{}
	err = c.Request("DeleteCaster", req, resp)
	return
}

type DeleteCasterArgs struct {
	SecurityToken string
	CasterId      string
	OwnerId       int64
	Version       string
}
type DeleteCasterResponse struct {
	RequestId string
	CasterId  string
}

func (c *LiveClient) DescribeCasterStreamUrl(req *DescribeCasterStreamUrlArgs) (resp *DescribeCasterStreamUrlResponse, err error) {
	resp = &DescribeCasterStreamUrlResponse{}
	err = c.Request("DescribeCasterStreamUrl", req, resp)
	return
}

type DescribeCasterStreamUrlCasterStream struct {
	SceneId     string
	StreamUrl   string
	OutputType  int
	StreamInfos DescribeCasterStreamUrlStreamInfoList
}

type DescribeCasterStreamUrlStreamInfo struct {
	TranscodeConfig string
	VideoFormat     string
	OutputStreamUrl string
}
type DescribeCasterStreamUrlArgs struct {
	SecurityToken string
	CasterId      string
	OwnerId       int64
	Version       string
}
type DescribeCasterStreamUrlResponse struct {
	RequestId     string
	CasterId      string
	Total         int
	CasterStreams DescribeCasterStreamUrlCasterStreamList
}

type DescribeCasterStreamUrlStreamInfoList []DescribeCasterStreamUrlStreamInfo

func (list *DescribeCasterStreamUrlStreamInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterStreamUrlStreamInfo)
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

type DescribeCasterStreamUrlCasterStreamList []DescribeCasterStreamUrlCasterStream

func (list *DescribeCasterStreamUrlCasterStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterStreamUrlCasterStream)
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

func (c *LiveClient) DescribeLiveMixConfig(req *DescribeLiveMixConfigArgs) (resp *DescribeLiveMixConfigResponse, err error) {
	resp = &DescribeLiveMixConfigResponse{}
	err = c.Request("DescribeLiveMixConfig", req, resp)
	return
}

type DescribeLiveMixConfigMixConfig struct {
	DomainName string
	AppName    string
	Template   string
}
type DescribeLiveMixConfigArgs struct {
	SecurityToken string
	DomainName    string
	OwnerId       int64
}
type DescribeLiveMixConfigResponse struct {
	RequestId     string
	MixConfigList DescribeLiveMixConfigMixConfigList
}

type DescribeLiveMixConfigMixConfigList []DescribeLiveMixConfigMixConfig

func (list *DescribeLiveMixConfigMixConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveMixConfigMixConfig)
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

func (c *LiveClient) StopCaster(req *StopCasterArgs) (resp *StopCasterResponse, err error) {
	resp = &StopCasterResponse{}
	err = c.Request("StopCaster", req, resp)
	return
}

type StopCasterArgs struct {
	SecurityToken string
	CasterId      string
	OwnerId       int64
	Version       string
}
type StopCasterResponse struct {
	RequestId string
}

func (c *LiveClient) ModifyCasterComponent(req *ModifyCasterComponentArgs) (resp *ModifyCasterComponentResponse, err error) {
	resp = &ModifyCasterComponentResponse{}
	err = c.Request("ModifyCasterComponent", req, resp)
	return
}

type ModifyCasterComponentArgs struct {
	ComponentId       string
	ImageLayerContent string
	CasterId          string
	ComponentLayer    string
	ComponentName     string
	OwnerId           int64
	Version           string
	ComponentType     string
	SecurityToken     string
	Effect            string
	TextLayerContent  string
}
type ModifyCasterComponentResponse struct {
	RequestId   string
	ComponentId string
}

func (c *LiveClient) SetCasterConfig(req *SetCasterConfigArgs) (resp *SetCasterConfigResponse, err error) {
	resp = &SetCasterConfigResponse{}
	err = c.Request("SetCasterConfig", req, resp)
	return
}

type SetCasterConfigArgs struct {
	UrgentMaterialId string
	TranscodeConfig  string
	Delay            float32
	SecurityToken    string
	CasterName       string
	CasterId         string
	DomainName       string
	OwnerId          int64
	Version          string
	RecordConfig     string
}
type SetCasterConfigResponse struct {
	RequestId string
	CasterId  string
}

func (c *LiveClient) DescribeCasters(req *DescribeCastersArgs) (resp *DescribeCastersResponse, err error) {
	resp = &DescribeCastersResponse{}
	err = c.Request("DescribeCasters", req, resp)
	return
}

type DescribeCastersCaster struct {
	Status         int
	NormType       int
	CasterId       string
	CasterName     string
	CreateTime     string
	Period         int
	ChargeType     string
	CasterTemplate string
}
type DescribeCastersArgs struct {
	SecurityToken string
	CasterName    string
	CasterId      string
	PageSize      int
	EndTime       string
	StartTime     string
	OwnerId       int64
	PageNum       int
	Version       string
	Status        int
}
type DescribeCastersResponse struct {
	RequestId  string
	Total      int
	CasterList DescribeCastersCasterList
}

type DescribeCastersCasterList []DescribeCastersCaster

func (list *DescribeCastersCasterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCastersCaster)
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

func (c *LiveClient) StartMixStreamsService(req *StartMixStreamsServiceArgs) (resp *StartMixStreamsServiceResponse, err error) {
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

func (c *LiveClient) AddLiveMixNotifyConfig(req *AddLiveMixNotifyConfigArgs) (resp *AddLiveMixNotifyConfigResponse, err error) {
	resp = &AddLiveMixNotifyConfigResponse{}
	err = c.Request("AddLiveMixNotifyConfig", req, resp)
	return
}

type AddLiveMixNotifyConfigArgs struct {
	SecurityToken string
	DomainName    string
	NotifyUrl     string
	OwnerId       int64
}
type AddLiveMixNotifyConfigResponse struct {
	RequestId string
}

func (c *LiveClient) StopMixStreamsService(req *StopMixStreamsServiceArgs) (resp *StopMixStreamsServiceResponse, err error) {
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

func (c *LiveClient) DescribeLiveMixNotifyConfig(req *DescribeLiveMixNotifyConfigArgs) (resp *DescribeLiveMixNotifyConfigResponse, err error) {
	resp = &DescribeLiveMixNotifyConfigResponse{}
	err = c.Request("DescribeLiveMixNotifyConfig", req, resp)
	return
}

type DescribeLiveMixNotifyConfigArgs struct {
	SecurityToken string
	DomainName    string
	OwnerId       int64
}
type DescribeLiveMixNotifyConfigResponse struct {
	RequestId string
	NotifyUrl string
}

func (c *LiveClient) UpdateCasterSceneConfig(req *UpdateCasterSceneConfigArgs) (resp *UpdateCasterSceneConfigResponse, err error) {
	resp = &UpdateCasterSceneConfigResponse{}
	err = c.Request("UpdateCasterSceneConfig", req, resp)
	return
}

type UpdateCasterSceneConfigArgs struct {
	ComponentIds  UpdateCasterSceneConfigComponentIdList
	SecurityToken string
	CasterId      string
	SceneId       string
	OwnerId       int64
	Version       string
	LayoutId      string
}
type UpdateCasterSceneConfigResponse struct {
	RequestId string
}

type UpdateCasterSceneConfigComponentIdList []string

func (list *UpdateCasterSceneConfigComponentIdList) UnmarshalJSON(data []byte) error {
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

func (c *LiveClient) DescribeCasterComponents(req *DescribeCasterComponentsArgs) (resp *DescribeCasterComponentsResponse, err error) {
	resp = &DescribeCasterComponentsResponse{}
	err = c.Request("DescribeCasterComponents", req, resp)
	return
}

type DescribeCasterComponentsComponent struct {
	ComponentId       string
	ComponentName     string
	LocationId        string
	ComponentType     string
	Effect            string
	ComponentLayer    DescribeCasterComponentsComponentLayer
	TextLayerContent  DescribeCasterComponentsTextLayerContent
	ImageLayerContent DescribeCasterComponentsImageLayerContent
}

type DescribeCasterComponentsComponentLayer struct {
	HeightNormalized    float32
	WidthNormalized     float32
	PositionRefer       string
	PositionNormalizeds DescribeCasterComponentsPositionNormalizedList
}

type DescribeCasterComponentsTextLayerContent struct {
	Text                  string
	Color                 string
	FontName              string
	SizeNormalized        float32
	BorderWidthNormalized float32
	BorderColor           string
}

type DescribeCasterComponentsImageLayerContent struct {
	MaterialId string
}
type DescribeCasterComponentsArgs struct {
	ComponentId   string
	SecurityToken string
	CasterId      string
	OwnerId       int64
	Version       string
}
type DescribeCasterComponentsResponse struct {
	RequestId  string
	Total      int
	Components DescribeCasterComponentsComponentList
}

type DescribeCasterComponentsPositionNormalizedList []string

func (list *DescribeCasterComponentsPositionNormalizedList) UnmarshalJSON(data []byte) error {
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

type DescribeCasterComponentsComponentList []DescribeCasterComponentsComponent

func (list *DescribeCasterComponentsComponentList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterComponentsComponent)
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

func (c *LiveClient) DescribeLiveStreamRecordIndexFile(req *DescribeLiveStreamRecordIndexFileArgs) (resp *DescribeLiveStreamRecordIndexFileResponse, err error) {
	resp = &DescribeLiveStreamRecordIndexFileResponse{}
	err = c.Request("DescribeLiveStreamRecordIndexFile", req, resp)
	return
}

type DescribeLiveStreamRecordIndexFileRecordIndexInfo struct {
	RecordId    string
	RecordUrl   string
	DomainName  string
	AppName     string
	StreamName  string
	OssBucket   string
	OssEndpoint string
	OssObject   string
	StartTime   string
	EndTime     string
	Duration    float32
	Height      int
	Width       int
	CreateTime  string
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

func (c *LiveClient) DeleteCasterComponent(req *DeleteCasterComponentArgs) (resp *DeleteCasterComponentResponse, err error) {
	resp = &DeleteCasterComponentResponse{}
	err = c.Request("DeleteCasterComponent", req, resp)
	return
}

type DeleteCasterComponentArgs struct {
	ComponentId   string
	SecurityToken string
	CasterId      string
	OwnerId       int64
	Version       string
}
type DeleteCasterComponentResponse struct {
	RequestId   string
	CasterId    string
	ComponentId string
}

func (c *LiveClient) DescribeCasterConfig(req *DescribeCasterConfigArgs) (resp *DescribeCasterConfigResponse, err error) {
	resp = &DescribeCasterConfigResponse{}
	err = c.Request("DescribeCasterConfig", req, resp)
	return
}

type DescribeCasterConfigTranscodeConfig struct {
	CasterTemplate  string
	LiveTemplateIds DescribeCasterConfigLiveTemplateIdList
}

type DescribeCasterConfigRecordConfig struct {
	OssEndpoint  string
	OssBucket    string
	RecordFormat DescribeCasterConfigRecordFormatItemList
}

type DescribeCasterConfigRecordFormatItem struct {
	Format               string
	OssObjectPrefix      string
	SliceOssObjectPrefix string
	CycleDuration        int
}
type DescribeCasterConfigArgs struct {
	SecurityToken string
	CasterId      string
	OwnerId       int64
	Version       string
}
type DescribeCasterConfigResponse struct {
	RequestId        string
	CasterId         string
	CasterName       string
	DomainName       string
	Delay            float32
	UrgentMaterialId string
	TranscodeConfig  DescribeCasterConfigTranscodeConfig
	RecordConfig     DescribeCasterConfigRecordConfig
}

type DescribeCasterConfigLiveTemplateIdList []string

func (list *DescribeCasterConfigLiveTemplateIdList) UnmarshalJSON(data []byte) error {
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

type DescribeCasterConfigRecordFormatItemList []DescribeCasterConfigRecordFormatItem

func (list *DescribeCasterConfigRecordFormatItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterConfigRecordFormatItem)
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

func (c *LiveClient) UpdateLiveAppSnapshotConfig(req *UpdateLiveAppSnapshotConfigArgs) (resp *UpdateLiveAppSnapshotConfigResponse, err error) {
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

func (c *LiveClient) StopMultipleStreamMixService(req *StopMultipleStreamMixServiceArgs) (resp *StopMultipleStreamMixServiceResponse, err error) {
	resp = &StopMultipleStreamMixServiceResponse{}
	err = c.Request("StopMultipleStreamMixService", req, resp)
	return
}

type StopMultipleStreamMixServiceArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	OwnerId       int64
	StreamName    string
}
type StopMultipleStreamMixServiceResponse struct {
	RequestId string
}

func (c *LiveClient) DescribeLiveSnapshotConfig(req *DescribeLiveSnapshotConfigArgs) (resp *DescribeLiveSnapshotConfigResponse, err error) {
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

func (c *LiveClient) DescribeLiveStreamsControlHistory(req *DescribeLiveStreamsControlHistoryArgs) (resp *DescribeLiveStreamsControlHistoryResponse, err error) {
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

func (c *LiveClient) StartCasterScene(req *StartCasterSceneArgs) (resp *StartCasterSceneResponse, err error) {
	resp = &StartCasterSceneResponse{}
	err = c.Request("StartCasterScene", req, resp)
	return
}

type StartCasterSceneArgs struct {
	SecurityToken string
	CasterId      string
	SceneId       string
	OwnerId       int64
	Version       string
}
type StartCasterSceneResponse struct {
	RequestId string
	StreamUrl string
}

func (c *LiveClient) AddLiveAppRecordConfig(req *AddLiveAppRecordConfigArgs) (resp *AddLiveAppRecordConfigResponse, err error) {
	resp = &AddLiveAppRecordConfigResponse{}
	err = c.Request("AddLiveAppRecordConfig", req, resp)
	return
}

type AddLiveAppRecordConfigRecordFormat struct {
	Format               string
	OssObjectPrefix      string
	SliceOssObjectPrefix string
	CycleDuration        int
}
type AddLiveAppRecordConfigArgs struct {
	OssBucket     string
	AppName       string
	SecurityToken string
	RecordFormats AddLiveAppRecordConfigRecordFormatList
	DomainName    string
	OssEndpoint   string
	OwnerId       int64
}
type AddLiveAppRecordConfigResponse struct {
	RequestId string
}

type AddLiveAppRecordConfigRecordFormatList []AddLiveAppRecordConfigRecordFormat

func (list *AddLiveAppRecordConfigRecordFormatList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddLiveAppRecordConfigRecordFormat)
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

func (c *LiveClient) DeleteLiveStreamsNotifyUrlConfig(req *DeleteLiveStreamsNotifyUrlConfigArgs) (resp *DeleteLiveStreamsNotifyUrlConfigResponse, err error) {
	resp = &DeleteLiveStreamsNotifyUrlConfigResponse{}
	err = c.Request("DeleteLiveStreamsNotifyUrlConfig", req, resp)
	return
}

type DeleteLiveStreamsNotifyUrlConfigArgs struct {
	SecurityToken string
	DomainName    string
	OwnerId       int64
}
type DeleteLiveStreamsNotifyUrlConfigResponse struct {
	RequestId string
}

func (c *LiveClient) CopyCasterSceneConfig(req *CopyCasterSceneConfigArgs) (resp *CopyCasterSceneConfigResponse, err error) {
	resp = &CopyCasterSceneConfigResponse{}
	err = c.Request("CopyCasterSceneConfig", req, resp)
	return
}

type CopyCasterSceneConfigArgs struct {
	FromSceneId   string
	SecurityToken string
	CasterId      string
	OwnerId       int64
	Version       string
	ToSceneId     string
}
type CopyCasterSceneConfigResponse struct {
	RequestId string
}

func (c *LiveClient) DeleteCasterVideoResource(req *DeleteCasterVideoResourceArgs) (resp *DeleteCasterVideoResourceResponse, err error) {
	resp = &DeleteCasterVideoResourceResponse{}
	err = c.Request("DeleteCasterVideoResource", req, resp)
	return
}

type DeleteCasterVideoResourceArgs struct {
	ResourceId    string
	SecurityToken string
	CasterId      string
	OwnerId       int64
	Version       string
}
type DeleteCasterVideoResourceResponse struct {
	RequestId string
}

func (c *LiveClient) RemoveMultipleStreamMixService(req *RemoveMultipleStreamMixServiceArgs) (resp *RemoveMultipleStreamMixServiceResponse, err error) {
	resp = &RemoveMultipleStreamMixServiceResponse{}
	err = c.Request("RemoveMultipleStreamMixService", req, resp)
	return
}

type RemoveMultipleStreamMixServiceArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	MixStreamName string
	MixDomainName string
	OwnerId       int64
	MixAppName    string
	StreamName    string
}
type RemoveMultipleStreamMixServiceResponse struct {
	RequestId string
}

func (c *LiveClient) DescribeLiveStreamRecordContent(req *DescribeLiveStreamRecordContentArgs) (resp *DescribeLiveStreamRecordContentResponse, err error) {
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

func (c *LiveClient) ModifyCasterLayout(req *ModifyCasterLayoutArgs) (resp *ModifyCasterLayoutResponse, err error) {
	resp = &ModifyCasterLayoutResponse{}
	err = c.Request("ModifyCasterLayout", req, resp)
	return
}

type ModifyCasterLayoutAudioLayer struct {
	VolumeRate   float32
	ValidChannel string
}

type ModifyCasterLayoutVideoLayer struct {
	HeightNormalized    float32
	WidthNormalized     float32
	PositionRefer       string
	PositionNormalizeds ModifyCasterLayoutPositionNormalizedList
}
type ModifyCasterLayoutArgs struct {
	BlendLists    ModifyCasterLayoutBlendListList
	AudioLayers   ModifyCasterLayoutAudioLayerList
	SecurityToken string
	VideoLayers   ModifyCasterLayoutVideoLayerList
	CasterId      string
	MixLists      ModifyCasterLayoutMixListList
	OwnerId       int64
	Version       string
	LayoutId      string
}
type ModifyCasterLayoutResponse struct {
	RequestId string
	LayoutId  string
}

type ModifyCasterLayoutPositionNormalizedList []float32

func (list *ModifyCasterLayoutPositionNormalizedList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]float32)
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

type ModifyCasterLayoutBlendListList []string

func (list *ModifyCasterLayoutBlendListList) UnmarshalJSON(data []byte) error {
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

type ModifyCasterLayoutAudioLayerList []ModifyCasterLayoutAudioLayer

func (list *ModifyCasterLayoutAudioLayerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyCasterLayoutAudioLayer)
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

type ModifyCasterLayoutVideoLayerList []ModifyCasterLayoutVideoLayer

func (list *ModifyCasterLayoutVideoLayerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyCasterLayoutVideoLayer)
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

type ModifyCasterLayoutMixListList []string

func (list *ModifyCasterLayoutMixListList) UnmarshalJSON(data []byte) error {
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

func (c *LiveClient) DescribeCasterScenes(req *DescribeCasterScenesArgs) (resp *DescribeCasterScenesResponse, err error) {
	resp = &DescribeCasterScenesResponse{}
	err = c.Request("DescribeCasterScenes", req, resp)
	return
}

type DescribeCasterScenesScene struct {
	SceneId      string
	SceneName    string
	OutputType   string
	LayoutId     string
	StreamUrl    string
	Status       int
	StreamInfos  DescribeCasterScenesStreamInfoList
	ComponentIds DescribeCasterScenesComponentIdList
}

type DescribeCasterScenesStreamInfo struct {
	TranscodeConfig string
	VideoFormat     string
	OutputStreamUrl string
}
type DescribeCasterScenesArgs struct {
	SecurityToken string
	CasterId      string
	SceneId       string
	OwnerId       int64
	Version       string
}
type DescribeCasterScenesResponse struct {
	RequestId string
	Total     int
	SceneList DescribeCasterScenesSceneList
}

type DescribeCasterScenesStreamInfoList []DescribeCasterScenesStreamInfo

func (list *DescribeCasterScenesStreamInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterScenesStreamInfo)
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

type DescribeCasterScenesComponentIdList []string

func (list *DescribeCasterScenesComponentIdList) UnmarshalJSON(data []byte) error {
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

type DescribeCasterScenesSceneList []DescribeCasterScenesScene

func (list *DescribeCasterScenesSceneList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterScenesScene)
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

func (c *LiveClient) SetCasterSceneConfig(req *SetCasterSceneConfigArgs) (resp *SetCasterSceneConfigResponse, err error) {
	resp = &SetCasterSceneConfigResponse{}
	err = c.Request("SetCasterSceneConfig", req, resp)
	return
}

type SetCasterSceneConfigArgs struct {
	ComponentIds  SetCasterSceneConfigComponentIdList
	SecurityToken string
	CasterId      string
	SceneId       string
	OwnerId       int64
	Version       string
	LayoutId      string
}
type SetCasterSceneConfigResponse struct {
	RequestId string
}

type SetCasterSceneConfigComponentIdList []string

func (list *SetCasterSceneConfigComponentIdList) UnmarshalJSON(data []byte) error {
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

func (c *LiveClient) DescribeCasterVideoResources(req *DescribeCasterVideoResourcesArgs) (resp *DescribeCasterVideoResourcesResponse, err error) {
	resp = &DescribeCasterVideoResourcesResponse{}
	err = c.Request("DescribeCasterVideoResources", req, resp)
	return
}

type DescribeCasterVideoResourcesVideoResource struct {
	MaterialId    string
	ResourceId    string
	ResourceName  string
	LocationId    string
	LiveStreamUrl string
	RepeatNum     int
}
type DescribeCasterVideoResourcesArgs struct {
	SecurityToken string
	CasterId      string
	OwnerId       int64
	Version       string
}
type DescribeCasterVideoResourcesResponse struct {
	RequestId      string
	Total          int
	VideoResources DescribeCasterVideoResourcesVideoResourceList
}

type DescribeCasterVideoResourcesVideoResourceList []DescribeCasterVideoResourcesVideoResource

func (list *DescribeCasterVideoResourcesVideoResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterVideoResourcesVideoResource)
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

func (c *LiveClient) DescribeLiveStreamOnlineUserNum(req *DescribeLiveStreamOnlineUserNumArgs) (resp *DescribeLiveStreamOnlineUserNumResponse, err error) {
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

func (c *LiveClient) DeleteLiveAppSnapshotConfig(req *DeleteLiveAppSnapshotConfigArgs) (resp *DeleteLiveAppSnapshotConfigResponse, err error) {
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

func (c *LiveClient) UpdateLiveMixNotifyConfig(req *UpdateLiveMixNotifyConfigArgs) (resp *UpdateLiveMixNotifyConfigResponse, err error) {
	resp = &UpdateLiveMixNotifyConfigResponse{}
	err = c.Request("UpdateLiveMixNotifyConfig", req, resp)
	return
}

type UpdateLiveMixNotifyConfigArgs struct {
	SecurityToken string
	DomainName    string
	NotifyUrl     string
	OwnerId       int64
}
type UpdateLiveMixNotifyConfigResponse struct {
	RequestId string
}

func (c *LiveClient) AddLiveRecordNotifyConfig(req *AddLiveRecordNotifyConfigArgs) (resp *AddLiveRecordNotifyConfigResponse, err error) {
	resp = &AddLiveRecordNotifyConfigResponse{}
	err = c.Request("AddLiveRecordNotifyConfig", req, resp)
	return
}

type AddLiveRecordNotifyConfigArgs struct {
	SecurityToken    string
	DomainName       string
	NotifyUrl        string
	OwnerId          int64
	NeedStatusNotify bool
}
type AddLiveRecordNotifyConfigResponse struct {
	RequestId string
}

func (c *LiveClient) AddLiveDetectNotifyConfig(req *AddLiveDetectNotifyConfigArgs) (resp *AddLiveDetectNotifyConfigResponse, err error) {
	resp = &AddLiveDetectNotifyConfigResponse{}
	err = c.Request("AddLiveDetectNotifyConfig", req, resp)
	return
}

type AddLiveDetectNotifyConfigArgs struct {
	SecurityToken string
	DomainName    string
	NotifyUrl     string
	OwnerId       int64
}
type AddLiveDetectNotifyConfigResponse struct {
	RequestId string
}

func (c *LiveClient) DescribeLiveStreamBitRateData(req *DescribeLiveStreamBitRateDataArgs) (resp *DescribeLiveStreamBitRateDataResponse, err error) {
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

func (c *LiveClient) DeleteLiveStreamTranscode(req *DeleteLiveStreamTranscodeArgs) (resp *DeleteLiveStreamTranscodeResponse, err error) {
	resp = &DeleteLiveStreamTranscodeResponse{}
	err = c.Request("DeleteLiveStreamTranscode", req, resp)
	return
}

type DeleteLiveStreamTranscodeArgs struct {
	App           string
	Template      string
	SecurityToken string
	Domain        string
	OwnerId       int64
}
type DeleteLiveStreamTranscodeResponse struct {
	RequestId string
}

func (c *LiveClient) AddLiveSnapshotDetectPornConfig(req *AddLiveSnapshotDetectPornConfigArgs) (resp *AddLiveSnapshotDetectPornConfigResponse, err error) {
	resp = &AddLiveSnapshotDetectPornConfigResponse{}
	err = c.Request("AddLiveSnapshotDetectPornConfig", req, resp)
	return
}

type AddLiveSnapshotDetectPornConfigArgs struct {
	OssBucket     string
	AppName       string
	SecurityToken string
	DomainName    string
	OssEndpoint   string
	Interval      int
	OwnerId       int64
	OssObject     string
	Scenes        AddLiveSnapshotDetectPornConfigSceneList
}
type AddLiveSnapshotDetectPornConfigResponse struct {
	RequestId string
}

type AddLiveSnapshotDetectPornConfigSceneList []string

func (list *AddLiveSnapshotDetectPornConfigSceneList) UnmarshalJSON(data []byte) error {
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

func (c *LiveClient) DescribeLivePullStreamConfig(req *DescribeLivePullStreamConfigArgs) (resp *DescribeLivePullStreamConfigResponse, err error) {
	resp = &DescribeLivePullStreamConfigResponse{}
	err = c.Request("DescribeLivePullStreamConfig", req, resp)
	return
}

type DescribeLivePullStreamConfigLiveAppRecord struct {
	DomainName string
	AppName    string
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

func (c *LiveClient) ModifyCasterVideoResource(req *ModifyCasterVideoResourceArgs) (resp *ModifyCasterVideoResourceResponse, err error) {
	resp = &ModifyCasterVideoResourceResponse{}
	err = c.Request("ModifyCasterVideoResource", req, resp)
	return
}

type ModifyCasterVideoResourceArgs struct {
	ResourceId    string
	LiveStreamUrl string
	SecurityToken string
	CasterId      string
	ResourceName  string
	RepeatNum     int
	OwnerId       int64
	MaterialId    string
	Version       string
}
type ModifyCasterVideoResourceResponse struct {
	RequestId  string
	CasterId   string
	ResourceId string
}

func (c *LiveClient) EffectCasterVideoResource(req *EffectCasterVideoResourceArgs) (resp *EffectCasterVideoResourceResponse, err error) {
	resp = &EffectCasterVideoResourceResponse{}
	err = c.Request("EffectCasterVideoResource", req, resp)
	return
}

type EffectCasterVideoResourceArgs struct {
	ResourceId    string
	SecurityToken string
	CasterId      string
	SceneId       string
	OwnerId       int64
	Version       string
}
type EffectCasterVideoResourceResponse struct {
	RequestId string
}

func (c *LiveClient) UpdateLiveSnapshotDetectPornConfig(req *UpdateLiveSnapshotDetectPornConfigArgs) (resp *UpdateLiveSnapshotDetectPornConfigResponse, err error) {
	resp = &UpdateLiveSnapshotDetectPornConfigResponse{}
	err = c.Request("UpdateLiveSnapshotDetectPornConfig", req, resp)
	return
}

type UpdateLiveSnapshotDetectPornConfigArgs struct {
	OssBucket     string
	AppName       string
	SecurityToken string
	DomainName    string
	OssEndpoint   string
	Interval      int
	OwnerId       int64
	OssObject     string
	Scenes        UpdateLiveSnapshotDetectPornConfigSceneList
}
type UpdateLiveSnapshotDetectPornConfigResponse struct {
	RequestId string
}

type UpdateLiveSnapshotDetectPornConfigSceneList []string

func (list *UpdateLiveSnapshotDetectPornConfigSceneList) UnmarshalJSON(data []byte) error {
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

func (c *LiveClient) DeleteLiveMixNotifyConfig(req *DeleteLiveMixNotifyConfigArgs) (resp *DeleteLiveMixNotifyConfigResponse, err error) {
	resp = &DeleteLiveMixNotifyConfigResponse{}
	err = c.Request("DeleteLiveMixNotifyConfig", req, resp)
	return
}

type DeleteLiveMixNotifyConfigArgs struct {
	SecurityToken string
	DomainName    string
	OwnerId       int64
}
type DeleteLiveMixNotifyConfigResponse struct {
	RequestId string
}

func (c *LiveClient) DescribeLiveSnapshotDetectPornConfig(req *DescribeLiveSnapshotDetectPornConfigArgs) (resp *DescribeLiveSnapshotDetectPornConfigResponse, err error) {
	resp = &DescribeLiveSnapshotDetectPornConfigResponse{}
	err = c.Request("DescribeLiveSnapshotDetectPornConfig", req, resp)
	return
}

type DescribeLiveSnapshotDetectPornConfigLiveSnapshotDetectPornConfig struct {
	DomainName  string
	AppName     string
	OssEndpoint string
	OssBucket   string
	OssObject   string
	Interval    int
}
type DescribeLiveSnapshotDetectPornConfigArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	PageSize      int
	OwnerId       int64
	PageNum       int
	Order         string
}
type DescribeLiveSnapshotDetectPornConfigResponse struct {
	RequestId                        string
	PageNum                          int
	PageSize                         int
	Order                            string
	TotalNum                         int
	TotalPage                        int
	LiveSnapshotDetectPornConfigList DescribeLiveSnapshotDetectPornConfigLiveSnapshotDetectPornConfigList
}

type DescribeLiveSnapshotDetectPornConfigLiveSnapshotDetectPornConfigList []DescribeLiveSnapshotDetectPornConfigLiveSnapshotDetectPornConfig

func (list *DescribeLiveSnapshotDetectPornConfigLiveSnapshotDetectPornConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveSnapshotDetectPornConfigLiveSnapshotDetectPornConfig)
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

func (c *LiveClient) DescribeLiveStreamHistoryUserNum(req *DescribeLiveStreamHistoryUserNumArgs) (resp *DescribeLiveStreamHistoryUserNumResponse, err error) {
	resp = &DescribeLiveStreamHistoryUserNumResponse{}
	err = c.Request("DescribeLiveStreamHistoryUserNum", req, resp)
	return
}

type DescribeLiveStreamHistoryUserNumLiveStreamUserNumInfo struct {
	StreamTime string
	UserNum    string
}
type DescribeLiveStreamHistoryUserNumArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
	StreamName    string
}
type DescribeLiveStreamHistoryUserNumResponse struct {
	RequestId              string
	LiveStreamUserNumInfos DescribeLiveStreamHistoryUserNumLiveStreamUserNumInfoList
}

type DescribeLiveStreamHistoryUserNumLiveStreamUserNumInfoList []DescribeLiveStreamHistoryUserNumLiveStreamUserNumInfo

func (list *DescribeLiveStreamHistoryUserNumLiveStreamUserNumInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamHistoryUserNumLiveStreamUserNumInfo)
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

func (c *LiveClient) DescribeLiveStreamsBlockList(req *DescribeLiveStreamsBlockListArgs) (resp *DescribeLiveStreamsBlockListResponse, err error) {
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

func (c *LiveClient) DescribeLiveStreamsOnlineList(req *DescribeLiveStreamsOnlineListArgs) (resp *DescribeLiveStreamsOnlineListResponse, err error) {
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

func (c *LiveClient) CopyCaster(req *CopyCasterArgs) (resp *CopyCasterResponse, err error) {
	resp = &CopyCasterResponse{}
	err = c.Request("CopyCaster", req, resp)
	return
}

type CopyCasterArgs struct {
	SrcCasterId   string
	SecurityToken string
	CasterName    string
	ClientToken   string
	OwnerId       int64
	Version       string
}
type CopyCasterResponse struct {
	RequestId string
	CasterId  string
}

func (c *LiveClient) DescribeLiveRecordConfig(req *DescribeLiveRecordConfigArgs) (resp *DescribeLiveRecordConfigResponse, err error) {
	resp = &DescribeLiveRecordConfigResponse{}
	err = c.Request("DescribeLiveRecordConfig", req, resp)
	return
}

type DescribeLiveRecordConfigLiveAppRecord struct {
	DomainName       string
	AppName          string
	OssEndpoint      string
	OssBucket        string
	CreateTime       string
	RecordFormatList DescribeLiveRecordConfigRecordFormatList
}

type DescribeLiveRecordConfigRecordFormat struct {
	Format               string
	OssObjectPrefix      string
	SliceOssObjectPrefix string
	CycleDuration        int
}
type DescribeLiveRecordConfigArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	PageSize      int
	OwnerId       int64
	PageNum       int
	Order         string
}
type DescribeLiveRecordConfigResponse struct {
	RequestId         string
	PageNum           int
	PageSize          int
	Order             string
	TotalNum          int
	TotalPage         int
	LiveAppRecordList DescribeLiveRecordConfigLiveAppRecordList
}

type DescribeLiveRecordConfigRecordFormatList []DescribeLiveRecordConfigRecordFormat

func (list *DescribeLiveRecordConfigRecordFormatList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveRecordConfigRecordFormat)
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

func (c *LiveClient) AddMultipleStreamMixService(req *AddMultipleStreamMixServiceArgs) (resp *AddMultipleStreamMixServiceResponse, err error) {
	resp = &AddMultipleStreamMixServiceResponse{}
	err = c.Request("AddMultipleStreamMixService", req, resp)
	return
}

type AddMultipleStreamMixServiceArgs struct {
	AppName       string
	SecurityToken string
	DomainName    string
	MixStreamName string
	MixDomainName string
	OwnerId       int64
	MixAppName    string
	StreamName    string
}
type AddMultipleStreamMixServiceResponse struct {
	RequestId string
}

func (c *LiveClient) StopCasterScene(req *StopCasterSceneArgs) (resp *StopCasterSceneResponse, err error) {
	resp = &StopCasterSceneResponse{}
	err = c.Request("StopCasterScene", req, resp)
	return
}

type StopCasterSceneArgs struct {
	SecurityToken string
	CasterId      string
	SceneId       string
	OwnerId       int64
	Version       string
}
type StopCasterSceneResponse struct {
	RequestId string
}

func (c *LiveClient) DescribeLiveStreamsPublishList(req *DescribeLiveStreamsPublishListArgs) (resp *DescribeLiveStreamsPublishListResponse, err error) {
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

func (c *LiveClient) CreateCaster(req *CreateCasterArgs) (resp *CreateCasterResponse, err error) {
	resp = &CreateCasterResponse{}
	err = c.Request("CreateCaster", req, resp)
	return
}

type CreateCasterArgs struct {
	CasterTemplate string
	NormType       int
	Period         int
	SecurityToken  string
	CasterName     string
	ClientToken    string
	ChargeType     string
	OwnerId        int64
	Version        string
}
type CreateCasterResponse struct {
	RequestId string
	CasterId  string
}

func (c *LiveClient) DeleteLiveDetectNotifyConfig(req *DeleteLiveDetectNotifyConfigArgs) (resp *DeleteLiveDetectNotifyConfigResponse, err error) {
	resp = &DeleteLiveDetectNotifyConfigResponse{}
	err = c.Request("DeleteLiveDetectNotifyConfig", req, resp)
	return
}

type DeleteLiveDetectNotifyConfigArgs struct {
	SecurityToken string
	DomainName    string
	OwnerId       int64
}
type DeleteLiveDetectNotifyConfigResponse struct {
	RequestId string
}

func (c *LiveClient) DeleteLiveRecordNotifyConfig(req *DeleteLiveRecordNotifyConfigArgs) (resp *DeleteLiveRecordNotifyConfigResponse, err error) {
	resp = &DeleteLiveRecordNotifyConfigResponse{}
	err = c.Request("DeleteLiveRecordNotifyConfig", req, resp)
	return
}

type DeleteLiveRecordNotifyConfigArgs struct {
	SecurityToken string
	DomainName    string
	OwnerId       int64
}
type DeleteLiveRecordNotifyConfigResponse struct {
	RequestId string
}

func (c *LiveClient) DescribeLiveStreamsFrameRateAndBitRateData(req *DescribeLiveStreamsFrameRateAndBitRateDataArgs) (resp *DescribeLiveStreamsFrameRateAndBitRateDataResponse, err error) {
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

func (c *LiveClient) ForbidLiveStream(req *ForbidLiveStreamArgs) (resp *ForbidLiveStreamResponse, err error) {
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

func (c *LiveClient) AddCasterComponent(req *AddCasterComponentArgs) (resp *AddCasterComponentResponse, err error) {
	resp = &AddCasterComponentResponse{}
	err = c.Request("AddCasterComponent", req, resp)
	return
}

type AddCasterComponentArgs struct {
	ImageLayerContent string
	CasterId          string
	ComponentLayer    string
	ComponentName     string
	OwnerId           int64
	Version           string
	ComponentType     string
	SecurityToken     string
	LocationId        string
	Effect            string
	TextLayerContent  string
}
type AddCasterComponentResponse struct {
	RequestId   string
	ComponentId string
}

func (c *LiveClient) AddLivePullStreamInfoConfig(req *AddLivePullStreamInfoConfigArgs) (resp *AddLivePullStreamInfoConfigResponse, err error) {
	resp = &AddLivePullStreamInfoConfigResponse{}
	err = c.Request("AddLivePullStreamInfoConfig", req, resp)
	return
}

type AddLivePullStreamInfoConfigArgs struct {
	SourceUrl     string
	AppName       string
	SecurityToken string
	DomainName    string
	EndTime       string
	StartTime     string
	OwnerId       int64
	StreamName    string
}
type AddLivePullStreamInfoConfigResponse struct {
	RequestId string
}
