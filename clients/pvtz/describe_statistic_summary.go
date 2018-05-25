package pvtz

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeStatisticSummaryRequest struct {
	requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

func (req *DescribeStatisticSummaryRequest) Invoke(client *sdk.Client) (resp *DescribeStatisticSummaryResponse, err error) {
	req.InitWithApiInfo("pvtz", "2018-01-01", "DescribeStatisticSummary", "pvtz", "")
	resp = &DescribeStatisticSummaryResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeStatisticSummaryResponse struct {
	responses.BaseResponse
	RequestId       common.String
	TotalCount      common.Long
	ZoneRequestTops DescribeStatisticSummaryZoneRequestTopList
	VpcRequestTops  DescribeStatisticSummaryVpcRequestTopList
}

type DescribeStatisticSummaryZoneRequestTop struct {
	ZoneName     common.String
	RequestCount common.Long
}

type DescribeStatisticSummaryVpcRequestTop struct {
	RegionId     common.String
	VpcId        common.String
	TunnelId     common.String
	RequestCount common.Long
	RegionName   common.String
}

type DescribeStatisticSummaryZoneRequestTopList []DescribeStatisticSummaryZoneRequestTop

func (list *DescribeStatisticSummaryZoneRequestTopList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeStatisticSummaryZoneRequestTop)
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

type DescribeStatisticSummaryVpcRequestTopList []DescribeStatisticSummaryVpcRequestTop

func (list *DescribeStatisticSummaryVpcRequestTopList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeStatisticSummaryVpcRequestTop)
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
