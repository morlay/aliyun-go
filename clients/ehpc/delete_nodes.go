package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteNodesRequest struct {
	requests.RpcRequest
	ReleaseInstance string                   `position:"Query" name:"ReleaseInstance"`
	Instances       *DeleteNodesInstanceList `position:"Query" type:"Repeated" name:"Instance"`
	ClusterId       string                   `position:"Query" name:"ClusterId"`
}

func (req *DeleteNodesRequest) Invoke(client *sdk.Client) (resp *DeleteNodesResponse, err error) {
	req.InitWithApiInfo("EHPC", "2018-04-12", "DeleteNodes", "ehs", "")
	resp = &DeleteNodesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteNodesInstance struct {
	Id string `name:"Id"`
}

type DeleteNodesResponse struct {
	responses.BaseResponse
	RequestId common.String
}

type DeleteNodesInstanceList []DeleteNodesInstance

func (list *DeleteNodesInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DeleteNodesInstance)
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
