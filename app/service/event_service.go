package service

import (
	"github.com/daveearley/ticketbooth/app/api/pagination"
	"github.com/daveearley/ticketbooth/app/api/request"
	"github.com/daveearley/ticketbooth/app/attributes"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/daveearley/ticketbooth/app/repository"
	"github.com/volatiletech/null"
)

type EventService interface {
	Find(id int) (*models.Event, error)
	Create(event request.CreateEvent, user *models.User) (*models.Event, error)
	Delete(id int) error
	List(p *pagination.Params, authUser *models.User) ([]*models.Event, error)
}

type eventService struct {
	er repository.EventRepository
}

func NewEventService(repository repository.EventRepository) *eventService {
	return &eventService{repository}
}

func (s *eventService) Delete(id int) error {
	err := s.er.DeleteByID(id)

	return err
}

func (s *eventService) Find(id int) (*models.Event, error) {
	event, err := s.er.GetByID(id)

	if err != nil {
		return nil, err
	}

	return event, nil
}

func (s *eventService) Create(req request.CreateEvent, user *models.User) (*models.Event, error) {
	event, err := s.er.Store(&models.Event{
		Title:       req.Title,
		Description: null.NewString(req.Description, true),
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		UserID:      user.ID,
		AccountID:   user.AccountID,
	})

	if err != nil {
		return nil, err
	}

	if req.Attributes != nil {
		s.er.SetAttributes(event, attributes.MapToAttributes(&req.Attributes))
	}

	return event, nil
}

func (s *eventService) List(p *pagination.Params, authUser *models.User) ([]*models.Event, error) {
	events, err := s.er.List(p, authUser)

	if err != nil {
		return nil, err
	}

	return events, nil
}