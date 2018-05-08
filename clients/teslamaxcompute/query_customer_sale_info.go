package teslamaxcompute

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryCustomerSaleInfoRequest struct {
	requests.RpcRequest
	RegionName string `position:"Query" name:"RegionName"`
}

func (req *QueryCustomerSaleInfoRequest) Invoke(client *sdk.Client) (resp *QueryCustomerSaleInfoResponse, err error) {
	req.InitWithApiInfo("TeslaMaxCompute", "2018-01-04", "QueryCustomerSaleInfo", "", "")
	resp = &QueryCustomerSaleInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryCustomerSaleInfoResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
	Data      QueryCustomerSaleInfoData
}

type QueryCustomerSaleInfoData struct {
	LastUpdate common.String
	Clusters   QueryCustomerSaleInfoClusterList
}

type QueryCustomerSaleInfoCluster struct {
	Cluster     common.String
	Region      common.String
	MachineRoom common.String
	SaleInfos   QueryCustomerSaleInfoSaleInfoList
}

type QueryCustomerSaleInfoSaleInfo struct {
	SaleMode    common.String
	Uid         common.String
	Mem         common.Long
	Cpu         common.Long
	BizCategory common.String
	Owner       common.String
	QueryDate   common.String
}

type QueryCustomerSaleInfoClusterList []QueryCustomerSaleInfoCluster

func (list *QueryCustomerSaleInfoClusterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCustomerSaleInfoCluster)
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

type QueryCustomerSaleInfoSaleInfoList []QueryCustomerSaleInfoSaleInfo

func (list *QueryCustomerSaleInfoSaleInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCustomerSaleInfoSaleInfo)
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
