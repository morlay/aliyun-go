package tesladam

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type HostGetsRequest struct {
	requests.RpcRequest
	Query     string `position:"Query" name:"Query"`
	EndTime   int    `position:"Query" name:"EndTime"`
	StartTime int    `position:"Query" name:"StartTime"`
	QueryType string `position:"Query" name:"QueryType"`
}

func (req *HostGetsRequest) Invoke(client *sdk.Client) (resp *HostGetsResponse, err error) {
	req.InitWithApiInfo("TeslaDam", "2018-01-18", "HostGets", "", "")
	resp = &HostGetsResponse{}
	err = client.DoAction(req, resp)
	return
}

type HostGetsResponse struct {
	responses.BaseResponse
	Status  bool
	Message common.String
	Data    HostGetsDataItemList
}

type HostGetsDataItem struct {
	Hostname         common.String
	Ip               common.String
	AppCode          common.String
	ClusterCode      common.String
	SshStatus        common.Integer
	HeartStatus      common.Integer
	HealthScoreLast  common.Integer
	HealthReasonLast common.String
}

type HostGetsDataItemList []HostGetsDataItem

func (list *HostGetsDataItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]HostGetsDataItem)
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
