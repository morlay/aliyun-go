package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeUserConfigsRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Config        string `position:"Query" name:"Config"`
}

func (r DescribeUserConfigsRequest) Invoke(client *sdk.Client) (response *DescribeUserConfigsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeUserConfigsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeUserConfigs", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeUserConfigsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeUserConfigsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeUserConfigsResponse struct {
	RequestId string
	Configs   DescribeUserConfigsConfigs
}

type DescribeUserConfigsConfigs struct {
	OssLogConfig       DescribeUserConfigsOssLogConfig
	GreenManagerConfig DescribeUserConfigsGreenManagerConfig
}

type DescribeUserConfigsOssLogConfig struct {
	Enable string
	Bucket string
	Prefix string
}

type DescribeUserConfigsGreenManagerConfig struct {
	Quota string
	Ratio string
}
