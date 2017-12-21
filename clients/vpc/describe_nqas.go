package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeNqasRequest struct {
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

func (r DescribeNqasRequest) Invoke(client *sdk.Client) (response *DescribeNqasResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeNqasRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeNqas", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		DescribeNqasResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeNqasResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeNqasResponse struct {
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	Nqas       DescribeNqasNqaList
}

type DescribeNqasNqa struct {
	NqaId         string
	RegionId      string
	Status        string
	RouterId      string
	DestinationIp string
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
