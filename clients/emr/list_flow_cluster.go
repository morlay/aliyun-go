package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListFlowClusterRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int    `position:"Query" name:"PageSize"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
}

func (req *ListFlowClusterRequest) Invoke(client *sdk.Client) (resp *ListFlowClusterResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListFlowCluster", "", "")
	resp = &ListFlowClusterResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListFlowClusterResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TotalCount common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	Clusters   ListFlowClusterClusterInfoList
}

type ListFlowClusterClusterInfo struct {
	Id                  common.String
	Name                common.String
	Type                common.String
	CreateTime          common.Long
	RunningTime         common.Integer
	Status              common.String
	ChargeType          common.String
	ExpiredTime         common.Long
	Period              common.Integer
	HasUncompletedOrder bool
	OrderList           common.String
	CreateResource      common.String
	OrderTaskInfo       ListFlowClusterOrderTaskInfo
	FailReason          ListFlowClusterFailReason
}

type ListFlowClusterOrderTaskInfo struct {
	TargetCount  common.Integer
	CurrentCount common.Integer
	OrderIdList  common.String
}

type ListFlowClusterFailReason struct {
	ErrorCode common.String
	ErrorMsg  common.String
	RequestId common.String
}

type ListFlowClusterClusterInfoList []ListFlowClusterClusterInfo

func (list *ListFlowClusterClusterInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFlowClusterClusterInfo)
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
