package tickets

import (
	"database/sql"
	"github.com/daveearley/product/app/api/pagination"
	"github.com/daveearley/product/app/models/generated"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type Repository interface {
	GetById(id int) (*models.Ticket, error)
	Store(event *models.Ticket) (*models.Ticket, error)
	SetAttributes(event *models.Ticket, attr []*models.Attribute) error
	List(p *pagination.Params, authUser *models.User) ([]*models.Ticket, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) GetById(id int) (*models.Ticket, error) {
	event, err := models.FindTicket(r.db, id)

	if err != nil {
		return nil, err
	}

	return event, nil
}

func (r *repository) Store(ticket *models.Ticket) (*models.Ticket, error) {
	if err := ticket.Insert(r.db, boil.Infer()); err != nil {
		return nil, err
	}

	return ticket, nil
}

func (r *repository) SetAttributes(ticket *models.Ticket, attr []*models.Attribute) error {
	return ticket.SetAttributes(r.db, true, attr...)
}

func (r *repository) List(p *pagination.Params, authUser *models.User) ([]*models.Ticket, error) {
	queryMods := pagination.QueryMods(p, authUser)
	queryMods = append(queryMods, qm.Load("Attributes"))

	events, err := models.Tickets(queryMods...).All(r.db)

	if err != nil {
		return nil, err
	}

	return events, nil
}
