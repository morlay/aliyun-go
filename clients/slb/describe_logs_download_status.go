package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLogsDownloadStatusRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r DescribeLogsDownloadStatusRequest) Invoke(client *sdk.Client) (response *DescribeLogsDownloadStatusResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLogsDownloadStatusRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeLogsDownloadStatus", "slb", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLogsDownloadStatusResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLogsDownloadStatusResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLogsDownloadStatusResponse struct {
	RequestId          string
	LogsDownloadStatus string
}
