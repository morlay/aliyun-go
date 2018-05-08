package push

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListSummaryAppsRequest struct {
	requests.RpcRequest
}

func (req *ListSummaryAppsRequest) Invoke(client *sdk.Client) (resp *ListSummaryAppsResponse, err error) {
	req.InitWithApiInfo("Push", "2016-08-01", "ListSummaryApps", "", "")
	resp = &ListSummaryAppsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListSummaryAppsResponse struct {
	responses.BaseResponse
	RequestId       common.String
	SummaryAppInfos ListSummaryAppsSummaryAppInfoList
}

type ListSummaryAppsSummaryAppInfo struct {
	AppName common.String
	AppKey  common.Long
}

type ListSummaryAppsSummaryAppInfoList []ListSummaryAppsSummaryAppInfo

func (list *ListSummaryAppsSummaryAppInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListSummaryAppsSummaryAppInfo)
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
