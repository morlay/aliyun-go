package mts

import (
	"encoding/json"

	"github.com/morlay/aliyun-go/core"
)

func (c *MtsClient) QueryAnalysisJobList(req *QueryAnalysisJobListArgs) (resp *QueryAnalysisJobListResponse, err error) {
	resp = &QueryAnalysisJobListResponse{}
	err = c.Request("QueryAnalysisJobList", req, resp)
	return
}

type QueryAnalysisJobListAnalysisJob struct {
	Id               string
	UserData         string
	State            string
	Code             string
	Message          string
	Percent          int64
	CreationTime     string
	PipelineId       string
	Priority         string
	TemplateList     QueryAnalysisJobListTemplateList
	InputFile        QueryAnalysisJobListInputFile
	AnalysisConfig   QueryAnalysisJobListAnalysisConfig
	MNSMessageResult QueryAnalysisJobListMNSMessageResult
}

type QueryAnalysisJobListTemplate struct {
	Id          string
	Name        string
	State       string
	Container   QueryAnalysisJobListContainer
	Video       QueryAnalysisJobListVideo
	Audio       QueryAnalysisJobListAudio
	TransConfig QueryAnalysisJobListTransConfig
	MuxConfig   QueryAnalysisJobListMuxConfig
}

type QueryAnalysisJobListContainer struct {
	Format string
}

type QueryAnalysisJobListVideo struct {
	Codec      string
	Profile    string
	Bitrate    string
	Crf        string
	Width      string
	Height     string
	Fps        string
	Gop        string
	Preset     string
	ScanMode   string
	Bufsize    string
	Maxrate    string
	PixFmt     string
	Degrain    string
	Qscale     string
	BitrateBnd QueryAnalysisJobListBitrateBnd
}

type QueryAnalysisJobListBitrateBnd struct {
	Max string
	Min string
}

type QueryAnalysisJobListAudio struct {
	Codec      string
	Profile    string
	Samplerate string
	Bitrate    string
	Channels   string
	Qscale     string
}

type QueryAnalysisJobListTransConfig struct {
	TransMode string
}

type QueryAnalysisJobListMuxConfig struct {
	Segment QueryAnalysisJobListSegment
	Gif     QueryAnalysisJobListGif
}

type QueryAnalysisJobListSegment struct {
	Duration string
}

type QueryAnalysisJobListGif struct {
	Loop       string
	FinalDelay string
}

type QueryAnalysisJobListInputFile struct {
	Bucket   string
	Location string
	Object   string
}

type QueryAnalysisJobListAnalysisConfig struct {
	QualityControl    QueryAnalysisJobListQualityControl
	PropertiesControl QueryAnalysisJobListPropertiesControl
}

type QueryAnalysisJobListQualityControl struct {
	RateQuality     string
	MethodStreaming string
}

type QueryAnalysisJobListPropertiesControl struct {
	Deinterlace string
	Crop        QueryAnalysisJobListCrop
}

type QueryAnalysisJobListCrop struct {
	Mode   string
	Width  string
	Height string
	Top    string
	Left   string
}

type QueryAnalysisJobListMNSMessageResult struct {
	MessageId    string
	ErrorMessage string
	ErrorCode    string
}
type QueryAnalysisJobListArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	AnalysisJobIds       string
}
type QueryAnalysisJobListResponse struct {
	RequestId              string
	AnalysisJobList        QueryAnalysisJobListAnalysisJobList
	NonExistAnalysisJobIds QueryAnalysisJobListNonExistAnalysisJobIdList
}

type QueryAnalysisJobListTemplateList []QueryAnalysisJobListTemplate

func (list *QueryAnalysisJobListTemplateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAnalysisJobListTemplate)
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

type QueryAnalysisJobListAnalysisJobList []QueryAnalysisJobListAnalysisJob

func (list *QueryAnalysisJobListAnalysisJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAnalysisJobListAnalysisJob)
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

type QueryAnalysisJobListNonExistAnalysisJobIdList []string

func (list *QueryAnalysisJobListNonExistAnalysisJobIdList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) UnbindInputBucket(req *UnbindInputBucketArgs) (resp *UnbindInputBucketResponse, err error) {
	resp = &UnbindInputBucketResponse{}
	err = c.Request("UnbindInputBucket", req, resp)
	return
}

type UnbindInputBucketArgs struct {
	Bucket               string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	RoleArn              string
	OwnerAccount         string
	OwnerId              int64
}
type UnbindInputBucketResponse struct {
	RequestId string
}

func (c *MtsClient) QueryMediaWorkflowList(req *QueryMediaWorkflowListArgs) (resp *QueryMediaWorkflowListResponse, err error) {
	resp = &QueryMediaWorkflowListResponse{}
	err = c.Request("QueryMediaWorkflowList", req, resp)
	return
}

type QueryMediaWorkflowListMediaWorkflow struct {
	MediaWorkflowId string
	Name            string
	Topology        string
	State           string
	CreationTime    string
}
type QueryMediaWorkflowListArgs struct {
	MediaWorkflowIds     string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type QueryMediaWorkflowListResponse struct {
	RequestId                string
	MediaWorkflowList        QueryMediaWorkflowListMediaWorkflowList
	NonExistMediaWorkflowIds QueryMediaWorkflowListNonExistMediaWorkflowIdList
}

type QueryMediaWorkflowListMediaWorkflowList []QueryMediaWorkflowListMediaWorkflow

func (list *QueryMediaWorkflowListMediaWorkflowList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaWorkflowListMediaWorkflow)
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

type QueryMediaWorkflowListNonExistMediaWorkflowIdList []string

func (list *QueryMediaWorkflowListNonExistMediaWorkflowIdList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) QueryTerrorismPipelineList(req *QueryTerrorismPipelineListArgs) (resp *QueryTerrorismPipelineListResponse, err error) {
	resp = &QueryTerrorismPipelineListResponse{}
	err = c.Request("QueryTerrorismPipelineList", req, resp)
	return
}

type QueryTerrorismPipelineListPipeline struct {
	Id           string
	Name         string
	State        string
	Priority     string
	NotifyConfig QueryTerrorismPipelineListNotifyConfig
}

type QueryTerrorismPipelineListNotifyConfig struct {
	Topic string
	Queue string
}
type QueryTerrorismPipelineListArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	PipelineIds          string
	OwnerAccount         string
	OwnerId              int64
}
type QueryTerrorismPipelineListResponse struct {
	RequestId    string
	PipelineList QueryTerrorismPipelineListPipelineList
	NonExistIds  QueryTerrorismPipelineListNonExistIdList
}

type QueryTerrorismPipelineListPipelineList []QueryTerrorismPipelineListPipeline

func (list *QueryTerrorismPipelineListPipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTerrorismPipelineListPipeline)
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

type QueryTerrorismPipelineListNonExistIdList []string

func (list *QueryTerrorismPipelineListNonExistIdList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) SubmitPornJob(req *SubmitPornJobArgs) (resp *SubmitPornJobResponse, err error) {
	resp = &SubmitPornJobResponse{}
	err = c.Request("SubmitPornJob", req, resp)
	return
}

type SubmitPornJobArgs struct {
	Input                string
	UserData             string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	PornConfig           string
	PipelineId           string
}
type SubmitPornJobResponse struct {
	RequestId string
	JobId     string
}

func (c *MtsClient) QueryMediaInfoJobList(req *QueryMediaInfoJobListArgs) (resp *QueryMediaInfoJobListResponse, err error) {
	resp = &QueryMediaInfoJobListResponse{}
	err = c.Request("QueryMediaInfoJobList", req, resp)
	return
}

type QueryMediaInfoJobListMediaInfoJob struct {
	JobId            string
	UserData         string
	PipelineId       string
	State            string
	Code             string
	Message          string
	CreationTime     string
	Input            QueryMediaInfoJobListInput
	Properties       QueryMediaInfoJobListProperties
	MNSMessageResult QueryMediaInfoJobListMNSMessageResult
}

type QueryMediaInfoJobListInput struct {
	Bucket   string
	Location string
	Object   string
}

type QueryMediaInfoJobListProperties struct {
	Width      string
	Height     string
	Bitrate    string
	Duration   string
	Fps        string
	FileSize   string
	FileFormat string
	Streams    QueryMediaInfoJobListStreams
	Format     QueryMediaInfoJobListFormat
}

type QueryMediaInfoJobListStreams struct {
	VideoStreamList    QueryMediaInfoJobListVideoStreamList
	AudioStreamList    QueryMediaInfoJobListAudioStreamList
	SubtitleStreamList QueryMediaInfoJobListSubtitleStreamList
}

type QueryMediaInfoJobListVideoStream struct {
	Index          string
	CodecName      string
	CodecLongName  string
	Profile        string
	CodecTimeBase  string
	CodecTagString string
	CodecTag       string
	Width          string
	Height         string
	HasBFrames     string
	Sar            string
	Dar            string
	PixFmt         string
	Level          string
	Fps            string
	AvgFPS         string
	Timebase       string
	StartTime      string
	Duration       string
	Bitrate        string
	NumFrames      string
	Lang           string
	Rotate         string
	NetworkCost    QueryMediaInfoJobListNetworkCost
}

type QueryMediaInfoJobListNetworkCost struct {
	PreloadTime   string
	CostBandwidth string
	AvgBitrate    string
}

type QueryMediaInfoJobListAudioStream struct {
	Index          string
	CodecName      string
	CodecTimeBase  string
	CodecLongName  string
	CodecTagString string
	CodecTag       string
	SampleFmt      string
	Samplerate     string
	Channels       string
	ChannelLayout  string
	Timebase       string
	StartTime      string
	Duration       string
	Bitrate        string
	NumFrames      string
	Lang           string
}

type QueryMediaInfoJobListSubtitleStream struct {
	Index string
	Lang  string
}

type QueryMediaInfoJobListFormat struct {
	NumStreams     string
	NumPrograms    string
	FormatName     string
	FormatLongName string
	StartTime      string
	Duration       string
	Size           string
	Bitrate        string
}

type QueryMediaInfoJobListMNSMessageResult struct {
	MessageId    string
	ErrorMessage string
	ErrorCode    string
}
type QueryMediaInfoJobListArgs struct {
	ResourceOwnerId      int64
	MediaInfoJobIds      string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type QueryMediaInfoJobListResponse struct {
	RequestId               string
	MediaInfoJobList        QueryMediaInfoJobListMediaInfoJobList
	NonExistMediaInfoJobIds QueryMediaInfoJobListNonExistMediaInfoJobIdList
}

type QueryMediaInfoJobListVideoStreamList []QueryMediaInfoJobListVideoStream

func (list *QueryMediaInfoJobListVideoStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaInfoJobListVideoStream)
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

type QueryMediaInfoJobListAudioStreamList []QueryMediaInfoJobListAudioStream

func (list *QueryMediaInfoJobListAudioStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaInfoJobListAudioStream)
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

type QueryMediaInfoJobListSubtitleStreamList []QueryMediaInfoJobListSubtitleStream

func (list *QueryMediaInfoJobListSubtitleStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaInfoJobListSubtitleStream)
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

type QueryMediaInfoJobListMediaInfoJobList []QueryMediaInfoJobListMediaInfoJob

func (list *QueryMediaInfoJobListMediaInfoJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaInfoJobListMediaInfoJob)
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

type QueryMediaInfoJobListNonExistMediaInfoJobIdList []string

func (list *QueryMediaInfoJobListNonExistMediaInfoJobIdList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) QueryMediaListByURL(req *QueryMediaListByURLArgs) (resp *QueryMediaListByURLResponse, err error) {
	resp = &QueryMediaListByURLResponse{}
	err = c.Request("QueryMediaListByURL", req, resp)
	return
}

type QueryMediaListByURLMedia struct {
	MediaId      string
	Title        string
	Description  string
	CoverURL     string
	CateId       int64
	Duration     string
	Format       string
	Size         string
	Bitrate      string
	Width        string
	Height       string
	Fps          string
	PublishState string
	CreationTime string
	PlayList     QueryMediaListByURLPlayList
	SnapshotList QueryMediaListByURLSnapshotList
	Tags         QueryMediaListByURLTagList
	RunIdList    QueryMediaListByURLRunIdListList
	File         QueryMediaListByURLFile
	MediaInfo    QueryMediaListByURLMediaInfo
}

type QueryMediaListByURLPlay struct {
	ActivityName      string
	MediaWorkflowId   string
	MediaWorkflowName string
	Duration          string
	Format            string
	Size              string
	Bitrate           string
	Width             string
	Height            string
	Fps               string
	Encryption        string
	File1             QueryMediaListByURLFile1
}

type QueryMediaListByURLFile1 struct {
	URL   string
	State string
}

type QueryMediaListByURLSnapshot struct {
	Type              string
	MediaWorkflowId   string
	MediaWorkflowName string
	ActivityName      string
	Count             string
	File2             QueryMediaListByURLFile2
}

type QueryMediaListByURLFile2 struct {
	URL   string
	State string
}

type QueryMediaListByURLFile struct {
	URL   string
	State string
}

type QueryMediaListByURLMediaInfo struct {
	Streams QueryMediaListByURLStreams
	Format  QueryMediaListByURLFormat
}

type QueryMediaListByURLStreams struct {
	VideoStreamList    QueryMediaListByURLVideoStreamList
	AudioStreamList    QueryMediaListByURLAudioStreamList
	SubtitleStreamList QueryMediaListByURLSubtitleStreamList
}

type QueryMediaListByURLVideoStream struct {
	Index          string
	CodecName      string
	CodecLongName  string
	Profile        string
	CodecTimeBase  string
	CodecTagString string
	CodecTag       string
	Width          string
	Height         string
	HasBFrames     string
	Sar            string
	Dar            string
	PixFmt         string
	Level          string
	Fps            string
	AvgFPS         string
	Timebase       string
	StartTime      string
	Duration       string
	Bitrate        string
	NumFrames      string
	Lang           string
	Rotate         string
	NetworkCost    QueryMediaListByURLNetworkCost
}

type QueryMediaListByURLNetworkCost struct {
	PreloadTime   string
	CostBandwidth string
	AvgBitrate    string
}

type QueryMediaListByURLAudioStream struct {
	Index          string
	CodecName      string
	CodecTimeBase  string
	CodecLongName  string
	CodecTagString string
	CodecTag       string
	SampleFmt      string
	Samplerate     string
	Channels       string
	ChannelLayout  string
	Timebase       string
	StartTime      string
	Duration       string
	Bitrate        string
	NumFrames      string
	Lang           string
}

type QueryMediaListByURLSubtitleStream struct {
	Index string
	Lang  string
}

type QueryMediaListByURLFormat struct {
	NumStreams     string
	NumPrograms    string
	FormatName     string
	FormatLongName string
	StartTime      string
	Duration       string
	Size           string
	Bitrate        string
}
type QueryMediaListByURLArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	IncludeSnapshotList  core.Bool
	FileURLs             string
	OwnerAccount         string
	OwnerId              int64
	IncludePlayList      core.Bool
	IncludeMediaInfo     core.Bool
}
type QueryMediaListByURLResponse struct {
	RequestId        string
	MediaList        QueryMediaListByURLMediaList
	NonExistFileURLs QueryMediaListByURLNonExistFileURLList
}

type QueryMediaListByURLPlayList []QueryMediaListByURLPlay

func (list *QueryMediaListByURLPlayList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListByURLPlay)
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

type QueryMediaListByURLSnapshotList []QueryMediaListByURLSnapshot

func (list *QueryMediaListByURLSnapshotList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListByURLSnapshot)
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

type QueryMediaListByURLTagList []string

func (list *QueryMediaListByURLTagList) UnmarshalJSON(data []byte) error {
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

type QueryMediaListByURLRunIdListList []string

func (list *QueryMediaListByURLRunIdListList) UnmarshalJSON(data []byte) error {
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

type QueryMediaListByURLVideoStreamList []QueryMediaListByURLVideoStream

func (list *QueryMediaListByURLVideoStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListByURLVideoStream)
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

type QueryMediaListByURLAudioStreamList []QueryMediaListByURLAudioStream

func (list *QueryMediaListByURLAudioStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListByURLAudioStream)
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

type QueryMediaListByURLSubtitleStreamList []QueryMediaListByURLSubtitleStream

func (list *QueryMediaListByURLSubtitleStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListByURLSubtitleStream)
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

type QueryMediaListByURLMediaList []QueryMediaListByURLMedia

func (list *QueryMediaListByURLMediaList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListByURLMedia)
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

type QueryMediaListByURLNonExistFileURLList []string

func (list *QueryMediaListByURLNonExistFileURLList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) QueryCoverJobList(req *QueryCoverJobListArgs) (resp *QueryCoverJobListResponse, err error) {
	resp = &QueryCoverJobListResponse{}
	err = c.Request("QueryCoverJobList", req, resp)
	return
}

type QueryCoverJobListCoverJob struct {
	Id             string
	UserData       string
	PipelineId     string
	State          string
	Code           string
	Message        string
	CreationTime   string
	CoverImageList QueryCoverJobListCoverImageList
	Input          QueryCoverJobListInput
	CoverConfig    QueryCoverJobListCoverConfig
}

type QueryCoverJobListCoverImage struct {
	Score string
	Url   string
	Time  string
}

type QueryCoverJobListInput struct {
	Bucket   string
	Location string
	Object   string
}

type QueryCoverJobListCoverConfig struct {
	OutputFile QueryCoverJobListOutputFile
}

type QueryCoverJobListOutputFile struct {
	Bucket   string
	Location string
	Object   string
}
type QueryCoverJobListArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	CoverJobIds          string
}
type QueryCoverJobListResponse struct {
	RequestId    string
	CoverJobList QueryCoverJobListCoverJobList
	NonExistIds  QueryCoverJobListNonExistIdList
}

type QueryCoverJobListCoverImageList []QueryCoverJobListCoverImage

func (list *QueryCoverJobListCoverImageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCoverJobListCoverImage)
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

type QueryCoverJobListCoverJobList []QueryCoverJobListCoverJob

func (list *QueryCoverJobListCoverJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCoverJobListCoverJob)
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

type QueryCoverJobListNonExistIdList []string

func (list *QueryCoverJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) ListTerrorismPipeline(req *ListTerrorismPipelineArgs) (resp *ListTerrorismPipelineResponse, err error) {
	resp = &ListTerrorismPipelineResponse{}
	err = c.Request("ListTerrorismPipeline", req, resp)
	return
}

type ListTerrorismPipelinePipeline struct {
	Id           string
	Name         string
	State        string
	Priority     string
	NotifyConfig ListTerrorismPipelineNotifyConfig
}

type ListTerrorismPipelineNotifyConfig struct {
	Topic string
	Queue string
}
type ListTerrorismPipelineArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int64
	State                string
	OwnerId              int64
	PageNumber           int64
}
type ListTerrorismPipelineResponse struct {
	RequestId    string
	TotalCount   int64
	PageNumber   int64
	PageSize     int64
	PipelineList ListTerrorismPipelinePipelineList
}

type ListTerrorismPipelinePipelineList []ListTerrorismPipelinePipeline

func (list *ListTerrorismPipelinePipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTerrorismPipelinePipeline)
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

func (c *MtsClient) ReportTerrorismJobResult(req *ReportTerrorismJobResultArgs) (resp *ReportTerrorismJobResultResponse, err error) {
	resp = &ReportTerrorismJobResultResponse{}
	err = c.Request("ReportTerrorismJobResult", req, resp)
	return
}

type ReportTerrorismJobResultArgs struct {
	JobId                string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	Label                string
	Detail               string
	OwnerId              int64
}
type ReportTerrorismJobResultResponse struct {
	RequestId string
	JobId     string
}

func (c *MtsClient) BindInputBucket(req *BindInputBucketArgs) (resp *BindInputBucketResponse, err error) {
	resp = &BindInputBucketResponse{}
	err = c.Request("BindInputBucket", req, resp)
	return
}

type BindInputBucketArgs struct {
	Bucket               string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	RoleArn              string
	OwnerAccount         string
	OwnerId              int64
}
type BindInputBucketResponse struct {
	RequestId string
}

func (c *MtsClient) DeleteTemplate(req *DeleteTemplateArgs) (resp *DeleteTemplateResponse, err error) {
	resp = &DeleteTemplateResponse{}
	err = c.Request("DeleteTemplate", req, resp)
	return
}

type DeleteTemplateArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	TemplateId           string
}
type DeleteTemplateResponse struct {
	RequestId  string
	TemplateId string
}

func (c *MtsClient) QueryFacerecogJobList(req *QueryFacerecogJobListArgs) (resp *QueryFacerecogJobListResponse, err error) {
	resp = &QueryFacerecogJobListResponse{}
	err = c.Request("QueryFacerecogJobList", req, resp)
	return
}

type QueryFacerecogJobListFacerecogJob struct {
	Id                   string
	UserData             string
	PipelineId           string
	State                string
	Code                 string
	Message              string
	CreationTime         string
	Input                QueryFacerecogJobListInput
	VideoFacerecogResult QueryFacerecogJobListVideoFacerecogResult
}

type QueryFacerecogJobListInput struct {
	Bucket   string
	Location string
	Object   string
}

type QueryFacerecogJobListVideoFacerecogResult struct {
	Facerecogs QueryFacerecogJobListFacerecogList
}

type QueryFacerecogJobListFacerecog struct {
	Time  string
	Faces QueryFacerecogJobListFaceList
}

type QueryFacerecogJobListFace struct {
	Name   string
	Score  string
	Target string
}
type QueryFacerecogJobListArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	FacerecogJobIds      string
	OwnerId              int64
}
type QueryFacerecogJobListResponse struct {
	RequestId        string
	FacerecogJobList QueryFacerecogJobListFacerecogJobList
	NonExistIds      QueryFacerecogJobListNonExistIdList
}

