package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListTagsRequest struct {
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	Lang      string `position:"Query" name:"Lang"`
}

func (r ListTagsRequest) Invoke(client *sdk.Client) (response *ListTagsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListTagsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListTags", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		ListTagsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListTagsResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListTagsResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Tags      ListTagsTagList
}

type ListTagsTag struct {
	Id        int64
	Name      string
	IsSubTag  bool
	ParentTag string
	Cover     ListTagsCover
}

type ListTagsCover struct {
	Id      int64
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
