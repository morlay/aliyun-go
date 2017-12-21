package push

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListSummaryAppsRequest struct {
}

func (r ListSummaryAppsRequest) Invoke(client *sdk.Client) (response *ListSummaryAppsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListSummaryAppsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Push", "2016-08-01", "ListSummaryApps", "", "")

	resp := struct {
		*responses.BaseResponse
		ListSummaryAppsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListSummaryAppsResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListSummaryAppsResponse struct {
	RequestId       string
	SummaryAppInfos ListSummaryAppsSummaryAppInfoList
}

type ListSummaryAppsSummaryAppInfo struct {
	AppName string
	AppKey  int64
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
