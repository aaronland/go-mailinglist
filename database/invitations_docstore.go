package database

import (
	"context"

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

}

func (db *InvitationsDocstoreDatabase) GetInvitationWithInvitee(ctx context.Context, iv string) (*invitation.Invitation, error) {

}

func (db *InvitationsDocstoreDatabase) ListInvitations(ctx context.Context, cb ListInvitationsFunc) error {

}

func (db *InvitationsDocstoreDatabase) ListInvitationsWithInviter(context.Context, ListInvitationsFunc, *subscription.Subscription) error {

}

func (db *InvitationsDocstoreDatabase) Close() error {
	return db.collection.Close()
}
