package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryClusterOrdersRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	Status          string `position:"Query" name:"Status"`
}

func (req *QueryClusterOrdersRequest) Invoke(client *sdk.Client) (resp *QueryClusterOrdersResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "QueryClusterOrders", "", "")
	resp = &QueryClusterOrdersResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryClusterOrdersResponse struct {
	responses.BaseResponse
	RequestId string
	OrderList QueryClusterOrdersOrderList
}

type QueryClusterOrdersOrder struct {
	OrderId         string
	CreateTimestamp int64
}

type QueryClusterOrdersOrderList []QueryClusterOrdersOrder

func (list *QueryClusterOrdersOrderList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryClusterOrdersOrder)
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
