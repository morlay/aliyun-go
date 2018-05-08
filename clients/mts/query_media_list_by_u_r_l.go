package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryMediaListByURLRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	IncludeSnapshotList  string `position:"Query" name:"IncludeSnapshotList"`
	FileURLs             string `position:"Query" name:"FileURLs"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	IncludePlayList      string `position:"Query" name:"IncludePlayList"`
	IncludeMediaInfo     string `position:"Query" name:"IncludeMediaInfo"`
}

func (req *QueryMediaListByURLRequest) Invoke(client *sdk.Client) (resp *QueryMediaListByURLResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryMediaListByURL", "mts", "")
	resp = &QueryMediaListByURLResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryMediaListByURLResponse struct {
	responses.BaseResponse
	RequestId        common.String
	MediaList        QueryMediaListByURLMediaList
	NonExistFileURLs QueryMediaListByURLNonExistFileURLList
}

type QueryMediaListByURLMedia struct {
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
	PlayList     QueryMediaListByURLPlayList
	SnapshotList QueryMediaListByURLSnapshotList
	Tags         QueryMediaListByURLTagList
	RunIdList    QueryMediaListByURLRunIdListList
	File         QueryMediaListByURLFile
	MediaInfo    QueryMediaListByURLMediaInfo
}

type QueryMediaListByURLPlay struct {
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
	File1             QueryMediaListByURLFile1
}

type QueryMediaListByURLFile1 struct {
	URL   common.String
	State common.String
}

type QueryMediaListByURLSnapshot struct {
	Type              common.String
	MediaWorkflowId   common.String
	MediaWorkflowName common.String
	ActivityName      common.String
	Count             common.String
	File2             QueryMediaListByURLFile2
}

type QueryMediaListByURLFile2 struct {
	URL   common.String
	State common.String
}

type QueryMediaListByURLFile struct {
	URL   common.String
	State common.String
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
	NetworkCost    QueryMediaListByURLNetworkCost
}

type QueryMediaListByURLNetworkCost struct {
	PreloadTime   common.String
	CostBandwidth common.String
	AvgBitrate    common.String
}

type QueryMediaListByURLAudioStream struct {
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

type QueryMediaListByURLSubtitleStream struct {
	Index common.String
	Lang  common.String
}

type QueryMediaListByURLFormat struct {
	NumStreams     common.String
	NumPrograms    common.String
	FormatName     common.String
	FormatLongName common.String
	StartTime      common.String
	Duration       common.String
	Size           common.String
	Bitrate        common.String
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

type QueryMediaListByURLNonExistFileURLList []common.String

func (list *QueryMediaListByURLNonExistFileURLList) UnmarshalJSON(data []byte) error {
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

type QueryMediaListByURLTagList []common.String

func (list *QueryMediaListByURLTagList) UnmarshalJSON(data []byte) error {
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

type QueryMediaListByURLRunIdListList []common.String

func (list *QueryMediaListByURLRunIdListList) UnmarshalJSON(data []byte) error {
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
