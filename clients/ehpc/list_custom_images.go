package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListCustomImagesRequest struct {
	requests.RpcRequest
	BaseOsTag       string `position:"Query" name:"BaseOsTag"`
	ImageOwnerAlias string `position:"Query" name:"ImageOwnerAlias"`
}

func (req *ListCustomImagesRequest) Invoke(client *sdk.Client) (resp *ListCustomImagesResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "ListCustomImages", "ehs", "")
	resp = &ListCustomImagesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListCustomImagesResponse struct {
	responses.BaseResponse
	RequestId string
	Images    ListCustomImagesImageInfoList
}

type ListCustomImagesImageInfo struct {
	ImageId         string
	ImageName       string
	ImageOwnerAlias string
	Description     string
	BaseOsTag       ListCustomImagesBaseOsTag
}

type ListCustomImagesBaseOsTag struct {
	OsTag        string
	Platform     string
	Version      string
	Architecture string
}

type ListCustomImagesImageInfoList []ListCustomImagesImageInfo

func (list *ListCustomImagesImageInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListCustomImagesImageInfo)
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
