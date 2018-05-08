package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryJobListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string `position:"Query" name:"JobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QueryJobListRequest) Invoke(client *sdk.Client) (resp *QueryJobListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryJobList", "mts", "")
	resp = &QueryJobListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryJobListResponse struct {
	responses.BaseResponse
	RequestId      common.String
	JobList        QueryJobListJobList
	NonExistJobIds QueryJobListNonExistJobIdList
}

type QueryJobListJob struct {
	JobId            common.String
	State            common.String
	Code             common.String
	Message          common.String
	Percent          common.Long
	PipelineId       common.String
	CreationTime     common.String
	FinishTime       common.String
	Input            QueryJobListInput
	Output           QueryJobListOutput
	MNSMessageResult QueryJobListMNSMessageResult
}

type QueryJobListInput struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type QueryJobListOutput struct {
	TemplateId             common.String
	UserData               common.String
	Rotate                 common.String
	VideoStreamMap         common.String
	AudioStreamMap         common.String
	DeWatermark            common.String
	Priority               common.String
	WaterMarkConfigUrl     common.String
	MergeConfigUrl         common.String
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
	WaterMarkTemplateId common.String
	Width               common.String
	Height              common.String
	Dx                  common.String
	Dy                  common.String
	ReferPos            common.String
	Type                common.String
	InputFile           QueryJobListInputFile
}

type QueryJobListInputFile struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type QueryJobListMerge struct {
	MergeURL common.String
	Start    common.String
	Duration common.String
	RoleArn  common.String
}

type QueryJobListOpening struct {
	OpenUrl common.String
	Start   common.String
	Width   common.String
	Height  common.String
}

type QueryJobListTailSlate struct {
	TailUrl       common.String
	Start         common.String
	BlendDuration common.String
	Width         common.String
	Height        common.String
	IsMergeAudio  bool
	BgColor       common.String
}

type QueryJobListOutputFile struct {
	Bucket   common.String
	Location common.String
	Object   common.String
	RoleArn  common.String
}

type QueryJobListM3U8NonStandardSupport struct {
	TS QueryJobListTS
}

type QueryJobListTS struct {
	Md5Support  bool
	SizeSupport bool
}

type QueryJobListProperties struct {
	Width       common.String
	Height      common.String
	Bitrate     common.String
	Duration    common.String
	Fps         common.String
	FileSize    common.String
	FileFormat  common.String
	SourceLogos QueryJobListSourceLogoList
	Streams     QueryJobListStreams
	Format      QueryJobListFormat
}

type QueryJobListSourceLogo struct {
	Source common.String
}

type QueryJobListStreams struct {
	VideoStreamList    QueryJobListVideoStreamList
	AudioStreamList    QueryJobListAudioStreamList
	SubtitleStreamList QueryJobListSubtitleStreamList
}

type QueryJobListVideoStream struct {
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
	NetworkCost    QueryJobListNetworkCost
}

type QueryJobListNetworkCost struct {
	PreloadTime   common.String
	CostBandwidth common.String
	AvgBitrate    common.String
}

type QueryJobListAudioStream struct {
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

type QueryJobListSubtitleStream struct {
	Index common.String
	Lang  common.String
}

type QueryJobListFormat struct {
	NumStreams     common.String
	NumPrograms    common.String
	FormatName     common.String
	FormatLongName common.String
	StartTime      common.String
	Duration       common.String
	Size           common.String
	Bitrate        common.String
}

type QueryJobListClip struct {
	TimeSpan QueryJobListTimeSpan
}

type QueryJobListTimeSpan struct {
	Seek     common.String
	Duration common.String
}

type QueryJobListSuperReso struct {
	IsHalfSample common.String
}

type QueryJobListSubtitleConfig struct {
	SubtitleList    QueryJobListSubtitleList
	ExtSubtitleList QueryJobListExtSubtitleList
}

type QueryJobListSubtitle struct {
	Map common.String
}

type QueryJobListExtSubtitle struct {
	FontName common.String
	CharEnc  common.String
	Input1   QueryJobListInput1
}

type QueryJobListInput1 struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type QueryJobListTransConfig struct {
	TransMode               common.String
	IsCheckReso             common.String
	IsCheckResoFail         common.String
	IsCheckVideoBitrate     common.String
	IsCheckAudioBitrate     common.String
	AdjDarMethod            common.String
	IsCheckVideoBitrateFail common.String
	IsCheckAudioBitrateFail common.String
}

type QueryJobListMuxConfig struct {
	Segment QueryJobListSegment
	Gif     QueryJobListGif
}

type QueryJobListSegment struct {
	Duration common.String
}

type QueryJobListGif struct {
	Loop            common.String
	FinalDelay      common.String
	IsCustomPalette common.String
	DitherMode      common.String
}

type QueryJobListAudio struct {
	Codec      common.String
	Profile    common.String
	Samplerate common.String
	Bitrate    common.String
	Channels   common.String
	Qscale     common.String
	Volume     QueryJobListVolume
}

type QueryJobListVolume struct {
	Level  common.String
	Method common.String
}

type QueryJobListVideo struct {
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
	BitrateBnd QueryJobListBitrateBnd
}

type QueryJobListBitrateBnd struct {
	Max common.String
	Min common.String
}

type QueryJobListContainer struct {
	Format common.String
}

type QueryJobListEncryption struct {
	Type    common.String
	Id      common.String
	Key     common.String
	KeyUri  common.String
	KeyType common.String
	SkipCnt common.String
}

type QueryJobListMNSMessageResult struct {
	MessageId    common.String
	ErrorMessage common.String
	ErrorCode    common.String
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

type QueryJobListNonExistJobIdList []common.String

func (list *QueryJobListNonExistJobIdList) UnmarshalJSON(data []byte) error {
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
