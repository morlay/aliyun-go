package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId     common.String
	JobResultList SubmitEditingJobsJobResultList
}

type SubmitEditingJobsJobResult struct {
	Success bool
	Code    common.String
	Message common.String
	Job     SubmitEditingJobsJob
}

type SubmitEditingJobsJob struct {
	JobId            common.String
	State            common.String
	Code             common.String
	Message          common.String
	Percent          common.Long
	PipelineId       common.String
	CreationTime     common.String
	FinishTime       common.String
	EditingInputs    SubmitEditingJobsEditingInputList
	EditingConfig    SubmitEditingJobsEditingConfig
	MNSMessageResult SubmitEditingJobsMNSMessageResult
}

type SubmitEditingJobsEditingInput struct {
	Id          common.String
	InputFile   SubmitEditingJobsInputFile
	InputConfig SubmitEditingJobsInputConfig
}

type SubmitEditingJobsInputFile struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type SubmitEditingJobsInputConfig struct {
	DeinterlaceMethod common.String
	IsNormalSar       common.String
}

type SubmitEditingJobsEditingConfig struct {
	TemplateId             common.String
	UserData               common.String
	Rotate                 common.String
	VideoStreamMap         common.String
	AudioStreamMap         common.String
	DeWatermark            common.String
	Priority               common.String
	WaterMarkConfigUrl     common.String
	MergeConfigUrl         common.String
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
	WaterMarkTemplateId common.String
	Width               common.String
	Height              common.String
	Dx                  common.String
	Dy                  common.String
	ReferPos            common.String
	Type                common.String
	InputFile1          SubmitEditingJobsInputFile1
}

type SubmitEditingJobsInputFile1 struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type SubmitEditingJobsMerge struct {
	MergeURL common.String
	Start    common.String
	Duration common.String
	RoleArn  common.String
}

type SubmitEditingJobsDigiWaterMark struct {
	Type       common.String
	Alpha      common.String
	InputFile2 SubmitEditingJobsInputFile2
}

type SubmitEditingJobsInputFile2 struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type SubmitEditingJobsOutputFile struct {
	Bucket   common.String
	Location common.String
	Object   common.String
	RoleArn  common.String
}

type SubmitEditingJobsM3U8NonStandardSupport struct {
	TS SubmitEditingJobsTS
}

type SubmitEditingJobsTS struct {
	Md5Support  bool
	SizeSupport bool
}

type SubmitEditingJobsProperties struct {
	Width      common.String
	Height     common.String
	Bitrate    common.String
	Duration   common.String
	Fps        common.String
	FileSize   common.String
	FileFormat common.String
	Streams    SubmitEditingJobsStreams
	Format     SubmitEditingJobsFormat
}

type SubmitEditingJobsStreams struct {
	VideoStreamList    SubmitEditingJobsVideoStreamList
	AudioStreamList    SubmitEditingJobsAudioStreamList
	SubtitleStreamList SubmitEditingJobsSubtitleStreamList
}

type SubmitEditingJobsVideoStream struct {
	Index          common.String
	CodecName      common.String
	CodecLongName  common.String
	Profile        common.String
	CodecTimeBase  common.String
	CodecTagString common.String
	CodecTag       common.String
	Width          common.String
	Height         common.String
	HasBFrames     common.String
	Sar            common.String
	Dar            common.String
	PixFmt         common.String
	Level          common.String
	Fps            common.String
	AvgFPS         common.String
	Timebase       common.String
	StartTime      common.String
	Duration       common.String
	Bitrate        common.String
	NumFrames      common.String
	Lang           common.String
	NetworkCost    SubmitEditingJobsNetworkCost
}

type SubmitEditingJobsNetworkCost struct {
	PreloadTime   common.String
	CostBandwidth common.String
	AvgBitrate    common.String
}

type SubmitEditingJobsAudioStream struct {
	Index          common.String
	CodecName      common.String
	CodecTimeBase  common.String
	CodecLongName  common.String
	CodecTagString common.String
	CodecTag       common.String
	SampleFmt      common.String
	Samplerate     common.String
	Channels       common.String
	ChannelLayout  common.String
	Timebase       common.String
	StartTime      common.String
	Duration       common.String
	Bitrate        common.String
	NumFrames      common.String
	Lang           common.String
}

type SubmitEditingJobsSubtitleStream struct {
	Index common.String
	Lang  common.String
}

type SubmitEditingJobsFormat struct {
	NumStreams     common.String
	NumPrograms    common.String
	FormatName     common.String
	FormatLongName common.String
	StartTime      common.String
	Duration       common.String
	Size           common.String
	Bitrate        common.String
}

