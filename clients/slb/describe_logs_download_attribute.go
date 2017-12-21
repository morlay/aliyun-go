package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLogsDownloadAttributeRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *DescribeLogsDownloadAttributeRequest) Invoke(client *sdk.Client) (resp *DescribeLogsDownloadAttributeResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeLogsDownloadAttribute", "slb", "")
	resp = &DescribeLogsDownloadAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLogsDownloadAttributeResponse struct {
	responses.BaseResponse
	RequestId     string
	LogType       string
	OSSBucketName string
	RoleName      string
}
