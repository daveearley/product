package ticket

import (
	"github.com/daveearley/product/app/attribute"
	"github.com/daveearley/product/app/models/generated"
	"github.com/daveearley/product/app/pagination"
	"github.com/daveearley/product/app/request"
	"github.com/volatiletech/null"
)

type Service interface {
	Find(id int) (*models.Ticket, error)
	Create(ticket request.CreateTicket, user *models.Event) (*models.Ticket, error)
	List(p *pagination.Params, authUser *models.Event) ([]*models.Ticket, error)
}

type service struct {
	er Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Find(id int) (*models.Ticket, error) {
	ticket, err := s.er.GetById(id)

	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (s *service) Create(req request.CreateTicket, event *models.Event) (*models.Ticket, error) {
	ticket, err := s.er.Store(&models.Ticket{
		Title:                    req.Title,
		SaleStartDate:            null.NewTime(req.SaleStartDate, true),
		SaleEndDate:              null.NewTime(req.SaleEndDate, true),
		IntitalQuantityAvailable: req.Quantity,
		EventID:                  event.ID,
	})

	if err != nil {
		return nil, err
	}

	if req.Attributes != nil {
		s.er.SetAttributes(ticket, attribute.MapToAttributes(&req.Attributes))
	}

	return ticket, nil
}

func (s *service) List(p *pagination.Params, authUser *models.User) ([]*models.Ticket, error) {
	tickets, err := s.er.List(p, authUser)

	if err != nil {
		return nil, err
	}

	return tickets, nil
}