package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeFlowNodeInstanceContainerLogRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Offset          int    `position:"Query" name:"Offset"`
	AppId           string `position:"Query" name:"AppId"`
	Length          int    `position:"Query" name:"Length"`
	ContainerId     string `position:"Query" name:"ContainerId"`
	NodeInstanceId  string `position:"Query" name:"NodeInstanceId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
}

func (req *DescribeFlowNodeInstanceContainerLogRequest) Invoke(client *sdk.Client) (resp *DescribeFlowNodeInstanceContainerLogResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeFlowNodeInstanceContainerLog", "", "")
	resp = &DescribeFlowNodeInstanceContainerLogResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeFlowNodeInstanceContainerLogResponse struct {
	responses.BaseResponse
	RequestId string
	Logs      DescribeFlowNodeInstanceContainerLogLogList
}

type DescribeFlowNodeInstanceContainerLogLog struct {
	LogEntry DescribeFlowNodeInstanceContainerLogLogEntry
}

type DescribeFlowNodeInstanceContainerLogLogEntry struct {
	Content string
}

type DescribeFlowNodeInstanceContainerLogLogList []DescribeFlowNodeInstanceContainerLogLog

func (list *DescribeFlowNodeInstanceContainerLogLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeFlowNodeInstanceContainerLogLog)
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
