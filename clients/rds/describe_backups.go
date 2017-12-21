package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeBackupsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	BackupId             string `position:"Query" name:"BackupId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	BackupStatus         string `position:"Query" name:"BackupStatus"`
	BackupLocation       string `position:"Query" name:"BackupLocation"`
	PageSize             int    `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	BackupMode           string `position:"Query" name:"BackupMode"`
}

func (r DescribeBackupsRequest) Invoke(client *sdk.Client) (response *DescribeBackupsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeBackupsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeBackups", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeBackupsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeBackupsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeBackupsResponse struct {
	RequestId        string
	TotalRecordCount string
	PageNumber       string
	PageRecordCount  string
	TotalBackupSize  int64
	Items            DescribeBackupsBackupList
}

type DescribeBackupsBackup struct {
	BackupId                  string
	DBInstanceId              string
	BackupStatus              string
	BackupStartTime           string
	BackupEndTime             string
	BackupType                string
	BackupMode                string
	BackupMethod              string
	BackupDownloadURL         string
	BackupIntranetDownloadURL string
	BackupLocation            string
	BackupExtractionStatus    string
	BackupScale               string
	BackupDBNames             string
	TotalBackupSize           int64
	BackupSize                int64
	HostInstanceID            string
	StoreStatus               string
}

type DescribeBackupsBackupList []DescribeBackupsBackup

func (list *DescribeBackupsBackupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBackupsBackup)
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
