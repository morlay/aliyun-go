package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddNodesRequest struct {
	requests.RpcRequest
	AutoRenewPeriod       int    `position:"Query" name:"AutoRenewPeriod"`
	Period                int    `position:"Query" name:"Period"`
	ImageId               string `position:"Query" name:"ImageId"`
	Count                 int    `position:"Query" name:"Count"`
	ClusterId             string `position:"Query" name:"ClusterId"`
	ComputeSpotStrategy   string `position:"Query" name:"ComputeSpotStrategy"`
	ImageOwnerAlias       string `position:"Query" name:"ImageOwnerAlias"`
	PeriodUnit            string `position:"Query" name:"PeriodUnit"`
	AutoRenew             string `position:"Query" name:"AutoRenew"`
	EcsChargeType         string `position:"Query" name:"EcsChargeType"`
	ComputeSpotPriceLimit string `position:"Query" name:"ComputeSpotPriceLimit"`
}

func (req *AddNodesRequest) Invoke(client *sdk.Client) (resp *AddNodesResponse, err error) {
	req.InitWithApiInfo("EHPC", "2018-04-12", "AddNodes", "ehs", "")
	resp = &AddNodesResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddNodesResponse struct {
	responses.BaseResponse
	RequestId   common.String
	InstanceIds AddNodesInstanceIdList
}

type AddNodesInstanceIdList []common.String

func (list *AddNodesInstanceIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
