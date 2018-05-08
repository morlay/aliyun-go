package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListPhotoTagsRequest struct {
	requests.RpcRequest
	LibraryId string `position:"Query" name:"LibraryId"`
	PhotoId   int64  `position:"Query" name:"PhotoId"`
	StoreName string `position:"Query" name:"StoreName"`
	Lang      string `position:"Query" name:"Lang"`
}

func (req *ListPhotoTagsRequest) Invoke(client *sdk.Client) (resp *ListPhotoTagsResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListPhotoTags", "cloudphoto", "")
	resp = &ListPhotoTagsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListPhotoTagsResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
	Tags      ListPhotoTagsTagList
}

type ListPhotoTagsTag struct {
	Id        common.Long
	IdStr     common.String
	IsSubTag  bool
	Name      common.String
	ParentTag common.String
}

type ListPhotoTagsTagList []ListPhotoTagsTag

func (list *ListPhotoTagsTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPhotoTagsTag)
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
