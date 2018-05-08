package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeRegionsRequest struct {
	requests.RpcRequest
}

func (req *DescribeRegionsRequest) Invoke(client *sdk.Client) (resp *DescribeRegionsResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "DescribeRegions", "imm", "")
	resp = &DescribeRegionsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeRegionsResponse struct {
	responses.BaseResponse
	RequestId common.String
	Regions   DescribeRegionsRegionsItemList
}

type DescribeRegionsRegionsItem struct {
	Region   common.String
	Status   common.String
	ShowName common.String
}

type DescribeRegionsRegionsItemList []DescribeRegionsRegionsItem

func (list *DescribeRegionsRegionsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRegionsRegionsItem)
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
