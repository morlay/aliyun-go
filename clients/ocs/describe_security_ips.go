package ocs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeSecurityIpsRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
}

func (r DescribeSecurityIpsRequest) Invoke(client *sdk.Client) (response *DescribeSecurityIpsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeSecurityIpsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ocs", "2015-03-01", "DescribeSecurityIps", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeSecurityIpsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeSecurityIpsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeSecurityIpsResponse struct {
	RequestId                         string
	DescribeOcsSecurityIpsResponseDTO DescribeSecurityIpsDescribeOcsSecurityIpsResponseDTO
}

type DescribeSecurityIpsDescribeOcsSecurityIpsResponseDTO struct {
	SecurityIps string
}
