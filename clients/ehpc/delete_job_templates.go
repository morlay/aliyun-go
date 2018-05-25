package ehpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteJobTemplatesRequest struct {
	requests.RpcRequest
	Templates string `position:"Query" name:"Templates"`
}

func (req *DeleteJobTemplatesRequest) Invoke(client *sdk.Client) (resp *DeleteJobTemplatesResponse, err error) {
	req.InitWithApiInfo("EHPC", "2018-04-12", "DeleteJobTemplates", "ehs", "")
	resp = &DeleteJobTemplatesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteJobTemplatesResponse struct {
	responses.BaseResponse
	RequestId common.String
}
