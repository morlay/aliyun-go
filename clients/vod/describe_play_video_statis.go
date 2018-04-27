package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribePlayVideoStatisRequest struct {
	requests.RpcRequest
	EndTime   string `position:"Query" name:"EndTime"`
	VideoId   string `position:"Query" name:"VideoId"`
	StartTime string `position:"Query" name:"StartTime"`
	OwnerId   int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribePlayVideoStatisRequest) Invoke(client *sdk.Client) (resp *DescribePlayVideoStatisResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "DescribePlayVideoStatis", "vod", "")
	resp = &DescribePlayVideoStatisResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribePlayVideoStatisResponse struct {
	responses.BaseResponse
	RequestId              string
	VideoPlayStatisDetails DescribePlayVideoStatisVideoPlayStatisDetailList
}

type DescribePlayVideoStatisVideoPlayStatisDetail struct {
	Date         string
	PlayDuration string
	VV           string
	UV           string
	PlayRange    string
	Title        string
}

type DescribePlayVideoStatisVideoPlayStatisDetailList []DescribePlayVideoStatisVideoPlayStatisDetail

func (list *DescribePlayVideoStatisVideoPlayStatisDetailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePlayVideoStatisVideoPlayStatisDetail)
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
