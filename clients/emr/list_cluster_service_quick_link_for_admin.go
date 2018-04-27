package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListClusterServiceQuickLinkForAdminRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *ListClusterServiceQuickLinkForAdminRequest) Invoke(client *sdk.Client) (resp *ListClusterServiceQuickLinkForAdminResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClusterServiceQuickLinkForAdmin", "", "")
	resp = &ListClusterServiceQuickLinkForAdminResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClusterServiceQuickLinkForAdminResponse struct {
	responses.BaseResponse
	RequestId     string
	QuickLinkList ListClusterServiceQuickLinkForAdminQuickLinkList
}

type ListClusterServiceQuickLinkForAdminQuickLink struct {
	ServiceName        string
	ServiceDisplayName string
	QuickLinkAddress   string
	Protocol           string
}

type ListClusterServiceQuickLinkForAdminQuickLinkList []ListClusterServiceQuickLinkForAdminQuickLink

func (list *ListClusterServiceQuickLinkForAdminQuickLinkList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceQuickLinkForAdminQuickLink)
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
