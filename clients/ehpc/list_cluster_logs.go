package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListClusterLogsRequest struct {
	requests.RpcRequest
	PageSize   int    `position:"Query" name:"PageSize"`
	ClusterId  string `position:"Query" name:"ClusterId"`
	PageNumber int    `position:"Query" name:"PageNumber"`
}

func (req *ListClusterLogsRequest) Invoke(client *sdk.Client) (resp *ListClusterLogsResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "ListClusterLogs", "ehs", "")
	resp = &ListClusterLogsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClusterLogsResponse struct {
	responses.BaseResponse
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	ClusterId  string
	Logs       ListClusterLogsLogInfoList
}

type ListClusterLogsLogInfo struct {
	Operation  string
	Level      string
	Message    string
	CreateTime string
}

type ListClusterLogsLogInfoList []ListClusterLogsLogInfo

func (list *ListClusterLogsLogInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterLogsLogInfo)
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
