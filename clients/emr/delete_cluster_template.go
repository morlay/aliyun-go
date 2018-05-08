package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteClusterTemplateRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
}

func (req *DeleteClusterTemplateRequest) Invoke(client *sdk.Client) (resp *DeleteClusterTemplateResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DeleteClusterTemplate", "", "")
	resp = &DeleteClusterTemplateResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteClusterTemplateResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   common.String
	ErrCode   common.String
	ErrMsg    common.String
}
