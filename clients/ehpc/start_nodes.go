package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type StartNodesRequest struct {
	requests.RpcRequest
	Role      string                  `position:"Query" name:"Role"`
	Instances *StartNodesInstanceList `position:"Query" type:"Repeated" name:"Instance"`
	ClusterId string                  `position:"Query" name:"ClusterId"`
}

func (req *StartNodesRequest) Invoke(client *sdk.Client) (resp *StartNodesResponse, err error) {
	req.InitWithApiInfo("EHPC", "2018-04-12", "StartNodes", "ehs", "")
	resp = &StartNodesResponse{}
	err = client.DoAction(req, resp)
	return
}

type StartNodesInstance struct {
	Id string `name:"Id"`
}

type StartNodesResponse struct {
	responses.BaseResponse
	RequestId common.String
}

type StartNodesInstanceList []StartNodesInstance

func (list *StartNodesInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]StartNodesInstance)
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
