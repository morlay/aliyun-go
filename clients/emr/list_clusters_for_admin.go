package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId  common.String
	TotalCount common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	Clusters   ListClustersForAdminClusterInfoList
}

type ListClustersForAdminClusterInfo struct {
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
	OrderTaskInfo       ListClustersForAdminOrderTaskInfo
	FailReason          ListClustersForAdminFailReason
}

type ListClustersForAdminOrderTaskInfo struct {
	TargetCount  common.Integer
	CurrentCount common.Integer
	OrderIdList  common.String
}

type ListClustersForAdminFailReason struct {
	ErrorCode common.String
	ErrorMsg  common.String
	RequestId common.String
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
