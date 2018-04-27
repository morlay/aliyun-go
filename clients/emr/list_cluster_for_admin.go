package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId   string
	TotalCount  int
	PageNumber  int
	PageSize    int
	ClusterList ListClusterForAdminClusterList
}

type ListClusterForAdminCluster struct {
	Id                     string
	Name                   string
	BizId                  string
	LusterType             string
	CreateTime             int64
	RunningTime            int64
	Status                 string
	PayType                string
	RegionId               string
	EasEnable              bool
	DepositType            string
	UseLocalMetadb         bool
	HighAvailabilityEnable bool
	NodeCount              int
	ExpiredTime            int64
	NetType                string
	HasUncompletedOrder    bool
	OrderList              string
	CreateResource         string
	UserId                 string
	EcmClusterId           string
	EmrVersion             string
	VpcId                  string
	VSwitchId              string
	OrderTaskInfo          ListClusterForAdminOrderTaskInfo
	FailReason             ListClusterForAdminFailReason
}

type ListClusterForAdminOrderTaskInfo struct {
	TargetCount  int
	CurrentCount int
	OrderIdList  string
}

type ListClusterForAdminFailReason struct {
	ErrorCode string
	ErrorMsg  string
	RequestId string
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
