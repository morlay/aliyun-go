package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetPrivateAccessUrlsRequest struct {
	requests.RpcRequest
	LibraryId string                           `position:"Query" name:"LibraryId"`
	PhotoIds  *GetPrivateAccessUrlsPhotoIdList `position:"Query" type:"Repeated" name:"PhotoId"`
	StoreName string                           `position:"Query" name:"StoreName"`
	ZoomType  string                           `position:"Query" name:"ZoomType"`
}

func (req *GetPrivateAccessUrlsRequest) Invoke(client *sdk.Client) (resp *GetPrivateAccessUrlsResponse, err error) {
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "GetPrivateAccessUrls", "cloudphoto", "")
	resp = &GetPrivateAccessUrlsResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetPrivateAccessUrlsResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	RequestId common.String
	Action    common.String
	Results   GetPrivateAccessUrlsResultList
}

type GetPrivateAccessUrlsResult struct {
	Code       common.String
	Message    common.String
	PhotoId    common.Long
	PhotoIdStr common.String
	AccessUrl  common.String
}

type GetPrivateAccessUrlsPhotoIdList []int64

func (list *GetPrivateAccessUrlsPhotoIdList) UnmarshalJSON(data []byte) error {
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

type GetPrivateAccessUrlsResultList []GetPrivateAccessUrlsResult

func (list *GetPrivateAccessUrlsResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetPrivateAccessUrlsResult)
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
