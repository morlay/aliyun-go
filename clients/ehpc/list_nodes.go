package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListNodesRequest struct {
	requests.RpcRequest
	HostName   string `position:"Query" name:"HostName"`
	Role       string `position:"Query" name:"Role"`
	PageSize   int    `position:"Query" name:"PageSize"`
	ClusterId  string `position:"Query" name:"ClusterId"`
	PageNumber int    `position:"Query" name:"PageNumber"`
}

func (req *ListNodesRequest) Invoke(client *sdk.Client) (resp *ListNodesResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "ListNodes", "ehs", "")
	resp = &ListNodesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListNodesResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TotalCount common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	Nodes      ListNodesNodeInfoList
}

type ListNodesNodeInfo struct {
	Id              common.String
	RegionId        common.String
	Status          common.String
	CreatedByEhpc   bool
	Role            common.String
	AddTime         common.String
	Expired         bool
	ExpiredTime     common.String
	SpotStrategy    common.String
	LockReason      common.String
	ImageOwnerAlias common.String
	ImageId         common.String
	TotalResources  ListNodesTotalResources
	UsedResources   ListNodesUsedResources
}

type ListNodesTotalResources struct {
	Cpu    common.Integer
	Memory common.Integer
	Gpu    common.Integer
}

type ListNodesUsedResources struct {
	Cpu    common.Integer
	Memory common.Integer
	Gpu    common.Integer
}

type ListNodesNodeInfoList []ListNodesNodeInfo

func (list *ListNodesNodeInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListNodesNodeInfo)
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
