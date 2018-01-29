package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubmitEditingJobsRequest struct {
	requests.RpcRequest
	OutputBucket         string `position:"Query" name:"OutputBucket"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	EditingJobOutputs    string `position:"Query" name:"EditingJobOutputs"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OutputLocation       string `position:"Query" name:"OutputLocation"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	EditingInputs        string `position:"Query" name:"EditingInputs"`
	PipelineId           string `position:"Query" name:"PipelineId"`
}

func (req *SubmitEditingJobsRequest) Invoke(client *sdk.Client) (resp *SubmitEditingJobsResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "SubmitEditingJobs", "mts", "")
	resp = &SubmitEditingJobsResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitEditingJobsResponse struct {
	responses.BaseResponse
	RequestId     string
	JobResultList SubmitEditingJobsJobResultList
}

type SubmitEditingJobsJobResult struct {
	Success bool
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
	Md5Support  bool
	SizeSupport bool
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
