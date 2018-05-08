package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeBackupTasksRequest struct {
	requests.RpcRequest
	BackupJobId          string `position:"Query" name:"BackupJobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Flag                 string `position:"Query" name:"Flag"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	BackupMode           string `position:"Query" name:"BackupMode"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	BackupJobStatus      string `position:"Query" name:"BackupJobStatus"`
}

func (req *DescribeBackupTasksRequest) Invoke(client *sdk.Client) (resp *DescribeBackupTasksResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeBackupTasks", "rds", "")
	resp = &DescribeBackupTasksResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeBackupTasksResponse struct {
	responses.BaseResponse
	RequestId common.String
	Items     DescribeBackupTasksBackupJobList
}

type DescribeBackupTasksBackupJob struct {
	BackupProgressStatus common.String
	JobMode              common.String
	Process              common.String
	TaskAction           common.String
	BackupjobId          common.String
}

type DescribeBackupTasksBackupJobList []DescribeBackupTasksBackupJob

func (list *DescribeBackupTasksBackupJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBackupTasksBackupJob)
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
