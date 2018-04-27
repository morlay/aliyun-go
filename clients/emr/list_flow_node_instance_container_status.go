package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListFlowNodeInstanceContainerStatusRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int    `position:"Query" name:"PageSize"`
	NodeInstanceId  string `position:"Query" name:"NodeInstanceId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
}

func (req *ListFlowNodeInstanceContainerStatusRequest) Invoke(client *sdk.Client) (resp *ListFlowNodeInstanceContainerStatusResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListFlowNodeInstanceContainerStatus", "", "")
	resp = &ListFlowNodeInstanceContainerStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListFlowNodeInstanceContainerStatusResponse struct {
	responses.BaseResponse
	RequestId           string
	ContainerStatusList ListFlowNodeInstanceContainerStatusContainerStatusList
}

type ListFlowNodeInstanceContainerStatusContainerStatus struct {
	ApplicationId string
	ContainerId   string
	HostName      string
	Status        string
}

type ListFlowNodeInstanceContainerStatusContainerStatusList []ListFlowNodeInstanceContainerStatusContainerStatus

func (list *ListFlowNodeInstanceContainerStatusContainerStatusList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFlowNodeInstanceContainerStatusContainerStatus)
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
