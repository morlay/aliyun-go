package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId     common.String
	ChangeLogList DescribeClusterOpLogChangeLogList
}

type DescribeClusterOpLogChangeLog struct {
	Id          common.Long
	GmtCreate   common.String
	GmtModified common.String
	TargetKey   common.String
	Status      common.String
	ChangeType  common.String
	Message     common.String
	TargetType  common.String
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
