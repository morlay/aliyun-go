package teslamaxcompute

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	Code      int
	Message   string
	RequestId string
	Data      QueryResourceInventoryData
}

type QueryResourceInventoryData struct {
	LastUpdate string
	Clusters   QueryResourceInventoryClusterList
}

type QueryResourceInventoryCluster struct {
	Status              string
	Cluster             string
	MachineRoom         string
	Region              string
	ResourceParameters  QueryResourceInventoryResourceParameterList
	ResourceInventories QueryResourceInventoryResourceInventoryList
}

type QueryResourceInventoryResourceParameter struct {
	ParaName  string
	ParaValue string
}

type QueryResourceInventoryResourceInventory struct {
	Total        int64
	Available    int64
	Used         int64
	ResourceType string
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
