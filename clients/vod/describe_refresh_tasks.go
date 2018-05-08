package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeRefreshTasksRequest struct {
	requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ObjectPath           string `position:"Query" name:"ObjectPath"`
	DomainName           string `position:"Query" name:"DomainName"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	PageSize             int    `position:"Query" name:"PageSize"`
	ObjectType           string `position:"Query" name:"ObjectType"`
	TaskId               string `position:"Query" name:"TaskId"`
	Status               string `position:"Query" name:"Status"`
}

func (req *DescribeRefreshTasksRequest) Invoke(client *sdk.Client) (resp *DescribeRefreshTasksResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "DescribeRefreshTasks", "vod", "")
	resp = &DescribeRefreshTasksResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeRefreshTasksResponse struct {
	responses.BaseResponse
	RequestId  common.String
	PageSize   common.Integer
	PageNumber common.Integer
	TotalCount common.Integer
	Tasks      DescribeRefreshTasksTaskList
}

type DescribeRefreshTasksTask struct {
	TaskId       common.String
	ObjectPath   common.String
	Status       common.String
	Process      common.String
	ObjectType   common.String
	CreationTime common.String
	Description  common.String
}

type DescribeRefreshTasksTaskList []DescribeRefreshTasksTask

func (list *DescribeRefreshTasksTaskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRefreshTasksTask)
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
