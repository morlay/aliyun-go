package ess

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCapacityHistoryRequest struct {
	requests.RpcRequest
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	PageSize             int    `position:"Query" name:"PageSize"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeCapacityHistoryRequest) Invoke(client *sdk.Client) (resp *DescribeCapacityHistoryResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "DescribeCapacityHistory", "ess", "")
	resp = &DescribeCapacityHistoryResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCapacityHistoryResponse struct {
	responses.BaseResponse
	TotalCount           common.Integer
	PageNumber           common.Integer
	PageSize             common.Integer
	CapacityHistoryItems DescribeCapacityHistoryCapacityHistoryModelList
}

type DescribeCapacityHistoryCapacityHistoryModel struct {
	ScalingGroupId      common.String
	TotalCapacity       common.Integer
	AttachedCapacity    common.Integer
	AutoCreatedCapacity common.Integer
	Timestamp           common.String
}

type DescribeCapacityHistoryCapacityHistoryModelList []DescribeCapacityHistoryCapacityHistoryModel

func (list *DescribeCapacityHistoryCapacityHistoryModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCapacityHistoryCapacityHistoryModel)
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
