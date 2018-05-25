package pvtz

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeUserServiceStatusRequest struct {
	requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

func (req *DescribeUserServiceStatusRequest) Invoke(client *sdk.Client) (resp *DescribeUserServiceStatusResponse, err error) {
	req.InitWithApiInfo("pvtz", "2018-01-01", "DescribeUserServiceStatus", "pvtz", "")
	resp = &DescribeUserServiceStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeUserServiceStatusResponse struct {
	responses.BaseResponse
	RequestId common.String
	Status    common.String
}
