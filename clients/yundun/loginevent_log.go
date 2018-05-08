package yundun

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId  common.String
	StartTime  common.String
	EndTime    common.String
	PageNumber common.Integer
	PageSize   common.Integer
	TotalCount common.Integer
	LogList    LogineventLogLoginEventLogList
}

type LogineventLogLoginEventLog struct {
	BlockTimes common.Integer
	SourceIp   common.String
	Status     common.Integer
	Time       common.String
	Location   common.String
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