type QueryFacerecogJobListFacerecogList []QueryFacerecogJobListFacerecog

func (list *QueryFacerecogJobListFacerecogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryFacerecogJobListFacerecog)
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

type QueryFacerecogJobListFaceList []QueryFacerecogJobListFace

func (list *QueryFacerecogJobListFaceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryFacerecogJobListFace)
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

type QueryFacerecogJobListFacerecogJobList []QueryFacerecogJobListFacerecogJob

func (list *QueryFacerecogJobListFacerecogJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryFacerecogJobListFacerecogJob)
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

type QueryFacerecogJobListNonExistIdList []string

func (list *QueryFacerecogJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) UpdateMediaCategory(req *UpdateMediaCategoryArgs) (resp *UpdateMediaCategoryResponse, err error) {
	resp = &UpdateMediaCategoryResponse{}
	err = c.Request("UpdateMediaCategory", req, resp)
	return
}

type UpdateMediaCategoryArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	CateId               int64
	OwnerAccount         string
	OwnerId              int64
	MediaId              string
}
type UpdateMediaCategoryResponse struct {
	RequestId string
}

func (c *MtsClient) SubmitJobs(req *SubmitJobsArgs) (resp *SubmitJobsResponse, err error) {
	resp = &SubmitJobsResponse{}
	err = c.Request("SubmitJobs", req, resp)
	return
}

type SubmitJobsJobResult struct {
	Success core.Bool
	Code    string
	Message string
	Job     SubmitJobsJob
}

type SubmitJobsJob struct {
	JobId            string
	State            string
	Code             string
	Message          string
	Percent          int64
	PipelineId       string
	CreationTime     string
	FinishTime       string
	Input            SubmitJobsInput
	Output           SubmitJobsOutput
	MNSMessageResult SubmitJobsMNSMessageResult
}

type SubmitJobsInput struct {
	Bucket   string
	Location string
	Object   string
}

type SubmitJobsOutput struct {
	TemplateId             string
	UserData               string
	Rotate                 string
	VideoStreamMap         string
	AudioStreamMap         string
	DeWatermark            string
	Priority               string
	WaterMarkConfigUrl     string
	MergeConfigUrl         string
	WaterMarkList          SubmitJobsWaterMarkList
	MergeList              SubmitJobsMergeList
	OpeningList            SubmitJobsOpeningList
	TailSlateList          SubmitJobsTailSlateList
	DigiWaterMark          SubmitJobsDigiWaterMark
	OutputFile             SubmitJobsOutputFile
	M3U8NonStandardSupport SubmitJobsM3U8NonStandardSupport
	Properties             SubmitJobsProperties
	Clip                   SubmitJobsClip
	SuperReso              SubmitJobsSuperReso
	SubtitleConfig         SubmitJobsSubtitleConfig
	TransConfig            SubmitJobsTransConfig
	MuxConfig              SubmitJobsMuxConfig
	Audio                  SubmitJobsAudio
	Video                  SubmitJobsVideo
	Container              SubmitJobsContainer
	Encryption             SubmitJobsEncryption
}

type SubmitJobsWaterMark struct {
	WaterMarkTemplateId string
	Width               string
	Height              string
	Dx                  string
	Dy                  string
	ReferPos            string
	Type                string
	InputFile           SubmitJobsInputFile
}

type SubmitJobsInputFile struct {
	Bucket   string
	Location string
	Object   string
}

type SubmitJobsMerge struct {
	MergeURL string
	Start    string
	Duration string
	RoleArn  string
}

type SubmitJobsOpening struct {
	OpenUrl string
	Start   string
	Width   string
	Height  string
}

type SubmitJobsTailSlate struct {
	TailUrl       string
	Start         string
	BlendDuration string
	Width         string
	Height        string
	IsMergeAudio  core.Bool
	BgColor       string
}

type SubmitJobsDigiWaterMark struct {
	Type       string
	Alpha      string
	InputFile1 SubmitJobsInputFile1
}

type SubmitJobsInputFile1 struct {
	Bucket   string
	Location string
	Object   string
}

type SubmitJobsOutputFile struct {
	Bucket   string
	Location string
	Object   string
	RoleArn  string
}

type SubmitJobsM3U8NonStandardSupport struct {
	TS SubmitJobsTS
}

type SubmitJobsTS struct {
	Md5Support  core.Bool
	SizeSupport core.Bool
}

type SubmitJobsProperties struct {
	Width      string
	Height     string
	Bitrate    string
	Duration   string
	Fps        string
	FileSize   string
	FileFormat string
	Streams    SubmitJobsStreams
	Format     SubmitJobsFormat
}

type SubmitJobsStreams struct {
	VideoStreamList    SubmitJobsVideoStreamList
	AudioStreamList    SubmitJobsAudioStreamList
	SubtitleStreamList SubmitJobsSubtitleStreamList
}

type SubmitJobsVideoStream struct {
	Index          string
	CodecName      string
	CodecLongName  string
	Profile        string
	CodecTimeBase  string
	CodecTagString string
	CodecTag       string
	Width          string
	Height         string
	HasBFrames     string
	Sar            string
	Dar            string
	PixFmt         string
	Level          string
	Fps            string
	AvgFPS         string
	Timebase       string
	StartTime      string
	Duration       string
	Bitrate        string
	NumFrames      string
	Lang           string
	NetworkCost    SubmitJobsNetworkCost
}

type SubmitJobsNetworkCost struct {
	PreloadTime   string
	CostBandwidth string
	AvgBitrate    string
}

type SubmitJobsAudioStream struct {
	Index          string
	CodecName      string
	CodecTimeBase  string
	CodecLongName  string
	CodecTagString string
	CodecTag       string
	SampleFmt      string
	Samplerate     string
	Channels       string
	ChannelLayout  string
	Timebase       string
	StartTime      string
	Duration       string
	Bitrate        string
	NumFrames      string
	Lang           string
}

type SubmitJobsSubtitleStream struct {
	Index string
	Lang  string
}

type SubmitJobsFormat struct {
	NumStreams     string
	NumPrograms    string
	FormatName     string
	FormatLongName string
	StartTime      string
	Duration       string
	Size           string
	Bitrate        string
}

type SubmitJobsClip struct {
	TimeSpan SubmitJobsTimeSpan
}

type SubmitJobsTimeSpan struct {
	Seek     string
	Duration string
}

type SubmitJobsSuperReso struct {
	IsHalfSample string
}

type SubmitJobsSubtitleConfig struct {
	SubtitleList    SubmitJobsSubtitleList
	ExtSubtitleList SubmitJobsExtSubtitleList
}

type SubmitJobsSubtitle struct {
	Map string
}

type SubmitJobsExtSubtitle struct {
	FontName string
	CharEnc  string
	Input2   SubmitJobsInput2
}

type SubmitJobsInput2 struct {
	Bucket   string
	Location string
	Object   string
}

type SubmitJobsTransConfig struct {
	TransMode               string
	IsCheckReso             string
	IsCheckResoFail         string
	IsCheckVideoBitrate     string
	IsCheckAudioBitrate     string
	AdjDarMethod            string
	IsCheckVideoBitrateFail string
	IsCheckAudioBitrateFail string
}

type SubmitJobsMuxConfig struct {
	Segment SubmitJobsSegment
	Gif     SubmitJobsGif
}

type SubmitJobsSegment struct {
	Duration string
}

type SubmitJobsGif struct {
	Loop            string
	FinalDelay      string
	IsCustomPalette string
	DitherMode      string
}

type SubmitJobsAudio struct {
	Codec      string
	Profile    string
	Samplerate string
	Bitrate    string
	Channels   string
	Qscale     string
	Volume     SubmitJobsVolume
}

type SubmitJobsVolume struct {
	Level  string
	Method string
}

type SubmitJobsVideo struct {
	Codec      string
	Profile    string
	Bitrate    string
	Crf        string
	Width      string
	Height     string
	Fps        string
	Gop        string
	Preset     string
	ScanMode   string
	Bufsize    string
	Maxrate    string
	PixFmt     string
	Degrain    string
	Qscale     string
	Crop       string
	Pad        string
	MaxFps     string
	BitrateBnd SubmitJobsBitrateBnd
}

type SubmitJobsBitrateBnd struct {
	Max string
	Min string
}

type SubmitJobsContainer struct {
	Format string
}

type SubmitJobsEncryption struct {
	Type    string
	Id      string
	Key     string
	KeyUri  string
	KeyType string
	SkipCnt string
}

type SubmitJobsMNSMessageResult struct {
	MessageId    string
	ErrorMessage string
	ErrorCode    string
}
type SubmitJobsArgs struct {
	Outputs              string
	Input                string
	OutputBucket         string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OutputLocation       string
	OwnerId              int64
	PipelineId           string
}
type SubmitJobsResponse struct {
	RequestId     string
	JobResultList SubmitJobsJobResultList
}

type SubmitJobsWaterMarkList []SubmitJobsWaterMark

func (list *SubmitJobsWaterMarkList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsWaterMark)
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

type SubmitJobsMergeList []SubmitJobsMerge

func (list *SubmitJobsMergeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsMerge)
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

type SubmitJobsOpeningList []SubmitJobsOpening

func (list *SubmitJobsOpeningList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsOpening)
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

type SubmitJobsTailSlateList []SubmitJobsTailSlate

func (list *SubmitJobsTailSlateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsTailSlate)
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

type SubmitJobsVideoStreamList []SubmitJobsVideoStream

func (list *SubmitJobsVideoStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsVideoStream)
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

type SubmitJobsAudioStreamList []SubmitJobsAudioStream

func (list *SubmitJobsAudioStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsAudioStream)
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

type SubmitJobsSubtitleStreamList []SubmitJobsSubtitleStream

func (list *SubmitJobsSubtitleStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsSubtitleStream)
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

type SubmitJobsSubtitleList []SubmitJobsSubtitle

func (list *SubmitJobsSubtitleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsSubtitle)
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

type SubmitJobsExtSubtitleList []SubmitJobsExtSubtitle

func (list *SubmitJobsExtSubtitleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsExtSubtitle)
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

type SubmitJobsJobResultList []SubmitJobsJobResult

func (list *SubmitJobsJobResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitJobsJobResult)
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

func (c *MtsClient) SubmitAnnotationJob(req *SubmitAnnotationJobArgs) (resp *SubmitAnnotationJobResponse, err error) {
	resp = &SubmitAnnotationJobResponse{}
	err = c.Request("SubmitAnnotationJob", req, resp)
	return
}

type SubmitAnnotationJobArgs struct {
	Input                string
	UserData             string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	AnnotationConfig     string
	OwnerId              int64
	PipelineId           string
}
type SubmitAnnotationJobResponse struct {
	RequestId string
	JobId     string
}

func (c *MtsClient) QueryCensorJobList(req *QueryCensorJobListArgs) (resp *QueryCensorJobListResponse, err error) {
	resp = &QueryCensorJobListResponse{}
	err = c.Request("QueryCensorJobList", req, resp)
	return
}

type QueryCensorJobListCensorJob struct {
	Id                    string
	UserData              string
	PipelineId            string
	State                 string
	Code                  string
	Message               string
	CreationTime          string
	Input                 QueryCensorJobListInput
	CensorConfig          QueryCensorJobListCensorConfig
	CensorPornResult      QueryCensorJobListCensorPornResult
	CensorTerrorismResult QueryCensorJobListCensorTerrorismResult
}

type QueryCensorJobListInput struct {
	Bucket   string
	Location string
	Object   string
}

type QueryCensorJobListCensorConfig struct {
	Interval   string
	BizType    string
	OutputFile QueryCensorJobListOutputFile
}

type QueryCensorJobListOutputFile struct {
	Bucket   string
	Location string
	Object   string
}

type QueryCensorJobListCensorPornResult struct {
	Label           string
	Suggestion      string
	MaxScore        string
	AverageScore    string
	PornCounterList QueryCensorJobListCounterList
	PornTopList     QueryCensorJobListTopList
}

type QueryCensorJobListCounter struct {
	Count int
	Label string
}

type QueryCensorJobListTop struct {
	Label     string
	Score     string
	Timestamp string
	Index     string
	Object    string
}

type QueryCensorJobListCensorTerrorismResult struct {
	Label                string
	Suggestion           string
	MaxScore             string
	AverageScore         string
	TerrorismCounterList QueryCensorJobListCounter1List
	TerrorismTopList     QueryCensorJobListTop2List
}

type QueryCensorJobListCounter1 struct {
	Count int
	Label string
}

type QueryCensorJobListTop2 struct {
	Label     string
	Score     string
	Timestamp string
	Index     string
	Object    string
}
type QueryCensorJobListArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	JobIds               string
	OwnerAccount         string
	OwnerId              int64
}
type QueryCensorJobListResponse struct {
	RequestId     string
	CensorJobList QueryCensorJobListCensorJobList
	NonExistIds   QueryCensorJobListNonExistIdList
}

type QueryCensorJobListCounterList []QueryCensorJobListCounter

func (list *QueryCensorJobListCounterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCensorJobListCounter)
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

type QueryCensorJobListTopList []QueryCensorJobListTop

func (list *QueryCensorJobListTopList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCensorJobListTop)
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

type QueryCensorJobListCounter1List []QueryCensorJobListCounter1

func (list *QueryCensorJobListCounter1List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCensorJobListCounter1)
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

type QueryCensorJobListTop2List []QueryCensorJobListTop2

func (list *QueryCensorJobListTop2List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCensorJobListTop2)
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

type QueryCensorJobListCensorJobList []QueryCensorJobListCensorJob

func (list *QueryCensorJobListCensorJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCensorJobListCensorJob)
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

type QueryCensorJobListNonExistIdList []string

func (list *QueryCensorJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) QueryAsrPipelineList(req *QueryAsrPipelineListArgs) (resp *QueryAsrPipelineListResponse, err error) {
	resp = &QueryAsrPipelineListResponse{}
	err = c.Request("QueryAsrPipelineList", req, resp)
	return
}

type QueryAsrPipelineListPipeline struct {
	Id           string
	Name         string
	State        string
	Priority     string
	NotifyConfig QueryAsrPipelineListNotifyConfig
}

type QueryAsrPipelineListNotifyConfig struct {
	Topic     string
	QueueName string
}
type QueryAsrPipelineListArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	PipelineIds          string
	OwnerAccount         string
	OwnerId              int64
}
type QueryAsrPipelineListResponse struct {
	RequestId    string
	PipelineList QueryAsrPipelineListPipelineList
	NonExistIds  QueryAsrPipelineListNonExistIdList
}

type QueryAsrPipelineListPipelineList []QueryAsrPipelineListPipeline

func (list *QueryAsrPipelineListPipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAsrPipelineListPipeline)
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

type QueryAsrPipelineListNonExistIdList []string

func (list *QueryAsrPipelineListNonExistIdList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) CategoryTree(req *CategoryTreeArgs) (resp *CategoryTreeResponse, err error) {
	resp = &CategoryTreeResponse{}
	err = c.Request("CategoryTree", req, resp)
	return
}

type CategoryTreeArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type CategoryTreeResponse struct {
	RequestId    string
	CategoryTree string
}

func (c *MtsClient) UpdateCategoryName(req *UpdateCategoryNameArgs) (resp *UpdateCategoryNameResponse, err error) {
	resp = &UpdateCategoryNameResponse{}
	err = c.Request("UpdateCategoryName", req, resp)
	return
}

type UpdateCategoryNameArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	CateId               string
	OwnerAccount         string
	OwnerId              int64
	CateName             string
}
type UpdateCategoryNameResponse struct {
	RequestId string
}

func (c *MtsClient) DeleteMedia(req *DeleteMediaArgs) (resp *DeleteMediaResponse, err error) {
	resp = &DeleteMediaResponse{}
	err = c.Request("DeleteMedia", req, resp)
	return
}

type DeleteMediaArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	MediaIds             string
	OwnerId              int64
}
type DeleteMediaResponse struct {
	RequestId string
}

func (c *MtsClient) AddWaterMarkTemplate(req *AddWaterMarkTemplateArgs) (resp *AddWaterMarkTemplateResponse, err error) {
	resp = &AddWaterMarkTemplateResponse{}
	err = c.Request("AddWaterMarkTemplate", req, resp)
	return
}

type AddWaterMarkTemplateWaterMarkTemplate struct {
	Id         string
	Name       string
	Width      string
	Height     string
	Dx         string
	Dy         string
	ReferPos   string
	Type       string
	State      string
	Timeline   AddWaterMarkTemplateTimeline
	RatioRefer AddWaterMarkTemplateRatioRefer
}

type AddWaterMarkTemplateTimeline struct {
	Start    string
	Duration string
}

type AddWaterMarkTemplateRatioRefer struct {
	Dx     string
	Dy     string
	Width  string
	Height string
}
type AddWaterMarkTemplateArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	Name                 string
	OwnerId              int64
	Config               string
}
type AddWaterMarkTemplateResponse struct {
	RequestId         string
	WaterMarkTemplate AddWaterMarkTemplateWaterMarkTemplate
}

func (c *MtsClient) CancelJob(req *CancelJobArgs) (resp *CancelJobResponse, err error) {
	resp = &CancelJobResponse{}
	err = c.Request("CancelJob", req, resp)
	return
}

type CancelJobArgs struct {
	JobId                string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type CancelJobResponse struct {
	RequestId string
	JobId     string
}

func (c *MtsClient) ActivateMediaWorkflow(req *ActivateMediaWorkflowArgs) (resp *ActivateMediaWorkflowResponse, err error) {
	resp = &ActivateMediaWorkflowResponse{}
	err = c.Request("ActivateMediaWorkflow", req, resp)
	return
}

type ActivateMediaWorkflowMediaWorkflow struct {
	MediaWorkflowId string
	Name            string
	Topology        string
	State           string
	CreationTime    string
}
type ActivateMediaWorkflowArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	MediaWorkflowId      string
	OwnerId              int64
}
type ActivateMediaWorkflowResponse struct {
	RequestId     string
	MediaWorkflow ActivateMediaWorkflowMediaWorkflow
}

func (c *MtsClient) QueryVideoSummaryPipelineList(req *QueryVideoSummaryPipelineListArgs) (resp *QueryVideoSummaryPipelineListResponse, err error) {
	resp = &QueryVideoSummaryPipelineListResponse{}
	err = c.Request("QueryVideoSummaryPipelineList", req, resp)
	return
}

type QueryVideoSummaryPipelineListPipeline struct {
	Id           string
	Name         string
	State        string
	Priority     string
	NotifyConfig QueryVideoSummaryPipelineListNotifyConfig
}

type QueryVideoSummaryPipelineListNotifyConfig struct {
	Topic     string
	QueueName string
}
type QueryVideoSummaryPipelineListArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	PipelineIds          string
	OwnerAccount         string
	OwnerId              int64
}
type QueryVideoSummaryPipelineListResponse struct {
	RequestId    string
	PipelineList QueryVideoSummaryPipelineListPipelineList
	NonExistIds  QueryVideoSummaryPipelineListNonExistIdList
}

type QueryVideoSummaryPipelineListPipelineList []QueryVideoSummaryPipelineListPipeline

func (list *QueryVideoSummaryPipelineListPipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryVideoSummaryPipelineListPipeline)
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

type QueryVideoSummaryPipelineListNonExistIdList []string

func (list *QueryVideoSummaryPipelineListNonExistIdList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) SubmitMediaInfoJob(req *SubmitMediaInfoJobArgs) (resp *SubmitMediaInfoJobResponse, err error) {
	resp = &SubmitMediaInfoJobResponse{}
	err = c.Request("SubmitMediaInfoJob", req, resp)
	return
}

type SubmitMediaInfoJobMediaInfoJob struct {
	JobId            string
	UserData         string
	PipelineId       string
	State            string
	Code             string
	Message          string
	CreationTime     string
	Input            SubmitMediaInfoJobInput
	Properties       SubmitMediaInfoJobProperties
	MNSMessageResult SubmitMediaInfoJobMNSMessageResult
}

type SubmitMediaInfoJobInput struct {
	Bucket   string
	Location string
	Object   string
}

type SubmitMediaInfoJobProperties struct {
	Width      string
	Height     string
	Bitrate    string
	Duration   string
	Fps        string
	FileSize   string
	FileFormat string
	Streams    SubmitMediaInfoJobStreams
	Format     SubmitMediaInfoJobFormat
}

type SubmitMediaInfoJobStreams struct {
	VideoStreamList    SubmitMediaInfoJobVideoStreamList
	AudioStreamList    SubmitMediaInfoJobAudioStreamList
	SubtitleStreamList SubmitMediaInfoJobSubtitleStreamList
}

type SubmitMediaInfoJobVideoStream struct {
	Index          string
	CodecName      string
	CodecLongName  string
	Profile        string
	CodecTimeBase  string
	CodecTagString string
	CodecTag       string
	Width          string
	Height         string
	HasBFrames     string
	Sar            string
	Dar            string
	PixFmt         string
	Level          string
	Fps            string
	AvgFPS         string
	Timebase       string
	StartTime      string
	Duration       string
	Bitrate        string
	NumFrames      string
	Lang           string
	Rotate         string
	NetworkCost    SubmitMediaInfoJobNetworkCost
}

type SubmitMediaInfoJobNetworkCost struct {
	PreloadTime   string
	CostBandwidth string
	AvgBitrate    string
}

type SubmitMediaInfoJobAudioStream struct {
	Index          string
	CodecName      string
	CodecTimeBase  string
	CodecLongName  string
	CodecTagString string
	CodecTag       string
	SampleFmt      string
	Samplerate     string
	Channels       string
	ChannelLayout  string
	Timebase       string
	StartTime      string
	Duration       string
	Bitrate        string
	NumFrames      string
	Lang           string
}

