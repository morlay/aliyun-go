package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListPhotoTagsRequest struct {
	LibraryId string `position:"Query" name:"LibraryId"`
	PhotoId   int64  `position:"Query" name:"PhotoId"`
	StoreName string `position:"Query" name:"StoreName"`
	Lang      string `position:"Query" name:"Lang"`
}

func (r ListPhotoTagsRequest) Invoke(client *sdk.Client) (response *ListPhotoTagsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListPhotoTagsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "ListPhotoTags", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		ListPhotoTagsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListPhotoTagsResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListPhotoTagsResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
	Tags      ListPhotoTagsTagList
}

type ListPhotoTagsTag struct {
	Id        int64
	IsSubTag  bool
	Name      string
	ParentTag string
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
