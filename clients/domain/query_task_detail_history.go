package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId         common.String
	PageSize          common.Integer
	Objects           QueryTaskDetailHistoryTaskDetailHistoryList
	CurrentPageCursor QueryTaskDetailHistoryCurrentPageCursor
	NextPageCursor    QueryTaskDetailHistoryNextPageCursor
	PrePageCursor     QueryTaskDetailHistoryPrePageCursor
}

type QueryTaskDetailHistoryTaskDetailHistory struct {
	TaskNo              common.String
	TaskDetailNo        common.String
	TaskType            common.String
	InstanceId          common.String
	DomainName          common.String
	TaskStatus          common.String
	UpdateTime          common.String
	CreateTime          common.String
	TryCount            common.Integer
	ErrorMsg            common.String
	TaskStatusCode      common.Integer
	TaskTypeDescription common.String
}

type QueryTaskDetailHistoryCurrentPageCursor struct {
	TaskNo              common.String
	TaskDetailNo        common.String
	TaskType            common.String
	InstanceId          common.String
	DomainName          common.String
	TaskStatus          common.String
	UpdateTime          common.String
	CreateTime          common.String
	TryCount            common.Integer
	ErrorMsg            common.String
	TaskStatusCode      common.Integer
	TaskTypeDescription common.String
}

type QueryTaskDetailHistoryNextPageCursor struct {
	TaskNo              common.String
	TaskDetailNo        common.String
	TaskType            common.String
	InstanceId          common.String
	DomainName          common.String
	TaskStatus          common.String
	UpdateTime          common.String
	CreateTime          common.String
	TryCount            common.Integer
	ErrorMsg            common.String
	TaskStatusCode      common.Integer
	TaskTypeDescription common.String
}

type QueryTaskDetailHistoryPrePageCursor struct {
	TaskNo              common.String
	TaskDetailNo        common.String
	TaskType            common.String
	InstanceId          common.String
	DomainName          common.String
	TaskStatus          common.String
	UpdateTime          common.String
	CreateTime          common.String
	TryCount            common.Integer
	ErrorMsg            common.String
	TaskStatusCode      common.Integer
	TaskTypeDescription common.String
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
