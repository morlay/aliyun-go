package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribePlayUserTotalRequest struct {
	requests.RpcRequest
	EndTime   string `position:"Query" name:"EndTime"`
	StartTime string `position:"Query" name:"StartTime"`
	OwnerId   int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribePlayUserTotalRequest) Invoke(client *sdk.Client) (resp *DescribePlayUserTotalResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "DescribePlayUserTotal", "vod", "")
	resp = &DescribePlayUserTotalResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribePlayUserTotalResponse struct {
	responses.BaseResponse
	RequestId            string
	UserPlayStatisTotals DescribePlayUserTotalUserPlayStatisTotalList
}

type DescribePlayUserTotalUserPlayStatisTotal struct {
	Date         string
	PlayDuration string
	PlayRange    string
	VV           DescribePlayUserTotalVV
	UV           DescribePlayUserTotalUV
}

type DescribePlayUserTotalVV struct {
	Android string
	IOS     string
	Flash   string
	HTML5   string
}

type DescribePlayUserTotalUV struct {
	Android string
	IOS     string
	Flash   string
	HTML5   string
}

type DescribePlayUserTotalUserPlayStatisTotalList []DescribePlayUserTotalUserPlayStatisTotal

func (list *DescribePlayUserTotalUserPlayStatisTotalList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePlayUserTotalUserPlayStatisTotal)
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
