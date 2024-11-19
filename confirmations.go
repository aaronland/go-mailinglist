package mailinglist

import (
	"time"
	"context"

	"github.com/aaronland/go-mailinglist/v2/database"
	"github.com/aaronland/go-mailinglist/v2/confirmation"
)

func AddConfirmation(ctx context.Context, db database.ConfirmationsDatabase, c *confirmation.Confirmation) (*confirmation.Confirmation, error) {

	now := time.Now()
	ts := now.Unix()

	c.Created = ts

	err := db.AddConfirmation(ctx, c)

	if err != nil {
		return nil, err
	}

	return c, nil
}
