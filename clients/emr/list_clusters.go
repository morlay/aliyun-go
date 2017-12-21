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
	CreateType       string                           `position:"Query" name:"CreateType"`
	IsDesc           string                           `position:"Query" name:"IsDesc"`
	PageNumber       int                              `position:"Query" name:"PageNumber"`
	PageSize         int                              `position:"Query" name:"PageSize"`
	DefaultStatus    string                           `position:"Query" name:"DefaultStatus"`
	ClusterTypeLists *ListClustersClusterTypeListList `position:"Query" type:"Repeated" name:"ClusterTypeList"`
	StatusLists      *ListClustersStatusListList      `position:"Query" type:"Repeated" name:"StatusList"`
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
	FailReason          ListClustersFailReason
}

type ListClustersFailReason struct {
	ErrorCode string
	ErrorMsg  string
	RequestId string
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
