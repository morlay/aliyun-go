package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListFlowJobHistoryRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int    `position:"Query" name:"PageSize"`
	Id              string `position:"Query" name:"Id"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
}

func (req *ListFlowJobHistoryRequest) Invoke(client *sdk.Client) (resp *ListFlowJobHistoryResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListFlowJobHistory", "", "")
	resp = &ListFlowJobHistoryResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListFlowJobHistoryResponse struct {
	responses.BaseResponse
	RequestId     string
	PageNumber    int
	PageSize      int
	Total         int
	NodeInstances ListFlowJobHistoryNodeInstanceList
}

type ListFlowJobHistoryNodeInstance struct {
	Id             string
	GmtCreate      int64
	GmtModified    int64
	Type           string
	Status         string
	JobId          string
	JobName        string
	JobType        string
	JobParams      string
	FailAct        string
	MaxRetry       int
	RetryInterval  int64
	NodeName       string
	ClusterId      string
	HostName       string
	ProjectId      string
	StartTime      int64
	EndTime        int64
	Pending        bool
	Retries        int
	ExternalId     string
	ExternalStatus string
	ExternalInfo   string
	ParamConf      string
	EnvConf        string
	RunConf        string
}

type ListFlowJobHistoryNodeInstanceList []ListFlowJobHistoryNodeInstance

func (list *ListFlowJobHistoryNodeInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFlowJobHistoryNodeInstance)
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
