package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeBandwidthLimitationRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceType         string `position:"Query" name:"InstanceType"`
	InstanceChargeType   string `position:"Query" name:"InstanceChargeType"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OperationType        string `position:"Query" name:"OperationType"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SpotStrategy         string `position:"Query" name:"SpotStrategy"`
}

func (req *DescribeBandwidthLimitationRequest) Invoke(client *sdk.Client) (resp *DescribeBandwidthLimitationResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeBandwidthLimitation", "ecs", "")
	resp = &DescribeBandwidthLimitationResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeBandwidthLimitationResponse struct {
	responses.BaseResponse
	RequestId  common.String
	Bandwidths DescribeBandwidthLimitationBandwidthList
}

type DescribeBandwidthLimitationBandwidth struct {
	InternetChargeType common.String
	Min                common.Integer
	Max                common.Integer
	Unit               common.String
}

type DescribeBandwidthLimitationBandwidthList []DescribeBandwidthLimitationBandwidth

func (list *DescribeBandwidthLimitationBandwidthList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBandwidthLimitationBandwidth)
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
