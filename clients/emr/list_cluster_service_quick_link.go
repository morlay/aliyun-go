package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId     common.String
	QuickLinkList ListClusterServiceQuickLinkQuickLinkList
}

type ListClusterServiceQuickLinkQuickLink struct {
	ServiceName        common.String
	ServiceDisplayName common.String
	QuickLinkAddress   common.String
	Protocol           common.String
	Port               common.String
	Type               common.String
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
