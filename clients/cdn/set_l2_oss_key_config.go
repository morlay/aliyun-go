package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetL2OssKeyConfigRequest struct {
	requests.RpcRequest
	SecurityToken  string `position:"Query" name:"SecurityToken"`
	DomainName     string `position:"Query" name:"DomainName"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
	PrivateOssAuth string `position:"Query" name:"PrivateOssAuth"`
}

func (req *SetL2OssKeyConfigRequest) Invoke(client *sdk.Client) (resp *SetL2OssKeyConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetL2OssKeyConfig", "", "")
	resp = &SetL2OssKeyConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetL2OssKeyConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
