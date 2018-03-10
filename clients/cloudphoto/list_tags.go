package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	Code      string
	Message   string
	RequestId string
	Action    string
	Tags      ListTagsTagList
}

type ListTagsTag struct {
	Id        int64
	IdStr     string
	Name      string
	IsSubTag  bool
	ParentTag string
	Cover     ListTagsCover
}

type ListTagsCover struct {
	Id      int64
	IdStr   string
	Title   string
	FileId  string
	State   string
	Md5     string
	IsVideo bool
	Remark  string
	Width   int64
	Height  int64
	Ctime   int64
	Mtime   int64
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
