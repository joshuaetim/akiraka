package infrastructure_test

import (
	"testing"

	"github.com/joshuaetim/quiz/domain/model"
	"github.com/joshuaetim/quiz/factory"
	"github.com/joshuaetim/quiz/infrastructure"
	"github.com/stretchr/testify/assert"
)

func TestStaffSave(t *testing.T) {
	initTest(t)
	db := infrastructure.DB()
	user, err := factory.SeedUser(db)
	assert.Nil(t, err)

	staff, err := factory.SeedStaff(db, user.ID)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, staff.ID)
}

func TestStaffGet(t *testing.T) {
	initTest(t)
	db := infrastructure.DB()
	user, err := factory.SeedUser(db)
	assert.Nil(t, err)

	staff, err := factory.SeedStaff(db, user.ID)
	assert.Nil(t, err)

	sr := infrastructure.NewStaffRepository(db)
	staffGet, err := sr.GetStaff(staff.ID)
	assert.Nil(t, err)

	assert.EqualValues(t, staff.ID, staffGet.ID)
	// assert.EqualValues(t, user.ID, staffGet.User.ID)

}

func TestStaffGetAll(t *testing.T) {
	initTest(t)
	db := infrastructure.DB()
	user, err := factory.SeedUser(db)
	assert.Nil(t, err)

	var allStaff []model.Staff

	for i := 1; i <= 4; i++ {
		staff, err := factory.SeedStaff(db, user.ID)
		assert.Nil(t, err)
		allStaff = append(allStaff, staff)
	}

	sr := infrastructure.NewStaffRepository(db)
	allStaffGet, err := sr.GetAllStaff()
	assert.Nil(t, err)

	assert.EqualValues(t, len(allStaff), len(allStaffGet))
}

func TestStaffUpdate(t *testing.T) {
	initTest(t)
	db := infrastructure.DB()
	user, err := factory.SeedUser(db)
	assert.Nil(t, err)

	staff, err := factory.SeedStaff(db, user.ID)
	assert.Nil(t, err)

	sr := infrastructure.NewStaffRepository(db)
	s1, err := sr.GetStaff(staff.ID)
	assert.Nil(t, err)

	s1.Firstname = "Johnny Update"

	s2, err := sr.UpdateStaff(s1)
	assert.Nil(t, err)

	assert.EqualValues(t, "Johnny Update", s2.Firstname)
}

func TestStaffDelete(t *testing.T) {
	initTest(t)
	db := infrastructure.DB()
	user, err := factory.SeedUser(db)
	assert.Nil(t, err)

	staff, err := factory.SeedStaff(db, user.ID)
	assert.Nil(t, err)

	sr := infrastructure.NewStaffRepository(db)
	err = sr.DeleteStaff(staff)
	assert.Nil(t, err)

	_, err = sr.GetStaff(staff.ID)
	assert.NotNil(t, err)
}
