package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListClusterForInternalRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64                                    `position:"Query" name:"ResourceOwnerId"`
	ClusterIdLists  *ListClusterForInternalClusterIdListList `position:"Query" type:"Repeated" name:"ClusterIdList"`
	UserId          string                                   `position:"Query" name:"UserId"`
}

func (req *ListClusterForInternalRequest) Invoke(client *sdk.Client) (resp *ListClusterForInternalResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClusterForInternal", "", "")
	resp = &ListClusterForInternalResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClusterForInternalResponse struct {
	responses.BaseResponse
	RequestId                   string
	DescribeClusterResponseList ListClusterForInternalDescribeClusterResponseList
}

type ListClusterForInternalDescribeClusterResponse struct {
	RequestId   string
	Id          string
	BizId       string
	Name        string
	StartTime   int64
	StopTime    int64
	LogEnable   bool
	LogPath     string
	UserId      string
	RunningTime int
	Status      string
	ExpiredTime int64
	FailReason  ListClusterForInternalFailReason
}

type ListClusterForInternalFailReason struct {
	ErrorCode string
	ErrorMsg  string
	RequestId string
}

type ListClusterForInternalClusterIdListList []string

func (list *ListClusterForInternalClusterIdListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type ListClusterForInternalDescribeClusterResponseList []ListClusterForInternalDescribeClusterResponse

func (list *ListClusterForInternalDescribeClusterResponseList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterForInternalDescribeClusterResponse)
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
