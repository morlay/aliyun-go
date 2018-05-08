package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
	Results   GetPublicAccessUrlsResultList
}

type GetPublicAccessUrlsResult struct {
	Code       common.String
	Message    common.String
	PhotoId    common.Long
	PhotoIdStr common.String
	AccessUrl  common.String
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
