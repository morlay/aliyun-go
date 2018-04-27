package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteClusterScriptRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
}

func (req *DeleteClusterScriptRequest) Invoke(client *sdk.Client) (resp *DeleteClusterScriptResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DeleteClusterScript", "", "")
	resp = &DeleteClusterScriptResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteClusterScriptResponse struct {
	responses.BaseResponse
	RequestId string
}
