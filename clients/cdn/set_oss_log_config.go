package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetOssLogConfigRequest struct {
	requests.RpcRequest
	Bucket        string `position:"Query" name:"Bucket"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Enable        string `position:"Query" name:"Enable"`
	Prefix        string `position:"Query" name:"Prefix"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *SetOssLogConfigRequest) Invoke(client *sdk.Client) (resp *SetOssLogConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetOssLogConfig", "", "")
	resp = &SetOssLogConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetOssLogConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
