package crm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyContacterRequest struct {
	requests.RpcRequest
	ContacterId       int64  `position:"Query" name:"ContacterId"`
	ContacterName     string `position:"Query" name:"ContacterName"`
	ContacterEmail    string `position:"Query" name:"ContacterEmail"`
	ContacterMobile   string `position:"Query" name:"ContacterMobile"`
	ContacterPosition string `position:"Query" name:"ContacterPosition"`
}

func (req *ModifyContacterRequest) Invoke(client *sdk.Client) (resp *ModifyContacterResponse, err error) {
	req.InitWithApiInfo("Crm", "2015-03-24", "ModifyContacter", "", "")
	resp = &ModifyContacterResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyContacterResponse struct {
	responses.BaseResponse
	Success       bool
	ResultCode    common.String
	ResultMessage common.String
}
