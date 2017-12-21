package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeBinlogFilesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	EndTime              string `position:"Query" name:"EndTime"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeBinlogFilesRequest) Invoke(client *sdk.Client) (resp *DescribeBinlogFilesResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeBinlogFiles", "rds", "")
	resp = &DescribeBinlogFilesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeBinlogFilesResponse struct {
	responses.BaseResponse
	RequestId        string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	TotalFileSize    int64
	Items            DescribeBinlogFilesBinLogFileList
}

type DescribeBinlogFilesBinLogFile struct {
	FileSize             int64
	LogBeginTime         string
	LogEndTime           string
	DownloadLink         string
	IntranetDownloadLink string
	LinkExpiredTime      string
	Checksum             string
	HostInstanceID       string
}

type DescribeBinlogFilesBinLogFileList []DescribeBinlogFilesBinLogFile

func (list *DescribeBinlogFilesBinLogFileList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBinlogFilesBinLogFile)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