type SubmitMediaInfoJobSubtitleStream struct {
	Index string
	Lang  string
}

type SubmitMediaInfoJobFormat struct {
	NumStreams     string
	NumPrograms    string
	FormatName     string
	FormatLongName string
	StartTime      string
	Duration       string
	Size           string
	Bitrate        string
}

type SubmitMediaInfoJobMNSMessageResult struct {
	MessageId    string
	ErrorMessage string
	ErrorCode    string
}
type SubmitMediaInfoJobArgs struct {
	Input                string
	UserData             string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	PipelineId           string
}
type SubmitMediaInfoJobResponse struct {
	RequestId    string
	MediaInfoJob SubmitMediaInfoJobMediaInfoJob
}

type SubmitMediaInfoJobVideoStreamList []SubmitMediaInfoJobVideoStream

func (list *SubmitMediaInfoJobVideoStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitMediaInfoJobVideoStream)
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

type SubmitMediaInfoJobAudioStreamList []SubmitMediaInfoJobAudioStream

func (list *SubmitMediaInfoJobAudioStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitMediaInfoJobAudioStream)
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

type SubmitMediaInfoJobSubtitleStreamList []SubmitMediaInfoJobSubtitleStream

func (list *SubmitMediaInfoJobSubtitleStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitMediaInfoJobSubtitleStream)
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

func (c *MtsClient) ListAllMediaBucket(req *ListAllMediaBucketArgs) (resp *ListAllMediaBucketResponse, err error) {
	resp = &ListAllMediaBucketResponse{}
	err = c.Request("ListAllMediaBucket", req, resp)
	return
}

type ListAllMediaBucketMediaBucket struct {
	Bucket string
	Type   string
}
type ListAllMediaBucketArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type ListAllMediaBucketResponse struct {
	RequestId       string
	MediaBucketList ListAllMediaBucketMediaBucketList
}

type ListAllMediaBucketMediaBucketList []ListAllMediaBucketMediaBucket

func (list *ListAllMediaBucketMediaBucketList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAllMediaBucketMediaBucket)
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

func (c *MtsClient) ListVideoSummaryPipeline(req *ListVideoSummaryPipelineArgs) (resp *ListVideoSummaryPipelineResponse, err error) {
	resp = &ListVideoSummaryPipelineResponse{}
	err = c.Request("ListVideoSummaryPipeline", req, resp)
	return
}

type ListVideoSummaryPipelinePipeline struct {
	Id           string
	Name         string
	State        string
	Priority     string
	NotifyConfig ListVideoSummaryPipelineNotifyConfig
}

type ListVideoSummaryPipelineNotifyConfig struct {
	Topic     string
	QueueName string
}
type ListVideoSummaryPipelineArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int64
	State                string
	OwnerId              int64
	PageNumber           int64
}
type ListVideoSummaryPipelineResponse struct {
	RequestId    string
	TotalCount   int64
	PageNumber   int64
	PageSize     int64
	PipelineList ListVideoSummaryPipelinePipelineList
}

type ListVideoSummaryPipelinePipelineList []ListVideoSummaryPipelinePipeline

func (list *ListVideoSummaryPipelinePipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListVideoSummaryPipelinePipeline)
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

func (c *MtsClient) ReportCensorJobResult(req *ReportCensorJobResultArgs) (resp *ReportCensorJobResultResponse, err error) {
	resp = &ReportCensorJobResultResponse{}
	err = c.Request("ReportCensorJobResult", req, resp)
	return
}

type ReportCensorJobResultArgs struct {
	JobId                string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	Label                string
	Detail               string
	OwnerId              int64
}
type ReportCensorJobResultResponse struct {
	RequestId string
	JobId     string
}

func (c *MtsClient) AddPornPipeline(req *AddPornPipelineArgs) (resp *AddPornPipelineResponse, err error) {
	resp = &AddPornPipelineResponse{}
	err = c.Request("AddPornPipeline", req, resp)
	return
}

type AddPornPipelinePipeline struct {
	Id           string
	Name         string
	Priority     int
	State        string
	NotifyConfig AddPornPipelineNotifyConfig
}

type AddPornPipelineNotifyConfig struct {
	Topic string
	Queue string
}
type AddPornPipelineArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	Name                 string
	NotifyConfig         string
	OwnerId              int64
	Priority             int
}
type AddPornPipelineResponse struct {
	RequestId string
	Pipeline  AddPornPipelinePipeline
}

func (c *MtsClient) AddTerrorismPipeline(req *AddTerrorismPipelineArgs) (resp *AddTerrorismPipelineResponse, err error) {
	resp = &AddTerrorismPipelineResponse{}
	err = c.Request("AddTerrorismPipeline", req, resp)
	return
}

type AddTerrorismPipelinePipeline struct {
	Id           string
	Name         string
	Priority     int
	State        string
	NotifyConfig AddTerrorismPipelineNotifyConfig
}

type AddTerrorismPipelineNotifyConfig struct {
	Topic string
	Queue string
}
type AddTerrorismPipelineArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	Name                 string
	NotifyConfig         string
	OwnerId              int64
	Priority             int
}
type AddTerrorismPipelineResponse struct {
	RequestId string
	Pipeline  AddTerrorismPipelinePipeline
}

func (c *MtsClient) AddTemplate(req *AddTemplateArgs) (resp *AddTemplateResponse, err error) {
	resp = &AddTemplateResponse{}
	err = c.Request("AddTemplate", req, resp)
	return
}

type AddTemplateTemplate struct {
	Id          string
	Name        string
	State       string
	Container   AddTemplateContainer
	Video       AddTemplateVideo
	Audio       AddTemplateAudio
	TransConfig AddTemplateTransConfig
	MuxConfig   AddTemplateMuxConfig
}

type AddTemplateContainer struct {
	Format string
}

type AddTemplateVideo struct {
	Codec      string
	Profile    string
	Bitrate    string
	Crf        string
	Width      string
	Height     string
	Fps        string
	Gop        string
	Preset     string
	ScanMode   string
	Bufsize    string
	Maxrate    string
	PixFmt     string
	Degrain    string
	Qscale     string
	Remove     string
	Crop       string
	Pad        string
	MaxFps     string
	BitrateBnd AddTemplateBitrateBnd
}

type AddTemplateBitrateBnd struct {
	Max string
	Min string
}

type AddTemplateAudio struct {
	Codec      string
	Profile    string
	Samplerate string
	Bitrate    string
	Channels   string
	Qscale     string
	Remove     string
	Volume     AddTemplateVolume
}

type AddTemplateVolume struct {
	Level  string
	Method string
}

type AddTemplateTransConfig struct {
	TransMode               string
	IsCheckReso             string
	IsCheckResoFail         string
	IsCheckVideoBitrate     string
	IsCheckAudioBitrate     string
	AdjDarMethod            string
	IsCheckVideoBitrateFail string
	IsCheckAudioBitrateFail string
}

type AddTemplateMuxConfig struct {
	Segment AddTemplateSegment
	Gif     AddTemplateGif
}

type AddTemplateSegment struct {
	Duration string
}

type AddTemplateGif struct {
	Loop            string
	FinalDelay      string
	IsCustomPalette string
	DitherMode      string
}
type AddTemplateArgs struct {
	Container            string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	Name                 string
	TransConfig          string
	MuxConfig            string
	Video                string
	Audio                string
	OwnerId              int64
}
type AddTemplateResponse struct {
	RequestId string
	Template  AddTemplateTemplate
}

func (c *MtsClient) UpdateTemplate(req *UpdateTemplateArgs) (resp *UpdateTemplateResponse, err error) {
	resp = &UpdateTemplateResponse{}
	err = c.Request("UpdateTemplate", req, resp)
	return
}

type UpdateTemplateTemplate struct {
	Id          string
	Name        string
	State       string
	Container   UpdateTemplateContainer
	Video       UpdateTemplateVideo
	Audio       UpdateTemplateAudio
	TransConfig UpdateTemplateTransConfig
	MuxConfig   UpdateTemplateMuxConfig
}

type UpdateTemplateContainer struct {
	Format string
}

type UpdateTemplateVideo struct {
	Codec      string
	Profile    string
	Bitrate    string
	Crf        string
	Width      string
	Height     string
	Fps        string
	Gop        string
	Preset     string
	ScanMode   string
	Bufsize    string
	Maxrate    string
	PixFmt     string
	Degrain    string
	Qscale     string
	Remove     string
	Crop       string
	Pad        string
	MaxFps     string
	BitrateBnd UpdateTemplateBitrateBnd
}

type UpdateTemplateBitrateBnd struct {
	Max string
	Min string
}

type UpdateTemplateAudio struct {
	Codec      string
	Profile    string
	Samplerate string
	Bitrate    string
	Channels   string
	Qscale     string
	Remove     string
}

type UpdateTemplateTransConfig struct {
	TransMode               string
	IsCheckReso             string
	IsCheckResoFail         string
	IsCheckVideoBitrate     string
	IsCheckAudioBitrate     string
	AdjDarMethod            string
	IsCheckVideoBitrateFail string
	IsCheckAudioBitrateFail string
}

type UpdateTemplateMuxConfig struct {
	Segment UpdateTemplateSegment
	Gif     UpdateTemplateGif
}

type UpdateTemplateSegment struct {
	Duration string
}

type UpdateTemplateGif struct {
	Loop            string
	FinalDelay      string
	IsCustomPalette string
	DitherMode      string
}
type UpdateTemplateArgs struct {
	Container            string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	MuxConfig            string
	Video                string
	OwnerId              int64
	TemplateId           string
	Name                 string
	TransConfig          string
	Audio                string
}
type UpdateTemplateResponse struct {
	RequestId string
	Template  UpdateTemplateTemplate
}

func (c *MtsClient) QueryEditingJobList(req *QueryEditingJobListArgs) (resp *QueryEditingJobListResponse, err error) {
	resp = &QueryEditingJobListResponse{}
	err = c.Request("QueryEditingJobList", req, resp)
	return
}

type QueryEditingJobListJob struct {
	JobId            string
	State            string
	Code             string
	Message          string
	Percent          int64
	PipelineId       string
	CreationTime     string
	FinishTime       string
	EditingInputs    QueryEditingJobListEditingInputList
	EditingConfig    QueryEditingJobListEditingConfig
	MNSMessageResult QueryEditingJobListMNSMessageResult
}

type QueryEditingJobListEditingInput struct {
	Id          string
	InputFile   QueryEditingJobListInputFile
	InputConfig QueryEditingJobListInputConfig
}

type QueryEditingJobListInputFile struct {
	Bucket   string
	Location string
	Object   string
}

type QueryEditingJobListInputConfig struct {
	DeinterlaceMethod string
	IsNormalSar       string
}

type QueryEditingJobListEditingConfig struct {
	TemplateId             string
	UserData               string
	Rotate                 string
	VideoStreamMap         string
	AudioStreamMap         string
	DeWatermark            string
	Priority               string
	WaterMarkConfigUrl     string
	MergeConfigUrl         string
	WaterMarkList          QueryEditingJobListWaterMarkList
	MergeList              QueryEditingJobListMergeList
	DigiWaterMark          QueryEditingJobListDigiWaterMark
	OutputFile             QueryEditingJobListOutputFile
	M3U8NonStandardSupport QueryEditingJobListM3U8NonStandardSupport
	Properties             QueryEditingJobListProperties
	Clip                   QueryEditingJobListClip
	SuperReso              QueryEditingJobListSuperReso
	SubtitleConfig         QueryEditingJobListSubtitleConfig
	TransConfig            QueryEditingJobListTransConfig
	MuxConfig              QueryEditingJobListMuxConfig
	Audio                  QueryEditingJobListAudio
	Video                  QueryEditingJobListVideo
	Container              QueryEditingJobListContainer
	Encryption             QueryEditingJobListEncryption
	Editing                QueryEditingJobListEditing
}

type QueryEditingJobListWaterMark struct {
	WaterMarkTemplateId string
	Width               string
	Height              string
	Dx                  string
	Dy                  string
	ReferPos            string
	Type                string
	InputFile1          QueryEditingJobListInputFile1
}

type QueryEditingJobListInputFile1 struct {
	Bucket   string
	Location string
	Object   string
}

type QueryEditingJobListMerge struct {
	MergeURL string
	Start    string
	Duration string
	RoleArn  string
}

type QueryEditingJobListDigiWaterMark struct {
	Type       string
	Alpha      string
	InputFile2 QueryEditingJobListInputFile2
}

type QueryEditingJobListInputFile2 struct {
	Bucket   string
	Location string
	Object   string
}

type QueryEditingJobListOutputFile struct {
	Bucket   string
	Location string
	Object   string
	RoleArn  string
}

type QueryEditingJobListM3U8NonStandardSupport struct {
	TS QueryEditingJobListTS
}

type QueryEditingJobListTS struct {
	Md5Support  core.Bool
	SizeSupport core.Bool
}

type QueryEditingJobListProperties struct {
	Width      string
	Height     string
	Bitrate    string
	Duration   string
	Fps        string
	FileSize   string
	FileFormat string
	Streams    QueryEditingJobListStreams
	Format     QueryEditingJobListFormat
}

type QueryEditingJobListStreams struct {
	VideoStreamList    QueryEditingJobListVideoStreamList
	AudioStreamList    QueryEditingJobListAudioStreamList
	SubtitleStreamList QueryEditingJobListSubtitleStreamList
}

type QueryEditingJobListVideoStream struct {
	Index          string
	CodecName      string
	CodecLongName  string
	Profile        string
	CodecTimeBase  string
	CodecTagString string
	CodecTag       string
	Width          string
	Height         string
	HasBFrames     string
	Sar            string
	Dar            string
	PixFmt         string
	Level          string
	Fps            string
	AvgFPS         string
	Timebase       string
	StartTime      string
	Duration       string
	Bitrate        string
	NumFrames      string
	Lang           string
	NetworkCost    QueryEditingJobListNetworkCost
}

type QueryEditingJobListNetworkCost struct {
	PreloadTime   string
	CostBandwidth string
	AvgBitrate    string
}

type QueryEditingJobListAudioStream struct {
	Index          string
	CodecName      string
	CodecTimeBase  string
	CodecLongName  string
	CodecTagString string
	CodecTag       string
	SampleFmt      string
	Samplerate     string
	Channels       string
	ChannelLayout  string
	Timebase       string
	StartTime      string
	Duration       string
	Bitrate        string
	NumFrames      string
	Lang           string
}

type QueryEditingJobListSubtitleStream struct {
	Index string
	Lang  string
}

type QueryEditingJobListFormat struct {
	NumStreams     string
	NumPrograms    string
	FormatName     string
	FormatLongName string
	StartTime      string
	Duration       string
	Size           string
	Bitrate        string
}

type QueryEditingJobListClip struct {
	TimeSpan QueryEditingJobListTimeSpan
}

type QueryEditingJobListTimeSpan struct {
	Seek     string
	Duration string
}

type QueryEditingJobListSuperReso struct {
	IsHalfSample string
}

type QueryEditingJobListSubtitleConfig struct {
	SubtitleList    QueryEditingJobListSubtitleList
	ExtSubtitleList QueryEditingJobListExtSubtitleList
}

type QueryEditingJobListSubtitle struct {
	Map string
}

type QueryEditingJobListExtSubtitle struct {
	FontName string
	CharEnc  string
	Input    QueryEditingJobListInput
}

type QueryEditingJobListInput struct {
	Bucket   string
	Location string
	Object   string
}

type QueryEditingJobListTransConfig struct {
	TransMode               string
	IsCheckReso             string
	IsCheckResoFail         string
	IsCheckVideoBitrate     string
	IsCheckAudioBitrate     string
	AdjDarMethod            string
	IsCheckVideoBitrateFail string
	IsCheckAudioBitrateFail string
	Duration                string
}

type QueryEditingJobListMuxConfig struct {
	Segment QueryEditingJobListSegment
	Gif     QueryEditingJobListGif
}

type QueryEditingJobListSegment struct {
	Duration string
}

type QueryEditingJobListGif struct {
	Loop            string
	FinalDelay      string
	IsCustomPalette string
	DitherMode      string
}

type QueryEditingJobListAudio struct {
	Codec      string
	Profile    string
	Samplerate string
	Bitrate    string
	Channels   string
	Qscale     string
	Volume     QueryEditingJobListVolume
}

type QueryEditingJobListVolume struct {
	Level  string
	Method string
}

type QueryEditingJobListVideo struct {
	Codec      string
	Profile    string
	Bitrate    string
	Crf        string
	Width      string
	Height     string
	Fps        string
	Gop        string
	Preset     string
	ScanMode   string
	Bufsize    string
	Maxrate    string
	PixFmt     string
	Degrain    string
	Qscale     string
	Crop       string
	Pad        string
	MaxFps     string
	BitrateBnd QueryEditingJobListBitrateBnd
}

type QueryEditingJobListBitrateBnd struct {
	Max string
	Min string
}

type QueryEditingJobListContainer struct {
	Format string
}

type QueryEditingJobListEncryption struct {
	Type    string
	Id      string
	Key     string
	KeyUri  string
	KeyType string
	SkipCnt string
}

type QueryEditingJobListEditing struct {
	ClipList QueryEditingJobListClip3List
	Timeline QueryEditingJobListTimeline
}

type QueryEditingJobListClip3 struct {
	Id            string
	Type          string
	SourceType    string
	SourceID      string
	SourceStrmMap string
	In            string
	Out           string
	Effects       QueryEditingJobListEffectList
}

type QueryEditingJobListEffect struct {
	Effect       string
	EffectConfig string
}

type QueryEditingJobListTimeline struct {
	TrackList      QueryEditingJobListTrackList
	TimelineConfig QueryEditingJobListTimelineConfig
}

type QueryEditingJobListTrack struct {
	Id    string
	Type  string
	Order string
	Clips QueryEditingJobListClip4List
}

type QueryEditingJobListClip4 struct {
	ClipID      string
	In          string
	Out         string
	ClipsConfig QueryEditingJobListClipsConfig
}

type QueryEditingJobListClipsConfig struct {
	ClipsConfigVideo QueryEditingJobListClipsConfigVideo
}

type QueryEditingJobListClipsConfigVideo struct {
	L string
	T string
}

type QueryEditingJobListTimelineConfig struct {
	TimelineConfigVideo QueryEditingJobListTimelineConfigVideo
	TimelineConfigAudio QueryEditingJobListTimelineConfigAudio
}

type QueryEditingJobListTimelineConfigVideo struct {
	Width   string
	Height  string
	BgColor string
	Fps     string
}

type QueryEditingJobListTimelineConfigAudio struct {
	Samplerate    string
	ChannelLayout string
	Channels      string
}

type QueryEditingJobListMNSMessageResult struct {
	MessageId    string
	ErrorMessage string
	ErrorCode    string
}
type QueryEditingJobListArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	JobIds               string
	OwnerAccount         string
	OwnerId              int64
}
type QueryEditingJobListResponse struct {
	RequestId      string
	JobList        QueryEditingJobListJobList
	NonExistJobIds QueryEditingJobListNonExistJobIdList
}

type QueryEditingJobListEditingInputList []QueryEditingJobListEditingInput

func (list *QueryEditingJobListEditingInputList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEditingJobListEditingInput)
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

type QueryEditingJobListWaterMarkList []QueryEditingJobListWaterMark

func (list *QueryEditingJobListWaterMarkList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEditingJobListWaterMark)
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

type QueryEditingJobListMergeList []QueryEditingJobListMerge

func (list *QueryEditingJobListMergeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEditingJobListMerge)
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

type QueryEditingJobListVideoStreamList []QueryEditingJobListVideoStream

func (list *QueryEditingJobListVideoStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEditingJobListVideoStream)
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

type QueryEditingJobListAudioStreamList []QueryEditingJobListAudioStream

func (list *QueryEditingJobListAudioStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEditingJobListAudioStream)
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

type QueryEditingJobListSubtitleStreamList []QueryEditingJobListSubtitleStream

func (list *QueryEditingJobListSubtitleStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEditingJobListSubtitleStream)
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

type QueryEditingJobListSubtitleList []QueryEditingJobListSubtitle

func (list *QueryEditingJobListSubtitleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEditingJobListSubtitle)
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

type QueryEditingJobListExtSubtitleList []QueryEditingJobListExtSubtitle

func (list *QueryEditingJobListExtSubtitleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEditingJobListExtSubtitle)
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

type QueryEditingJobListClip3List []QueryEditingJobListClip3

func (list *QueryEditingJobListClip3List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEditingJobListClip3)
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

type QueryEditingJobListEffectList []QueryEditingJobListEffect

func (list *QueryEditingJobListEffectList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEditingJobListEffect)
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

type QueryEditingJobListTrackList []QueryEditingJobListTrack

func (list *QueryEditingJobListTrackList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEditingJobListTrack)
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

type QueryEditingJobListClip4List []QueryEditingJobListClip4

func (list *QueryEditingJobListClip4List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEditingJobListClip4)
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

type QueryEditingJobListJobList []QueryEditingJobListJob

func (list *QueryEditingJobListJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryEditingJobListJob)
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

type QueryEditingJobListNonExistJobIdList []string

func (list *QueryEditingJobListNonExistJobIdList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) QueryCoverPipelineList(req *QueryCoverPipelineListArgs) (resp *QueryCoverPipelineListResponse, err error) {
	resp = &QueryCoverPipelineListResponse{}
	err = c.Request("QueryCoverPipelineList", req, resp)
	return
}

type QueryCoverPipelineListPipeline struct {
	Id           string
	Name         string
	State        string
	Priority     string
	Role         string
	NotifyConfig QueryCoverPipelineListNotifyConfig
}

