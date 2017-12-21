package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryBatchTaskListRequest struct {
	requests.RpcRequest
	BeginCreateTime string `position:"Query" name:"BeginCreateTime"`
	EndCreateTime   string `position:"Query" name:"EndCreateTime"`
	UserClientIp    string `position:"Query" name:"UserClientIp"`
	PageSize        int    `position:"Query" name:"PageSize"`
	Lang            string `position:"Query" name:"Lang"`
	PageNum         int    `position:"Query" name:"PageNum"`
}

func (req *QueryBatchTaskListRequest) Invoke(client *sdk.Client) (resp *QueryBatchTaskListResponse, err error) {
	req.InitWithApiInfo("Domain", "2016-05-11", "QueryBatchTaskList", "", "")
	resp = &QueryBatchTaskListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryBatchTaskListResponse struct {
	responses.BaseResponse
	RequestId      string
	TotalItemNum   int
	CurrentPageNum int
	TotalPageNum   int
	PageSize       int
	PrePage        bool
	NextPage       bool
	Data           QueryBatchTaskListTaskInfoList
}

type QueryBatchTaskListTaskInfo struct {
	TaskType   string
	TaskNum    int
	TaskStatus string
	CreateTime string
	Clientip   string
	TaskNo     string
}

type QueryBatchTaskListTaskInfoList []QueryBatchTaskListTaskInfo

func (list *QueryBatchTaskListTaskInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryBatchTaskListTaskInfo)
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
