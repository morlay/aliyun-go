package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListClusterNodesRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *ListClusterNodesRequest) Invoke(client *sdk.Client) (resp *ListClusterNodesResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClusterNodes", "", "")
	resp = &ListClusterNodesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClusterNodesResponse struct {
	responses.BaseResponse
	RequestId    common.String
	ClusterNodes ListClusterNodesClusterNodeList
}

type ListClusterNodesClusterNode struct {
	NodeId common.String
	NodeIp common.String
	Status common.String
}

type ListClusterNodesClusterNodeList []ListClusterNodesClusterNode

func (list *ListClusterNodesClusterNodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterNodesClusterNode)
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
