package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListClusterServiceStatusRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *ListClusterServiceStatusRequest) Invoke(client *sdk.Client) (resp *ListClusterServiceStatusResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClusterServiceStatus", "", "")
	resp = &ListClusterServiceStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClusterServiceStatusResponse struct {
	responses.BaseResponse
	RequestId                common.String
	ClusterServiceStatusList ListClusterServiceStatusClusterServiceStatusList
}

type ListClusterServiceStatusClusterServiceStatus struct {
	NodeIp common.String
	Status common.String
}

type ListClusterServiceStatusClusterServiceStatusList []ListClusterServiceStatusClusterServiceStatus

func (list *ListClusterServiceStatusClusterServiceStatusList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceStatusClusterServiceStatus)
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
