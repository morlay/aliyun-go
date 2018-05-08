package ocs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeSecurityIpsRequest struct {
	requests.RpcRequest
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
}

func (req *DescribeSecurityIpsRequest) Invoke(client *sdk.Client) (resp *DescribeSecurityIpsResponse, err error) {
	req.InitWithApiInfo("Ocs", "2015-03-01", "DescribeSecurityIps", "", "")
	resp = &DescribeSecurityIpsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSecurityIpsResponse struct {
	responses.BaseResponse
	RequestId                         common.String
	DescribeOcsSecurityIpsResponseDTO DescribeSecurityIpsDescribeOcsSecurityIpsResponseDTO
}

type DescribeSecurityIpsDescribeOcsSecurityIpsResponseDTO struct {
	SecurityIps common.String
}
