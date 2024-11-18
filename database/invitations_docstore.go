package database

import (
	"context"
	"fmt"
	"io"

	"github.com/aaronland/go-mailinglist/v2/invitation"
	"github.com/aaronland/go-mailinglist/v2/subscription"
	aa_docstore "github.com/aaronland/gocloud-docstore"
	"gocloud.dev/docstore"
)

type InvitationsDocstoreDatabase struct {
	InvitationsDatabase
	collection *docstore.Collection
}

func init() {

	ctx := context.Background()

	err := RegisterInvitationsDatabase(ctx, "awsdynamodb", NewInvitationsDocstoreDatabase)

	if err != nil {
		panic(err)
	}

	for _, scheme := range docstore.DefaultURLMux().CollectionSchemes() {

		err := RegisterInvitationsDatabase(ctx, scheme, NewInvitationsDocstoreDatabase)

		if err != nil {
			panic(err)
		}

	}

}

func NewInvitationsDocstoreDatabase(ctx context.Context, uri string) (InvitationsDatabase, error) {

	col, err := aa_docstore.OpenCollection(ctx, uri)

	if err != nil {
		return nil, err
	}

	db := &InvitationsDocstoreDatabase{
		collection: col,
	}

	return db, nil
}

func (db *InvitationsDocstoreDatabase) AddInvitation(ctx context.Context, iv *invitation.Invitation) error {
	return db.collection.Put(ctx, iv)
}

func (db *InvitationsDocstoreDatabase) RemoveInvitation(ctx context.Context, iv *invitation.Invitation) error {
	return db.collection.Delete(ctx, iv)
}

func (db *InvitationsDocstoreDatabase) UpdateInvitation(ctx context.Context, iv *invitation.Invitation) error {
	return db.collection.Replace(ctx, iv)
}

func (db *InvitationsDocstoreDatabase) GetInvitationWithCode(ctx context.Context, code string) (*invitation.Invitation, error) {
	q := db.collection.Query()
	q = q.Where("code", "=", code)

	return db.getInvitationWithQuery(ctx, q)
}

func (db *InvitationsDocstoreDatabase) GetInvitationWithInvitee(ctx context.Context, iv string) (*invitation.Invitation, error) {

	q := db.collection.Query()
	q = q.Where("invitee", "=", iv)

	return db.getInvitationWithQuery(ctx, q)
}

func (db *InvitationsDocstoreDatabase) ListInvitationsWithInviter(ctx context.Context, sub *subscription.Subscription, cb ListInvitationsFunc) error {

	q := db.collection.Query()
	q = q.Where("inviter", "=", sub.Address)

	return db.getInvitationsWithCallback(ctx, q, cb)
}

func (db *InvitationsDocstoreDatabase) ListInvitations(ctx context.Context, cb ListInvitationsFunc) error {
	q := db.collection.Query()
	return db.getInvitationsWithCallback(ctx, q, cb)
}

func (db *InvitationsDocstoreDatabase) Close() error {
	return db.collection.Close()
}

func (db *InvitationsDocstoreDatabase) getInvitationWithQuery(ctx context.Context, q *docstore.Query) (*invitation.Invitation, error) {

	iter := q.Get(ctx)
	defer iter.Stop()

	var iv invitation.Invitation
	err := iter.Next(ctx, &iv)

	if err == io.EOF {
		return nil, NoRecordError("")
	} else if err != nil {
		return nil, fmt.Errorf("Failed to interate, %w", err)
	} else {
		return &iv, nil
	}
}

func (db *InvitationsDocstoreDatabase) getInvitationsWithCallback(ctx context.Context, q *docstore.Query, cb ListInvitationsFunc) error {

	iter := q.Get(ctx)
	defer iter.Stop()

	for {

		var iv invitation.Invitation
		err := iter.Next(ctx, &iv)

		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("Failed to interate, %w", err)
		} else {

			err := cb(ctx, &iv)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
