package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId            common.String
	UserPlayStatisTotals DescribePlayUserTotalUserPlayStatisTotalList
}

type DescribePlayUserTotalUserPlayStatisTotal struct {
	Date         common.String
	PlayDuration common.String
	PlayRange    common.String
	VV           DescribePlayUserTotalVV
	UV           DescribePlayUserTotalUV
}

type DescribePlayUserTotalVV struct {
	Android common.String
	IOS     common.String
	Flash   common.String
	HTML5   common.String
}

type DescribePlayUserTotalUV struct {
	Android common.String
	IOS     common.String
	Flash   common.String
	HTML5   common.String
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
