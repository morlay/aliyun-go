package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId  common.String
	TotalCount common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	Nodes      ListNodesNoPagingNodeInfoList
}

type ListNodesNoPagingNodeInfo struct {
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
	TotalResources  ListNodesNoPagingTotalResources
	UsedResources   ListNodesNoPagingUsedResources
}

type ListNodesNoPagingTotalResources struct {
	Cpu    common.Integer
	Memory common.Integer
	Gpu    common.Integer
}

type ListNodesNoPagingUsedResources struct {
	Cpu    common.Integer
	Memory common.Integer
	Gpu    common.Integer
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
