package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	Nodes      ListNodesNodeInfoList
}

type ListNodesNodeInfo struct {
	Id              string
	RegionId        string
	Status          string
	CreatedByEhpc   bool
	Role            string
	AddTime         string
	Expired         bool
	ExpiredTime     string
	SpotStrategy    string
	LockReason      string
	ImageOwnerAlias string
	ImageId         string
	TotalResources  ListNodesTotalResources
	UsedResources   ListNodesUsedResources
}

type ListNodesTotalResources struct {
	Cpu    int
	Memory int
	Gpu    int
}

type ListNodesUsedResources struct {
	Cpu    int
	Memory int
	Gpu    int
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
