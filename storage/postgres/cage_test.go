package postgres

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/danielnegri/jurassic-park-go/model"
	"github.com/danielnegri/jurassic-park-go/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestPostgres_CreateCage(t *testing.T) {
	if shouldSkip() {
		t.SkipNow()
	}

	setup(t)

	uid := uuid.MustNextID()
	p1 := &model.Cage{
		ID:       model.NewCageID(uid),
		Capacity: gofakeit.IntRange(0, 10000),
		Active:   gofakeit.Bool(),
	}

	ctx := context.Background()
	if err := postgres.CreateCage(ctx, p1); err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, p1.CreatedAt)
	assert.NotNil(t, p1.UpdatedAt)

	p2, err := postgres.getCage(ctx, p1.ID, errors.Op("TestPostgres_CreateCage"))
	if err != nil {
		t.Fatal(err)
	}

	j1, err := json.Marshal(p1)
	if err != nil {
		t.Fatal(err)
	}

	j2, err := json.Marshal(p2)
	if err != nil {
		t.Fatal(err)
	}

	assert.JSONEq(t, string(j1), string(j2))
}

func TestPostgres_UpdateCage(t *testing.T) {
	if shouldSkip() {
		t.SkipNow()
	}

	setup(t)

	uid := uuid.MustNextID()
	p0 := &model.Cage{
		ID:       model.NewCageID(uid),
		Capacity: gofakeit.IntRange(0, 10000),
		Active:   false,
	}

	ctx := context.Background()
	if err := postgres.CreateCage(ctx, p0); err != nil {
		t.Fatal(err)
	}

	p1, err := postgres.getCage(ctx, p0.ID, errors.Op("TestPostgres_UpdateCage"))
	if err != nil {
		t.Fatal(err)
	}

	updater := func(old *model.Cage) (*model.Cage, error) {
		old.Capacity = gofakeit.IntRange(0, 10000)
		old.Active = true
		return old, nil
	}

	if err := postgres.UpdateCage(ctx, p0.ID, updater); err != nil {
		t.Fatal(err)
	}

	p2, err := postgres.getCage(ctx, p0.ID, errors.Op("TestPostgres_UpdateCage"))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, p1.ID, p2.ID)
	assert.NotEqual(t, p1.Capacity, p2.Capacity)
	assert.False(t, p1.Active)
	assert.True(t, p2.Active)
	assert.Equal(t, p1.CreatedAt.UTC(), p2.CreatedAt.UTC())
	assert.GreaterOrEqual(t, p2.UpdatedAt.UTC(), p1.UpdatedAt.UTC())
}

func TestPostgres_GetCageByHandleNotFound(t *testing.T) {
	if shouldSkip() {
		t.SkipNow()
	}

	setup(t)

	p, err := postgres.GetCage(context.Background(), "foo")
	assert.True(t, errors.Is(err, errors.KindNotFound))
	assert.Nil(t, p)
}
