package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryTaskDetailListRequest struct {
	requests.RpcRequest
	TaskStatus   int    `position:"Query" name:"TaskStatus"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	TaskNo       string `position:"Query" name:"TaskNo"`
	DomainName   string `position:"Query" name:"DomainName"`
	PageSize     int    `position:"Query" name:"PageSize"`
	Lang         string `position:"Query" name:"Lang"`
	PageNum      int    `position:"Query" name:"PageNum"`
}

func (req *QueryTaskDetailListRequest) Invoke(client *sdk.Client) (resp *QueryTaskDetailListResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "QueryTaskDetailList", "", "")
	resp = &QueryTaskDetailListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryTaskDetailListResponse struct {
	responses.BaseResponse
	RequestId      common.String
	TotalItemNum   common.Integer
	CurrentPageNum common.Integer
	TotalPageNum   common.Integer
	PageSize       common.Integer
	PrePage        bool
	NextPage       bool
	Data           QueryTaskDetailListTaskDetailList
}

type QueryTaskDetailListTaskDetail struct {
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
	TaskResult          common.String
	TaskTypeDescription common.String
}

type QueryTaskDetailListTaskDetailList []QueryTaskDetailListTaskDetail

func (list *QueryTaskDetailListTaskDetailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTaskDetailListTaskDetail)
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
