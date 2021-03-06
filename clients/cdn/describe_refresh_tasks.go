package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeRefreshTasksRequest struct {
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

func (req *DescribeRefreshTasksRequest) Invoke(client *sdk.Client) (resp *DescribeRefreshTasksResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeRefreshTasks", "", "")
	resp = &DescribeRefreshTasksResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeRefreshTasksResponse struct {
	responses.BaseResponse
	RequestId  common.String
	PageNumber common.Long
	PageSize   common.Long
	TotalCount common.Long
	Tasks      DescribeRefreshTasksCDNTaskList
}

type DescribeRefreshTasksCDNTask struct {
	TaskId       common.String
	ObjectPath   common.String
	Process      common.String
	Status       common.String
	CreationTime common.String
	Description  common.String
	ObjectType   common.String
}

type DescribeRefreshTasksCDNTaskList []DescribeRefreshTasksCDNTask

func (list *DescribeRefreshTasksCDNTaskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRefreshTasksCDNTask)
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
