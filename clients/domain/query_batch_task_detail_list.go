package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryBatchTaskDetailListRequest struct {
	requests.RpcRequest
	TaskStatus   int    `position:"Query" name:"TaskStatus"`
	SaleId       string `position:"Query" name:"SaleId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	TaskNo       string `position:"Query" name:"TaskNo"`
	DomainName   string `position:"Query" name:"DomainName"`
	PageSize     int    `position:"Query" name:"PageSize"`
	Lang         string `position:"Query" name:"Lang"`
	PageNum      int    `position:"Query" name:"PageNum"`
}

func (req *QueryBatchTaskDetailListRequest) Invoke(client *sdk.Client) (resp *QueryBatchTaskDetailListResponse, err error) {
	req.InitWithApiInfo("Domain", "2016-05-11", "QueryBatchTaskDetailList", "", "")
	resp = &QueryBatchTaskDetailListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryBatchTaskDetailListResponse struct {
	responses.BaseResponse
	RequestId      string
	TotalItemNum   int
	CurrentPageNum int
	TotalPageNum   int
	PageSize       int
	PrePage        bool
	NextPage       bool
	Data           QueryBatchTaskDetailListTaskDetailList
}

type QueryBatchTaskDetailListTaskDetail struct {
	TaskNo     string
	TaskType   string
	DomainName string
	TaskStatus string
	UpdateTime string
	TryCount   int
	ErrorMsg   string
}

type QueryBatchTaskDetailListTaskDetailList []QueryBatchTaskDetailListTaskDetail

func (list *QueryBatchTaskDetailListTaskDetailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryBatchTaskDetailListTaskDetail)
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
