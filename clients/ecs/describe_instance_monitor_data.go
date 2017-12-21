package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeInstanceMonitorDataRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int    `position:"Query" name:"Period"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeInstanceMonitorDataRequest) Invoke(client *sdk.Client) (resp *DescribeInstanceMonitorDataResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceMonitorData", "ecs", "")
	resp = &DescribeInstanceMonitorDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeInstanceMonitorDataResponse struct {
	responses.BaseResponse
	RequestId   string
	MonitorData DescribeInstanceMonitorDataInstanceMonitorDataList
}

type DescribeInstanceMonitorDataInstanceMonitorData struct {
	InstanceId        string
	CPU               int
	IntranetRX        int
	IntranetTX        int
	IntranetBandwidth int
	InternetRX        int
	InternetTX        int
	InternetBandwidth int
	IOPSRead          int
	IOPSWrite         int
	BPSRead           int
	BPSWrite          int
	CPUCreditUsage    float32
	CPUCreditBalance  float32
	TimeStamp         string
}

type DescribeInstanceMonitorDataInstanceMonitorDataList []DescribeInstanceMonitorDataInstanceMonitorData

func (list *DescribeInstanceMonitorDataInstanceMonitorDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceMonitorDataInstanceMonitorData)
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
