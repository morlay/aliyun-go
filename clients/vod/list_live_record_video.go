package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId           common.String
	Total               common.Integer
	LiveRecordVideoList ListLiveRecordVideoLiveRecordVideoList
}

type ListLiveRecordVideoLiveRecordVideo struct {
	StreamName      common.String
	DomainName      common.String
	AppName         common.String
	PlaylistId      common.String
	RecordStartTime common.String
	RecordEndTime   common.String
	PlayInfoList    ListLiveRecordVideoPlayInfoList
	Video           ListLiveRecordVideoVideo
}

type ListLiveRecordVideoPlayInfo struct {
	Width      common.Long
	Height     common.Long
	Size       common.Long
	PlayURL    common.String
	Bitrate    common.String
	Definition common.String
	Duration   common.String
	Format     common.String
	Fps        common.String
	Encrypt    common.Long
	Plaintext  common.String
	Complexity common.String
	StreamType common.String
	Rand       common.String
	JobId      common.String
}

type ListLiveRecordVideoVideo struct {
	VideoId         common.String
	Title           common.String
	Tags            common.String
	Status          common.String
	Size            common.Long
	Privilege       common.Integer
	Duration        common.Float
	Description     common.String
	CustomerId      common.Long
	CreateTime      common.String
	CreationTime    common.String
	ModifyTime      common.String
	CoverURL        common.String
	CateId          common.Integer
	CateName        common.String
	DownloadSwitch  common.String
	TemplateGroupId common.String
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

type ListLiveRecordVideoSnapshotList []common.String

func (list *ListLiveRecordVideoSnapshotList) UnmarshalJSON(data []byte) error {
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
