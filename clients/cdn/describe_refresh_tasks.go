package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeRefreshTasksRequest struct {
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

func (r DescribeRefreshTasksRequest) Invoke(client *sdk.Client) (response *DescribeRefreshTasksResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeRefreshTasksRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeRefreshTasks", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeRefreshTasksResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeRefreshTasksResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeRefreshTasksResponse struct {
	RequestId  string
	PageNumber int64
	PageSize   int64
	TotalCount int64
	Tasks      DescribeRefreshTasksCDNTaskList
}

type DescribeRefreshTasksCDNTask struct {
	TaskId       string
	ObjectPath   string
	Process      string
	Status       string
	CreationTime string
	Description  string
	ObjectType   string
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
