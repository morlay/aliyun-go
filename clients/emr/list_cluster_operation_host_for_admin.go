package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListClusterOperationHostForAdminRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int    `position:"Query" name:"PageSize"`
	OperationId     string `position:"Query" name:"OperationId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	UserId          string `position:"Query" name:"UserId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
	Status          string `position:"Query" name:"Status"`
}

func (req *ListClusterOperationHostForAdminRequest) Invoke(client *sdk.Client) (resp *ListClusterOperationHostForAdminResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClusterOperationHostForAdmin", "", "")
	resp = &ListClusterOperationHostForAdminResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClusterOperationHostForAdminResponse struct {
	responses.BaseResponse
	RequestId                string
	TotalCount               int
	PageNumber               int
	PageSize                 int
	ClusterOperationHostList ListClusterOperationHostForAdminClusterOperationHostList
}

type ListClusterOperationHostForAdminClusterOperationHost struct {
	HostId     string
	HostName   string
	Status     string
	Percentage string
}

type ListClusterOperationHostForAdminClusterOperationHostList []ListClusterOperationHostForAdminClusterOperationHost

func (list *ListClusterOperationHostForAdminClusterOperationHostList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterOperationHostForAdminClusterOperationHost)
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
