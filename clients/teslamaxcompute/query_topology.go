package teslamaxcompute

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryTopologyRequest struct {
	requests.RpcRequest
}

func (req *QueryTopologyRequest) Invoke(client *sdk.Client) (resp *QueryTopologyResponse, err error) {
	req.InitWithApiInfo("TeslaMaxCompute", "2018-01-04", "QueryTopology", "", "")
	resp = &QueryTopologyResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryTopologyResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
	Result    QueryTopologyResultItemList
}

type QueryTopologyResultItem struct {
	LastUpdate common.String
	Regions    QueryTopologyRegionItemList
}

type QueryTopologyRegionItem struct {
	Region       common.String
	RegionEnName common.String
	RegionCnName common.String
	Clusters     QueryTopologyClusterItemList
}

type QueryTopologyClusterItem struct {
	Cluster      common.String
	ProductLine  common.String
	ProductClass common.String
	NetCode      common.String
	Business     common.String
	MachineRoom  common.String
	NetArch      common.String
}

type QueryTopologyResultItemList []QueryTopologyResultItem

func (list *QueryTopologyResultItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTopologyResultItem)
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

type QueryTopologyRegionItemList []QueryTopologyRegionItem

func (list *QueryTopologyRegionItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTopologyRegionItem)
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

type QueryTopologyClusterItemList []QueryTopologyClusterItem

func (list *QueryTopologyClusterItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTopologyClusterItem)
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
