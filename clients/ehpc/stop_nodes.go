package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type StopNodesRequest struct {
	requests.RpcRequest
	Role      string                 `position:"Query" name:"Role"`
	Instances *StopNodesInstanceList `position:"Query" type:"Repeated" name:"Instance"`
	ClusterId string                 `position:"Query" name:"ClusterId"`
}

func (req *StopNodesRequest) Invoke(client *sdk.Client) (resp *StopNodesResponse, err error) {
	req.InitWithApiInfo("EHPC", "2018-04-12", "StopNodes", "ehs", "")
	resp = &StopNodesResponse{}
	err = client.DoAction(req, resp)
	return
}

type StopNodesInstance struct {
	Id string `name:"Id"`
}

type StopNodesResponse struct {
	responses.BaseResponse
	RequestId common.String
}

type StopNodesInstanceList []StopNodesInstance

func (list *StopNodesInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]StopNodesInstance)
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
