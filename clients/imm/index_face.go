package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type IndexFaceRequest struct {
	requests.RpcRequest
	SrcUris string `position:"Query" name:"SrcUris"`
	Project string `position:"Query" name:"Project"`
	SetId   string `position:"Query" name:"SetId"`
	Force   string `position:"Query" name:"Force"`
}

func (req *IndexFaceRequest) Invoke(client *sdk.Client) (resp *IndexFaceResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "IndexFace", "imm", "")
	resp = &IndexFaceResponse{}
	err = client.DoAction(req, resp)
	return
}

type IndexFaceResponse struct {
	responses.BaseResponse
	RequestId      string
	SetId          string
	SuccessDetails IndexFaceSuccessDetailsItemList
	FailDetails    IndexFaceFailDetailsItemList
	SrcUris        IndexFaceSrcUriList
}

type IndexFaceSuccessDetailsItem struct {
	SrcUri  string
	PhotoId string
	Faces   IndexFaceFacesItemList
}

type IndexFaceFacesItem struct {
	FaceId        string
	FaceRectangle IndexFaceFaceRectangle
	FaceAttribute IndexFaceFaceAttribute
}

type IndexFaceFaceRectangle struct {
	Top    string
	Left   string
	Width  string
	Height string
}

type IndexFaceFaceAttribute struct {
	Gender      IndexFaceGender
	Age         IndexFaceAge
	Headpose    IndexFaceHeadpose
	Eyestatus   IndexFaceEyestatus
	Blur        IndexFaceBlur
	Facequality IndexFaceFacequality
}

type IndexFaceGender struct {
	Value string
}

type IndexFaceAge struct {
	Value int
}

type IndexFaceHeadpose struct {
	Pitch_angle float32
	Roll_angle  float32
	Yaw_angle   float32
}

type IndexFaceEyestatus struct {
	Left_eye_status  IndexFaceLeft_eye_status
	Right_eye_status IndexFaceRight_eye_status
}

type IndexFaceLeft_eye_status struct {
	Normal_glass_eye_open  float32
	No_glass_eye_close     float32
	Occlusion              float32
	No_glass_eye_open      float32
	Normal_glass_eye_close float32
	Dark_glasses           float32
}

type IndexFaceRight_eye_status struct {
	Normal_glass_eye_open  float32
	No_glass_eye_close     float32
	Occlusion              float32
	No_glass_eye_open      float32
	Normal_glass_eye_close float32
	Dark_glasses           float32
}

type IndexFaceBlur struct {
	Blurness IndexFaceBlurness
}

type IndexFaceBlurness struct {
	Value     float32
	Threshold float32
}

type IndexFaceFacequality struct {
	Value     float32
	Threshold float32
}

type IndexFaceFailDetailsItem struct {
	SrcUri string
	Reason string
}

type IndexFaceSuccessDetailsItemList []IndexFaceSuccessDetailsItem

func (list *IndexFaceSuccessDetailsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]IndexFaceSuccessDetailsItem)
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

type IndexFaceFailDetailsItemList []IndexFaceFailDetailsItem

func (list *IndexFaceFailDetailsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]IndexFaceFailDetailsItem)
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

type IndexFaceSrcUriList []string

func (list *IndexFaceSrcUriList) UnmarshalJSON(data []byte) error {
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

type IndexFaceFacesItemList []IndexFaceFacesItem

func (list *IndexFaceFacesItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]IndexFaceFacesItem)
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
