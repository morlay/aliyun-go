package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeKeyPairsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	KeyPairFingerPrint   string `position:"Query" name:"KeyPairFingerPrint"`
	PageSize             int    `position:"Query" name:"PageSize"`
	KeyPairName          string `position:"Query" name:"KeyPairName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeKeyPairsRequest) Invoke(client *sdk.Client) (resp *DescribeKeyPairsResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeKeyPairs", "ecs", "")
	resp = &DescribeKeyPairsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeKeyPairsResponse struct {
	responses.BaseResponse
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	KeyPairs   DescribeKeyPairsKeyPairList
}

type DescribeKeyPairsKeyPair struct {
	KeyPairName        string
	KeyPairFingerPrint string
}

type DescribeKeyPairsKeyPairList []DescribeKeyPairsKeyPair

func (list *DescribeKeyPairsKeyPairList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeKeyPairsKeyPair)
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
