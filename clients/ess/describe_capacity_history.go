package ess

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCapacityHistoryRequest struct {
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	PageSize             int    `position:"Query" name:"PageSize"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (r DescribeCapacityHistoryRequest) Invoke(client *sdk.Client) (response *DescribeCapacityHistoryResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeCapacityHistoryRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ess", "2014-08-28", "DescribeCapacityHistory", "ess", "")

	resp := struct {
		*responses.BaseResponse
		DescribeCapacityHistoryResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeCapacityHistoryResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeCapacityHistoryResponse struct {
	TotalCount           int
	PageNumber           int
	PageSize             int
	CapacityHistoryItems DescribeCapacityHistoryCapacityHistoryModelList
}

type DescribeCapacityHistoryCapacityHistoryModel struct {
	ScalingGroupId      string
	TotalCapacity       int
	AttachedCapacity    int
	AutoCreatedCapacity int
	Timestamp           string
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
