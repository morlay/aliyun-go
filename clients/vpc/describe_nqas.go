package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeNqasRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	RouterId             string `position:"Query" name:"RouterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	NqaId                string `position:"Query" name:"NqaId"`
	IsDefault            string `position:"Query" name:"IsDefault"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeNqasRequest) Invoke(client *sdk.Client) (resp *DescribeNqasResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeNqas", "vpc", "")
	resp = &DescribeNqasResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeNqasResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TotalCount common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	Nqas       DescribeNqasNqaList
}

type DescribeNqasNqa struct {
	NqaId         common.String
	RegionId      common.String
	Status        common.String
	RouterId      common.String
	DestinationIp common.String
}

type DescribeNqasNqaList []DescribeNqasNqa

func (list *DescribeNqasNqaList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeNqasNqa)
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
