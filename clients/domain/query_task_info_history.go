package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId         common.String
	PageSize          common.Integer
	Objects           QueryTaskInfoHistoryTaskInfoHistoryList
	CurrentPageCursor QueryTaskInfoHistoryCurrentPageCursor
	NextPageCursor    QueryTaskInfoHistoryNextPageCursor
	PrePageCursor     QueryTaskInfoHistoryPrePageCursor
}

type QueryTaskInfoHistoryTaskInfoHistory struct {
	TaskType            common.String
	TaskNum             common.Integer
	TaskStatus          common.String
	CreateTime          common.String
	Clientip            common.String
	TaskNo              common.String
	CreateTimeLong      common.Long
	TaskStatusCode      common.Integer
	TaskTypeDescription common.String
}

type QueryTaskInfoHistoryCurrentPageCursor struct {
	TaskType            common.String
	TaskNum             common.Integer
	TaskStatus          common.String
	CreateTime          common.String
	Clientip            common.String
	TaskNo              common.String
	CreateTimeLong      common.Long
	TaskStatusCode      common.Integer
	TaskTypeDescription common.String
}

type QueryTaskInfoHistoryNextPageCursor struct {
	TaskType            common.String
	TaskNum             common.Integer
	TaskStatus          common.String
	CreateTime          common.String
	Clientip            common.String
	TaskNo              common.String
	CreateTimeLong      common.Long
	TaskStatusCode      common.Integer
	TaskTypeDescription common.String
}

type QueryTaskInfoHistoryPrePageCursor struct {
	TaskType            common.String
	TaskNum             common.Integer
	TaskStatus          common.String
	CreateTime          common.String
	Clientip            common.String
	TaskNo              common.String
	CreateTimeLong      common.Long
	TaskStatusCode      common.Integer
	TaskTypeDescription common.String
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
