package yundun

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type LogineventLogRequest struct {
	requests.RpcRequest
	JstOwnerId int64  `position:"Query" name:"JstOwnerId"`
	PageNumber int    `position:"Query" name:"PageNumber"`
	PageSize   int    `position:"Query" name:"PageSize"`
	InstanceId string `position:"Query" name:"InstanceId"`
	RecordType int    `position:"Query" name:"RecordType"`
}

func (req *LogineventLogRequest) Invoke(client *sdk.Client) (resp *LogineventLogResponse, err error) {
	req.InitWithApiInfo("Yundun", "2015-04-16", "LogineventLog", "", "")
	resp = &LogineventLogResponse{}
	err = client.DoAction(req, resp)
	return
}

type LogineventLogResponse struct {
	responses.BaseResponse
	RequestId  string
	StartTime  string
	EndTime    string
	PageNumber int
	PageSize   int
	TotalCount int
	LogList    LogineventLogLoginEventLogList
}

type LogineventLogLoginEventLog struct {
	BlockTimes int
	SourceIp   string
	Status     int
	Time       string
	Location   string
}

type LogineventLogLoginEventLogList []LogineventLogLoginEventLog

func (list *LogineventLogLoginEventLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]LogineventLogLoginEventLog)
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
