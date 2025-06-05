package checkins

import (
	"github.com/Tutitoos/statping-ng/database"
	"github.com/Tutitoos/statping-ng/types/metrics"
	"github.com/Tutitoos/statping-ng/utils"
)

var db database.Database
var dbHits database.Database

func SetDB(database database.Database) {
	db = database.Model(&Checkin{})
	dbHits = database.Model(&CheckinHit{})
}

func (c *Checkin) AfterFind() {
	c.AllHits = c.Hits()
	c.AllFailures = c.Failures().LastAmount(32)
	if last := c.LastHit(); last != nil {
		c.LastHitTime = last.CreatedAt
	}
	metrics.Query("checkin", "find")
}

func Find(id int64, permalink ...string) (*Checkin, error) {
	var checkin Checkin

	// Try to find by ID first
	q := db.Where("id = ?", id).Find(&checkin)

	// Try permalink if ID search had an error or found nothing, and permalink is provided
	if (q.Error() != nil || q.RowsAffected() == 0) && len(permalink) > 0 && permalink[0] != "" {
		q = db.Where("permalink = ?", permalink[0]).Find(&checkin)
	}

	return &checkin, q.Error()
}

func FindByAPI(key string) (*Checkin, error) {
	var checkin Checkin
	q := db.Where("api_key = ?", key).Find(&checkin)
	return &checkin, q.Error()
}

func All() []*Checkin {
	var checkins []*Checkin
	db.Find(&checkins)
	return checkins
}

func (c *Checkin) Create() error {
	if c.ApiKey == "" {
		c.ApiKey = utils.RandomString(32)
	}
	q := db.Create(c)
	return q.Error()
}

func (c *Checkin) Update() error {
	q := db.Update(c)
	return q.Error()
}

func (c *Checkin) Delete() error {
	c.Close()
	q := dbHits.Where("checkin = ?", c.Id).Delete(&CheckinHit{})
	if err := q.Error(); err != nil {
		return err
	}
	q = db.Model(&Checkin{}).Delete(c)
	return q.Error()
}