type QueryCoverPipelineListNotifyConfig struct {
	Topic string
	Queue string
}
type QueryCoverPipelineListArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	PipelineIds          string
	OwnerAccount         string
	OwnerId              int64
}
type QueryCoverPipelineListResponse struct {
	RequestId    string
	PipelineList QueryCoverPipelineListPipelineList
	NonExistIds  QueryCoverPipelineListNonExistIdList
}

type QueryCoverPipelineListPipelineList []QueryCoverPipelineListPipeline

func (list *QueryCoverPipelineListPipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCoverPipelineListPipeline)
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

type QueryCoverPipelineListNonExistIdList []string

func (list *QueryCoverPipelineListNonExistIdList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) BindOutputBucket(req *BindOutputBucketArgs) (resp *BindOutputBucketResponse, err error) {
	resp = &BindOutputBucketResponse{}
	err = c.Request("BindOutputBucket", req, resp)
	return
}

type BindOutputBucketArgs struct {
	Bucket               string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	RoleArn              string
	OwnerAccount         string
	OwnerId              int64
}
type BindOutputBucketResponse struct {
	RequestId string
}

func (c *MtsClient) UpdateMediaCover(req *UpdateMediaCoverArgs) (resp *UpdateMediaCoverResponse, err error) {
	resp = &UpdateMediaCoverResponse{}
	err = c.Request("UpdateMediaCover", req, resp)
	return
}

type UpdateMediaCoverArgs struct {
	CoverURL             string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	MediaId              string
}
type UpdateMediaCoverResponse struct {
	RequestId string
}

func (c *MtsClient) ListCensorPipeline(req *ListCensorPipelineArgs) (resp *ListCensorPipelineResponse, err error) {
	resp = &ListCensorPipelineResponse{}
	err = c.Request("ListCensorPipeline", req, resp)
	return
}

type ListCensorPipelinePipeline struct {
	Id           string
	Name         string
	State        string
	Priority     string
	NotifyConfig ListCensorPipelineNotifyConfig
}

type ListCensorPipelineNotifyConfig struct {
	Topic string
	Queue string
}
type ListCensorPipelineArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int64
	State                string
	OwnerId              int64
	PageNumber           int64
}
type ListCensorPipelineResponse struct {
	RequestId    string
	TotalCount   int64
	PageNumber   int64
	PageSize     int64
	PipelineList ListCensorPipelinePipelineList
}

type ListCensorPipelinePipelineList []ListCensorPipelinePipeline

func (list *ListCensorPipelinePipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListCensorPipelinePipeline)
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

func (c *MtsClient) QueryPornPipelineList(req *QueryPornPipelineListArgs) (resp *QueryPornPipelineListResponse, err error) {
	resp = &QueryPornPipelineListResponse{}
	err = c.Request("QueryPornPipelineList", req, resp)
	return
}

type QueryPornPipelineListPipeline struct {
	Id           string
	Name         string
	State        string
	Priority     string
	NotifyConfig QueryPornPipelineListNotifyConfig
}

type QueryPornPipelineListNotifyConfig struct {
	Topic string
	Queue string
}
type QueryPornPipelineListArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	PipelineIds          string
	OwnerAccount         string
	OwnerId              int64
}
type QueryPornPipelineListResponse struct {
	RequestId    string
	PipelineList QueryPornPipelineListPipelineList
	NonExistIds  QueryPornPipelineListNonExistIdList
}

type QueryPornPipelineListPipelineList []QueryPornPipelineListPipeline

func (list *QueryPornPipelineListPipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPornPipelineListPipeline)
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

type QueryPornPipelineListNonExistIdList []string

func (list *QueryPornPipelineListNonExistIdList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) AddMediaWorkflow(req *AddMediaWorkflowArgs) (resp *AddMediaWorkflowResponse, err error) {
	resp = &AddMediaWorkflowResponse{}
	err = c.Request("AddMediaWorkflow", req, resp)
	return
}

type AddMediaWorkflowMediaWorkflow struct {
	MediaWorkflowId string
	Name            string
	Topology        string
	State           string
	CreationTime    string
}
type AddMediaWorkflowArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	Topology             string
	OwnerAccount         string
	Name                 string
	OwnerId              int64
}
type AddMediaWorkflowResponse struct {
	RequestId     string
	MediaWorkflow AddMediaWorkflowMediaWorkflow
}

func (c *MtsClient) DeleteMediaWorkflow(req *DeleteMediaWorkflowArgs) (resp *DeleteMediaWorkflowResponse, err error) {
	resp = &DeleteMediaWorkflowResponse{}
	err = c.Request("DeleteMediaWorkflow", req, resp)
	return
}

type DeleteMediaWorkflowMediaWorkflow struct {
	MediaWorkflowId string
	Name            string
	Topology        string
	State           string
	CreationTime    string
}
type DeleteMediaWorkflowArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	MediaWorkflowId      string
	OwnerId              int64
}
type DeleteMediaWorkflowResponse struct {
	RequestId     string
	MediaWorkflow DeleteMediaWorkflowMediaWorkflow
}

func (c *MtsClient) SubmitAsrJob(req *SubmitAsrJobArgs) (resp *SubmitAsrJobResponse, err error) {
	resp = &SubmitAsrJobResponse{}
	err = c.Request("SubmitAsrJob", req, resp)
	return
}

type SubmitAsrJobArgs struct {
	Input                string
	UserData             string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	AsrConfig            string
	OwnerId              int64
	PipelineId           string
}
type SubmitAsrJobResponse struct {
	RequestId string
	JobId     string
}

func (c *MtsClient) PlayerAuth(req *PlayerAuthArgs) (resp *PlayerAuthResponse, err error) {
	resp = &PlayerAuthResponse{}
	err = c.Request("PlayerAuth", req, resp)
	return
}

type PlayerAuth_Switch struct {
	State        string
	FunctionId   string
	SwitchId     string
	FunctionName string
}
type PlayerAuthArgs struct {
	ResourceOwnerId      string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              string
}
type PlayerAuthResponse struct {
	RequestId  string
	LogURL     string
	SwitchList PlayerAuth_SwitchList
}

type PlayerAuth_SwitchList []PlayerAuth_Switch

func (list *PlayerAuth_SwitchList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]PlayerAuth_Switch)
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

func (c *MtsClient) AddVideoSummaryPipeline(req *AddVideoSummaryPipelineArgs) (resp *AddVideoSummaryPipelineResponse, err error) {
	resp = &AddVideoSummaryPipelineResponse{}
	err = c.Request("AddVideoSummaryPipeline", req, resp)
	return
}

type AddVideoSummaryPipelinePipeline struct {
	Id           string
	Name         string
	Priority     int
	State        string
	NotifyConfig AddVideoSummaryPipelineNotifyConfig
}

type AddVideoSummaryPipelineNotifyConfig struct {
	Topic string
	Queue string
}
type AddVideoSummaryPipelineArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	Name                 string
	NotifyConfig         string
	OwnerId              int64
	Priority             int
}
type AddVideoSummaryPipelineResponse struct {
	RequestId string
	Pipeline  AddVideoSummaryPipelinePipeline
}

func (c *MtsClient) SubmitCoverJob(req *SubmitCoverJobArgs) (resp *SubmitCoverJobResponse, err error) {
	resp = &SubmitCoverJobResponse{}
	err = c.Request("SubmitCoverJob", req, resp)
	return
}

type SubmitCoverJobArgs struct {
	Input                string
	UserData             string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	CoverConfig          string
	OwnerId              int64
	PipelineId           string
}
type SubmitCoverJobResponse struct {
	RequestId string
	JobId     string
}

func (c *MtsClient) UpdateMediaWorkflow(req *UpdateMediaWorkflowArgs) (resp *UpdateMediaWorkflowResponse, err error) {
	resp = &UpdateMediaWorkflowResponse{}
	err = c.Request("UpdateMediaWorkflow", req, resp)
	return
}

type UpdateMediaWorkflowMediaWorkflow struct {
	MediaWorkflowId string
	Name            string
	Topology        string
	State           string
	CreationTime    string
}
type UpdateMediaWorkflowArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	Topology             string
	OwnerAccount         string
	MediaWorkflowId      string
	OwnerId              int64
}
type UpdateMediaWorkflowResponse struct {
	RequestId     string
	MediaWorkflow UpdateMediaWorkflowMediaWorkflow
}

func (c *MtsClient) ReportCoverJobResult(req *ReportCoverJobResultArgs) (resp *ReportCoverJobResultResponse, err error) {
	resp = &ReportCoverJobResultResponse{}
	err = c.Request("ReportCoverJobResult", req, resp)
	return
}

type ReportCoverJobResultArgs struct {
	Result               string
	JobId                string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type ReportCoverJobResultResponse struct {
	RequestId string
	JobId     string
}

func (c *MtsClient) AddCensorPipeline(req *AddCensorPipelineArgs) (resp *AddCensorPipelineResponse, err error) {
	resp = &AddCensorPipelineResponse{}
	err = c.Request("AddCensorPipeline", req, resp)
	return
}

type AddCensorPipelinePipeline struct {
	Id           string
	Name         string
	Priority     int
	State        string
	NotifyConfig AddCensorPipelineNotifyConfig
}

type AddCensorPipelineNotifyConfig struct {
	Topic string
	Queue string
}
type AddCensorPipelineArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	Name                 string
	NotifyConfig         string
	OwnerId              int64
	Priority             int
}
type AddCensorPipelineResponse struct {
	RequestId string
	Pipeline  AddCensorPipelinePipeline
}

func (c *MtsClient) ReportAnnotationJobResult(req *ReportAnnotationJobResultArgs) (resp *ReportAnnotationJobResultResponse, err error) {
	resp = &ReportAnnotationJobResultResponse{}
	err = c.Request("ReportAnnotationJobResult", req, resp)
	return
}

type ReportAnnotationJobResultArgs struct {
	Annotation           string
	JobId                string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	Details              string
	OwnerId              int64
}
type ReportAnnotationJobResultResponse struct {
	RequestId string
	JobId     string
}

func (c *MtsClient) QueryTagJobList(req *QueryTagJobListArgs) (resp *QueryTagJobListResponse, err error) {
	resp = &QueryTagJobListResponse{}
	err = c.Request("QueryTagJobList", req, resp)
	return
}

type QueryTagJobListTagJob struct {
	Id             string
	UserData       string
	PipelineId     string
	State          string
	Code           string
	Message        string
	CreationTime   string
	Input          QueryTagJobListInput
	VideoTagResult QueryTagJobListVideoTagResult
}

type QueryTagJobListInput struct {
	Bucket   string
	Location string
	Object   string
}

type QueryTagJobListVideoTagResult struct {
	Details      string
	TagAnResults QueryTagJobListTagAnResultList
	TagFrResults QueryTagJobListTagFrResultList
}

type QueryTagJobListTagAnResult struct {
	Label string
	Score string
}

type QueryTagJobListTagFrResult struct {
	Time     string
	TagFaces QueryTagJobListTagFaceList
}

type QueryTagJobListTagFace struct {
	Name   string
	Score  string
	Target string
}
type QueryTagJobListArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	TagJobIds            string
	OwnerAccount         string
	OwnerId              int64
}
type QueryTagJobListResponse struct {
	RequestId   string
	TagJobList  QueryTagJobListTagJobList
	NonExistIds QueryTagJobListNonExistIdList
}

type QueryTagJobListTagAnResultList []QueryTagJobListTagAnResult

func (list *QueryTagJobListTagAnResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTagJobListTagAnResult)
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

type QueryTagJobListTagFrResultList []QueryTagJobListTagFrResult

func (list *QueryTagJobListTagFrResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTagJobListTagFrResult)
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

type QueryTagJobListTagFaceList []QueryTagJobListTagFace

func (list *QueryTagJobListTagFaceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTagJobListTagFace)
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

type QueryTagJobListTagJobList []QueryTagJobListTagJob

func (list *QueryTagJobListTagJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTagJobListTagJob)
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

type QueryTagJobListNonExistIdList []string

func (list *QueryTagJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) SearchTemplate(req *SearchTemplateArgs) (resp *SearchTemplateResponse, err error) {
	resp = &SearchTemplateResponse{}
	err = c.Request("SearchTemplate", req, resp)
	return
}

type SearchTemplateTemplate struct {
	Id          string
	Name        string
	State       string
	Container   SearchTemplateContainer
	Video       SearchTemplateVideo
	Audio       SearchTemplateAudio
	TransConfig SearchTemplateTransConfig
	MuxConfig   SearchTemplateMuxConfig
}

type SearchTemplateContainer struct {
	Format string
}

type SearchTemplateVideo struct {
	Codec      string
	Profile    string
	Bitrate    string
	Crf        string
	Width      string
	Height     string
	Fps        string
	Gop        string
	Preset     string
	ScanMode   string
	Bufsize    string
	Maxrate    string
	PixFmt     string
	Degrain    string
	Qscale     string
	Remove     string
	Crop       string
	Pad        string
	MaxFps     string
	BitrateBnd SearchTemplateBitrateBnd
}

type SearchTemplateBitrateBnd struct {
	Max string
	Min string
}

type SearchTemplateAudio struct {
	Codec      string
	Profile    string
	Samplerate string
	Bitrate    string
	Channels   string
	Qscale     string
	Remove     string
}

type SearchTemplateTransConfig struct {
	TransMode               string
	IsCheckReso             string
	IsCheckResoFail         string
	IsCheckVideoBitrate     string
	IsCheckAudioBitrate     string
	AdjDarMethod            string
	IsCheckVideoBitrateFail string
	IsCheckAudioBitrateFail string
}

type SearchTemplateMuxConfig struct {
	Segment SearchTemplateSegment
	Gif     SearchTemplateGif
}

type SearchTemplateSegment struct {
	Duration string
}

type SearchTemplateGif struct {
	Loop            string
	FinalDelay      string
	IsCustomPalette string
	DitherMode      string
}
type SearchTemplateArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int64
	State                string
	OwnerId              int64
	PageNumber           int64
}
type SearchTemplateResponse struct {
	RequestId    string
	TotalCount   int64
	PageNumber   int64
	PageSize     int64
	TemplateList SearchTemplateTemplateList
}

type SearchTemplateTemplateList []SearchTemplateTemplate

func (list *SearchTemplateTemplateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SearchTemplateTemplate)
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

func (c *MtsClient) SearchMedia(req *SearchMediaArgs) (resp *SearchMediaResponse, err error) {
	resp = &SearchMediaResponse{}
	err = c.Request("SearchMedia", req, resp)
	return
}

type SearchMediaMedia struct {
	MediaId      string
	Title        string
	Description  string
	CoverURL     string
	CateId       int64
	Duration     string
	Format       string
	Size         string
	Bitrate      string
	Width        string
	Height       string
	Fps          string
	PublishState string
	CreationTime string
	Tags         SearchMediaTagList
	RunIdList    SearchMediaRunIdListList
	File         SearchMediaFile
}

type SearchMediaFile struct {
	URL   string
	State string
}
type SearchMediaArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	Description          string
	OwnerId              int64
	Title                string
	PageNumber           int64
	CateId               string
	PageSize             int64
	From                 string
	SortBy               string
	To                   string
	Tag                  string
	KeyWord              string
}
type SearchMediaResponse struct {
	RequestId  string
	TotalNum   int64
	PageNumber int64
	PageSize   int64
	MediaList  SearchMediaMediaList
}

type SearchMediaTagList []string

func (list *SearchMediaTagList) UnmarshalJSON(data []byte) error {
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

type SearchMediaRunIdListList []string

func (list *SearchMediaRunIdListList) UnmarshalJSON(data []byte) error {
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

type SearchMediaMediaList []SearchMediaMedia

func (list *SearchMediaMediaList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SearchMediaMedia)
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

func (c *MtsClient) PlayInfo(req *PlayInfoArgs) (resp *PlayInfoResponse, err error) {
	resp = &PlayInfoResponse{}
	err = c.Request("PlayInfo", req, resp)
	return
}

type PlayInfoPlayInfo struct {
	Url          string
	Duration     string
	Size         string
	Width        string
	Height       string
	Bitrate      string
	Fps          string
	Format       string
	Definition   string
	Encryption   string
	Rand         string
	Plaintext    string
	Complexity   string
	ActivityName string
}
type PlayInfoArgs struct {
	PlayDomain           string
	ResourceOwnerId      string
	Formats              string
	ResourceOwnerAccount string
	OwnerAccount         string
	HlsUriToken          string
	OwnerId              string
	MediaId              string
	Rand                 string
	AuthTimeout          int64
	AuthInfo             string
}
type PlayInfoResponse struct {
	RequestId         string
	PlayInfoList      PlayInfoPlayInfoList
	NotFoundCDNDomain PlayInfoNotFoundCDNDomainList
}

type PlayInfoPlayInfoList []PlayInfoPlayInfo

func (list *PlayInfoPlayInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]PlayInfoPlayInfo)
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

type PlayInfoNotFoundCDNDomainList []string

func (list *PlayInfoNotFoundCDNDomainList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) DecryptKey(req *DecryptKeyArgs) (resp *DecryptKeyResponse, err error) {
	resp = &DecryptKeyResponse{}
	err = c.Request("DecryptKey", req, resp)
	return
}

type DecryptKeyArgs struct {
	Rand                 string
	ResourceOwnerId      string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              string
	CiphertextBlob       string
}
type DecryptKeyResponse struct {
	RequestId string
	Plaintext string
	Rand      string
}

func (c *MtsClient) SubmitFacerecogJob(req *SubmitFacerecogJobArgs) (resp *SubmitFacerecogJobResponse, err error) {
	resp = &SubmitFacerecogJobResponse{}
	err = c.Request("SubmitFacerecogJob", req, resp)
	return
}

type SubmitFacerecogJobArgs struct {
	Input                string
	UserData             string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	FacerecogConfig      string
	PipelineId           string
}
type SubmitFacerecogJobResponse struct {
	RequestId string
	JobId     string
}

func (c *MtsClient) DeactivateMediaWorkflow(req *DeactivateMediaWorkflowArgs) (resp *DeactivateMediaWorkflowResponse, err error) {
	resp = &DeactivateMediaWorkflowResponse{}
	err = c.Request("DeactivateMediaWorkflow", req, resp)
	return
}

type DeactivateMediaWorkflowMediaWorkflow struct {
	MediaWorkflowId string
	Name            string
	Topology        string
	State           string
	CreationTime    string
}
type DeactivateMediaWorkflowArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	MediaWorkflowId      string
	OwnerId              int64
}
type DeactivateMediaWorkflowResponse struct {
	RequestId     string
	MediaWorkflow DeactivateMediaWorkflowMediaWorkflow
}

func (c *MtsClient) ListPornPipeline(req *ListPornPipelineArgs) (resp *ListPornPipelineResponse, err error) {
	resp = &ListPornPipelineResponse{}
	err = c.Request("ListPornPipeline", req, resp)
	return
}

type ListPornPipelinePipeline struct {
	Id           string
	Name         string
	State        string
	Priority     string
	NotifyConfig ListPornPipelineNotifyConfig
}

type ListPornPipelineNotifyConfig struct {
	Topic string
	Queue string
}
type ListPornPipelineArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int64
	State                string
	OwnerId              int64
	PageNumber           int64
}
type ListPornPipelineResponse struct {
	RequestId    string
	TotalCount   int64
	PageNumber   int64
	PageSize     int64
	PipelineList ListPornPipelinePipelineList
}

type ListPornPipelinePipelineList []ListPornPipelinePipeline

func (list *ListPornPipelinePipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPornPipelinePipeline)
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

func (c *MtsClient) QueryPipelineList(req *QueryPipelineListArgs) (resp *QueryPipelineListResponse, err error) {
	resp = &QueryPipelineListResponse{}
	err = c.Request("QueryPipelineList", req, resp)
	return
}

type QueryPipelineListPipeline struct {
	Id           string
	Name         string
	State        string
	Speed        string
	SpeedLevel   int64
	Role         string
	NotifyConfig QueryPipelineListNotifyConfig
}

type QueryPipelineListNotifyConfig struct {
	Topic     string
	QueueName string
}
type QueryPipelineListArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	PipelineIds          string
	OwnerAccount         string
	OwnerId              int64
}
type QueryPipelineListResponse struct {
	RequestId    string
	PipelineList QueryPipelineListPipelineList
	NonExistPids QueryPipelineListNonExistPidList
}

type QueryPipelineListPipelineList []QueryPipelineListPipeline

func (list *QueryPipelineListPipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPipelineListPipeline)
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

type QueryPipelineListNonExistPidList []string

func (list *QueryPipelineListNonExistPidList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) AddCoverPipeline(req *AddCoverPipelineArgs) (resp *AddCoverPipelineResponse, err error) {
	resp = &AddCoverPipelineResponse{}
	err = c.Request("AddCoverPipeline", req, resp)
	return
}

type AddCoverPipelinePipeline struct {
	Id           string
	Name         string
	Priority     string
	State        string
	Role         string
	NotifyConfig AddCoverPipelineNotifyConfig
}

type AddCoverPipelineNotifyConfig struct {
	Topic string
	Queue string
}
type AddCoverPipelineArgs struct {
	ResourceOwnerId      int64
	Role                 string
	ResourceOwnerAccount string
	OwnerAccount         string
	Name                 string
	NotifyConfig         string
	OwnerId              int64
	Priority             string
}
type AddCoverPipelineResponse struct {
	RequestId string
	Pipeline  AddCoverPipelinePipeline
}

func (c *MtsClient) ListMedia(req *ListMediaArgs) (resp *ListMediaResponse, err error) {
	resp = &ListMediaResponse{}
	err = c.Request("ListMedia", req, resp)
	return
}

type ListMediaMedia struct {
	MediaId      string
	Title        string
	Description  string
	CoverURL     string
	CateId       int64
	Duration     string
	Format       string
	Size         string
	Bitrate      string
	Width        string
	Height       string
	Fps          string
	PublishState string
	CreationTime string
	Tags         ListMediaTagList
	RunIdList    ListMediaRunIdListList
	File         ListMediaFile
}

