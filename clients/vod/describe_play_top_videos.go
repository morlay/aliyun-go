package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribePlayTopVideosRequest struct {
	requests.RpcRequest
	BizDate  string `position:"Query" name:"BizDate"`
	PageNo   int64  `position:"Query" name:"PageNo"`
	PageSize int64  `position:"Query" name:"PageSize"`
	OwnerId  int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribePlayTopVideosRequest) Invoke(client *sdk.Client) (resp *DescribePlayTopVideosResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "DescribePlayTopVideos", "vod", "")
	resp = &DescribePlayTopVideosResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribePlayTopVideosResponse struct {
	responses.BaseResponse
	RequestId     string
	PageNo        int64
	PageSize      int64
	TotalNum      int64
	TopPlayVideos DescribePlayTopVideosTopPlayVideoStatisList
}

type DescribePlayTopVideosTopPlayVideoStatis struct {
	PlayDuration string
	VV           string
	UV           string
	VideoId      string
	Title        string
}

type DescribePlayTopVideosTopPlayVideoStatisList []DescribePlayTopVideosTopPlayVideoStatis

func (list *DescribePlayTopVideosTopPlayVideoStatisList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePlayTopVideosTopPlayVideoStatis)
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
