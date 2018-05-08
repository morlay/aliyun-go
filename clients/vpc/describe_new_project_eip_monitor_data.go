package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeNewProjectEipMonitorDataRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int    `position:"Query" name:"Period"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	AllocationId         string `position:"Query" name:"AllocationId"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeNewProjectEipMonitorDataRequest) Invoke(client *sdk.Client) (resp *DescribeNewProjectEipMonitorDataResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeNewProjectEipMonitorData", "vpc", "")
	resp = &DescribeNewProjectEipMonitorDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeNewProjectEipMonitorDataResponse struct {
	responses.BaseResponse
	RequestId       common.String
	EipMonitorDatas DescribeNewProjectEipMonitorDataEipMonitorDataList
}

type DescribeNewProjectEipMonitorDataEipMonitorData struct {
	EipRX        common.Integer
	EipTX        common.Integer
	EipFlow      common.Integer
	EipBandwidth common.Integer
	EipPackets   common.Integer
	TimeStamp    common.String
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