type ListMediaFile struct {
	URL   string
	State string
}
type ListMediaArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	NextPageToken        string
	OwnerAccount         string
	MaximumPageSize      int64
	From                 string
	To                   string
	OwnerId              int64
}
type ListMediaResponse struct {
	RequestId     string
	NextPageToken string
	MediaList     ListMediaMediaList
}

type ListMediaTagList []string

func (list *ListMediaTagList) UnmarshalJSON(data []byte) error {
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

type ListMediaRunIdListList []string

func (list *ListMediaRunIdListList) UnmarshalJSON(data []byte) error {
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

type ListMediaMediaList []ListMediaMedia

func (list *ListMediaMediaList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListMediaMedia)
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

func (c *MtsClient) AddMedia(req *AddMediaArgs) (resp *AddMediaResponse, err error) {
	resp = &AddMediaResponse{}
	err = c.Request("AddMedia", req, resp)
	return
}

type AddMediaMedia struct {
	MediaId      string
	Title        string
	Description  string
	CoverURL     string
	CateId       int64
	Duration     string
	Format       string
	Size         string
	Bitrate      string
	Width        string
	Height       string
	Fps          string
	PublishState string
	CreationTime string
	Tags         AddMediaTagList
	RunIdList    AddMediaRunIdListList
	File         AddMediaFile
}

type AddMediaFile struct {
	URL   string
	State string
}
type AddMediaArgs struct {
	ResourceOwnerId       int64
	ResourceOwnerAccount  string
	OwnerAccount          string
	Description           string
	OwnerId               int64
	Title                 string
	Tags                  string
	CoverURL              string
	CateId                int64
	FileURL               string
	MediaWorkflowId       string
	MediaWorkflowUserData string
}
type AddMediaResponse struct {
	RequestId string
	Media     AddMediaMedia
}

type AddMediaTagList []string

func (list *AddMediaTagList) UnmarshalJSON(data []byte) error {
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

type AddMediaRunIdListList []string

func (list *AddMediaRunIdListList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) AddMediaTag(req *AddMediaTagArgs) (resp *AddMediaTagResponse, err error) {
	resp = &AddMediaTagResponse{}
	err = c.Request("AddMediaTag", req, resp)
	return
}

type AddMediaTagArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	Tag                  string
	OwnerId              int64
	MediaId              string
}
type AddMediaTagResponse struct {
	RequestId string
}

func (c *MtsClient) UpdateMediaPublishState(req *UpdateMediaPublishStateArgs) (resp *UpdateMediaPublishStateResponse, err error) {
	resp = &UpdateMediaPublishStateResponse{}
	err = c.Request("UpdateMediaPublishState", req, resp)
	return
}

type UpdateMediaPublishStateArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	Publish              core.Bool
	OwnerAccount         string
	OwnerId              int64
	MediaId              string
}
type UpdateMediaPublishStateResponse struct {
	RequestId string
}

func (c *MtsClient) UpdateCoverPipeline(req *UpdateCoverPipelineArgs) (resp *UpdateCoverPipelineResponse, err error) {
	resp = &UpdateCoverPipelineResponse{}
	err = c.Request("UpdateCoverPipeline", req, resp)
	return
}

type UpdateCoverPipelinePipeline struct {
	Id           string
	Name         string
	State        string
	Priority     int
	Role         string
	NotifyConfig UpdateCoverPipelineNotifyConfig
}

type UpdateCoverPipelineNotifyConfig struct {
	Topic string
	Queue string
}
type UpdateCoverPipelineArgs struct {
	ResourceOwnerId      int64
	Role                 string
	ResourceOwnerAccount string
	OwnerAccount         string
	Name                 string
	State                string
	NotifyConfig         string
	OwnerId              int64
	Priority             int
	PipelineId           string
}
type UpdateCoverPipelineResponse struct {
	RequestId string
	Pipeline  UpdateCoverPipelinePipeline
}

func (c *MtsClient) ListAllCategory(req *ListAllCategoryArgs) (resp *ListAllCategoryResponse, err error) {
	resp = &ListAllCategoryResponse{}
	err = c.Request("ListAllCategory", req, resp)
	return
}

type ListAllCategoryCategory struct {
	CateId   string
	CateName string
	ParentId string
	Level    string
}
type ListAllCategoryArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type ListAllCategoryResponse struct {
	RequestId    string
	CategoryList ListAllCategoryCategoryList
}

type ListAllCategoryCategoryList []ListAllCategoryCategory

func (list *ListAllCategoryCategoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAllCategoryCategory)
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

func (c *MtsClient) ListJob(req *ListJobArgs) (resp *ListJobResponse, err error) {
	resp = &ListJobResponse{}
	err = c.Request("ListJob", req, resp)
	return
}

type ListJobJob struct {
	JobId            string
	State            string
	Code             string
	Message          string
	Percent          int64
	PipelineId       string
	CreationTime     string
	FinishTime       string
	Input            ListJobInput
	Output           ListJobOutput
	MNSMessageResult ListJobMNSMessageResult
}

type ListJobInput struct {
	Bucket   string
	Location string
	Object   string
}

type ListJobOutput struct {
	TemplateId             string
	UserData               string
	Rotate                 string
	VideoStreamMap         string
	AudioStreamMap         string
	DeWatermark            string
	Priority               string
	WaterMarkConfigUrl     string
	MergeConfigUrl         string
	WaterMarkList          ListJobWaterMarkList
	MergeList              ListJobMergeList
	OpeningList            ListJobMerge1List
	TailSlateList          ListJobMerge2List
	OutputFile             ListJobOutputFile
	M3U8NonStandardSupport ListJobM3U8NonStandardSupport
	Properties             ListJobProperties
	Clip                   ListJobClip
	SuperReso              ListJobSuperReso
	SubtitleConfig         ListJobSubtitleConfig
	TransConfig            ListJobTransConfig
	MuxConfig              ListJobMuxConfig
	Audio                  ListJobAudio
	Video                  ListJobVideo
	Container              ListJobContainer
	Encryption             ListJobEncryption
}

type ListJobWaterMark struct {
	WaterMarkTemplateId string
	Width               string
	Height              string
	Dx                  string
	Dy                  string
	ReferPos            string
	Type                string
	InputFile           ListJobInputFile
}

type ListJobInputFile struct {
	Bucket   string
	Location string
	Object   string
}

type ListJobMerge struct {
	MergeURL string
	Start    string
	Duration string
	RoleArn  string
}

type ListJobMerge1 struct {
	OpenUrl string
	Start   string
	Width   string
	Height  string
}

type ListJobMerge2 struct {
	TailUrl       string
	Start         string
	BlendDuration string
	Width         string
	Height        string
	IsMergeAudio  core.Bool
	BgColor       string
}

type ListJobOutputFile struct {
	Bucket   string
	Location string
	Object   string
	RoleArn  string
}

type ListJobM3U8NonStandardSupport struct {
	TS ListJobTS
}

type ListJobTS struct {
	Md5Support  core.Bool
	SizeSupport core.Bool
}

type ListJobProperties struct {
	Width      string
	Height     string
	Bitrate    string
	Duration   string
	Fps        string
	FileSize   string
	FileFormat string
	Streams    ListJobStreams
	Format     ListJobFormat
}

type ListJobStreams struct {
	VideoStreamList    ListJobVideoStreamList
	AudioStreamList    ListJobAudioStreamList
	SubtitleStreamList ListJobSubtitleStreamList
}

type ListJobVideoStream struct {
	Index          string
	CodecName      string
	CodecLongName  string
	Profile        string
	CodecTimeBase  string
	CodecTagString string
	CodecTag       string
	Width          string
	Height         string
	HasBFrames     string
	Sar            string
	Dar            string
	PixFmt         string
	Level          string
	Fps            string
	AvgFPS         string
	Timebase       string
	StartTime      string
	Duration       string
	Bitrate        string
	NumFrames      string
	Lang           string
	NetworkCost    ListJobNetworkCost
}

type ListJobNetworkCost struct {
	PreloadTime   string
	CostBandwidth string
	AvgBitrate    string
}

type ListJobAudioStream struct {
	Index          string
	CodecName      string
	CodecTimeBase  string
	CodecLongName  string
	CodecTagString string
	CodecTag       string
	SampleFmt      string
	Samplerate     string
	Channels       string
	ChannelLayout  string
	Timebase       string
	StartTime      string
	Duration       string
	Bitrate        string
	NumFrames      string
	Lang           string
}

type ListJobSubtitleStream struct {
	Index string
	Lang  string
}

type ListJobFormat struct {
	NumStreams     string
	NumPrograms    string
	FormatName     string
	FormatLongName string
	StartTime      string
	Duration       string
	Size           string
	Bitrate        string
}

type ListJobClip struct {
	TimeSpan ListJobTimeSpan
}

type ListJobTimeSpan struct {
	Seek     string
	Duration string
}

type ListJobSuperReso struct {
	IsHalfSample string
}

type ListJobSubtitleConfig struct {
	SubtitleList    ListJobSubtitleList
	ExtSubtitleList ListJobExtSubtitleList
}

type ListJobSubtitle struct {
	Map string
}

type ListJobExtSubtitle struct {
	FontName string
	CharEnc  string
	Input3   ListJobInput3
}

type ListJobInput3 struct {
	Bucket   string
	Location string
	Object   string
}

type ListJobTransConfig struct {
	TransMode               string
	IsCheckReso             string
	IsCheckResoFail         string
	IsCheckVideoBitrate     string
	IsCheckAudioBitrate     string
	AdjDarMethod            string
	IsCheckVideoBitrateFail string
	IsCheckAudioBitrateFail string
}

type ListJobMuxConfig struct {
	Segment ListJobSegment
	Gif     ListJobGif
}

type ListJobSegment struct {
	Duration string
}

type ListJobGif struct {
	Loop            string
	FinalDelay      string
	IsCustomPalette string
	DitherMode      string
}

type ListJobAudio struct {
	Codec      string
	Profile    string
	Samplerate string
	Bitrate    string
	Channels   string
	Qscale     string
	Volume     ListJobVolume
}

type ListJobVolume struct {
	Level  string
	Method string
}

type ListJobVideo struct {
	Codec      string
	Profile    string
	Bitrate    string
	Crf        string
	Width      string
	Height     string
	Fps        string
	Gop        string
	Preset     string
	ScanMode   string
	Bufsize    string
	Maxrate    string
	PixFmt     string
	Degrain    string
	Qscale     string
	Crop       string
	Pad        string
	MaxFps     string
	BitrateBnd ListJobBitrateBnd
}

type ListJobBitrateBnd struct {
	Max string
	Min string
}

type ListJobContainer struct {
	Format string
}

type ListJobEncryption struct {
	Type    string
	Id      string
	Key     string
	KeyUri  string
	KeyType string
	SkipCnt string
}

type ListJobMNSMessageResult struct {
	MessageId    string
	ErrorMessage string
	ErrorCode    string
}
type ListJobArgs struct {
	ResourceOwnerId            int64
	ResourceOwnerAccount       string
	NextPageToken              string
	StartOfJobCreatedTimeRange string
	OwnerAccount               string
	MaximumPageSize            int64
	State                      string
	OwnerId                    int64
	EndOfJobCreatedTimeRange   string
	PipelineId                 string
}
type ListJobResponse struct {
	RequestId     string
	NextPageToken string
	JobList       ListJobJobList
}

type ListJobWaterMarkList []ListJobWaterMark

func (list *ListJobWaterMarkList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobWaterMark)
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

type ListJobMergeList []ListJobMerge

func (list *ListJobMergeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobMerge)
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

type ListJobMerge1List []ListJobMerge1

func (list *ListJobMerge1List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobMerge1)
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

type ListJobMerge2List []ListJobMerge2

func (list *ListJobMerge2List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobMerge2)
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

type ListJobVideoStreamList []ListJobVideoStream

func (list *ListJobVideoStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobVideoStream)
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

type ListJobAudioStreamList []ListJobAudioStream

func (list *ListJobAudioStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobAudioStream)
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

type ListJobSubtitleStreamList []ListJobSubtitleStream

func (list *ListJobSubtitleStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobSubtitleStream)
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

type ListJobSubtitleList []ListJobSubtitle

func (list *ListJobSubtitleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobSubtitle)
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

type ListJobExtSubtitleList []ListJobExtSubtitle

func (list *ListJobExtSubtitleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobExtSubtitle)
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

type ListJobJobList []ListJobJob

func (list *ListJobJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobJob)
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

func (c *MtsClient) AddCategory(req *AddCategoryArgs) (resp *AddCategoryResponse, err error) {
	resp = &AddCategoryResponse{}
	err = c.Request("AddCategory", req, resp)
	return
}

type AddCategoryCategory struct {
	CateId   string
	CateName string
	ParentId string
	Level    string
}
type AddCategoryArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	ParentId             int64
	CateName             string
}
type AddCategoryResponse struct {
	RequestId string
	Category  AddCategoryCategory
}

func (c *MtsClient) UpdateTerrorismPipeline(req *UpdateTerrorismPipelineArgs) (resp *UpdateTerrorismPipelineResponse, err error) {
	resp = &UpdateTerrorismPipelineResponse{}
	err = c.Request("UpdateTerrorismPipeline", req, resp)
	return
}

type UpdateTerrorismPipelinePipeline struct {
	Id           string
	Name         string
	State        string
	Priority     int
	NotifyConfig UpdateTerrorismPipelineNotifyConfig
}

type UpdateTerrorismPipelineNotifyConfig struct {
	Topic string
	Queue string
}
type UpdateTerrorismPipelineArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	Name                 string
	State                string
	NotifyConfig         string
	OwnerId              int64
	Priority             int
	PipelineId           string
}
type UpdateTerrorismPipelineResponse struct {
	RequestId string
	Pipeline  UpdateTerrorismPipelinePipeline
}

func (c *MtsClient) SubmitCensorJob(req *SubmitCensorJobArgs) (resp *SubmitCensorJobResponse, err error) {
	resp = &SubmitCensorJobResponse{}
	err = c.Request("SubmitCensorJob", req, resp)
	return
}

type SubmitCensorJobArgs struct {
	Input                string
	UserData             string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	CensorConfig         string
	PipelineId           string
}
type SubmitCensorJobResponse struct {
	RequestId string
	JobId     string
}

func (c *MtsClient) DeletePipeline(req *DeletePipelineArgs) (resp *DeletePipelineResponse, err error) {
	resp = &DeletePipelineResponse{}
	err = c.Request("DeletePipeline", req, resp)
	return
}

type DeletePipelineArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	PipelineId           string
}
type DeletePipelineResponse struct {
	RequestId  string
	PipelineId string
}

func (c *MtsClient) QueryMediaWorkflowExecutionList(req *QueryMediaWorkflowExecutionListArgs) (resp *QueryMediaWorkflowExecutionListResponse, err error) {
	resp = &QueryMediaWorkflowExecutionListResponse{}
	err = c.Request("QueryMediaWorkflowExecutionList", req, resp)
	return
}

type QueryMediaWorkflowExecutionListMediaWorkflowExecution struct {
	RunId           string
	MediaWorkflowId string
	Name            string
	State           string
	MediaId         string
	CreationTime    string
	ActivityList    QueryMediaWorkflowExecutionListActivityList
	Input           QueryMediaWorkflowExecutionListInput
}

type QueryMediaWorkflowExecutionListActivity struct {
	Name             string
	Type             string
	JobId            string
	State            string
	Code             string
	Message          string
	StartTime        string
	EndTime          string
	MNSMessageResult QueryMediaWorkflowExecutionListMNSMessageResult
}

type QueryMediaWorkflowExecutionListMNSMessageResult struct {
	MessageId    string
	ErrorMessage string
	ErrorCode    string
}

type QueryMediaWorkflowExecutionListInput struct {
	UserData  string
	InputFile QueryMediaWorkflowExecutionListInputFile
}

type QueryMediaWorkflowExecutionListInputFile struct {
	Bucket   string
	Location string
	Object   string
}
type QueryMediaWorkflowExecutionListArgs struct {
	RunIds               string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type QueryMediaWorkflowExecutionListResponse struct {
	RequestId                  string
	MediaWorkflowExecutionList QueryMediaWorkflowExecutionListMediaWorkflowExecutionList
	NonExistRunIds             QueryMediaWorkflowExecutionListNonExistRunIdList
}

type QueryMediaWorkflowExecutionListActivityList []QueryMediaWorkflowExecutionListActivity

func (list *QueryMediaWorkflowExecutionListActivityList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaWorkflowExecutionListActivity)
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

type QueryMediaWorkflowExecutionListMediaWorkflowExecutionList []QueryMediaWorkflowExecutionListMediaWorkflowExecution

func (list *QueryMediaWorkflowExecutionListMediaWorkflowExecutionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaWorkflowExecutionListMediaWorkflowExecution)
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

type QueryMediaWorkflowExecutionListNonExistRunIdList []string

func (list *QueryMediaWorkflowExecutionListNonExistRunIdList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) ReportPornJobResult(req *ReportPornJobResultArgs) (resp *ReportPornJobResultResponse, err error) {
	resp = &ReportPornJobResultResponse{}
	err = c.Request("ReportPornJobResult", req, resp)
	return
}

type ReportPornJobResultArgs struct {
	JobId                string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	Label                string
	Detail               string
	OwnerId              int64
}
type ReportPornJobResultResponse struct {
	RequestId string
	JobId     string
}

func (c *MtsClient) UpdateMedia(req *UpdateMediaArgs) (resp *UpdateMediaResponse, err error) {
	resp = &UpdateMediaResponse{}
	err = c.Request("UpdateMedia", req, resp)
	return
}

type UpdateMediaMedia struct {
	MediaId      string
	Title        string
	Description  string
	CoverURL     string
	CateId       int64
	Duration     string
	Format       string
	Size         string
	Bitrate      string
	Width        string
	Height       string
	Fps          string
	PublishState string
	CreationTime string
	Tags         UpdateMediaTagList
	RunIdList    UpdateMediaRunIdListList
	File         UpdateMediaFile
}

type UpdateMediaFile struct {
	URL   string
	State string
}
type UpdateMediaArgs struct {
	CoverURL             string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	CateId               int64
	OwnerAccount         string
	Description          string
	OwnerId              int64
	MediaId              string
	Title                string
	Tags                 string
}
type UpdateMediaResponse struct {
	RequestId string
	Media     UpdateMediaMedia
}

type UpdateMediaTagList []string

func (list *UpdateMediaTagList) UnmarshalJSON(data []byte) error {
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

type UpdateMediaRunIdListList []string

func (list *UpdateMediaRunIdListList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) QueryAnnotationJobList(req *QueryAnnotationJobListArgs) (resp *QueryAnnotationJobListResponse, err error) {
	resp = &QueryAnnotationJobListResponse{}
	err = c.Request("QueryAnnotationJobList", req, resp)
	return
}

type QueryAnnotationJobListAnnotationJob struct {
	Id                    string
	UserData              string
	PipelineId            string
	State                 string
	Code                  string
	Message               string
	CreationTime          string
	Input                 QueryAnnotationJobListInput
	VideoAnnotationResult QueryAnnotationJobListVideoAnnotationResult
}

type QueryAnnotationJobListInput struct {
	Bucket   string
	Location string
	Object   string
}

type QueryAnnotationJobListVideoAnnotationResult struct {
	Details     string
	Annotations QueryAnnotationJobListAnnotationList
}

type QueryAnnotationJobListAnnotation struct {
	Label string
	Score string
}
type QueryAnnotationJobListArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	AnnotationJobIds     string
}
type QueryAnnotationJobListResponse struct {
	RequestId         string
	AnnotationJobList QueryAnnotationJobListAnnotationJobList
	NonExistIds       QueryAnnotationJobListNonExistIdList
}

type QueryAnnotationJobListAnnotationList []QueryAnnotationJobListAnnotation

func (list *QueryAnnotationJobListAnnotationList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAnnotationJobListAnnotation)
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

type QueryAnnotationJobListAnnotationJobList []QueryAnnotationJobListAnnotationJob

func (list *QueryAnnotationJobListAnnotationJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAnnotationJobListAnnotationJob)
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

type QueryAnnotationJobListNonExistIdList []string

func (list *QueryAnnotationJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) QueryVideoSummaryJobList(req *QueryVideoSummaryJobListArgs) (resp *QueryVideoSummaryJobListResponse, err error) {
	resp = &QueryVideoSummaryJobListResponse{}
	err = c.Request("QueryVideoSummaryJobList", req, resp)
	return
}

type QueryVideoSummaryJobListJob struct {
	Id                 string
	UserData           string
	PipelineId         string
	State              string
	Code               string
	Message            string
	CreationTime       string
	Input              QueryVideoSummaryJobListInput
	VideoSummaryResult QueryVideoSummaryJobListVideoSummaryResult
}

type QueryVideoSummaryJobListInput struct {
	Bucket   string
	Location string
	Object   string
}

type QueryVideoSummaryJobListVideoSummaryResult struct {
	VideoSummaryList QueryVideoSummaryJobListVideoSummaryList
}

type QueryVideoSummaryJobListVideoSummary struct {
	StartTime string
	EndTime   string
}
type QueryVideoSummaryJobListArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	JobIds               string
	OwnerAccount         string
	OwnerId              int64
}
type QueryVideoSummaryJobListResponse struct {
	RequestId   string
	JobList     QueryVideoSummaryJobListJobList
	NonExistIds QueryVideoSummaryJobListNonExistIdList
}

type QueryVideoSummaryJobListVideoSummaryList []QueryVideoSummaryJobListVideoSummary

func (list *QueryVideoSummaryJobListVideoSummaryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryVideoSummaryJobListVideoSummary)
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

