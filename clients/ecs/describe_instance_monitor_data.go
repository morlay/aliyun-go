package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId   common.String
	MonitorData DescribeInstanceMonitorDataInstanceMonitorDataList
}

type DescribeInstanceMonitorDataInstanceMonitorData struct {
	InstanceId                   common.String
	CPU                          common.Integer
	IntranetRX                   common.Integer
	IntranetTX                   common.Integer
	IntranetBandwidth            common.Integer
	InternetRX                   common.Integer
	InternetTX                   common.Integer
	InternetBandwidth            common.Integer
	IOPSRead                     common.Integer
	IOPSWrite                    common.Integer
	BPSRead                      common.Integer
	BPSWrite                     common.Integer
	CPUCreditUsage               common.Float
	CPUCreditBalance             common.Float
	CPUAdvanceCreditBalance      common.Float
	CPUNotpaidSurplusCreditUsage common.Float
	TimeStamp                    common.String
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
