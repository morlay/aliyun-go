package teslamaxcompute

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryResourceInventoryRequest struct {
	requests.RpcRequest
}

func (req *QueryResourceInventoryRequest) Invoke(client *sdk.Client) (resp *QueryResourceInventoryResponse, err error) {
	req.InitWithApiInfo("TeslaMaxCompute", "2018-01-04", "QueryResourceInventory", "", "")
	resp = &QueryResourceInventoryResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryResourceInventoryResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
	Data      QueryResourceInventoryData
}

type QueryResourceInventoryData struct {
	LastUpdate common.String
	Clusters   QueryResourceInventoryClusterList
}

type QueryResourceInventoryCluster struct {
	Status              common.String
	Cluster             common.String
	MachineRoom         common.String
	Region              common.String
	ResourceParameters  QueryResourceInventoryResourceParameterList
	ResourceInventories QueryResourceInventoryResourceInventoryList
}

type QueryResourceInventoryResourceParameter struct {
	ParaName  common.String
	ParaValue common.String
}

type QueryResourceInventoryResourceInventory struct {
	Total        common.Long
	Available    common.Long
	Used         common.Long
	ResourceType common.String
}

type QueryResourceInventoryClusterList []QueryResourceInventoryCluster

func (list *QueryResourceInventoryClusterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryResourceInventoryCluster)
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

type QueryResourceInventoryResourceParameterList []QueryResourceInventoryResourceParameter

func (list *QueryResourceInventoryResourceParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryResourceInventoryResourceParameter)
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

type QueryResourceInventoryResourceInventoryList []QueryResourceInventoryResourceInventory

func (list *QueryResourceInventoryResourceInventoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryResourceInventoryResourceInventory)
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
