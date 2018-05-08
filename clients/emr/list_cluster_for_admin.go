package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListClusterForAdminRequest struct {
	requests.RpcRequest
	ResourceOwnerId   int64                                    `position:"Query" name:"ResourceOwnerId"`
	StatusLists       *ListClusterForAdminStatusListList       `position:"Query" type:"Repeated" name:"StatusList"`
	FuzzyName         string                                   `position:"Query" name:"FuzzyName"`
	UserId            string                                   `position:"Query" name:"UserId"`
	PageNumber        int                                      `position:"Query" name:"PageNumber"`
	EcmClusterIdLists *ListClusterForAdminEcmClusterIdListList `position:"Query" type:"Repeated" name:"EcmClusterIdList"`
	ClusterIdLists    *ListClusterForAdminClusterIdListList    `position:"Query" type:"Repeated" name:"ClusterIdList"`
	PayTypeLists      *ListClusterForAdminPayTypeListList      `position:"Query" type:"Repeated" name:"PayTypeList"`
	Name              string                                   `position:"Query" name:"Name"`
	PageSize          int                                      `position:"Query" name:"PageSize"`
	EmrVersion        string                                   `position:"Query" name:"EmrVersion"`
	Resize            string                                   `position:"Query" name:"Resize"`
	ClusterTypeLists  *ListClusterForAdminClusterTypeListList  `position:"Query" type:"Repeated" name:"ClusterTypeList"`
}

func (req *ListClusterForAdminRequest) Invoke(client *sdk.Client) (resp *ListClusterForAdminResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClusterForAdmin", "", "")
	resp = &ListClusterForAdminResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClusterForAdminResponse struct {
	responses.BaseResponse
	RequestId   common.String
	TotalCount  common.Integer
	PageNumber  common.Integer
	PageSize    common.Integer
	ClusterList ListClusterForAdminClusterList
}

type ListClusterForAdminCluster struct {
	Id                     common.String
	Name                   common.String
	BizId                  common.String
	LusterType             common.String
	CreateTime             common.Long
	RunningTime            common.Long
	Status                 common.String
	PayType                common.String
	RegionId               common.String
	EasEnable              bool
	DepositType            common.String
	UseLocalMetadb         bool
	HighAvailabilityEnable bool
	NodeCount              common.Integer
	ExpiredTime            common.Long
	NetType                common.String
	HasUncompletedOrder    bool
	OrderList              common.String
	CreateResource         common.String
	UserId                 common.String
	EcmClusterId           common.String
	EmrVersion             common.String
	VpcId                  common.String
	VSwitchId              common.String
	OrderTaskInfo          ListClusterForAdminOrderTaskInfo
	FailReason             ListClusterForAdminFailReason
}

type ListClusterForAdminOrderTaskInfo struct {
	TargetCount  common.Integer
	CurrentCount common.Integer
	OrderIdList  common.String
}

type ListClusterForAdminFailReason struct {
	ErrorCode common.String
	ErrorMsg  common.String
	RequestId common.String
}

type ListClusterForAdminStatusListList []string

func (list *ListClusterForAdminStatusListList) UnmarshalJSON(data []byte) error {
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

type ListClusterForAdminEcmClusterIdListList []int64

func (list *ListClusterForAdminEcmClusterIdListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type ListClusterForAdminClusterIdListList []string

func (list *ListClusterForAdminClusterIdListList) UnmarshalJSON(data []byte) error {
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

type ListClusterForAdminPayTypeListList []string

func (list *ListClusterForAdminPayTypeListList) UnmarshalJSON(data []byte) error {
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

type ListClusterForAdminClusterTypeListList []string

func (list *ListClusterForAdminClusterTypeListList) UnmarshalJSON(data []byte) error {
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

type ListClusterForAdminClusterList []ListClusterForAdminCluster

func (list *ListClusterForAdminClusterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterForAdminCluster)
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
