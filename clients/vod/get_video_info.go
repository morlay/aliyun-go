package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetVideoInfoRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VideoId              string `position:"Query" name:"VideoId"`
	ResultTypes          string `position:"Query" name:"ResultTypes"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *GetVideoInfoRequest) Invoke(client *sdk.Client) (resp *GetVideoInfoResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "GetVideoInfo", "vod", "")
	resp = &GetVideoInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetVideoInfoResponse struct {
	responses.BaseResponse
	RequestId string
	Video     GetVideoInfoVideo
}

type GetVideoInfoVideo struct {
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
	Snapshots    GetVideoInfoSnapshotList
}

type GetVideoInfoSnapshotList []string

func (list *GetVideoInfoSnapshotList) UnmarshalJSON(data []byte) error {
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
