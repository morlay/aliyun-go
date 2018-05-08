package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListTagsRequest struct {
	requests.RpcRequest
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	Lang      string `position:"Query" name:"Lang"`
}

func (req *ListTagsRequest) Invoke(client *sdk.Client) (resp *ListTagsResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListTags", "cloudphoto", "")
	resp = &ListTagsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListTagsResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
	Tags      ListTagsTagList
}

type ListTagsTag struct {
	Id        common.Long
	IdStr     common.String
	Name      common.String
	IsSubTag  bool
	ParentTag common.String
	Cover     ListTagsCover
}

type ListTagsCover struct {
	Id      common.Long
	IdStr   common.String
	Title   common.String
	FileId  common.String
	State   common.String
	Md5     common.String
	IsVideo bool
	Remark  common.String
	Width   common.Long
	Height  common.Long
	Ctime   common.Long
	Mtime   common.Long
}

type ListTagsTagList []ListTagsTag

func (list *ListTagsTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTagsTag)
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
