package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListLiveRecordVideoRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DomainName           string `position:"Query" name:"DomainName"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AppName              string `position:"Query" name:"AppName"`
	PageNo               int    `position:"Query" name:"PageNo"`
	PageSize             int    `position:"Query" name:"PageSize"`
	SortBy               string `position:"Query" name:"SortBy"`
	StreamName           string `position:"Query" name:"StreamName"`
	QueryType            string `position:"Query" name:"QueryType"`
}

func (req *ListLiveRecordVideoRequest) Invoke(client *sdk.Client) (resp *ListLiveRecordVideoResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "ListLiveRecordVideo", "vod", "")
	resp = &ListLiveRecordVideoResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListLiveRecordVideoResponse struct {
	responses.BaseResponse
	RequestId           string
	Total               int
	LiveRecordVideoList ListLiveRecordVideoLiveRecordVideoList
}

type ListLiveRecordVideoLiveRecordVideo struct {
	StreamName      string
	DomainName      string
	AppName         string
	PlaylistId      string
	RecordStartTime string
	RecordEndTime   string
	PlayInfoList    ListLiveRecordVideoPlayInfoList
	Video           ListLiveRecordVideoVideo
}

type ListLiveRecordVideoPlayInfo struct {
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

type ListLiveRecordVideoVideo struct {
	VideoId         string
	Title           string
	Tags            string
	Status          string
	Size            int64
	Privilege       int
	Duration        float32
	Description     string
	CustomerId      int64
	CreateTime      string
	ModifyTime      string
	CoverURL        string
	CateId          int
	CateName        string
	DownloadSwitch  string
	TemplateGroupId string
	Snapshots       ListLiveRecordVideoSnapshotList
}

type ListLiveRecordVideoLiveRecordVideoList []ListLiveRecordVideoLiveRecordVideo

func (list *ListLiveRecordVideoLiveRecordVideoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListLiveRecordVideoLiveRecordVideo)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type ListLiveRecordVideoPlayInfoList []ListLiveRecordVideoPlayInfo

func (list *ListLiveRecordVideoPlayInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListLiveRecordVideoPlayInfo)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type ListLiveRecordVideoSnapshotList []string

func (list *ListLiveRecordVideoSnapshotList) UnmarshalJSON(data []byte) error {
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
