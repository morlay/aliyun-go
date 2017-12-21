package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeIpInfoRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	IP            string `position:"Query" name:"IP"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeIpInfoRequest) Invoke(client *sdk.Client) (response *DescribeIpInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeIpInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeIpInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeIpInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeIpInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeIpInfoResponse struct {
	RequestId   string
	CdnIp       string
	ISP         string
	IspEname    string
	Region      string
	RegionEname string
}
