package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetVideoListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CateId               int    `position:"Query" name:"CateId"`
	PageNo               int    `position:"Query" name:"PageNo"`
	PageSize             int    `position:"Query" name:"PageSize"`
	EndTime              string `position:"Query" name:"EndTime"`
	SortBy               string `position:"Query" name:"SortBy"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Status               string `position:"Query" name:"Status"`
}

func (r GetVideoListRequest) Invoke(client *sdk.Client) (response *GetVideoListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetVideoListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "GetVideoList", "", "")

	resp := struct {
		*responses.BaseResponse
		GetVideoListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetVideoListResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetVideoListResponse struct {
	RequestId string
	Total     int
	VideoList GetVideoListVideoList
}

type GetVideoListVideo struct {
	VideoId      string
	Title        string
	Tags         string
	Status       string
	Size         int64
	Duration     float32
	Description  string
	CreateTime   string
	CreationTime string
	ModifyTime   string
	CoverURL     string
	CateId       int
	CateName     string
	Snapshots    GetVideoListSnapshotList
}

type GetVideoListVideoList []GetVideoListVideo

func (list *GetVideoListVideoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetVideoListVideo)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type GetVideoListSnapshotList []string

func (list *GetVideoListSnapshotList) UnmarshalJSON(data []byte) error {
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
