package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListClustersForAdminRequest struct {
	requests.RpcRequest
	ResourceOwnerId  int64                                    `position:"Query" name:"ResourceOwnerId"`
	StatusLists      *ListClustersForAdminStatusListList      `position:"Query" type:"Repeated" name:"StatusList"`
	PageSize         int                                      `position:"Query" name:"PageSize"`
	ClusterTypeLists *ListClustersForAdminClusterTypeListList `position:"Query" type:"Repeated" name:"ClusterTypeList"`
	IsDesc           string                                   `position:"Query" name:"IsDesc"`
	UserId           string                                   `position:"Query" name:"UserId"`
	CreateType       string                                   `position:"Query" name:"CreateType"`
	DefaultStatus    string                                   `position:"Query" name:"DefaultStatus"`
	PageNumber       int                                      `position:"Query" name:"PageNumber"`
}

func (req *ListClustersForAdminRequest) Invoke(client *sdk.Client) (resp *ListClustersForAdminResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClustersForAdmin", "", "")
	resp = &ListClustersForAdminResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClustersForAdminResponse struct {
	responses.BaseResponse
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	Clusters   ListClustersForAdminClusterInfoList
}

type ListClustersForAdminClusterInfo struct {
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
	OrderTaskInfo       ListClustersForAdminOrderTaskInfo
	FailReason          ListClustersForAdminFailReason
}

type ListClustersForAdminOrderTaskInfo struct {
	TargetCount  int
	CurrentCount int
	OrderIdList  string
}

type ListClustersForAdminFailReason struct {
	ErrorCode string
	ErrorMsg  string
	RequestId string
}

type ListClustersForAdminStatusListList []string

func (list *ListClustersForAdminStatusListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type ListClustersForAdminClusterTypeListList []string

func (list *ListClustersForAdminClusterTypeListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type ListClustersForAdminClusterInfoList []ListClustersForAdminClusterInfo

func (list *ListClustersForAdminClusterInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClustersForAdminClusterInfo)
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
