package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribePlayUserAvgRequest struct {
	requests.RpcRequest
	EndTime   string `position:"Query" name:"EndTime"`
	StartTime string `position:"Query" name:"StartTime"`
	OwnerId   int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribePlayUserAvgRequest) Invoke(client *sdk.Client) (resp *DescribePlayUserAvgResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "DescribePlayUserAvg", "vod", "")
	resp = &DescribePlayUserAvgResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribePlayUserAvgResponse struct {
	responses.BaseResponse
	RequestId          string
	UserPlayStatisAvgs DescribePlayUserAvgUserPlayStatisAvgList
}

type DescribePlayUserAvgUserPlayStatisAvg struct {
	Date            string
	AvgPlayDuration string
	AvgPlayCount    string
}

type DescribePlayUserAvgUserPlayStatisAvgList []DescribePlayUserAvgUserPlayStatisAvg

func (list *DescribePlayUserAvgUserPlayStatisAvgList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePlayUserAvgUserPlayStatisAvg)
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
