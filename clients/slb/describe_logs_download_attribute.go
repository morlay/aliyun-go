package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLogsDownloadAttributeRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r DescribeLogsDownloadAttributeRequest) Invoke(client *sdk.Client) (response *DescribeLogsDownloadAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLogsDownloadAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeLogsDownloadAttribute", "slb", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLogsDownloadAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLogsDownloadAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLogsDownloadAttributeResponse struct {
	RequestId     string
	LogType       string
	OSSBucketName string
	RoleName      string
}
