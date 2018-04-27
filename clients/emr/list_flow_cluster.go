package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	Clusters   ListFlowClusterClusterInfoList
}

type ListFlowClusterClusterInfo struct {
	Id                  string
	Name                string
	Type                string
	CreateTime          int64
	RunningTime         int
	Status              string
	ChargeType          string
	ExpiredTime         int64
	Period              int
	HasUncompletedOrder bool
	OrderList           string
	CreateResource      string
	OrderTaskInfo       ListFlowClusterOrderTaskInfo
	FailReason          ListFlowClusterFailReason
}

type ListFlowClusterOrderTaskInfo struct {
	TargetCount  int
	CurrentCount int
	OrderIdList  string
}

type ListFlowClusterFailReason struct {
	ErrorCode string
	ErrorMsg  string
	RequestId string
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
