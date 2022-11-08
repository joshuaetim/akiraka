package infrastructure_test

import (
	"testing"

	"github.com/joshuaetim/quiz/domain/model"
	"github.com/joshuaetim/quiz/factory"
	"github.com/joshuaetim/quiz/infrastructure"
	"github.com/stretchr/testify/assert"
)

func TestVisitorSave(t *testing.T) {
	initTest(t)
	db := infrastructure.DB()
	user, err := factory.SeedUser(db)
	assert.Nil(t, err)

	staff, err := factory.SeedStaff(db, user.ID)
	assert.Nil(t, err)

	visitor, err := factory.SeedVisitor(db, user.ID, staff.ID)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, visitor.ID)
}

func TestVisitorGet(t *testing.T) {
	initTest(t)
	db := infrastructure.DB()
	user, err := factory.SeedUser(db)
	assert.Nil(t, err)

	staff, err := factory.SeedStaff(db, user.ID)
	assert.Nil(t, err)

	visitor, err := factory.SeedVisitor(db, user.ID, staff.ID)
	assert.Nil(t, err)

	vr := infrastructure.NewVisitorRepository(db)
	v, err := vr.GetVisitor(visitor.ID)
	assert.Nil(t, err)
	assert.EqualValues(t, visitor.ID, v.ID)
}

func TestVisitorGetUserVisitor(t *testing.T) {
	initTest(t)
	db := infrastructure.DB()
	user, err := factory.SeedUser(db)
	assert.Nil(t, err)

	staff, err := factory.SeedStaff(db, user.ID)
	assert.Nil(t, err)

	visitor, err := factory.SeedVisitor(db, user.ID, staff.ID)
	assert.Nil(t, err)

	vr := infrastructure.NewVisitorRepository(db)
	v, err := vr.GetUserVisitor(visitor.ID, user.ID)
	assert.Nil(t, err)
	assert.EqualValues(t, visitor.ID, v.ID)
}

func TestVisitorGetAll(t *testing.T) {
	initTest(t)
	db := infrastructure.DB()
	user, err := factory.SeedUser(db)
	assert.Nil(t, err)

	staff, err := factory.SeedStaff(db, user.ID)
	assert.Nil(t, err)

	var visitors []model.Visitor
	for i := 0; i < 4; i++ {
		visitor, err := factory.SeedVisitor(db, user.ID, staff.ID)
		assert.Nil(t, err)
		visitors = append(visitors, visitor)
	}

	sr := infrastructure.NewVisitorRepository(db)
	visitorsGet, err := sr.GetAllVisitor()
	assert.Nil(t, err)

	assert.EqualValues(t, len(visitors), len(visitorsGet))
}

func TestVisitorUpdate(t *testing.T) {
	initTest(t)
	db := infrastructure.DB()
	user, err := factory.SeedUser(db)
	assert.Nil(t, err)

	staff, err := factory.SeedStaff(db, user.ID)
	assert.Nil(t, err)

	visitor, err := factory.SeedVisitor(db, user.ID, staff.ID)
	assert.Nil(t, err)

	vr := infrastructure.NewVisitorRepository(db)
	visitor.Firstname = "Name Updated"
	v2, err := vr.UpdateVisitor(visitor)
	assert.Nil(t, err)

	assert.EqualValues(t, "Name Updated", v2.Firstname)
}

func TestVisitorDelete(t *testing.T) {
	initTest(t)
	db := infrastructure.DB()
	user, err := factory.SeedUser(db)
	assert.Nil(t, err)

	staff, err := factory.SeedStaff(db, user.ID)
	assert.Nil(t, err)

	visitor, err := factory.SeedVisitor(db, user.ID, staff.ID)
	assert.Nil(t, err)

	vr := infrastructure.NewVisitorRepository(db)
	err = vr.DeleteVisitor(visitor)
	assert.Nil(t, err)

	_, err = vr.GetVisitor(visitor.ID)
	assert.NotNil(t, err)
}
