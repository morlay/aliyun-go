package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCustomLogConfigRequest struct {
	requests.RpcRequest
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ConfigId      string `position:"Query" name:"ConfigId"`
}

func (req *DescribeCustomLogConfigRequest) Invoke(client *sdk.Client) (resp *DescribeCustomLogConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeCustomLogConfig", "", "")
	resp = &DescribeCustomLogConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCustomLogConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
	Remark    common.String
	Sample    common.String
	Tag       common.String
}
