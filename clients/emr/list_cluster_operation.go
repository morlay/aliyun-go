package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListClusterOperationRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int    `position:"Query" name:"PageSize"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
	Status          string `position:"Query" name:"Status"`
}

func (req *ListClusterOperationRequest) Invoke(client *sdk.Client) (resp *ListClusterOperationResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClusterOperation", "", "")
	resp = &ListClusterOperationResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClusterOperationResponse struct {
	responses.BaseResponse
	RequestId            string
	TotalCount           int
	PageNumber           int
	PageSize             int
	ClusterOperationList ListClusterOperationClusterOperationList
}

type ListClusterOperationClusterOperation struct {
	OperationId   string
	OperationName string
	StartTime     string
	Duration      string
	Status        string
	Percentage    string
	Comment       string
}

type ListClusterOperationClusterOperationList []ListClusterOperationClusterOperation

func (list *ListClusterOperationClusterOperationList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterOperationClusterOperation)
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
