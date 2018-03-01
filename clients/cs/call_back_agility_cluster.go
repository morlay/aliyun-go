package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CallBackAgilityClusterRequest struct {
	requests.RoaRequest
	ReqOnce string `position:"Path" name:"ReqOnce"`
	Token   string `position:"Path" name:"Token"`
}

func (req *CallBackAgilityClusterRequest) Invoke(client *sdk.Client) (resp *CallBackAgilityClusterResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "CallBackAgilityCluster", "/agility/token/[Token]/req_once/[ReqOnce]/callback", "", "")
	req.Method = "POST"

	resp = &CallBackAgilityClusterResponse{}
	err = client.DoAction(req, resp)
	return
}

type CallBackAgilityClusterResponse struct {
	responses.BaseResponse
}
