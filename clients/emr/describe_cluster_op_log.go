package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterOpLogRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
}

func (req *DescribeClusterOpLogRequest) Invoke(client *sdk.Client) (resp *DescribeClusterOpLogResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeClusterOpLog", "", "")
	resp = &DescribeClusterOpLogResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterOpLogResponse struct {
	responses.BaseResponse
	RequestId     string
	ChangeLogList DescribeClusterOpLogChangeLogList
}

type DescribeClusterOpLogChangeLog struct {
	Id          int64
	GmtCreate   string
	GmtModified string
	TargetKey   string
	Status      string
	ChangeType  string
	Message     string
	TargetType  string
}

type DescribeClusterOpLogChangeLogList []DescribeClusterOpLogChangeLog

func (list *DescribeClusterOpLogChangeLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterOpLogChangeLog)
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
