package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryMediaListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	IncludeSnapshotList  string `position:"Query" name:"IncludeSnapshotList"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MediaIds             string `position:"Query" name:"MediaIds"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	IncludePlayList      string `position:"Query" name:"IncludePlayList"`
	IncludeMediaInfo     string `position:"Query" name:"IncludeMediaInfo"`
}

func (req *QueryMediaListRequest) Invoke(client *sdk.Client) (resp *QueryMediaListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryMediaList", "mts", "")
	resp = &QueryMediaListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryMediaListResponse struct {
	responses.BaseResponse
	RequestId        common.String
	MediaList        QueryMediaListMediaList
	NonExistMediaIds QueryMediaListNonExistMediaIdList
}

type QueryMediaListMedia struct {
	MediaId      common.String
	Title        common.String
	Description  common.String
	CoverURL     common.String
	CateId       common.Long
	Duration     common.String
	Format       common.String
	Size         common.String
	Bitrate      common.String
	Width        common.String
	Height       common.String
	Fps          common.String
	PublishState common.String
	CreationTime common.String
	PlayList     QueryMediaListPlayList
	SnapshotList QueryMediaListSnapshotList
	Tags         QueryMediaListTagList
	RunIdList    QueryMediaListRunIdListList
	File         QueryMediaListFile
	MediaInfo    QueryMediaListMediaInfo
}

type QueryMediaListPlay struct {
	ActivityName      common.String
	MediaWorkflowId   common.String
	MediaWorkflowName common.String
	Duration          common.String
	Format            common.String
	Size              common.String
	Bitrate           common.String
	Width             common.String
	Height            common.String
	Fps               common.String
	Encryption        common.String
	File1             QueryMediaListFile1
}

type QueryMediaListFile1 struct {
	URL   common.String
	State common.String
}

type QueryMediaListSnapshot struct {
	Type              common.String
	MediaWorkflowId   common.String
	MediaWorkflowName common.String
	ActivityName      common.String
	Count             common.String
	File2             QueryMediaListFile2
}

type QueryMediaListFile2 struct {
	URL   common.String
	State common.String
}

type QueryMediaListFile struct {
	URL   common.String
	State common.String
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
	Rotate         common.String
	NetworkCost    QueryMediaListNetworkCost
}

type QueryMediaListNetworkCost struct {
	PreloadTime   common.String
	CostBandwidth common.String
	AvgBitrate    common.String
}

type QueryMediaListAudioStream struct {
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

type QueryMediaListSubtitleStream struct {
	Index common.String
	Lang  common.String
}

type QueryMediaListFormat struct {
	NumStreams     common.String
	NumPrograms    common.String
	FormatName     common.String
	FormatLongName common.String
	StartTime      common.String
	Duration       common.String
	Size           common.String
	Bitrate        common.String
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

type QueryMediaListNonExistMediaIdList []common.String

func (list *QueryMediaListNonExistMediaIdList) UnmarshalJSON(data []byte) error {
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

type QueryMediaListTagList []common.String

func (list *QueryMediaListTagList) UnmarshalJSON(data []byte) error {
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

type QueryMediaListRunIdListList []common.String

func (list *QueryMediaListRunIdListList) UnmarshalJSON(data []byte) error {
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
