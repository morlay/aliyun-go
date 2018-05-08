package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId            common.String
	TotalCount           common.Integer
	PageNumber           common.Integer
	PageSize             common.Integer
	ClusterOperationList ListClusterOperationForAdminClusterOperationList
}

type ListClusterOperationForAdminClusterOperation struct {
	OperationId   common.String
	OperationName common.String
	StartTime     common.String
	Duration      common.String
	Status        common.String
	Percentage    common.String
	Comment       common.String
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
