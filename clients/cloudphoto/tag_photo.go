package cloudphoto

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type TagPhotoRequest struct {
	LibraryId   string                  `position:"Query" name:"LibraryId"`
	Confidences *TagPhotoConfidenceList `position:"Query" type:"Repeated" name:"Confidence"`
	StoreName   string                  `position:"Query" name:"StoreName"`
	PhotoId     int64                   `position:"Query" name:"PhotoId"`
	TagKeys     *TagPhotoTagKeyList     `position:"Query" type:"Repeated" name:"TagKey"`
}

func (r TagPhotoRequest) Invoke(client *sdk.Client) (response *TagPhotoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		TagPhotoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudPhoto", "2017-07-11", "TagPhoto", "cloudphoto", "")

	resp := struct {
		*responses.BaseResponse
		TagPhotoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.TagPhotoResponse

	err = client.DoAction(&req, &resp)
	return
}

type TagPhotoResponse struct {
	Code      string
	Message   string
	RequestId string
	Action    string
}

type TagPhotoConfidenceList []float32

func (list *TagPhotoConfidenceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]float32)
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

type TagPhotoTagKeyList []string

func (list *TagPhotoTagKeyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
