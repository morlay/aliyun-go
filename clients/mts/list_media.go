package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListMediaRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	NextPageToken        string `position:"Query" name:"NextPageToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MaximumPageSize      int64  `position:"Query" name:"MaximumPageSize"`
	From                 string `position:"Query" name:"From"`
	To                   string `position:"Query" name:"To"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ListMediaRequest) Invoke(client *sdk.Client) (response *ListMediaResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListMediaRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "ListMedia", "", "")

	resp := struct {
		*responses.BaseResponse
		ListMediaResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListMediaResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListMediaResponse struct {
	RequestId     string
	NextPageToken string
	MediaList     ListMediaMediaList
}

type ListMediaMedia struct {
	MediaId      string
	Title        string
	Description  string
	CoverURL     string
	CateId       int64
	Duration     string
	Format       string
	Size         string
	Bitrate      string
	Width        string
	Height       string
	Fps          string
	PublishState string
	CreationTime string
	Tags         ListMediaTagList
	RunIdList    ListMediaRunIdListList
	File         ListMediaFile
}

type ListMediaFile struct {
	URL   string
	State string
}

type ListMediaMediaList []ListMediaMedia

func (list *ListMediaMediaList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListMediaMedia)
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

type ListMediaTagList []string

func (list *ListMediaTagList) UnmarshalJSON(data []byte) error {
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

type ListMediaRunIdListList []string

func (list *ListMediaRunIdListList) UnmarshalJSON(data []byte) error {
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
