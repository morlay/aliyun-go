package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListFlowInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64                           `position:"Query" name:"ResourceOwnerId"`
	StatusLists     *ListFlowInstanceStatusListList `position:"Query" type:"Repeated" name:"StatusList"`
	PageSize        int                             `position:"Query" name:"PageSize"`
	FlowName        string                          `position:"Query" name:"FlowName"`
	Id              string                          `position:"Query" name:"Id"`
	FlowId          string                          `position:"Query" name:"FlowId"`
	ProjectId       string                          `position:"Query" name:"ProjectId"`
	PageNumber      int                             `position:"Query" name:"PageNumber"`
}

func (req *ListFlowInstanceRequest) Invoke(client *sdk.Client) (resp *ListFlowInstanceResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListFlowInstance", "", "")
	resp = &ListFlowInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListFlowInstanceResponse struct {
	responses.BaseResponse
	RequestId     common.String
	PageNumber    common.Integer
	PageSize      common.Integer
	Total         common.Integer
	FlowInstances ListFlowInstanceFlowInstanceList
}

type ListFlowInstanceFlowInstance struct {
	Id          common.String
	GmtCreate   common.Long
	GmtModified common.Long
	FlowId      common.String
	FlowName    common.String
	ProjectId   common.String
	Status      common.String
	ClusterId   common.String
	StartTime   common.Long
	EndTime     common.Long
}

type ListFlowInstanceStatusListList []string

func (list *ListFlowInstanceStatusListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type ListFlowInstanceFlowInstanceList []ListFlowInstanceFlowInstance

func (list *ListFlowInstanceFlowInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFlowInstanceFlowInstance)
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
