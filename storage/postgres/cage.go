package postgres

import (
	"context"

	"github.com/danielnegri/jurassic-park-go/model"
	"github.com/danielnegri/jurassic-park-go/pkg/errors"
	"github.com/danielnegri/jurassic-park-go/storage"
	"github.com/go-pg/pg/v10"
)

func (p *Postgres) CreateCage(ctx context.Context, cage *model.Cage) error {
	const op errors.Op = "postgres.CreateProduct"

	now := p.now().UTC()
	cage.CreatedAt = &now
	cage.UpdatedAt = &now

	_, err := p.db.WithContext(ctx).Model(cage).Insert()
	if err != nil {
		return errors.E(op, kind(err), err)
	}

	return nil
}

func (p *Postgres) UpdateCage(ctx context.Context, id model.ID, updater storage.CageUpdater) error {
	const op errors.Op = "postgres.UpdateCage"

	updateFn := func(tx *pg.Tx) error {
		old, err := p.getCage(ctx, id, op)
		if err != nil {
			return err
		}

		cage, err := updater(old)
		if err != nil {
			return err
		}

		now := p.now().UTC()
		cage.UpdatedAt = &now

		if _, err := p.db.WithContext(ctx).
			Model(cage).
			WherePK().
			Update(); err != nil {
			return errors.E(op, kind(err), err)
		}

		return nil
	}

	return p.ExecTx(ctx, updateFn)
}

func (p *Postgres) GetCage(ctx context.Context, id model.ID) (*model.Cage, error) {
	const op errors.Op = "postgres.GetCage"
	return p.getCage(ctx, id, op)
}

func (p *Postgres) getCage(ctx context.Context, id model.ID, op errors.Op) (*model.Cage, error) {
	var cage model.Cage
	err := p.db.WithContext(ctx).
		Model((*model.Cage)(nil)).
		Where("id = ?", string(id)).
		Select(&cage)
	if err != nil {
		return nil, errors.E(op, kind(err), err)
	}

	return &cage, nil
}

func (p *Postgres) ListCages(ctx context.Context, params storage.ListCageParams) ([]*model.Cage, error) {
	return nil, nil
}
