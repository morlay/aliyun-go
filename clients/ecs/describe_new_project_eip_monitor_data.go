package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeNewProjectEipMonitorDataRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int    `position:"Query" name:"Period"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	AllocationId         string `position:"Query" name:"AllocationId"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeNewProjectEipMonitorDataRequest) Invoke(client *sdk.Client) (response *DescribeNewProjectEipMonitorDataResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeNewProjectEipMonitorDataRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeNewProjectEipMonitorData", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeNewProjectEipMonitorDataResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeNewProjectEipMonitorDataResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeNewProjectEipMonitorDataResponse struct {
	RequestId       string
	EipMonitorDatas DescribeNewProjectEipMonitorDataEipMonitorDataList
}

type DescribeNewProjectEipMonitorDataEipMonitorData struct {
	EipRX        int
	EipTX        int
	EipFlow      int
	EipBandwidth int
	EipPackets   int
	TimeStamp    string
}

type DescribeNewProjectEipMonitorDataEipMonitorDataList []DescribeNewProjectEipMonitorDataEipMonitorData

func (list *DescribeNewProjectEipMonitorDataEipMonitorDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeNewProjectEipMonitorDataEipMonitorData)
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
