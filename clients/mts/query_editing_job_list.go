package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryEditingJobListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string `position:"Query" name:"JobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QueryEditingJobListRequest) Invoke(client *sdk.Client) (resp *QueryEditingJobListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryEditingJobList", "", "")
	resp = &QueryEditingJobListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryEditingJobListResponse struct {
	responses.BaseResponse
	RequestId      string
	JobList        QueryEditingJobListJobList
	NonExistJobIds QueryEditingJobListNonExistJobIdList
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
	Md5Support  bool
	SizeSupport bool
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
