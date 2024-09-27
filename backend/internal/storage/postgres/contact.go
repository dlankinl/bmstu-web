package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"ppo/domain"
	"strings"
)

type ContactRepository struct {
	db *pgxpool.Pool
}

func NewContactRepository(db *pgxpool.Pool) domain.IContactsRepository {
	return &ContactRepository{
		db: db,
	}
}

func (r *ContactRepository) Create(ctx context.Context, contact *domain.Contact) (err error) {
	query := `insert into ppo.contacts(owner_id, name, value) 
	values ($1, $2, $3)`

	_, err = r.db.Exec(
		ctx,
		query,
		contact.OwnerID,
		contact.Name,
		contact.Value,
	)
	if err != nil {
		return fmt.Errorf("создание средства связи: %w", err)
	}

	return nil
}

func (r *ContactRepository) GetById(ctx context.Context, id uuid.UUID) (contact *domain.Contact, err error) {
	query := `select owner_id, name, value from ppo.contacts where id = $1`

	contact = new(domain.Contact)
	err = r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&contact.OwnerID,
		&contact.Name,
		&contact.Value,
	)
	if err != nil {
		return nil, fmt.Errorf("получение средства связи по id: %w", err)
	}

	contact.ID = id
	return contact, nil
}

func (r *ContactRepository) GetByOwnerId(ctx context.Context, id uuid.UUID) (contacts []*domain.Contact, err error) {
	query := `
		select 
		    id,
		    name,
		    value 
		from ppo.contacts 
		where owner_id = $1`

	rows, err := r.db.Query(
		ctx,
		query,
		id,
	)
	if err != nil {
		return nil, fmt.Errorf("получение средств связи: %w", err)
	}

	contacts = make([]*domain.Contact, 0)
	for rows.Next() {
		tmp := new(domain.Contact)

		err = rows.Scan(
			&tmp.ID,
			&tmp.Name,
			&tmp.Value,
		)
		tmp.OwnerID = id

		if err != nil {
			return nil, fmt.Errorf("сканирование полученных строк: %w", err)
		}
		contacts = append(contacts, tmp)
	}

	return contacts, nil
}

func (r *ContactRepository) Update(ctx context.Context, contact *domain.Contact) (err error) {
	queryArgs := make([]any, 0)
	queryElems := make([]string, 0)
	query := "update ppo.contacts set "

	i := 1
	if contact.OwnerID.ID() != 0 {
		queryElems = append(queryElems, fmt.Sprintf("owner_id = $%d", i))
		queryArgs = append(queryArgs, contact.OwnerID)
		i++
	}
	if contact.Name != "" {
		queryElems = append(queryElems, fmt.Sprintf("name = $%d", i))
		queryArgs = append(queryArgs, contact.Name)
		i++
	}
	if contact.Value != "" {
		queryElems = append(queryElems, fmt.Sprintf("value = $%d", i))
		queryArgs = append(queryArgs, contact.Value)
		i++
	}
	query += strings.Join(queryElems, ", ")
	query += fmt.Sprintf(" where id = $%d", i)
	queryArgs = append(queryArgs, contact.ID)

	_, err = r.db.Exec(
		ctx,
		query,
		queryArgs...,
	)
	if err != nil {
		return fmt.Errorf("обновление информации о средстве связи: %w", err)
	}

	return nil
}

func (r *ContactRepository) DeleteById(ctx context.Context, id uuid.UUID) (err error) {
	query := `delete from ppo.contacts where id = $1`

	_, err = r.db.Exec(
		ctx,
		query,
		id,
	)
	if err != nil {
		return fmt.Errorf("удаление средства связи по id: %w", err)
	}

	return nil
}