type QueryVideoSummaryJobListJobList []QueryVideoSummaryJobListJob

func (list *QueryVideoSummaryJobListJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryVideoSummaryJobListJob)
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

type QueryVideoSummaryJobListNonExistIdList []string

func (list *QueryVideoSummaryJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) QueryTemplateList(req *QueryTemplateListArgs) (resp *QueryTemplateListResponse, err error) {
	resp = &QueryTemplateListResponse{}
	err = c.Request("QueryTemplateList", req, resp)
	return
}

type QueryTemplateListTemplate struct {
	Id          string
	Name        string
	State       string
	Container   QueryTemplateListContainer
	Video       QueryTemplateListVideo
	Audio       QueryTemplateListAudio
	TransConfig QueryTemplateListTransConfig
	MuxConfig   QueryTemplateListMuxConfig
}

type QueryTemplateListContainer struct {
	Format string
}

type QueryTemplateListVideo struct {
	Codec      string
	Profile    string
	Bitrate    string
	Crf        string
	Width      string
	Height     string
	Fps        string
	Gop        string
	Preset     string
	ScanMode   string
	Bufsize    string
	Maxrate    string
	PixFmt     string
	Degrain    string
	Qscale     string
	Remove     string
	Crop       string
	Pad        string
	MaxFps     string
	BitrateBnd QueryTemplateListBitrateBnd
}

type QueryTemplateListBitrateBnd struct {
	Max string
	Min string
}

type QueryTemplateListAudio struct {
	Codec      string
	Profile    string
	Samplerate string
	Bitrate    string
	Channels   string
	Qscale     string
	Remove     string
}

type QueryTemplateListTransConfig struct {
	TransMode               string
	IsCheckReso             string
	IsCheckResoFail         string
	IsCheckVideoBitrate     string
	IsCheckAudioBitrate     string
	AdjDarMethod            string
	IsCheckVideoBitrateFail string
	IsCheckAudioBitrateFail string
}

type QueryTemplateListMuxConfig struct {
	Segment QueryTemplateListSegment
	Gif     QueryTemplateListGif
}

type QueryTemplateListSegment struct {
	Duration string
}

type QueryTemplateListGif struct {
	Loop            string
	FinalDelay      string
	IsCustomPalette string
	DitherMode      string
}
type QueryTemplateListArgs struct {
	ResourceOwnerId      int64
	TemplateIds          string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type QueryTemplateListResponse struct {
	RequestId    string
	TemplateList QueryTemplateListTemplateList
	NonExistTids QueryTemplateListNonExistTidList
}

type QueryTemplateListTemplateList []QueryTemplateListTemplate

func (list *QueryTemplateListTemplateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTemplateListTemplate)
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

type QueryTemplateListNonExistTidList []string

func (list *QueryTemplateListNonExistTidList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) QuerySnapshotJobList(req *QuerySnapshotJobListArgs) (resp *QuerySnapshotJobListResponse, err error) {
	resp = &QuerySnapshotJobListResponse{}
	err = c.Request("QuerySnapshotJobList", req, resp)
	return
}

type QuerySnapshotJobListSnapshotJob struct {
	Id               string
	UserData         string
	PipelineId       string
	State            string
	Code             string
	Count            string
	TileCount        string
	Message          string
	CreationTime     string
	Input            QuerySnapshotJobListInput
	SnapshotConfig   QuerySnapshotJobListSnapshotConfig
	MNSMessageResult QuerySnapshotJobListMNSMessageResult
}

type QuerySnapshotJobListInput struct {
	Bucket   string
	Location string
	Object   string
	RoleArn  string
}

type QuerySnapshotJobListSnapshotConfig struct {
	Time           string
	Interval       string
	Num            string
	Width          string
	Height         string
	FrameType      string
	OutputFile     QuerySnapshotJobListOutputFile
	TileOutputFile QuerySnapshotJobListTileOutputFile
	TileOut        QuerySnapshotJobListTileOut
}

type QuerySnapshotJobListOutputFile struct {
	Bucket   string
	Location string
	Object   string
	RoleArn  string
}

type QuerySnapshotJobListTileOutputFile struct {
	Bucket   string
	Location string
	Object   string
	RoleArn  string
}

type QuerySnapshotJobListTileOut struct {
	Lines         string
	Columns       string
	CellWidth     string
	CellHeight    string
	Margin        string
	Padding       string
	Color         string
	IsKeepCellPic string
}

type QuerySnapshotJobListMNSMessageResult struct {
	MessageId    string
	ErrorMessage string
	ErrorCode    string
}
type QuerySnapshotJobListArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	SnapshotJobIds       string
	OwnerAccount         string
	OwnerId              int64
}
type QuerySnapshotJobListResponse struct {
	RequestId              string
	SnapshotJobList        QuerySnapshotJobListSnapshotJobList
	NonExistSnapshotJobIds QuerySnapshotJobListNonExistSnapshotJobIdList
}

type QuerySnapshotJobListSnapshotJobList []QuerySnapshotJobListSnapshotJob

func (list *QuerySnapshotJobListSnapshotJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QuerySnapshotJobListSnapshotJob)
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

type QuerySnapshotJobListNonExistSnapshotJobIdList []string

func (list *QuerySnapshotJobListNonExistSnapshotJobIdList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) UpdateAsrPipeline(req *UpdateAsrPipelineArgs) (resp *UpdateAsrPipelineResponse, err error) {
	resp = &UpdateAsrPipelineResponse{}
	err = c.Request("UpdateAsrPipeline", req, resp)
	return
}

type UpdateAsrPipelinePipeline struct {
	Id           string
	Name         string
	State        string
	Priority     int
	NotifyConfig UpdateAsrPipelineNotifyConfig
}

type UpdateAsrPipelineNotifyConfig struct {
	Topic     string
	QueueName string
}
type UpdateAsrPipelineArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	Name                 string
	State                string
	NotifyConfig         string
	OwnerId              int64
	Priority             int
	PipelineId           string
}
type UpdateAsrPipelineResponse struct {
	RequestId string
	Pipeline  UpdateAsrPipelinePipeline
}

func (c *MtsClient) RefreshCdnDomainConfigsCache(req *RefreshCdnDomainConfigsCacheArgs) (resp *RefreshCdnDomainConfigsCacheResponse, err error) {
	resp = &RefreshCdnDomainConfigsCacheResponse{}
	err = c.Request("RefreshCdnDomainConfigsCache", req, resp)
	return
}

type RefreshCdnDomainConfigsCacheArgs struct {
	ResourceOwnerId      string
	ResourceOwnerAccount string
	OwnerAccount         string
	Domains              string
	OwnerId              string
}
type RefreshCdnDomainConfigsCacheResponse struct {
	RequestId     string
	SucessDomains RefreshCdnDomainConfigsCacheSucessDomainList
	FailedDomains RefreshCdnDomainConfigsCacheFailedDomainList
}

type RefreshCdnDomainConfigsCacheSucessDomainList []string

func (list *RefreshCdnDomainConfigsCacheSucessDomainList) UnmarshalJSON(data []byte) error {
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

type RefreshCdnDomainConfigsCacheFailedDomainList []string

func (list *RefreshCdnDomainConfigsCacheFailedDomainList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) SubmitEditingJobs(req *SubmitEditingJobsArgs) (resp *SubmitEditingJobsResponse, err error) {
	resp = &SubmitEditingJobsResponse{}
	err = c.Request("SubmitEditingJobs", req, resp)
	return
}

type SubmitEditingJobsJobResult struct {
	Success core.Bool
	Code    string
	Message string
	Job     SubmitEditingJobsJob
}

type SubmitEditingJobsJob struct {
	JobId            string
	State            string
	Code             string
	Message          string
	Percent          int64
	PipelineId       string
	CreationTime     string
	FinishTime       string
	EditingInputs    SubmitEditingJobsEditingInputList
	EditingConfig    SubmitEditingJobsEditingConfig
	MNSMessageResult SubmitEditingJobsMNSMessageResult
}

type SubmitEditingJobsEditingInput struct {
	Id          string
	InputFile   SubmitEditingJobsInputFile
	InputConfig SubmitEditingJobsInputConfig
}

type SubmitEditingJobsInputFile struct {
	Bucket   string
	Location string
	Object   string
}

type SubmitEditingJobsInputConfig struct {
	DeinterlaceMethod string
	IsNormalSar       string
}

type SubmitEditingJobsEditingConfig struct {
	TemplateId             string
	UserData               string
	Rotate                 string
	VideoStreamMap         string
	AudioStreamMap         string
	DeWatermark            string
	Priority               string
	WaterMarkConfigUrl     string
	MergeConfigUrl         string
	WaterMarkList          SubmitEditingJobsWaterMarkList
	MergeList              SubmitEditingJobsMergeList
	DigiWaterMark          SubmitEditingJobsDigiWaterMark
	OutputFile             SubmitEditingJobsOutputFile
	M3U8NonStandardSupport SubmitEditingJobsM3U8NonStandardSupport
	Properties             SubmitEditingJobsProperties
	Clip                   SubmitEditingJobsClip
	SuperReso              SubmitEditingJobsSuperReso
	SubtitleConfig         SubmitEditingJobsSubtitleConfig
	TransConfig            SubmitEditingJobsTransConfig
	MuxConfig              SubmitEditingJobsMuxConfig
	Audio                  SubmitEditingJobsAudio
	Video                  SubmitEditingJobsVideo
	Container              SubmitEditingJobsContainer
	Encryption             SubmitEditingJobsEncryption
	Editing                SubmitEditingJobsEditing
}

type SubmitEditingJobsWaterMark struct {
	WaterMarkTemplateId string
	Width               string
	Height              string
	Dx                  string
	Dy                  string
	ReferPos            string
	Type                string
	InputFile1          SubmitEditingJobsInputFile1
}

type SubmitEditingJobsInputFile1 struct {
	Bucket   string
	Location string
	Object   string
}

type SubmitEditingJobsMerge struct {
	MergeURL string
	Start    string
	Duration string
	RoleArn  string
}

type SubmitEditingJobsDigiWaterMark struct {
	Type       string
	Alpha      string
	InputFile2 SubmitEditingJobsInputFile2
}

type SubmitEditingJobsInputFile2 struct {
	Bucket   string
	Location string
	Object   string
}

type SubmitEditingJobsOutputFile struct {
	Bucket   string
	Location string
	Object   string
	RoleArn  string
}

type SubmitEditingJobsM3U8NonStandardSupport struct {
	TS SubmitEditingJobsTS
}

type SubmitEditingJobsTS struct {
	Md5Support  core.Bool
	SizeSupport core.Bool
}

type SubmitEditingJobsProperties struct {
	Width      string
	Height     string
	Bitrate    string
	Duration   string
	Fps        string
	FileSize   string
	FileFormat string
	Streams    SubmitEditingJobsStreams
	Format     SubmitEditingJobsFormat
}

type SubmitEditingJobsStreams struct {
	VideoStreamList    SubmitEditingJobsVideoStreamList
	AudioStreamList    SubmitEditingJobsAudioStreamList
	SubtitleStreamList SubmitEditingJobsSubtitleStreamList
}

type SubmitEditingJobsVideoStream struct {
	Index          string
	CodecName      string
	CodecLongName  string
	Profile        string
	CodecTimeBase  string
	CodecTagString string
	CodecTag       string
	Width          string
	Height         string
	HasBFrames     string
	Sar            string
	Dar            string
	PixFmt         string
	Level          string
	Fps            string
	AvgFPS         string
	Timebase       string
	StartTime      string
	Duration       string
	Bitrate        string
	NumFrames      string
	Lang           string
	NetworkCost    SubmitEditingJobsNetworkCost
}

type SubmitEditingJobsNetworkCost struct {
	PreloadTime   string
	CostBandwidth string
	AvgBitrate    string
}

type SubmitEditingJobsAudioStream struct {
	Index          string
	CodecName      string
	CodecTimeBase  string
	CodecLongName  string
	CodecTagString string
	CodecTag       string
	SampleFmt      string
	Samplerate     string
	Channels       string
	ChannelLayout  string
	Timebase       string
	StartTime      string
	Duration       string
	Bitrate        string
	NumFrames      string
	Lang           string
}

type SubmitEditingJobsSubtitleStream struct {
	Index string
	Lang  string
}

type SubmitEditingJobsFormat struct {
	NumStreams     string
	NumPrograms    string
	FormatName     string
	FormatLongName string
	StartTime      string
	Duration       string
	Size           string
	Bitrate        string
}

type SubmitEditingJobsClip struct {
	TimeSpan SubmitEditingJobsTimeSpan
}

type SubmitEditingJobsTimeSpan struct {
	Seek     string
	Duration string
}

type SubmitEditingJobsSuperReso struct {
	IsHalfSample string
}

type SubmitEditingJobsSubtitleConfig struct {
	SubtitleList    SubmitEditingJobsSubtitleList
	ExtSubtitleList SubmitEditingJobsExtSubtitleList
}

type SubmitEditingJobsSubtitle struct {
	Map string
}

type SubmitEditingJobsExtSubtitle struct {
	FontName string
	CharEnc  string
	Input    SubmitEditingJobsInput
}

type SubmitEditingJobsInput struct {
	Bucket   string
	Location string
	Object   string
}

type SubmitEditingJobsTransConfig struct {
	TransMode               string
	IsCheckReso             string
	IsCheckResoFail         string
	IsCheckVideoBitrate     string
	IsCheckAudioBitrate     string
	AdjDarMethod            string
	IsCheckVideoBitrateFail string
	IsCheckAudioBitrateFail string
	Duration                string
}

type SubmitEditingJobsMuxConfig struct {
	Segment SubmitEditingJobsSegment
	Gif     SubmitEditingJobsGif
}

type SubmitEditingJobsSegment struct {
	Duration string
}

type SubmitEditingJobsGif struct {
	Loop            string
	FinalDelay      string
	IsCustomPalette string
	DitherMode      string
}

type SubmitEditingJobsAudio struct {
	Codec      string
	Profile    string
	Samplerate string
	Bitrate    string
	Channels   string
	Qscale     string
	Volume     SubmitEditingJobsVolume
}

type SubmitEditingJobsVolume struct {
	Level  string
	Method string
}

type SubmitEditingJobsVideo struct {
	Codec      string
	Profile    string
	Bitrate    string
	Crf        string
	Width      string
	Height     string
	Fps        string
	Gop        string
	Preset     string
	ScanMode   string
	Bufsize    string
	Maxrate    string
	PixFmt     string
	Degrain    string
	Qscale     string
	Crop       string
	Pad        string
	MaxFps     string
	BitrateBnd SubmitEditingJobsBitrateBnd
}

type SubmitEditingJobsBitrateBnd struct {
	Max string
	Min string
}

type SubmitEditingJobsContainer struct {
	Format string
}

type SubmitEditingJobsEncryption struct {
	Type    string
	Id      string
	Key     string
	KeyUri  string
	KeyType string
	SkipCnt string
}

type SubmitEditingJobsEditing struct {
	ClipList SubmitEditingJobsClip3List
	Timeline SubmitEditingJobsTimeline
}

type SubmitEditingJobsClip3 struct {
	Id            string
	Type          string
	SourceType    string
	SourceID      string
	SourceStrmMap string
	In            string
	Out           string
	Effects       SubmitEditingJobsEffectList
}

type SubmitEditingJobsEffect struct {
	Effect       string
	EffectConfig string
}

type SubmitEditingJobsTimeline struct {
	TrackList      SubmitEditingJobsTrackList
	TimelineConfig SubmitEditingJobsTimelineConfig
}

type SubmitEditingJobsTrack struct {
	Id    string
	Type  string
	Order string
	Clips SubmitEditingJobsClip4List
}

type SubmitEditingJobsClip4 struct {
	ClipID      string
	In          string
	Out         string
	ClipsConfig SubmitEditingJobsClipsConfig
}

type SubmitEditingJobsClipsConfig struct {
	ClipsConfigVideo SubmitEditingJobsClipsConfigVideo
}

type SubmitEditingJobsClipsConfigVideo struct {
	L string
	T string
}

type SubmitEditingJobsTimelineConfig struct {
	TimelineConfigVideo SubmitEditingJobsTimelineConfigVideo
	TimelineConfigAudio SubmitEditingJobsTimelineConfigAudio
}

type SubmitEditingJobsTimelineConfigVideo struct {
	Width   string
	Height  string
	BgColor string
	Fps     string
}

type SubmitEditingJobsTimelineConfigAudio struct {
	Samplerate    string
	ChannelLayout string
	Channels      string
}

type SubmitEditingJobsMNSMessageResult struct {
	MessageId    string
	ErrorMessage string
	ErrorCode    string
}
type SubmitEditingJobsArgs struct {
	OutputBucket         string
	ResourceOwnerId      int64
	EditingJobOutputs    string
	ResourceOwnerAccount string
	OwnerAccount         string
	OutputLocation       string
	OwnerId              int64
	EditingInputs        string
	PipelineId           string
}
type SubmitEditingJobsResponse struct {
	RequestId     string
	JobResultList SubmitEditingJobsJobResultList
}

type SubmitEditingJobsEditingInputList []SubmitEditingJobsEditingInput

func (list *SubmitEditingJobsEditingInputList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitEditingJobsEditingInput)
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

type SubmitEditingJobsWaterMarkList []SubmitEditingJobsWaterMark

func (list *SubmitEditingJobsWaterMarkList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitEditingJobsWaterMark)
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

type SubmitEditingJobsMergeList []SubmitEditingJobsMerge

func (list *SubmitEditingJobsMergeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitEditingJobsMerge)
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

type SubmitEditingJobsVideoStreamList []SubmitEditingJobsVideoStream

func (list *SubmitEditingJobsVideoStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitEditingJobsVideoStream)
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

type SubmitEditingJobsAudioStreamList []SubmitEditingJobsAudioStream

func (list *SubmitEditingJobsAudioStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitEditingJobsAudioStream)
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

type SubmitEditingJobsSubtitleStreamList []SubmitEditingJobsSubtitleStream

func (list *SubmitEditingJobsSubtitleStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitEditingJobsSubtitleStream)
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

type SubmitEditingJobsSubtitleList []SubmitEditingJobsSubtitle

func (list *SubmitEditingJobsSubtitleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitEditingJobsSubtitle)
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

type SubmitEditingJobsExtSubtitleList []SubmitEditingJobsExtSubtitle

func (list *SubmitEditingJobsExtSubtitleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitEditingJobsExtSubtitle)
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

type SubmitEditingJobsClip3List []SubmitEditingJobsClip3

func (list *SubmitEditingJobsClip3List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitEditingJobsClip3)
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

type SubmitEditingJobsEffectList []SubmitEditingJobsEffect

func (list *SubmitEditingJobsEffectList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitEditingJobsEffect)
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

type SubmitEditingJobsTrackList []SubmitEditingJobsTrack

func (list *SubmitEditingJobsTrackList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitEditingJobsTrack)
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

type SubmitEditingJobsClip4List []SubmitEditingJobsClip4

func (list *SubmitEditingJobsClip4List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitEditingJobsClip4)
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

type SubmitEditingJobsJobResultList []SubmitEditingJobsJobResult

func (list *SubmitEditingJobsJobResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitEditingJobsJobResult)
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

func (c *MtsClient) QueryCensorPipelineList(req *QueryCensorPipelineListArgs) (resp *QueryCensorPipelineListResponse, err error) {
	resp = &QueryCensorPipelineListResponse{}
	err = c.Request("QueryCensorPipelineList", req, resp)
	return
}

type QueryCensorPipelineListPipeline struct {
	Id           string
	Name         string
	State        string
	Priority     string
	NotifyConfig QueryCensorPipelineListNotifyConfig
}

type QueryCensorPipelineListNotifyConfig struct {
	Topic string
	Queue string
}
type QueryCensorPipelineListArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	PipelineIds          string
	OwnerAccount         string
	OwnerId              int64
}
type QueryCensorPipelineListResponse struct {
	RequestId    string
	PipelineList QueryCensorPipelineListPipelineList
	NonExistIds  QueryCensorPipelineListNonExistIdList
}

type QueryCensorPipelineListPipelineList []QueryCensorPipelineListPipeline

func (list *QueryCensorPipelineListPipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCensorPipelineListPipeline)
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

type QueryCensorPipelineListNonExistIdList []string

func (list *QueryCensorPipelineListNonExistIdList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) SubmitAnalysisJob(req *SubmitAnalysisJobArgs) (resp *SubmitAnalysisJobResponse, err error) {
	resp = &SubmitAnalysisJobResponse{}
	err = c.Request("SubmitAnalysisJob", req, resp)
	return
}

type SubmitAnalysisJobAnalysisJob struct {
	Id               string
	UserData         string
	State            string
	Code             string
	Message          string
	Percent          int64
	CreationTime     string
	PipelineId       string
	Priority         string
	TemplateList     SubmitAnalysisJobTemplateList
	InputFile        SubmitAnalysisJobInputFile
	AnalysisConfig   SubmitAnalysisJobAnalysisConfig
	MNSMessageResult SubmitAnalysisJobMNSMessageResult
}

type SubmitAnalysisJobTemplate struct {
	Id          string
	Name        string
	State       string
	Container   SubmitAnalysisJobContainer
	Video       SubmitAnalysisJobVideo
	Audio       SubmitAnalysisJobAudio
	TransConfig SubmitAnalysisJobTransConfig
	MuxConfig   SubmitAnalysisJobMuxConfig
}

type SubmitAnalysisJobContainer struct {
	Format string
}

type SubmitAnalysisJobVideo struct {
	Codec      string
	Profile    string
	Bitrate    string
	Crf        string
	Width      string
	Height     string
	Fps        string
	Gop        string
	Preset     string
	ScanMode   string
	Bufsize    string
	Maxrate    string
	PixFmt     string
	Degrain    string
	Qscale     string
	BitrateBnd SubmitAnalysisJobBitrateBnd
}

