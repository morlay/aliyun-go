package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DetectFaceRequest struct {
	requests.RpcRequest
	SrcUris string `position:"Query" name:"SrcUris"`
	Project string `position:"Query" name:"Project"`
}

func (req *DetectFaceRequest) Invoke(client *sdk.Client) (resp *DetectFaceResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "DetectFace", "imm", "")
	resp = &DetectFaceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DetectFaceResponse struct {
	responses.BaseResponse
	RequestId      string
	SuccessDetails DetectFaceSuccessDetailsItemList
	FailDetails    DetectFaceFailDetailsItemList
	SrcUris        DetectFaceSrcUriList
}

type DetectFaceSuccessDetailsItem struct {
	SrcUri  string
	PhotoId string
	Faces   DetectFaceFacesItemList
}

type DetectFaceFacesItem struct {
	FaceId        string
	FaceRectangle DetectFaceFaceRectangle
	FaceAttribute DetectFaceFaceAttribute
}

type DetectFaceFaceRectangle struct {
	Top    string
	Left   string
	Width  string
	Height string
}

type DetectFaceFaceAttribute struct {
	Gender      DetectFaceGender
	Age         DetectFaceAge
	Headpose    DetectFaceHeadpose
	Eyestatus   DetectFaceEyestatus
	Blur        DetectFaceBlur
	Facequality DetectFaceFacequality
}

type DetectFaceGender struct {
	Value string
}

type DetectFaceAge struct {
	Value int
}

type DetectFaceHeadpose struct {
	Pitch_angle float32
	Roll_angle  float32
	Yaw_angle   float32
}

type DetectFaceEyestatus struct {
	Left_eye_status  DetectFaceLeft_eye_status
	Right_eye_status DetectFaceRight_eye_status
}

type DetectFaceLeft_eye_status struct {
	Normal_glass_eye_open  float32
	No_glass_eye_close     float32
	Occlusion              float32
	No_glass_eye_open      float32
	Normal_glass_eye_close float32
	Dark_glasses           float32
}

type DetectFaceRight_eye_status struct {
	Normal_glass_eye_open  float32
	No_glass_eye_close     float32
	Occlusion              float32
	No_glass_eye_open      float32
	Normal_glass_eye_close float32
	Dark_glasses           float32
}

type DetectFaceBlur struct {
	Blurness DetectFaceBlurness
}

type DetectFaceBlurness struct {
	Value     float32
	Threshold float32
}

type DetectFaceFacequality struct {
	Value     float32
	Threshold float32
}

type DetectFaceFailDetailsItem struct {
	SrcUri string
	Reason string
}

type DetectFaceSuccessDetailsItemList []DetectFaceSuccessDetailsItem

func (list *DetectFaceSuccessDetailsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DetectFaceSuccessDetailsItem)
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

type DetectFaceFailDetailsItemList []DetectFaceFailDetailsItem

func (list *DetectFaceFailDetailsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DetectFaceFailDetailsItem)
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

type DetectFaceSrcUriList []string

func (list *DetectFaceSrcUriList) UnmarshalJSON(data []byte) error {
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

type DetectFaceFacesItemList []DetectFaceFacesItem

func (list *DetectFaceFacesItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DetectFaceFacesItem)
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
