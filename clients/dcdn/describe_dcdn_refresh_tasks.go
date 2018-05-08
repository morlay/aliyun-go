package dcdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDcdnRefreshTasksRequest struct {
	requests.RpcRequest
	ObjectPath      string `position:"Query" name:"ObjectPath"`
	DomainName      string `position:"Query" name:"DomainName"`
	EndTime         string `position:"Query" name:"EndTime"`
	StartTime       string `position:"Query" name:"StartTime"`
	OwnerId         int64  `position:"Query" name:"OwnerId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SecurityToken   string `position:"Query" name:"SecurityToken"`
	PageSize        int    `position:"Query" name:"PageSize"`
	ObjectType      string `position:"Query" name:"ObjectType"`
	TaskId          string `position:"Query" name:"TaskId"`
	Status          string `position:"Query" name:"Status"`
}

func (req *DescribeDcdnRefreshTasksRequest) Invoke(client *sdk.Client) (resp *DescribeDcdnRefreshTasksResponse, err error) {
	req.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnRefreshTasks", "dcdn", "")
	resp = &DescribeDcdnRefreshTasksResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDcdnRefreshTasksResponse struct {
	responses.BaseResponse
	RequestId  string
	PageNumber int64
	PageSize   int64
	TotalCount int64
	Tasks      DescribeDcdnRefreshTasksTaskList
}

type DescribeDcdnRefreshTasksTask struct {
	TaskId       string
	ObjectPath   string
	Process      string
	Status       string
	CreationTime string
	Description  string
	ObjectType   string
}

type DescribeDcdnRefreshTasksTaskList []DescribeDcdnRefreshTasksTask

func (list *DescribeDcdnRefreshTasksTaskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnRefreshTasksTask)
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
