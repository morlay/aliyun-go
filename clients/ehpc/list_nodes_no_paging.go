package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListNodesNoPagingRequest struct {
	requests.RpcRequest
	HostName     string `position:"Query" name:"HostName"`
	Role         string `position:"Query" name:"Role"`
	ClusterId    string `position:"Query" name:"ClusterId"`
	OnlyDetached string `position:"Query" name:"OnlyDetached"`
}

func (req *ListNodesNoPagingRequest) Invoke(client *sdk.Client) (resp *ListNodesNoPagingResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "ListNodesNoPaging", "ehs", "")
	resp = &ListNodesNoPagingResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListNodesNoPagingResponse struct {
	responses.BaseResponse
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	Nodes      ListNodesNoPagingNodeInfoList
}

type ListNodesNoPagingNodeInfo struct {
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
	TotalResources  ListNodesNoPagingTotalResources
	UsedResources   ListNodesNoPagingUsedResources
}

type ListNodesNoPagingTotalResources struct {
	Cpu    int
	Memory int
	Gpu    int
}

type ListNodesNoPagingUsedResources struct {
	Cpu    int
	Memory int
	Gpu    int
}

type ListNodesNoPagingNodeInfoList []ListNodesNoPagingNodeInfo

func (list *ListNodesNoPagingNodeInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListNodesNoPagingNodeInfo)
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
