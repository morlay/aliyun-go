package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateClusterScriptRequest struct {
	requests.RpcRequest
	Args            string `position:"Query" name:"Args"`
	Path            string `position:"Query" name:"Path"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Name            string `position:"Query" name:"Name"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	NodeIdList      string `position:"Query" name:"NodeIdList"`
}

func (req *CreateClusterScriptRequest) Invoke(client *sdk.Client) (resp *CreateClusterScriptResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "CreateClusterScript", "", "")
	resp = &CreateClusterScriptResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateClusterScriptResponse struct {
	responses.BaseResponse
	RequestId string
	Id        string
}