type SubmitEditingJobsClip struct {
	TimeSpan SubmitEditingJobsTimeSpan
}

type SubmitEditingJobsTimeSpan struct {
	Seek     common.String
	Duration common.String
}

type SubmitEditingJobsSuperReso struct {
	IsHalfSample common.String
}

type SubmitEditingJobsSubtitleConfig struct {
	SubtitleList    SubmitEditingJobsSubtitleList
	ExtSubtitleList SubmitEditingJobsExtSubtitleList
}

type SubmitEditingJobsSubtitle struct {
	Map common.String
}

type SubmitEditingJobsExtSubtitle struct {
	FontName common.String
	CharEnc  common.String
	Input    SubmitEditingJobsInput
}

type SubmitEditingJobsInput struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type SubmitEditingJobsTransConfig struct {
	TransMode               common.String
	IsCheckReso             common.String
	IsCheckResoFail         common.String
	IsCheckVideoBitrate     common.String
	IsCheckAudioBitrate     common.String
	AdjDarMethod            common.String
	IsCheckVideoBitrateFail common.String
	IsCheckAudioBitrateFail common.String
	Duration                common.String
}

type SubmitEditingJobsMuxConfig struct {
	Segment SubmitEditingJobsSegment
	Gif     SubmitEditingJobsGif
}

type SubmitEditingJobsSegment struct {
	Duration common.String
}

type SubmitEditingJobsGif struct {
	Loop            common.String
	FinalDelay      common.String
	IsCustomPalette common.String
	DitherMode      common.String
}

type SubmitEditingJobsAudio struct {
	Codec      common.String
	Profile    common.String
	Samplerate common.String
	Bitrate    common.String
	Channels   common.String
	Qscale     common.String
	Volume     SubmitEditingJobsVolume
}

type SubmitEditingJobsVolume struct {
	Level  common.String
	Method common.String
}

type SubmitEditingJobsVideo struct {
	Codec      common.String
	Profile    common.String
	Bitrate    common.String
	Crf        common.String
	Width      common.String
	Height     common.String
	Fps        common.String
	Gop        common.String
	Preset     common.String
	ScanMode   common.String
	Bufsize    common.String
	Maxrate    common.String
	PixFmt     common.String
	Degrain    common.String
	Qscale     common.String
	Crop       common.String
	Pad        common.String
	MaxFps     common.String
	BitrateBnd SubmitEditingJobsBitrateBnd
}

type SubmitEditingJobsBitrateBnd struct {
	Max common.String
	Min common.String
}

type SubmitEditingJobsContainer struct {
	Format common.String
}

type SubmitEditingJobsEncryption struct {
	Type    common.String
	Id      common.String
	Key     common.String
	KeyUri  common.String
	KeyType common.String
	SkipCnt common.String
}

type SubmitEditingJobsEditing struct {
	ClipList SubmitEditingJobsClip3List
	Timeline SubmitEditingJobsTimeline
}

type SubmitEditingJobsClip3 struct {
	Id            common.String
	Type          common.String
	SourceType    common.String
	SourceID      common.String
	SourceStrmMap common.String
	In            common.String
	Out           common.String
	Effects       SubmitEditingJobsEffectList
}

type SubmitEditingJobsEffect struct {
	Effect       common.String
	EffectConfig common.String
}

type SubmitEditingJobsTimeline struct {
	TrackList      SubmitEditingJobsTrackList
	TimelineConfig SubmitEditingJobsTimelineConfig
}

type SubmitEditingJobsTrack struct {
	Id    common.String
	Type  common.String
	Order common.String
	Clips SubmitEditingJobsClip4List
}

type SubmitEditingJobsClip4 struct {
	ClipID      common.String
	In          common.String
	Out         common.String
	ClipsConfig SubmitEditingJobsClipsConfig
}

type SubmitEditingJobsClipsConfig struct {
	ClipsConfigVideo SubmitEditingJobsClipsConfigVideo
}

type SubmitEditingJobsClipsConfigVideo struct {
	L common.String
	T common.String
}

type SubmitEditingJobsTimelineConfig struct {
	TimelineConfigVideo SubmitEditingJobsTimelineConfigVideo
	TimelineConfigAudio SubmitEditingJobsTimelineConfigAudio
}

type SubmitEditingJobsTimelineConfigVideo struct {
	Width   common.String
	Height  common.String
	BgColor common.String
	Fps     common.String
}

type SubmitEditingJobsTimelineConfigAudio struct {
	Samplerate    common.String
	ChannelLayout common.String
	Channels      common.String
}

type SubmitEditingJobsMNSMessageResult struct {
	MessageId    common.String
	ErrorMessage common.String
	ErrorCode    common.String
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
