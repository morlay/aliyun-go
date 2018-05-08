package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListClusterNodeRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int    `position:"Query" name:"PageSize"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
}

func (req *ListClusterNodeRequest) Invoke(client *sdk.Client) (resp *ListClusterNodeResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClusterNode", "", "")
	resp = &ListClusterNodeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClusterNodeResponse struct {
	responses.BaseResponse
	RequestId       common.String
	ClusterNodeList ListClusterNodeClusterNodeList
}

type ListClusterNodeClusterNode struct {
	NodeId common.String
	NodeIp common.String
	Status common.String
}

type ListClusterNodeClusterNodeList []ListClusterNodeClusterNode

func (list *ListClusterNodeClusterNodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterNodeClusterNode)
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
