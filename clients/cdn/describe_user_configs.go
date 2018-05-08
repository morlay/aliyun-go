package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeUserConfigsRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Config        string `position:"Query" name:"Config"`
}

func (req *DescribeUserConfigsRequest) Invoke(client *sdk.Client) (resp *DescribeUserConfigsResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeUserConfigs", "", "")
	resp = &DescribeUserConfigsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeUserConfigsResponse struct {
	responses.BaseResponse
	RequestId common.String
	Configs   DescribeUserConfigsConfigs
}

type DescribeUserConfigsConfigs struct {
	OssLogConfig       DescribeUserConfigsOssLogConfig
	GreenManagerConfig DescribeUserConfigsGreenManagerConfig
}

type DescribeUserConfigsOssLogConfig struct {
	Enable common.String
	Bucket common.String
	Prefix common.String
}

type DescribeUserConfigsGreenManagerConfig struct {
	Quota common.String
	Ratio common.String
}
