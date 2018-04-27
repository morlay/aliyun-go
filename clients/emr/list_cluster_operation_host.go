package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListClusterOperationHostRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int    `position:"Query" name:"PageSize"`
	OperationId     string `position:"Query" name:"OperationId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
	Status          string `position:"Query" name:"Status"`
}

func (req *ListClusterOperationHostRequest) Invoke(client *sdk.Client) (resp *ListClusterOperationHostResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClusterOperationHost", "", "")
	resp = &ListClusterOperationHostResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClusterOperationHostResponse struct {
	responses.BaseResponse
	RequestId                string
	TotalCount               int
	PageNumber               int
	PageSize                 int
	ClusterOperationHostList ListClusterOperationHostClusterOperationHostList
}

type ListClusterOperationHostClusterOperationHost struct {
	HostId     string
	HostName   string
	Status     string
	Percentage string
}

type ListClusterOperationHostClusterOperationHostList []ListClusterOperationHostClusterOperationHost

func (list *ListClusterOperationHostClusterOperationHostList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterOperationHostClusterOperationHost)
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
