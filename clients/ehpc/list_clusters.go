package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId  common.String
	TotalCount common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	Clusters   ListClustersClusterInfoSimpleList
}

type ListClustersClusterInfoSimple struct {
	Id              common.String
	RegionId        common.String
	ZoneId          common.String
	Name            common.String
	Description     common.String
	Status          common.String
	OsTag           common.String
	AccountType     common.String
	SchedulerType   common.String
	Count           common.Integer
	InstanceType    common.String
	LoginNodes      common.String
	CreateTime      common.String
	ImageOwnerAlias common.String
	ImageId         common.String
	Managers        ListClustersManagers
	Computes        ListClustersComputes
	TotalResources  ListClustersTotalResources
	UsedResources   ListClustersUsedResources
}

type ListClustersManagers struct {
	Toatal         common.Integer
	NormalCount    common.Integer
	ExceptionCount common.Integer
}

type ListClustersComputes struct {
	Toatal         common.Integer
	NormalCount    common.Integer
	ExceptionCount common.Integer
}

type ListClustersTotalResources struct {
	Cpu    common.Integer
	Memory common.Integer
	Gpu    common.Integer
}

type ListClustersUsedResources struct {
	Cpu    common.Integer
	Memory common.Integer
	Gpu    common.Integer
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