type SubmitAnalysisJobBitrateBnd struct {
	Max string
	Min string
}

type SubmitAnalysisJobAudio struct {
	Codec      string
	Profile    string
	Samplerate string
	Bitrate    string
	Channels   string
	Qscale     string
}

type SubmitAnalysisJobTransConfig struct {
	TransMode string
}

type SubmitAnalysisJobMuxConfig struct {
	Segment SubmitAnalysisJobSegment
	Gif     SubmitAnalysisJobGif
}

type SubmitAnalysisJobSegment struct {
	Duration string
}

type SubmitAnalysisJobGif struct {
	Loop       string
	FinalDelay string
}

type SubmitAnalysisJobInputFile struct {
	Bucket   string
	Location string
	Object   string
}

type SubmitAnalysisJobAnalysisConfig struct {
	QualityControl    SubmitAnalysisJobQualityControl
	PropertiesControl SubmitAnalysisJobPropertiesControl
}

type SubmitAnalysisJobQualityControl struct {
	RateQuality     string
	MethodStreaming string
}

type SubmitAnalysisJobPropertiesControl struct {
	Deinterlace string
	Crop        SubmitAnalysisJobCrop
}

type SubmitAnalysisJobCrop struct {
	Mode   string
	Width  string
	Height string
	Top    string
	Left   string
}

type SubmitAnalysisJobMNSMessageResult struct {
	MessageId    string
	ErrorMessage string
	ErrorCode    string
}
type SubmitAnalysisJobArgs struct {
	Input                string
	UserData             string
	ResourceOwnerId      int64
	AnalysisConfig       string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Priority             string
	PipelineId           string
}
type SubmitAnalysisJobResponse struct {
	RequestId   string
	AnalysisJob SubmitAnalysisJobAnalysisJob
}

type SubmitAnalysisJobTemplateList []SubmitAnalysisJobTemplate

func (list *SubmitAnalysisJobTemplateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitAnalysisJobTemplate)
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

func (c *MtsClient) ListCoverPipeline(req *ListCoverPipelineArgs) (resp *ListCoverPipelineResponse, err error) {
	resp = &ListCoverPipelineResponse{}
	err = c.Request("ListCoverPipeline", req, resp)
	return
}

type ListCoverPipelinePipeline struct {
	UserId       int64
	PipelineId   string
	Name         string
	State        string
	Priority     string
	QuotaNum     int
	QuotaUsed    int
	NotifyConfig string
	Role         string
	ExtendConfig string
}
type ListCoverPipelineArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int64
	State                string
	OwnerId              int64
	PageNumber           int64
}
type ListCoverPipelineResponse struct {
	RequestId    string
	TotalCount   int64
	PageNumber   int64
	PageSize     int64
	PipelineList ListCoverPipelinePipelineList
}

type ListCoverPipelinePipelineList []ListCoverPipelinePipeline

func (list *ListCoverPipelinePipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListCoverPipelinePipeline)
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

func (c *MtsClient) UpdateCensorPipeline(req *UpdateCensorPipelineArgs) (resp *UpdateCensorPipelineResponse, err error) {
	resp = &UpdateCensorPipelineResponse{}
	err = c.Request("UpdateCensorPipeline", req, resp)
	return
}

type UpdateCensorPipelinePipeline struct {
	Id           string
	Name         string
	State        string
	Priority     int
	NotifyConfig UpdateCensorPipelineNotifyConfig
}

type UpdateCensorPipelineNotifyConfig struct {
	Topic string
	Queue string
}
type UpdateCensorPipelineArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	Name                 string
	State                string
	NotifyConfig         string
	OwnerId              int64
	Priority             int
	PipelineId           string
}
type UpdateCensorPipelineResponse struct {
	RequestId string
	Pipeline  UpdateCensorPipelinePipeline
}

func (c *MtsClient) DeleteCategory(req *DeleteCategoryArgs) (resp *DeleteCategoryResponse, err error) {
	resp = &DeleteCategoryResponse{}
	err = c.Request("DeleteCategory", req, resp)
	return
}

type DeleteCategoryArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	CateId               int64
	OwnerAccount         string
	OwnerId              int64
}
type DeleteCategoryResponse struct {
	RequestId string
}

func (c *MtsClient) UnbindOutputBucket(req *UnbindOutputBucketArgs) (resp *UnbindOutputBucketResponse, err error) {
	resp = &UnbindOutputBucketResponse{}
	err = c.Request("UnbindOutputBucket", req, resp)
	return
}

type UnbindOutputBucketArgs struct {
	Bucket               string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type UnbindOutputBucketResponse struct {
	RequestId string
}

func (c *MtsClient) QueryAsrJobList(req *QueryAsrJobListArgs) (resp *QueryAsrJobListResponse, err error) {
	resp = &QueryAsrJobListResponse{}
	err = c.Request("QueryAsrJobList", req, resp)
	return
}

type QueryAsrJobListJob struct {
	Id           string
	UserData     string
	PipelineId   string
	State        string
	Code         string
	Message      string
	CreationTime string
	Input        QueryAsrJobListInput
	AsrConfig    QueryAsrJobListAsrConfig
	AsrResult    QueryAsrJobListAsrResult
}

type QueryAsrJobListInput struct {
	Bucket   string
	Location string
	Object   string
}

type QueryAsrJobListAsrConfig struct {
	Scene string
}

type QueryAsrJobListAsrResult struct {
	Duration    string
	AsrTextList QueryAsrJobListAsrTextList
}

type QueryAsrJobListAsrText struct {
	StartTime  int
	EndTime    string
	ChannelId  string
	SpeechRate string
	Text       string
}
type QueryAsrJobListArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	JobIds               string
	OwnerAccount         string
	OwnerId              int64
}
type QueryAsrJobListResponse struct {
	RequestId   string
	JobList     QueryAsrJobListJobList
	NonExistIds QueryAsrJobListNonExistIdList
}

type QueryAsrJobListAsrTextList []QueryAsrJobListAsrText

func (list *QueryAsrJobListAsrTextList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAsrJobListAsrText)
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

type QueryAsrJobListJobList []QueryAsrJobListJob

func (list *QueryAsrJobListJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAsrJobListJob)
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

type QueryAsrJobListNonExistIdList []string

func (list *QueryAsrJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) SubmitTagJob(req *SubmitTagJobArgs) (resp *SubmitTagJobResponse, err error) {
	resp = &SubmitTagJobResponse{}
	err = c.Request("SubmitTagJob", req, resp)
	return
}

type SubmitTagJobArgs struct {
	Input                string
	UserData             string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	TagConfig            string
	OwnerId              int64
	PipelineId           string
}
type SubmitTagJobResponse struct {
	RequestId string
	JobId     string
}

func (c *MtsClient) UpdateWaterMarkTemplate(req *UpdateWaterMarkTemplateArgs) (resp *UpdateWaterMarkTemplateResponse, err error) {
	resp = &UpdateWaterMarkTemplateResponse{}
	err = c.Request("UpdateWaterMarkTemplate", req, resp)
	return
}

type UpdateWaterMarkTemplateWaterMarkTemplate struct {
	Id         string
	Name       string
	Width      string
	Height     string
	Dx         string
	Dy         string
	ReferPos   string
	Type       string
	State      string
	Timeline   UpdateWaterMarkTemplateTimeline
	RatioRefer UpdateWaterMarkTemplateRatioRefer
}

type UpdateWaterMarkTemplateTimeline struct {
	Start    string
	Duration string
}

type UpdateWaterMarkTemplateRatioRefer struct {
	Dx     string
	Dy     string
	Width  string
	Height string
}
type UpdateWaterMarkTemplateArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	Name                 string
	OwnerId              int64
	WaterMarkTemplateId  string
	Config               string
}
type UpdateWaterMarkTemplateResponse struct {
	RequestId         string
	WaterMarkTemplate UpdateWaterMarkTemplateWaterMarkTemplate
}

func (c *MtsClient) ListAsrPipeline(req *ListAsrPipelineArgs) (resp *ListAsrPipelineResponse, err error) {
	resp = &ListAsrPipelineResponse{}
	err = c.Request("ListAsrPipeline", req, resp)
	return
}

type ListAsrPipelinePipeline struct {
	Id           string
	Name         string
	State        string
	Priority     string
	NotifyConfig ListAsrPipelineNotifyConfig
}

type ListAsrPipelineNotifyConfig struct {
	Topic     string
	QueueName string
}
type ListAsrPipelineArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int64
	State                string
	OwnerId              int64
	PageNumber           int64
}
type ListAsrPipelineResponse struct {
	RequestId    string
	TotalCount   int64
	PageNumber   int64
	PageSize     int64
	PipelineList ListAsrPipelinePipelineList
}

type ListAsrPipelinePipelineList []ListAsrPipelinePipeline

func (list *ListAsrPipelinePipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAsrPipelinePipeline)
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

func (c *MtsClient) QueryMediaList(req *QueryMediaListArgs) (resp *QueryMediaListResponse, err error) {
	resp = &QueryMediaListResponse{}
	err = c.Request("QueryMediaList", req, resp)
	return
}

type QueryMediaListMedia struct {
	MediaId      string
	Title        string
	Description  string
	CoverURL     string
	CateId       int64
	Duration     string
	Format       string
	Size         string
	Bitrate      string
	Width        string
	Height       string
	Fps          string
	PublishState string
	CreationTime string
	PlayList     QueryMediaListPlayList
	SnapshotList QueryMediaListSnapshotList
	Tags         QueryMediaListTagList
	RunIdList    QueryMediaListRunIdListList
	File         QueryMediaListFile
	MediaInfo    QueryMediaListMediaInfo
}

type QueryMediaListPlay struct {
	ActivityName      string
	MediaWorkflowId   string
	MediaWorkflowName string
	Duration          string
	Format            string
	Size              string
	Bitrate           string
	Width             string
	Height            string
	Fps               string
	Encryption        string
	File1             QueryMediaListFile1
}

type QueryMediaListFile1 struct {
	URL   string
	State string
}

type QueryMediaListSnapshot struct {
	Type              string
	MediaWorkflowId   string
	MediaWorkflowName string
	ActivityName      string
	Count             string
	File2             QueryMediaListFile2
}

type QueryMediaListFile2 struct {
	URL   string
	State string
}

type QueryMediaListFile struct {
	URL   string
	State string
}

type QueryMediaListMediaInfo struct {
	Streams QueryMediaListStreams
	Format  QueryMediaListFormat
}

type QueryMediaListStreams struct {
	VideoStreamList    QueryMediaListVideoStreamList
	AudioStreamList    QueryMediaListAudioStreamList
	SubtitleStreamList QueryMediaListSubtitleStreamList
}

type QueryMediaListVideoStream struct {
	Index          string
	CodecName      string
	CodecLongName  string
	Profile        string
	CodecTimeBase  string
	CodecTagString string
	CodecTag       string
	Width          string
	Height         string
	HasBFrames     string
	Sar            string
	Dar            string
	PixFmt         string
	Level          string
	Fps            string
	AvgFPS         string
	Timebase       string
	StartTime      string
	Duration       string
	Bitrate        string
	NumFrames      string
	Lang           string
	Rotate         string
	NetworkCost    QueryMediaListNetworkCost
}

type QueryMediaListNetworkCost struct {
	PreloadTime   string
	CostBandwidth string
	AvgBitrate    string
}

type QueryMediaListAudioStream struct {
	Index          string
	CodecName      string
	CodecTimeBase  string
	CodecLongName  string
	CodecTagString string
	CodecTag       string
	SampleFmt      string
	Samplerate     string
	Channels       string
	ChannelLayout  string
	Timebase       string
	StartTime      string
	Duration       string
	Bitrate        string
	NumFrames      string
	Lang           string
}

type QueryMediaListSubtitleStream struct {
	Index string
	Lang  string
}

type QueryMediaListFormat struct {
	NumStreams     string
	NumPrograms    string
	FormatName     string
	FormatLongName string
	StartTime      string
	Duration       string
	Size           string
	Bitrate        string
}
type QueryMediaListArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	IncludeSnapshotList  core.Bool
	OwnerAccount         string
	MediaIds             string
	OwnerId              int64
	IncludePlayList      core.Bool
	IncludeMediaInfo     core.Bool
}
type QueryMediaListResponse struct {
	RequestId        string
	MediaList        QueryMediaListMediaList
	NonExistMediaIds QueryMediaListNonExistMediaIdList
}

type QueryMediaListPlayList []QueryMediaListPlay

func (list *QueryMediaListPlayList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListPlay)
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

type QueryMediaListSnapshotList []QueryMediaListSnapshot

func (list *QueryMediaListSnapshotList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListSnapshot)
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

type QueryMediaListTagList []string

func (list *QueryMediaListTagList) UnmarshalJSON(data []byte) error {
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

type QueryMediaListRunIdListList []string

func (list *QueryMediaListRunIdListList) UnmarshalJSON(data []byte) error {
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

type QueryMediaListVideoStreamList []QueryMediaListVideoStream

func (list *QueryMediaListVideoStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListVideoStream)
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

type QueryMediaListAudioStreamList []QueryMediaListAudioStream

func (list *QueryMediaListAudioStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListAudioStream)
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

type QueryMediaListSubtitleStreamList []QueryMediaListSubtitleStream

func (list *QueryMediaListSubtitleStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListSubtitleStream)
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

type QueryMediaListMediaList []QueryMediaListMedia

func (list *QueryMediaListMediaList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaListMedia)
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

type QueryMediaListNonExistMediaIdList []string

func (list *QueryMediaListNonExistMediaIdList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) QueryWaterMarkTemplateList(req *QueryWaterMarkTemplateListArgs) (resp *QueryWaterMarkTemplateListResponse, err error) {
	resp = &QueryWaterMarkTemplateListResponse{}
	err = c.Request("QueryWaterMarkTemplateList", req, resp)
	return
}

type QueryWaterMarkTemplateListWaterMarkTemplate struct {
	Id         string
	Name       string
	Width      string
	Height     string
	Dx         string
	Dy         string
	ReferPos   string
	Type       string
	State      string
	Timeline   QueryWaterMarkTemplateListTimeline
	RatioRefer QueryWaterMarkTemplateListRatioRefer
}

type QueryWaterMarkTemplateListTimeline struct {
	Start    string
	Duration string
}

type QueryWaterMarkTemplateListRatioRefer struct {
	Dx     string
	Dy     string
	Width  string
	Height string
}
type QueryWaterMarkTemplateListArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	WaterMarkTemplateIds string
}
type QueryWaterMarkTemplateListResponse struct {
	RequestId             string
	WaterMarkTemplateList QueryWaterMarkTemplateListWaterMarkTemplateList
	NonExistWids          QueryWaterMarkTemplateListNonExistWidList
}

type QueryWaterMarkTemplateListWaterMarkTemplateList []QueryWaterMarkTemplateListWaterMarkTemplate

func (list *QueryWaterMarkTemplateListWaterMarkTemplateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryWaterMarkTemplateListWaterMarkTemplate)
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

type QueryWaterMarkTemplateListNonExistWidList []string

func (list *QueryWaterMarkTemplateListNonExistWidList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) AddPipeline(req *AddPipelineArgs) (resp *AddPipelineResponse, err error) {
	resp = &AddPipelineResponse{}
	err = c.Request("AddPipeline", req, resp)
	return
}

type AddPipelinePipeline struct {
	Id           string
	Name         string
	State        string
	Speed        string
	SpeedLevel   int64
	Role         string
	NotifyConfig AddPipelineNotifyConfig
}

type AddPipelineNotifyConfig struct {
	Topic     string
	QueueName string
}
type AddPipelineArgs struct {
	ResourceOwnerId      int64
	Role                 string
	ResourceOwnerAccount string
	OwnerAccount         string
	Name                 string
	NotifyConfig         string
	OwnerId              int64
	SpeedLevel           int64
	Speed                string
}
type AddPipelineResponse struct {
	RequestId string
	Pipeline  AddPipelinePipeline
}

func (c *MtsClient) QueryAuthConfig(req *QueryAuthConfigArgs) (resp *QueryAuthConfigResponse, err error) {
	resp = &QueryAuthConfigResponse{}
	err = c.Request("QueryAuthConfig", req, resp)
	return
}

type QueryAuthConfigArgs struct {
	ResourceOwnerId      string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              string
}
type QueryAuthConfigResponse struct {
	RequestId string
	Key1      string
	Key2      string
}

func (c *MtsClient) UpdatePipeline(req *UpdatePipelineArgs) (resp *UpdatePipelineResponse, err error) {
	resp = &UpdatePipelineResponse{}
	err = c.Request("UpdatePipeline", req, resp)
	return
}

type UpdatePipelinePipeline struct {
	Id           string
	Name         string
	State        string
	Speed        string
	Role         string
	NotifyConfig UpdatePipelineNotifyConfig
}

type UpdatePipelineNotifyConfig struct {
	Topic     string
	QueueName string
}
type UpdatePipelineArgs struct {
	ResourceOwnerId      int64
	Role                 string
	ResourceOwnerAccount string
	OwnerAccount         string
	Name                 string
	State                string
	NotifyConfig         string
	OwnerId              int64
	PipelineId           string
}
type UpdatePipelineResponse struct {
	RequestId string
	Pipeline  UpdatePipelinePipeline
}

func (c *MtsClient) DeleteWaterMarkTemplate(req *DeleteWaterMarkTemplateArgs) (resp *DeleteWaterMarkTemplateResponse, err error) {
	resp = &DeleteWaterMarkTemplateResponse{}
	err = c.Request("DeleteWaterMarkTemplate", req, resp)
	return
}

type DeleteWaterMarkTemplateArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	WaterMarkTemplateId  string
}
type DeleteWaterMarkTemplateResponse struct {
	RequestId           string
	WaterMarkTemplateId string
}

func (c *MtsClient) SearchPipeline(req *SearchPipelineArgs) (resp *SearchPipelineResponse, err error) {
	resp = &SearchPipelineResponse{}
	err = c.Request("SearchPipeline", req, resp)
	return
}

type SearchPipelinePipeline struct {
	Id           string
	Name         string
	State        string
	Speed        string
	SpeedLevel   int64
	Role         string
	NotifyConfig SearchPipelineNotifyConfig
}

type SearchPipelineNotifyConfig struct {
	Topic     string
	QueueName string
}
type SearchPipelineArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int64
	State                string
	OwnerId              int64
	PageNumber           int64
}
type SearchPipelineResponse struct {
	RequestId    string
	TotalCount   int64
	PageNumber   int64
	PageSize     int64
	PipelineList SearchPipelinePipelineList
}

type SearchPipelinePipelineList []SearchPipelinePipeline

func (list *SearchPipelinePipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SearchPipelinePipeline)
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

func (c *MtsClient) SetAuthConfig(req *SetAuthConfigArgs) (resp *SetAuthConfigResponse, err error) {
	resp = &SetAuthConfigResponse{}
	err = c.Request("SetAuthConfig", req, resp)
	return
}

type SetAuthConfigArgs struct {
	Key1                 string
	Key2                 string
	ResourceOwnerId      string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              string
}
type SetAuthConfigResponse struct {
	RequestId string
	Key1      string
	Key2      string
}

func (c *MtsClient) SubmitTerrorismJob(req *SubmitTerrorismJobArgs) (resp *SubmitTerrorismJobResponse, err error) {
	resp = &SubmitTerrorismJobResponse{}
	err = c.Request("SubmitTerrorismJob", req, resp)
	return
}

type SubmitTerrorismJobArgs struct {
	Input                string
	UserData             string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	PipelineId           string
	TerrorismConfig      string
}
type SubmitTerrorismJobResponse struct {
	RequestId string
	JobId     string
}

func (c *MtsClient) ReportTagJobResult(req *ReportTagJobResultArgs) (resp *ReportTagJobResultResponse, err error) {
	resp = &ReportTagJobResultResponse{}
	err = c.Request("ReportTagJobResult", req, resp)
	return
}

type ReportTagJobResultArgs struct {
	Result               string
	JobId                string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	Tag                  string
	OwnerId              int64
}
type ReportTagJobResultResponse struct {
	RequestId string
	JobId     string
}

func (c *MtsClient) AddAsrPipeline(req *AddAsrPipelineArgs) (resp *AddAsrPipelineResponse, err error) {
	resp = &AddAsrPipelineResponse{}
	err = c.Request("AddAsrPipeline", req, resp)
	return
}

type AddAsrPipelinePipeline struct {
	Id           string
	Name         string
	Priority     int
	State        string
	NotifyConfig AddAsrPipelineNotifyConfig
}

type AddAsrPipelineNotifyConfig struct {
	Topic string
	Queue string
}
type AddAsrPipelineArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	Name                 string
	NotifyConfig         string
	OwnerId              int64
	Priority             int
}
type AddAsrPipelineResponse struct {
	RequestId string
	Pipeline  AddAsrPipelinePipeline
}

func (c *MtsClient) SearchWaterMarkTemplate(req *SearchWaterMarkTemplateArgs) (resp *SearchWaterMarkTemplateResponse, err error) {
	resp = &SearchWaterMarkTemplateResponse{}
	err = c.Request("SearchWaterMarkTemplate", req, resp)
	return
}

type SearchWaterMarkTemplateWaterMarkTemplate struct {
	Id         string
	Name       string
	Width      string
	Height     string
	Dx         string
	Dy         string
	ReferPos   string
	Type       string
	State      string
	Timeline   SearchWaterMarkTemplateTimeline
	RatioRefer SearchWaterMarkTemplateRatioRefer
}

type SearchWaterMarkTemplateTimeline struct {
	Start    string
	Duration string
}

