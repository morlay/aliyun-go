package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListClustersRequest struct {
	requests.RpcRequest
	ResourceOwnerId  int64                            `position:"Query" name:"ResourceOwnerId"`
	StatusLists      *ListClustersStatusListList      `position:"Query" type:"Repeated" name:"StatusList"`
	PageSize         int                              `position:"Query" name:"PageSize"`
	ClusterTypeLists *ListClustersClusterTypeListList `position:"Query" type:"Repeated" name:"ClusterTypeList"`
	IsDesc           string                           `position:"Query" name:"IsDesc"`
	CreateType       string                           `position:"Query" name:"CreateType"`
	DefaultStatus    string                           `position:"Query" name:"DefaultStatus"`
	PageNumber       int                              `position:"Query" name:"PageNumber"`
}

func (req *ListClustersRequest) Invoke(client *sdk.Client) (resp *ListClustersResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClusters", "", "")
	resp = &ListClustersResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClustersResponse struct {
	responses.BaseResponse
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	Clusters   ListClustersClusterInfoList
}

type ListClustersClusterInfo struct {
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
	OrderTaskInfo       ListClustersOrderTaskInfo
	FailReason          ListClustersFailReason
}

type ListClustersOrderTaskInfo struct {
	TargetCount  int
	CurrentCount int
	OrderIdList  string
}

type ListClustersFailReason struct {
	ErrorCode string
	ErrorMsg  string
	RequestId string
}

type ListClustersStatusListList []string

func (list *ListClustersStatusListList) UnmarshalJSON(data []byte) error {
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

type ListClustersClusterTypeListList []string

func (list *ListClustersClusterTypeListList) UnmarshalJSON(data []byte) error {
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

type ListClustersClusterInfoList []ListClustersClusterInfo

func (list *ListClustersClusterInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClustersClusterInfo)
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
