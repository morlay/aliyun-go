package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListClustersRequest struct {
	requests.RpcRequest
	PageSize   int `position:"Query" name:"PageSize"`
	PageNumber int `position:"Query" name:"PageNumber"`
}

func (req *ListClustersRequest) Invoke(client *sdk.Client) (resp *ListClustersResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "ListClusters", "ehs", "")
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
	Clusters   ListClustersClusterInfoSimpleList
}

type ListClustersClusterInfoSimple struct {
	Id              string
	RegionId        string
	ZoneId          string
	Name            string
	Description     string
	Status          string
	OsTag           string
	AccountType     string
	SchedulerType   string
	Count           int
	InstanceType    string
	LoginNodes      string
	CreateTime      string
	ImageOwnerAlias string
	ImageId         string
	Managers        ListClustersManagers
	Computes        ListClustersComputes
	TotalResources  ListClustersTotalResources
	UsedResources   ListClustersUsedResources
}

type ListClustersManagers struct {
	Toatal         int
	NormalCount    int
	ExceptionCount int
}

type ListClustersComputes struct {
	Toatal         int
	NormalCount    int
	ExceptionCount int
}

type ListClustersTotalResources struct {
	Cpu    int
	Memory int
	Gpu    int
}

type ListClustersUsedResources struct {
	Cpu    int
	Memory int
	Gpu    int
}

type ListClustersClusterInfoSimpleList []ListClustersClusterInfoSimple

func (list *ListClustersClusterInfoSimpleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClustersClusterInfoSimple)
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
