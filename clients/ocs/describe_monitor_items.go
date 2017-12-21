package ocs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeMonitorItemsRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
}

func (r DescribeMonitorItemsRequest) Invoke(client *sdk.Client) (response *DescribeMonitorItemsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeMonitorItemsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ocs", "2015-03-01", "DescribeMonitorItems", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeMonitorItemsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeMonitorItemsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeMonitorItemsResponse struct {
	RequestId                     string
	GetOcsMonitorItemsResponseDTO DescribeMonitorItemsGetOcsMonitorItemsResponseDTO
}

type DescribeMonitorItemsGetOcsMonitorItemsResponseDTO struct {
	MonitorItems DescribeMonitorItemsMonitorItemList
}

type DescribeMonitorItemsMonitorItem struct {
	MonitorKey string
	Unit       string
}

type DescribeMonitorItemsMonitorItemList []DescribeMonitorItemsMonitorItem

func (list *DescribeMonitorItemsMonitorItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMonitorItemsMonitorItem)
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
