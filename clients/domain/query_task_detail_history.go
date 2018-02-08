package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryTaskDetailHistoryRequest struct {
	requests.RpcRequest
	TaskStatus         int    `position:"Query" name:"TaskStatus"`
	UserClientIp       string `position:"Query" name:"UserClientIp"`
	TaskNo             string `position:"Query" name:"TaskNo"`
	DomainName         string `position:"Query" name:"DomainName"`
	PageSize           int    `position:"Query" name:"PageSize"`
	TaskDetailNoCursor string `position:"Query" name:"TaskDetailNoCursor"`
	Lang               string `position:"Query" name:"Lang"`
	DomainNameCursor   string `position:"Query" name:"DomainNameCursor"`
}

func (req *QueryTaskDetailHistoryRequest) Invoke(client *sdk.Client) (resp *QueryTaskDetailHistoryResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "QueryTaskDetailHistory", "", "")
	resp = &QueryTaskDetailHistoryResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryTaskDetailHistoryResponse struct {
	responses.BaseResponse
	RequestId         string
	PageSize          int
	Objects           QueryTaskDetailHistoryTaskDetailHistoryList
	CurrentPageCursor QueryTaskDetailHistoryCurrentPageCursor
	NextPageCursor    QueryTaskDetailHistoryNextPageCursor
	PrePageCursor     QueryTaskDetailHistoryPrePageCursor
}

type QueryTaskDetailHistoryTaskDetailHistory struct {
	TaskNo         string
	TaskDetailNo   string
	TaskType       string
	InstanceId     string
	DomainName     string
	TaskStatus     string
	UpdateTime     string
	CreateTime     string
	TryCount       int
	ErrorMsg       string
	TaskStatusCode int
}

type QueryTaskDetailHistoryCurrentPageCursor struct {
	TaskNo         string
	TaskDetailNo   string
	TaskType       string
	InstanceId     string
	DomainName     string
	TaskStatus     string
	UpdateTime     string
	CreateTime     string
	TryCount       int
	ErrorMsg       string
	TaskStatusCode int
}

type QueryTaskDetailHistoryNextPageCursor struct {
	TaskNo         string
	TaskDetailNo   string
	TaskType       string
	InstanceId     string
	DomainName     string
	TaskStatus     string
	UpdateTime     string
	CreateTime     string
	TryCount       int
	ErrorMsg       string
	TaskStatusCode int
}

type QueryTaskDetailHistoryPrePageCursor struct {
	TaskNo         string
	TaskDetailNo   string
	TaskType       string
	InstanceId     string
	DomainName     string
	TaskStatus     string
	UpdateTime     string
	CreateTime     string
	TryCount       int
	ErrorMsg       string
	TaskStatusCode int
}

type QueryTaskDetailHistoryTaskDetailHistoryList []QueryTaskDetailHistoryTaskDetailHistory

func (list *QueryTaskDetailHistoryTaskDetailHistoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTaskDetailHistoryTaskDetailHistory)
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
