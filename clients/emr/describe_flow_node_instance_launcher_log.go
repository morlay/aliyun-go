package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeFlowNodeInstanceLauncherLogRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Start           int    `position:"Query" name:"Start"`
	Lines           int    `position:"Query" name:"Lines"`
	NodeInstanceId  string `position:"Query" name:"NodeInstanceId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
}

func (req *DescribeFlowNodeInstanceLauncherLogRequest) Invoke(client *sdk.Client) (resp *DescribeFlowNodeInstanceLauncherLogResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeFlowNodeInstanceLauncherLog", "", "")
	resp = &DescribeFlowNodeInstanceLauncherLogResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeFlowNodeInstanceLauncherLogResponse struct {
	responses.BaseResponse
	RequestId string
	Logs      DescribeFlowNodeInstanceLauncherLogLogList
}

type DescribeFlowNodeInstanceLauncherLogLog struct {
	LogEntry DescribeFlowNodeInstanceLauncherLogLogEntry
}

type DescribeFlowNodeInstanceLauncherLogLogEntry struct {
	Content string
}

type DescribeFlowNodeInstanceLauncherLogLogList []DescribeFlowNodeInstanceLauncherLogLog

func (list *DescribeFlowNodeInstanceLauncherLogLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeFlowNodeInstanceLauncherLogLog)
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
