package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryEditingJobList", "mts", "")
	resp = &QueryEditingJobListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryEditingJobListResponse struct {
	responses.BaseResponse
	RequestId      common.String
	JobList        QueryEditingJobListJobList
	NonExistJobIds QueryEditingJobListNonExistJobIdList
}

type QueryEditingJobListJob struct {
	JobId            common.String
	State            common.String
	Code             common.String
	Message          common.String
	Percent          common.Long
	PipelineId       common.String
	CreationTime     common.String
	FinishTime       common.String
	EditingInputs    QueryEditingJobListEditingInputList
	EditingConfig    QueryEditingJobListEditingConfig
	MNSMessageResult QueryEditingJobListMNSMessageResult
}

type QueryEditingJobListEditingInput struct {
	Id          common.String
	InputFile   QueryEditingJobListInputFile
	InputConfig QueryEditingJobListInputConfig
}

type QueryEditingJobListInputFile struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type QueryEditingJobListInputConfig struct {
	DeinterlaceMethod common.String
	IsNormalSar       common.String
}

type QueryEditingJobListEditingConfig struct {
	TemplateId             common.String
	UserData               common.String
	Rotate                 common.String
	VideoStreamMap         common.String
	AudioStreamMap         common.String
	DeWatermark            common.String
	Priority               common.String
	WaterMarkConfigUrl     common.String
	MergeConfigUrl         common.String
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
	WaterMarkTemplateId common.String
	Width               common.String
	Height              common.String
	Dx                  common.String
	Dy                  common.String
	ReferPos            common.String
	Type                common.String
	InputFile1          QueryEditingJobListInputFile1
}

type QueryEditingJobListInputFile1 struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type QueryEditingJobListMerge struct {
	MergeURL common.String
	Start    common.String
	Duration common.String
	RoleArn  common.String
}

type QueryEditingJobListDigiWaterMark struct {
	Type       common.String
	Alpha      common.String
	InputFile2 QueryEditingJobListInputFile2
}

type QueryEditingJobListInputFile2 struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type QueryEditingJobListOutputFile struct {
	Bucket   common.String
	Location common.String
	Object   common.String
	RoleArn  common.String
}

type QueryEditingJobListM3U8NonStandardSupport struct {
	TS QueryEditingJobListTS
}

type QueryEditingJobListTS struct {
	Md5Support  bool
	SizeSupport bool
}

type QueryEditingJobListProperties struct {
	Width      common.String
	Height     common.String
	Bitrate    common.String
	Duration   common.String
	Fps        common.String
	FileSize   common.String
	FileFormat common.String
	Streams    QueryEditingJobListStreams
	Format     QueryEditingJobListFormat
}

type QueryEditingJobListStreams struct {
	VideoStreamList    QueryEditingJobListVideoStreamList
	AudioStreamList    QueryEditingJobListAudioStreamList
	SubtitleStreamList QueryEditingJobListSubtitleStreamList
}

type QueryEditingJobListVideoStream struct {
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
	NetworkCost    QueryEditingJobListNetworkCost
}

type QueryEditingJobListNetworkCost struct {
	PreloadTime   common.String
	CostBandwidth common.String
	AvgBitrate    common.String
}

type QueryEditingJobListAudioStream struct {
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

type QueryEditingJobListSubtitleStream struct {
	Index common.String
	Lang  common.String
}

type QueryEditingJobListFormat struct {
	NumStreams     common.String
	NumPrograms    common.String
	FormatName     common.String
	FormatLongName common.String
	StartTime      common.String
	Duration       common.String
	Size           common.String
	Bitrate        common.String
}

type QueryEditingJobListClip struct {
	TimeSpan QueryEditingJobListTimeSpan
}

type QueryEditingJobListTimeSpan struct {
	Seek     common.String
	Duration common.String
}

type QueryEditingJobListSuperReso struct {
	IsHalfSample common.String
}

type QueryEditingJobListSubtitleConfig struct {
	SubtitleList    QueryEditingJobListSubtitleList
	ExtSubtitleList QueryEditingJobListExtSubtitleList
}

type QueryEditingJobListSubtitle struct {
	Map common.String
}

type QueryEditingJobListExtSubtitle struct {
	FontName common.String
	CharEnc  common.String
	Input    QueryEditingJobListInput
}

type QueryEditingJobListInput struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type QueryEditingJobListTransConfig struct {
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

type QueryEditingJobListMuxConfig struct {
	Segment QueryEditingJobListSegment
	Gif     QueryEditingJobListGif
}

type QueryEditingJobListSegment struct {
	Duration common.String
}

type QueryEditingJobListGif struct {
	Loop            common.String
	FinalDelay      common.String
	IsCustomPalette common.String
	DitherMode      common.String
}

type QueryEditingJobListAudio struct {
	Codec      common.String
	Profile    common.String
	Samplerate common.String
	Bitrate    common.String
	Channels   common.String
	Qscale     common.String
	Volume     QueryEditingJobListVolume
}

type QueryEditingJobListVolume struct {
	Level  common.String
	Method common.String
}

type QueryEditingJobListVideo struct {
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
	BitrateBnd QueryEditingJobListBitrateBnd
}

type QueryEditingJobListBitrateBnd struct {
	Max common.String
	Min common.String
}

type QueryEditingJobListContainer struct {
	Format common.String
}

type QueryEditingJobListEncryption struct {
	Type    common.String
	Id      common.String
	Key     common.String
	KeyUri  common.String
	KeyType common.String
	SkipCnt common.String
}

type QueryEditingJobListEditing struct {
	ClipList QueryEditingJobListClip3List
	Timeline QueryEditingJobListTimeline
}

type QueryEditingJobListClip3 struct {
	Id            common.String
	Type          common.String
	SourceType    common.String
	SourceID      common.String
	SourceStrmMap common.String
	In            common.String
	Out           common.String
	Effects       QueryEditingJobListEffectList
}

type QueryEditingJobListEffect struct {
	Effect       common.String
	EffectConfig common.String
}

type QueryEditingJobListTimeline struct {
	TrackList      QueryEditingJobListTrackList
	TimelineConfig QueryEditingJobListTimelineConfig
}

type QueryEditingJobListTrack struct {
	Id    common.String
	Type  common.String
	Order common.String
	Clips QueryEditingJobListClip4List
}

type QueryEditingJobListClip4 struct {
	ClipID      common.String
	In          common.String
	Out         common.String
	ClipsConfig QueryEditingJobListClipsConfig
}

type QueryEditingJobListClipsConfig struct {
	ClipsConfigVideo QueryEditingJobListClipsConfigVideo
}

type QueryEditingJobListClipsConfigVideo struct {
	L common.String
	T common.String
}

type QueryEditingJobListTimelineConfig struct {
	TimelineConfigVideo QueryEditingJobListTimelineConfigVideo
	TimelineConfigAudio QueryEditingJobListTimelineConfigAudio
}

type QueryEditingJobListTimelineConfigVideo struct {
	Width   common.String
	Height  common.String
	BgColor common.String
	Fps     common.String
}

type QueryEditingJobListTimelineConfigAudio struct {
	Samplerate    common.String
	ChannelLayout common.String
	Channels      common.String
}

type QueryEditingJobListMNSMessageResult struct {
	MessageId    common.String
	ErrorMessage common.String
	ErrorCode    common.String
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

type QueryEditingJobListNonExistJobIdList []common.String

func (list *QueryEditingJobListNonExistJobIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
	err := json.Unmarshal(data, &m)
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
