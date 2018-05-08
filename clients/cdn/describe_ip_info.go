package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeIpInfoRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	IP            string `position:"Query" name:"IP"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeIpInfoRequest) Invoke(client *sdk.Client) (resp *DescribeIpInfoResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeIpInfo", "", "")
	resp = &DescribeIpInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeIpInfoResponse struct {
	responses.BaseResponse
	RequestId   common.String
	CdnIp       common.String
	ISP         common.String
	IspEname    common.String
	Region      common.String
	RegionEname common.String
}
