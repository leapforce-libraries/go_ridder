package ridder

import (
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Title struct {
	ID          int32   `json:"Id"`
	Code        string  `json:"Code" max:"3"`
	Description *string `json:"Description" max:"80"`
	Salutation  *string `json:"Salutation" max:"50"`
}

func (service *Service) GetTitles() (*[]Title, *errortools.Error) {
	var titles []Title

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url("titles"),
		ResponseModel: &titles,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &titles, nil
}

func (service *Service) GetTitleByID(titleID int32) (*Title, *errortools.Error) {
	return service.getTitle(fmt.Sprintf("titles/id/%v", titleID))
}

func (service *Service) GetTitleByCode(code string) (*Title, *errortools.Error) {
	return service.getTitle(fmt.Sprintf("titles/code/%s", code))
}

func (service *Service) getTitle(urlPath string) (*Title, *errortools.Error) {
	var title Title

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(urlPath),
		ResponseModel: &title,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &title, nil
}

func (service *Service) CreateTitle(title *Title) (*int32, *errortools.Error) {
	if title == nil {
		return nil, nil
	}

	var titleIDString string

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		Url:           service.url("titles"),
		BodyModel:     title,
		ResponseModel: &titleIDString,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return service.parseInt32String(titleIDString)
}
