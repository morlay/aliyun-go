package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddNodesRequest struct {
	requests.RpcRequest
	ImageId               string `position:"Query" name:"ImageId"`
	Count                 int    `position:"Query" name:"Count"`
	ClusterId             string `position:"Query" name:"ClusterId"`
	ComputeSpotStrategy   string `position:"Query" name:"ComputeSpotStrategy"`
	ComputeSpotPriceLimit string `position:"Query" name:"ComputeSpotPriceLimit"`
	ImageOwnerAlias       string `position:"Query" name:"ImageOwnerAlias"`
}

func (req *AddNodesRequest) Invoke(client *sdk.Client) (resp *AddNodesResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "AddNodes", "ehs", "")
	resp = &AddNodesResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddNodesResponse struct {
	responses.BaseResponse
	RequestId   string
	InstanceIds AddNodesInstanceIdList
}

type AddNodesInstanceIdList []string

func (list *AddNodesInstanceIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
