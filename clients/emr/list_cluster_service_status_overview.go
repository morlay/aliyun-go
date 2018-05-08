package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListClusterServiceStatusOverviewRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
}

func (req *ListClusterServiceStatusOverviewRequest) Invoke(client *sdk.Client) (resp *ListClusterServiceStatusOverviewResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClusterServiceStatusOverview", "", "")
	resp = &ListClusterServiceStatusOverviewResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClusterServiceStatusOverviewResponse struct {
	responses.BaseResponse
	RequestId                common.String
	ClusterServiceStatusList ListClusterServiceStatusOverviewClusterServiceStatusOverviewList
}

type ListClusterServiceStatusOverviewClusterServiceStatusOverview struct {
	ClusterId   common.String
	ServiceName common.String
	Status      common.String
}

type ListClusterServiceStatusOverviewClusterServiceStatusOverviewList []ListClusterServiceStatusOverviewClusterServiceStatusOverview

func (list *ListClusterServiceStatusOverviewClusterServiceStatusOverviewList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceStatusOverviewClusterServiceStatusOverview)
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
