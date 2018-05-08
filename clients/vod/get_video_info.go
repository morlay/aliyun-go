package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	AI        common.String
	Video     GetVideoInfoVideo
}

type GetVideoInfoVideo struct {
	VideoId      common.String
	Title        common.String
	Tags         common.String
	Status       common.String
	Size         common.Long
	Duration     common.Float
	Description  common.String
	CreateTime   common.String
	CreationTime common.String
	ModifyTime   common.String
	CoverURL     common.String
	CateId       common.Long
	CateName     common.String
	Snapshots    GetVideoInfoSnapshotList
}

type GetVideoInfoSnapshotList []common.String

func (list *GetVideoInfoSnapshotList) UnmarshalJSON(data []byte) error {
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
