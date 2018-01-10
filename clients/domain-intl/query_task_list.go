package domain_intl

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryTaskListRequest struct {
	requests.RpcRequest
	BeginCreateTime int64  `position:"Query" name:"BeginCreateTime"`
	EndCreateTime   int64  `position:"Query" name:"EndCreateTime"`
	UserClientIp    string `position:"Query" name:"UserClientIp"`
	PageSize        int    `position:"Query" name:"PageSize"`
	Lang            string `position:"Query" name:"Lang"`
	PageNum         int    `position:"Query" name:"PageNum"`
}

func (req *QueryTaskListRequest) Invoke(client *sdk.Client) (resp *QueryTaskListResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "QueryTaskList", "domain", "")
	resp = &QueryTaskListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryTaskListResponse struct {
	responses.BaseResponse
	RequestId      string
	TotalItemNum   int
	CurrentPageNum int
	TotalPageNum   int
	PageSize       int
	PrePage        bool
	NextPage       bool
	Data           QueryTaskListTaskInfoList
}

type QueryTaskListTaskInfo struct {
	TaskType   string
	TaskNum    int
	TaskStatus string
	CreateTime string
	Clientip   string
	TaskNo     string
}

type QueryTaskListTaskInfoList []QueryTaskListTaskInfo

func (list *QueryTaskListTaskInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTaskListTaskInfo)
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
