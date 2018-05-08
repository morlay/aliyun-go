package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId     common.String
	PageNumber    common.Integer
	PageSize      common.Integer
	Total         common.Integer
	NodeInstances ListFlowJobHistoryNodeInstanceList
}

type ListFlowJobHistoryNodeInstance struct {
	Id             common.String
	GmtCreate      common.Long
	GmtModified    common.Long
	Type           common.String
	Status         common.String
	JobId          common.String
	JobName        common.String
	JobType        common.String
	JobParams      common.String
	FailAct        common.String
	MaxRetry       common.Integer
	RetryInterval  common.Long
	NodeName       common.String
	ClusterId      common.String
	HostName       common.String
	ProjectId      common.String
	StartTime      common.Long
	EndTime        common.Long
	Pending        bool
	Retries        common.Integer
	ExternalId     common.String
	ExternalStatus common.String
	ExternalInfo   common.String
	ParamConf      common.String
	EnvConf        common.String
	RunConf        common.String
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
