package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListClusterServiceQuickLinkRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *ListClusterServiceQuickLinkRequest) Invoke(client *sdk.Client) (resp *ListClusterServiceQuickLinkResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClusterServiceQuickLink", "", "")
	resp = &ListClusterServiceQuickLinkResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClusterServiceQuickLinkResponse struct {
	responses.BaseResponse
	RequestId     string
	QuickLinkList ListClusterServiceQuickLinkQuickLinkList
}

type ListClusterServiceQuickLinkQuickLink struct {
	ServiceName        string
	ServiceDisplayName string
	QuickLinkAddress   string
	Protocol           string
	Port               string
	Type               string
}

type ListClusterServiceQuickLinkQuickLinkList []ListClusterServiceQuickLinkQuickLink

func (list *ListClusterServiceQuickLinkQuickLinkList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceQuickLinkQuickLink)
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