type SearchWaterMarkTemplateRatioRefer struct {
	Dx     string
	Dy     string
	Width  string
	Height string
}
type SearchWaterMarkTemplateArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int64
	State                string
	OwnerId              int64
	PageNumber           int64
}
type SearchWaterMarkTemplateResponse struct {
	RequestId             string
	TotalCount            int64
	PageNumber            int64
	PageSize              int64
	WaterMarkTemplateList SearchWaterMarkTemplateWaterMarkTemplateList
}

type SearchWaterMarkTemplateWaterMarkTemplateList []SearchWaterMarkTemplateWaterMarkTemplate

func (list *SearchWaterMarkTemplateWaterMarkTemplateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SearchWaterMarkTemplateWaterMarkTemplate)
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

func (c *MtsClient) QueryPornJobList(req *QueryPornJobListArgs) (resp *QueryPornJobListResponse, err error) {
	resp = &QueryPornJobListResponse{}
	err = c.Request("QueryPornJobList", req, resp)
	return
}

type QueryPornJobListPornJob struct {
	Id               string
	UserData         string
	PipelineId       string
	State            string
	Code             string
	Message          string
	CreationTime     string
	Input            QueryPornJobListInput
	PornConfig       QueryPornJobListPornConfig
	CensorPornResult QueryPornJobListCensorPornResult
}

type QueryPornJobListInput struct {
	Bucket   string
	Location string
	Object   string
}

type QueryPornJobListPornConfig struct {
	Interval   string
	BizType    string
	OutputFile QueryPornJobListOutputFile
}

type QueryPornJobListOutputFile struct {
	Bucket   string
	Location string
	Object   string
}

type QueryPornJobListCensorPornResult struct {
	Label           string
	Suggestion      string
	MaxScore        string
	AverageScore    string
	PornCounterList QueryPornJobListCounterList
	PornTopList     QueryPornJobListTopList
}

type QueryPornJobListCounter struct {
	Count int
	Label string
}

type QueryPornJobListTop struct {
	Label     string
	Score     string
	Timestamp string
	Index     string
	Object    string
}
type QueryPornJobListArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	JobIds               string
	OwnerAccount         string
	OwnerId              int64
}
type QueryPornJobListResponse struct {
	RequestId   string
	PornJobList QueryPornJobListPornJobList
	NonExistIds QueryPornJobListNonExistIdList
}

type QueryPornJobListCounterList []QueryPornJobListCounter

func (list *QueryPornJobListCounterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPornJobListCounter)
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

type QueryPornJobListTopList []QueryPornJobListTop

func (list *QueryPornJobListTopList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPornJobListTop)
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

type QueryPornJobListPornJobList []QueryPornJobListPornJob

func (list *QueryPornJobListPornJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPornJobListPornJob)
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

type QueryPornJobListNonExistIdList []string

func (list *QueryPornJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) UpdateVideoSummaryPipeline(req *UpdateVideoSummaryPipelineArgs) (resp *UpdateVideoSummaryPipelineResponse, err error) {
	resp = &UpdateVideoSummaryPipelineResponse{}
	err = c.Request("UpdateVideoSummaryPipeline", req, resp)
	return
}

type UpdateVideoSummaryPipelinePipeline struct {
	Id           string
	Name         string
	State        string
	Priority     int
	NotifyConfig UpdateVideoSummaryPipelineNotifyConfig
}

type UpdateVideoSummaryPipelineNotifyConfig struct {
	Topic     string
	QueueName string
}
type UpdateVideoSummaryPipelineArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	Name                 string
	State                string
	NotifyConfig         string
	OwnerId              int64
	Priority             int
	PipelineId           string
}
type UpdateVideoSummaryPipelineResponse struct {
	RequestId string
	Pipeline  UpdateVideoSummaryPipelinePipeline
}

func (c *MtsClient) ReportFacerecogJobResult(req *ReportFacerecogJobResultArgs) (resp *ReportFacerecogJobResultResponse, err error) {
	resp = &ReportFacerecogJobResultResponse{}
	err = c.Request("ReportFacerecogJobResult", req, resp)
	return
}

type ReportFacerecogJobResultArgs struct {
	JobId                string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	Facerecog            string
	OwnerAccount         string
	Details              string
	OwnerId              int64
}
type ReportFacerecogJobResultResponse struct {
	RequestId string
	JobId     string
}

func (c *MtsClient) DeleteMediaTag(req *DeleteMediaTagArgs) (resp *DeleteMediaTagResponse, err error) {
	resp = &DeleteMediaTagResponse{}
	err = c.Request("DeleteMediaTag", req, resp)
	return
}

type DeleteMediaTagArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	Tag                  string
	OwnerId              int64
	MediaId              string
}
type DeleteMediaTagResponse struct {
	RequestId string
}

func (c *MtsClient) SubmitSnapshotJob(req *SubmitSnapshotJobArgs) (resp *SubmitSnapshotJobResponse, err error) {
	resp = &SubmitSnapshotJobResponse{}
	err = c.Request("SubmitSnapshotJob", req, resp)
	return
}

type SubmitSnapshotJobSnapshotJob struct {
	Id               string
	UserData         string
	PipelineId       string
	State            string
	Code             string
	Count            string
	TileCount        string
	Message          string
	CreationTime     string
	Input            SubmitSnapshotJobInput
	SnapshotConfig   SubmitSnapshotJobSnapshotConfig
	MNSMessageResult SubmitSnapshotJobMNSMessageResult
}

type SubmitSnapshotJobInput struct {
	Bucket   string
	Location string
	Object   string
	RoleArn  string
}

type SubmitSnapshotJobSnapshotConfig struct {
	Time           string
	Interval       string
	Num            string
	Width          string
	Height         string
	FrameType      string
	OutputFile     SubmitSnapshotJobOutputFile
	TileOutputFile SubmitSnapshotJobTileOutputFile
	TileOut        SubmitSnapshotJobTileOut
}

type SubmitSnapshotJobOutputFile struct {
	Bucket   string
	Location string
	Object   string
	RoleArn  string
}

type SubmitSnapshotJobTileOutputFile struct {
	Bucket   string
	Location string
	Object   string
	RoleArn  string
}

type SubmitSnapshotJobTileOut struct {
	Lines         string
	Columns       string
	CellWidth     string
	CellHeight    string
	Margin        string
	Padding       string
	Color         string
	IsKeepCellPic string
}

type SubmitSnapshotJobMNSMessageResult struct {
	MessageId    string
	ErrorMessage string
	ErrorCode    string
}
type SubmitSnapshotJobArgs struct {
	Input                string
	UserData             string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	SnapshotConfig       string
	PipelineId           string
}
type SubmitSnapshotJobResponse struct {
	RequestId   string
	SnapshotJob SubmitSnapshotJobSnapshotJob
}

func (c *MtsClient) UpdatePornPipeline(req *UpdatePornPipelineArgs) (resp *UpdatePornPipelineResponse, err error) {
	resp = &UpdatePornPipelineResponse{}
	err = c.Request("UpdatePornPipeline", req, resp)
	return
}

type UpdatePornPipelinePipeline struct {
	Id           string
	Name         string
	State        string
	Priority     int
	NotifyConfig UpdatePornPipelineNotifyConfig
}

type UpdatePornPipelineNotifyConfig struct {
	Topic string
	Queue string
}
type UpdatePornPipelineArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	Name                 string
	State                string
	NotifyConfig         string
	OwnerId              int64
	Priority             int
	PipelineId           string
}
type UpdatePornPipelineResponse struct {
	RequestId string
	Pipeline  UpdatePornPipelinePipeline
}

func (c *MtsClient) QueryJobList(req *QueryJobListArgs) (resp *QueryJobListResponse, err error) {
	resp = &QueryJobListResponse{}
	err = c.Request("QueryJobList", req, resp)
	return
}

type QueryJobListJob struct {
	JobId            string
	State            string
	Code             string
	Message          string
	Percent          int64
	PipelineId       string
	CreationTime     string
	FinishTime       string
	Input            QueryJobListInput
	Output           QueryJobListOutput
	MNSMessageResult QueryJobListMNSMessageResult
}

type QueryJobListInput struct {
	Bucket   string
	Location string
	Object   string
}

type QueryJobListOutput struct {
	TemplateId             string
	UserData               string
	Rotate                 string
	VideoStreamMap         string
	AudioStreamMap         string
	DeWatermark            string
	Priority               string
	WaterMarkConfigUrl     string
	MergeConfigUrl         string
	WaterMarkList          QueryJobListWaterMarkList
	MergeList              QueryJobListMergeList
	OpeningList            QueryJobListOpeningList
	TailSlateList          QueryJobListTailSlateList
	OutputFile             QueryJobListOutputFile
	M3U8NonStandardSupport QueryJobListM3U8NonStandardSupport
	Properties             QueryJobListProperties
	Clip                   QueryJobListClip
	SuperReso              QueryJobListSuperReso
	SubtitleConfig         QueryJobListSubtitleConfig
	TransConfig            QueryJobListTransConfig
	MuxConfig              QueryJobListMuxConfig
	Audio                  QueryJobListAudio
	Video                  QueryJobListVideo
	Container              QueryJobListContainer
	Encryption             QueryJobListEncryption
}

type QueryJobListWaterMark struct {
	WaterMarkTemplateId string
	Width               string
	Height              string
	Dx                  string
	Dy                  string
	ReferPos            string
	Type                string
	InputFile           QueryJobListInputFile
}

type QueryJobListInputFile struct {
	Bucket   string
	Location string
	Object   string
}

type QueryJobListMerge struct {
	MergeURL string
	Start    string
	Duration string
	RoleArn  string
}

type QueryJobListOpening struct {
	OpenUrl string
	Start   string
	Width   string
	Height  string
}

type QueryJobListTailSlate struct {
	TailUrl       string
	Start         string
	BlendDuration string
	Width         string
	Height        string
	IsMergeAudio  core.Bool
	BgColor       string
}

type QueryJobListOutputFile struct {
	Bucket   string
	Location string
	Object   string
	RoleArn  string
}

type QueryJobListM3U8NonStandardSupport struct {
	TS QueryJobListTS
}

type QueryJobListTS struct {
	Md5Support  core.Bool
	SizeSupport core.Bool
}

type QueryJobListProperties struct {
	Width       string
	Height      string
	Bitrate     string
	Duration    string
	Fps         string
	FileSize    string
	FileFormat  string
	SourceLogos QueryJobListSourceLogoList
	Streams     QueryJobListStreams
	Format      QueryJobListFormat
}

type QueryJobListSourceLogo struct {
	Source string
}

type QueryJobListStreams struct {
	VideoStreamList    QueryJobListVideoStreamList
	AudioStreamList    QueryJobListAudioStreamList
	SubtitleStreamList QueryJobListSubtitleStreamList
}

type QueryJobListVideoStream struct {
	Index          string
	CodecName      string
	CodecLongName  string
	Profile        string
	CodecTimeBase  string
	CodecTagString string
	CodecTag       string
	Width          string
	Height         string
	HasBFrames     string
	Sar            string
	Dar            string
	PixFmt         string
	Level          string
	Fps            string
	AvgFPS         string
	Timebase       string
	StartTime      string
	Duration       string
	Bitrate        string
	NumFrames      string
	Lang           string
	NetworkCost    QueryJobListNetworkCost
}

type QueryJobListNetworkCost struct {
	PreloadTime   string
	CostBandwidth string
	AvgBitrate    string
}

type QueryJobListAudioStream struct {
	Index          string
	CodecName      string
	CodecTimeBase  string
	CodecLongName  string
	CodecTagString string
	CodecTag       string
	SampleFmt      string
	Samplerate     string
	Channels       string
	ChannelLayout  string
	Timebase       string
	StartTime      string
	Duration       string
	Bitrate        string
	NumFrames      string
	Lang           string
}

type QueryJobListSubtitleStream struct {
	Index string
	Lang  string
}

type QueryJobListFormat struct {
	NumStreams     string
	NumPrograms    string
	FormatName     string
	FormatLongName string
	StartTime      string
	Duration       string
	Size           string
	Bitrate        string
}

type QueryJobListClip struct {
	TimeSpan QueryJobListTimeSpan
}

type QueryJobListTimeSpan struct {
	Seek     string
	Duration string
}

type QueryJobListSuperReso struct {
	IsHalfSample string
}

type QueryJobListSubtitleConfig struct {
	SubtitleList    QueryJobListSubtitleList
	ExtSubtitleList QueryJobListExtSubtitleList
}

type QueryJobListSubtitle struct {
	Map string
}

type QueryJobListExtSubtitle struct {
	FontName string
	CharEnc  string
	Input1   QueryJobListInput1
}

type QueryJobListInput1 struct {
	Bucket   string
	Location string
	Object   string
}

type QueryJobListTransConfig struct {
	TransMode               string
	IsCheckReso             string
	IsCheckResoFail         string
	IsCheckVideoBitrate     string
	IsCheckAudioBitrate     string
	AdjDarMethod            string
	IsCheckVideoBitrateFail string
	IsCheckAudioBitrateFail string
}

type QueryJobListMuxConfig struct {
	Segment QueryJobListSegment
	Gif     QueryJobListGif
}

type QueryJobListSegment struct {
	Duration string
}

type QueryJobListGif struct {
	Loop            string
	FinalDelay      string
	IsCustomPalette string
	DitherMode      string
}

type QueryJobListAudio struct {
	Codec      string
	Profile    string
	Samplerate string
	Bitrate    string
	Channels   string
	Qscale     string
	Volume     QueryJobListVolume
}

type QueryJobListVolume struct {
	Level  string
	Method string
}

type QueryJobListVideo struct {
	Codec      string
	Profile    string
	Bitrate    string
	Crf        string
	Width      string
	Height     string
	Fps        string
	Gop        string
	Preset     string
	ScanMode   string
	Bufsize    string
	Maxrate    string
	PixFmt     string
	Degrain    string
	Qscale     string
	Crop       string
	Pad        string
	MaxFps     string
	BitrateBnd QueryJobListBitrateBnd
}

type QueryJobListBitrateBnd struct {
	Max string
	Min string
}

type QueryJobListContainer struct {
	Format string
}

type QueryJobListEncryption struct {
	Type    string
	Id      string
	Key     string
	KeyUri  string
	KeyType string
	SkipCnt string
}

type QueryJobListMNSMessageResult struct {
	MessageId    string
	ErrorMessage string
	ErrorCode    string
}
type QueryJobListArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	JobIds               string
	OwnerAccount         string
	OwnerId              int64
}
type QueryJobListResponse struct {
	RequestId      string
	JobList        QueryJobListJobList
	NonExistJobIds QueryJobListNonExistJobIdList
}

type QueryJobListWaterMarkList []QueryJobListWaterMark

func (list *QueryJobListWaterMarkList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListWaterMark)
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

type QueryJobListMergeList []QueryJobListMerge

func (list *QueryJobListMergeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListMerge)
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

type QueryJobListOpeningList []QueryJobListOpening

func (list *QueryJobListOpeningList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListOpening)
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

type QueryJobListTailSlateList []QueryJobListTailSlate

func (list *QueryJobListTailSlateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListTailSlate)
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

type QueryJobListSourceLogoList []QueryJobListSourceLogo

func (list *QueryJobListSourceLogoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListSourceLogo)
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

type QueryJobListVideoStreamList []QueryJobListVideoStream

func (list *QueryJobListVideoStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListVideoStream)
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

type QueryJobListAudioStreamList []QueryJobListAudioStream

func (list *QueryJobListAudioStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListAudioStream)
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

type QueryJobListSubtitleStreamList []QueryJobListSubtitleStream

func (list *QueryJobListSubtitleStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListSubtitleStream)
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

type QueryJobListSubtitleList []QueryJobListSubtitle

func (list *QueryJobListSubtitleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListSubtitle)
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

type QueryJobListExtSubtitleList []QueryJobListExtSubtitle

func (list *QueryJobListExtSubtitleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListExtSubtitle)
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

type QueryJobListJobList []QueryJobListJob

func (list *QueryJobListJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryJobListJob)
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

type QueryJobListNonExistJobIdList []string

func (list *QueryJobListNonExistJobIdList) UnmarshalJSON(data []byte) error {
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

func (c *MtsClient) ListMediaWorkflowExecutions(req *ListMediaWorkflowExecutionsArgs) (resp *ListMediaWorkflowExecutionsResponse, err error) {
	resp = &ListMediaWorkflowExecutionsResponse{}
	err = c.Request("ListMediaWorkflowExecutions", req, resp)
	return
}

type ListMediaWorkflowExecutionsMediaWorkflowExecution struct {
	RunId           string
	MediaWorkflowId string
	Name            string
	State           string
	MediaId         string
	CreationTime    string
	ActivityList    ListMediaWorkflowExecutionsActivityList
	Input           ListMediaWorkflowExecutionsInput
}

type ListMediaWorkflowExecutionsActivity struct {
	Name             string
	Type             string
	JobId            string
	State            string
	Code             string
	Message          string
	StartTime        string
	EndTime          string
	MNSMessageResult ListMediaWorkflowExecutionsMNSMessageResult
}

type ListMediaWorkflowExecutionsMNSMessageResult struct {
	MessageId    string
	ErrorMessage string
	ErrorCode    string
}

type ListMediaWorkflowExecutionsInput struct {
	UserData  string
	InputFile ListMediaWorkflowExecutionsInputFile
}

type ListMediaWorkflowExecutionsInputFile struct {
	Bucket   string
	Location string
	Object   string
}
type ListMediaWorkflowExecutionsArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	InputFileURL         string
	NextPageToken        string
	OwnerAccount         string
	MaximumPageSize      int64
	MediaWorkflowId      string
	OwnerId              int64
	MediaWorkflowName    string
}
type ListMediaWorkflowExecutionsResponse struct {
	RequestId                  string
	NextPageToken              string
	MediaWorkflowExecutionList ListMediaWorkflowExecutionsMediaWorkflowExecutionList
}

type ListMediaWorkflowExecutionsActivityList []ListMediaWorkflowExecutionsActivity

func (list *ListMediaWorkflowExecutionsActivityList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListMediaWorkflowExecutionsActivity)
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

type ListMediaWorkflowExecutionsMediaWorkflowExecutionList []ListMediaWorkflowExecutionsMediaWorkflowExecution

func (list *ListMediaWorkflowExecutionsMediaWorkflowExecutionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListMediaWorkflowExecutionsMediaWorkflowExecution)
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

func (c *MtsClient) SearchMediaWorkflow(req *SearchMediaWorkflowArgs) (resp *SearchMediaWorkflowResponse, err error) {
	resp = &SearchMediaWorkflowResponse{}
	err = c.Request("SearchMediaWorkflow", req, resp)
	return
}

type SearchMediaWorkflowMediaWorkflow struct {
	MediaWorkflowId string
	Name            string
	Topology        string
	State           string
	CreationTime    string
}
type SearchMediaWorkflowArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int64
	StateList            string
	OwnerId              int64
	PageNumber           int64
}
type SearchMediaWorkflowResponse struct {
	RequestId         string
	TotalCount        int64
	PageNumber        int64
	PageSize          int64
	MediaWorkflowList SearchMediaWorkflowMediaWorkflowList
}

type SearchMediaWorkflowMediaWorkflowList []SearchMediaWorkflowMediaWorkflow

func (list *SearchMediaWorkflowMediaWorkflowList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SearchMediaWorkflowMediaWorkflow)
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

func (c *MtsClient) QueryTerrorismJobList(req *QueryTerrorismJobListArgs) (resp *QueryTerrorismJobListResponse, err error) {
	resp = &QueryTerrorismJobListResponse{}
	err = c.Request("QueryTerrorismJobList", req, resp)
	return
}

type QueryTerrorismJobListTerrorismJob struct {
	Id                    string
	UserData              string
	PipelineId            string
	State                 string
	Code                  string
	Message               string
	CreationTime          string
	Input                 QueryTerrorismJobListInput
	TerrorismConfig       QueryTerrorismJobListTerrorismConfig
	CensorTerrorismResult QueryTerrorismJobListCensorTerrorismResult
}

type QueryTerrorismJobListInput struct {
	Bucket   string
	Location string
	Object   string
}

type QueryTerrorismJobListTerrorismConfig struct {
	Interval   string
	BizType    string
	OutputFile QueryTerrorismJobListOutputFile
}

type QueryTerrorismJobListOutputFile struct {
	Bucket   string
	Location string
	Object   string
}

type QueryTerrorismJobListCensorTerrorismResult struct {
	Label                string
	Suggestion           string
	MaxScore             string
	AverageScore         string
	TerrorismCounterList QueryTerrorismJobListCounterList
	TerrorismTopList     QueryTerrorismJobListTopList
}

type QueryTerrorismJobListCounter struct {
	Count int
	Label string
}

type QueryTerrorismJobListTop struct {
	Label     string
	Score     string
	Timestamp string
	Index     string
	Object    string
}
type QueryTerrorismJobListArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	JobIds               string
	OwnerAccount         string
	OwnerId              int64
}
type QueryTerrorismJobListResponse struct {
	RequestId        string
	TerrorismJobList QueryTerrorismJobListTerrorismJobList
	NonExistIds      QueryTerrorismJobListNonExistIdList
}

type QueryTerrorismJobListCounterList []QueryTerrorismJobListCounter

func (list *QueryTerrorismJobListCounterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTerrorismJobListCounter)
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

type QueryTerrorismJobListTopList []QueryTerrorismJobListTop

func (list *QueryTerrorismJobListTopList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTerrorismJobListTop)
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

type QueryTerrorismJobListTerrorismJobList []QueryTerrorismJobListTerrorismJob

func (list *QueryTerrorismJobListTerrorismJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTerrorismJobListTerrorismJob)
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

type QueryTerrorismJobListNonExistIdList []string

func (list *QueryTerrorismJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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
