package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ResetNodesRequest struct {
	requests.RpcRequest
	Instances *ResetNodesInstanceList `position:"Query" type:"Repeated" name:"Instance"`
	ClusterId string                  `position:"Query" name:"ClusterId"`
}

func (req *ResetNodesRequest) Invoke(client *sdk.Client) (resp *ResetNodesResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "ResetNodes", "ehs", "")
	resp = &ResetNodesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ResetNodesInstance struct {
	Id string `name:"Id"`
}

type ResetNodesResponse struct {
	responses.BaseResponse
	RequestId string
}

type ResetNodesInstanceList []ResetNodesInstance

func (list *ResetNodesInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ResetNodesInstance)
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
