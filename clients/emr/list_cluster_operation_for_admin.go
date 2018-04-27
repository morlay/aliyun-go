package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListClusterOperationForAdminRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int    `position:"Query" name:"PageSize"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	UserId          string `position:"Query" name:"UserId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
	Status          string `position:"Query" name:"Status"`
}

func (req *ListClusterOperationForAdminRequest) Invoke(client *sdk.Client) (resp *ListClusterOperationForAdminResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClusterOperationForAdmin", "", "")
	resp = &ListClusterOperationForAdminResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClusterOperationForAdminResponse struct {
	responses.BaseResponse
	RequestId            string
	TotalCount           int
	PageNumber           int
	PageSize             int
	ClusterOperationList ListClusterOperationForAdminClusterOperationList
}

type ListClusterOperationForAdminClusterOperation struct {
	OperationId   string
	OperationName string
	StartTime     string
	Duration      string
	Status        string
	Percentage    string
	Comment       string
}

type ListClusterOperationForAdminClusterOperationList []ListClusterOperationForAdminClusterOperation

func (list *ListClusterOperationForAdminClusterOperationList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterOperationForAdminClusterOperation)
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
