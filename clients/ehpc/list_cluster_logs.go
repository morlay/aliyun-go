package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId  common.String
	TotalCount common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	ClusterId  common.String
	Logs       ListClusterLogsLogInfoList
}

type ListClusterLogsLogInfo struct {
	Operation  common.String
	Level      common.String
	Message    common.String
	CreateTime common.String
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
