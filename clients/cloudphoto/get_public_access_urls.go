package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetPublicAccessUrlsRequest struct {
	requests.RpcRequest
	DomainType string                          `position:"Query" name:"DomainType"`
	LibraryId  string                          `position:"Query" name:"LibraryId"`
	PhotoIds   *GetPublicAccessUrlsPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName  string                          `position:"Query" name:"StoreName"`
	ZoomType   string                          `position:"Query" name:"ZoomType"`
}

func (req *GetPublicAccessUrlsRequest) Invoke(client *sdk.Client) (resp *GetPublicAccessUrlsResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetPublicAccessUrls", "cloudphoto", "")
	resp = &GetPublicAccessUrlsResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetPublicAccessUrlsResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	RequestId string
	Action    string
	Results   GetPublicAccessUrlsResultList
}

type GetPublicAccessUrlsResult struct {
	Code      string
	Message   string
	PhotoId   int64
	AccessUrl string
}

type GetPublicAccessUrlsPhotoIdList []int64

func (list *GetPublicAccessUrlsPhotoIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type GetPublicAccessUrlsResultList []GetPublicAccessUrlsResult

func (list *GetPublicAccessUrlsResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetPublicAccessUrlsResult)
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
