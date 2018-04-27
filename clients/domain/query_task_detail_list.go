package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId      string
	TotalItemNum   int
	CurrentPageNum int
	TotalPageNum   int
	PageSize       int
	PrePage        bool
	NextPage       bool
	Data           QueryTaskDetailListTaskDetailList
}

type QueryTaskDetailListTaskDetail struct {
	TaskNo              string
	TaskDetailNo        string
	TaskType            string
	InstanceId          string
	DomainName          string
	TaskStatus          string
	UpdateTime          string
	CreateTime          string
	TryCount            int
	ErrorMsg            string
	TaskStatusCode      int
	TaskResult          string
	TaskTypeDescription string
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
