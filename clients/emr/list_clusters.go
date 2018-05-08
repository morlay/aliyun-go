package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId  common.String
	TotalCount common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	Clusters   ListClustersClusterInfoList
}

type ListClustersClusterInfo struct {
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
	OrderTaskInfo       ListClustersOrderTaskInfo
	FailReason          ListClustersFailReason
}

type ListClustersOrderTaskInfo struct {
	TargetCount  common.Integer
	CurrentCount common.Integer
	OrderIdList  common.String
}

type ListClustersFailReason struct {
	ErrorCode common.String
	ErrorMsg  common.String
	RequestId common.String
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
