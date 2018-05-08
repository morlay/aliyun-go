package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type StartArchiveSQLLogRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Database             string `position:"Query" name:"Database"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	User                 string `position:"Query" name:"User"`
	QueryKeywords        string `position:"Query" name:"QueryKeywords"`
}

func (req *StartArchiveSQLLogRequest) Invoke(client *sdk.Client) (resp *StartArchiveSQLLogResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "StartArchiveSQLLog", "rds", "")
	resp = &StartArchiveSQLLogResponse{}
	err = client.DoAction(req, resp)
	return
}

type StartArchiveSQLLogResponse struct {
	responses.BaseResponse
	RequestId common.String
}
