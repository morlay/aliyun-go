package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryTaskInfoHistoryRequest struct {
	requests.RpcRequest
	BeginCreateTime  int64  `position:"Query" name:"BeginCreateTime"`
	EndCreateTime    int64  `position:"Query" name:"EndCreateTime"`
	TaskNoCursor     string `position:"Query" name:"TaskNoCursor"`
	UserClientIp     string `position:"Query" name:"UserClientIp"`
	PageSize         int    `position:"Query" name:"PageSize"`
	Lang             string `position:"Query" name:"Lang"`
	CreateTimeCursor int64  `position:"Query" name:"CreateTimeCursor"`
}

func (req *QueryTaskInfoHistoryRequest) Invoke(client *sdk.Client) (resp *QueryTaskInfoHistoryResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "QueryTaskInfoHistory", "", "")
	resp = &QueryTaskInfoHistoryResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryTaskInfoHistoryResponse struct {
	responses.BaseResponse
	RequestId         string
	PageSize          int
	Objects           QueryTaskInfoHistoryTaskInfoHistoryList
	CurrentPageCursor QueryTaskInfoHistoryCurrentPageCursor
	NextPageCursor    QueryTaskInfoHistoryNextPageCursor
	PrePageCursor     QueryTaskInfoHistoryPrePageCursor
}

type QueryTaskInfoHistoryTaskInfoHistory struct {
	TaskType            string
	TaskNum             int
	TaskStatus          string
	CreateTime          string
	Clientip            string
	TaskNo              string
	CreateTimeLong      int64
	TaskStatusCode      int
	TaskTypeDescription string
}

type QueryTaskInfoHistoryCurrentPageCursor struct {
	TaskType            string
	TaskNum             int
	TaskStatus          string
	CreateTime          string
	Clientip            string
	TaskNo              string
	CreateTimeLong      int64
	TaskStatusCode      int
	TaskTypeDescription string
}

type QueryTaskInfoHistoryNextPageCursor struct {
	TaskType            string
	TaskNum             int
	TaskStatus          string
	CreateTime          string
	Clientip            string
	TaskNo              string
	CreateTimeLong      int64
	TaskStatusCode      int
	TaskTypeDescription string
}

type QueryTaskInfoHistoryPrePageCursor struct {
	TaskType            string
	TaskNum             int
	TaskStatus          string
	CreateTime          string
	Clientip            string
	TaskNo              string
	CreateTimeLong      int64
	TaskStatusCode      int
	TaskTypeDescription string
}

type QueryTaskInfoHistoryTaskInfoHistoryList []QueryTaskInfoHistoryTaskInfoHistory

func (list *QueryTaskInfoHistoryTaskInfoHistoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTaskInfoHistoryTaskInfoHistory)
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
