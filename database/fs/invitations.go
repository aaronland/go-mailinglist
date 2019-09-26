package fs

import (
	"context"
	"fmt"
	"github.com/aaronland/go-mailinglist/database"
	"github.com/aaronland/go-mailinglist/invitation"
	"os"
)

type FSInvitationsDatabase struct {
	database.InvitationsDatabase
	root string
}

func NewFSInvitationsDatabase(root string) (database.InvitationsDatabase, error) {

	abs_root, err := ensureRoot(root)

	if err != nil {
		return nil, err
	}

	db := FSInvitationsDatabase{
		root: abs_root,
	}

	return &db, nil
}

func (db *FSInvitationsDatabase) AddInvitation(invite *invitation.Invitation) error {

	path := db.pathForInvitation(invite)

	_, err := os.Stat(path)

	if err == nil {
		return nil
	}

	return db.writeInvitation(invite, path)
}

func (db *FSInvitationsDatabase) RemoveInvitation(invite *invitation.Invitation) error {

	path := db.pathForInvitation(invite)

	_, err := os.Stat(path)

	if err != nil {

		if os.IsNotExist(err) {
			return nil
		}

		return err
	}

	return os.Remove(path)
}

func (db *FSInvitationsDatabase) UpdateInvitation(invite *invitation.Invitation) error {

	path := db.pathForInvitation(invite)

	_, err := os.Stat(path)

	if err == nil {
		return nil
	}

	return db.writeInvitation(invite, path)
}

// this won't work because the invitation won't have a FS path with the invitee address in it...

func (db *FSInvitationsDatabase) GetInvitationWithInvitee(addr string) (*invitation.Invitation, error) {

	path := pathForAddress(db.root, addr)

	_, err := os.Stat(path)

	if err != nil {

		if os.IsNotExist(err) {
			return nil, new(database.NoRecordError)
		}

		return nil, err
	}

	return db.readInvitation(path)
}

func (db *FSInvitationsDatabase) readInvitation(path string) (*invitation.Invitation, error) {

	invite, err := unmarshalData(path, "invitation")

	if err != nil {
		return nil, err
	}

	return invite.(*invitation.Invitation), nil
}

func (db *FSInvitationsDatabase) writeInvitation(invite *invitation.Invitation, path string) error {

	return marshalData(invite, path)
}

func (db *FSInvitationsDatabase) ListInvitations(ctx context.Context, cb database.ListInvitationsFunc) error {

	local_cb := func(ctx context.Context, invite *invitation.Invitation) error {

		select {
		case <-ctx.Done():
			return nil
		default:
			// pass
		}

		return cb(invite)
	}

	return db.crawlInvitations(ctx, local_cb)
}

func (db *FSInvitationsDatabase) ListInvitationsWithInviter(ctx context.Context, cb database.ListInvitationsFunc, inviter string) error {

	local_cb := func(ctx context.Context, invite *invitation.Invitation) error {

		select {
		case <-ctx.Done():
			return nil
		default:
			// pass
		}

		if invite.Inviter != inviter {
			return nil
                }

		return cb(invite)
	}

	return db.crawlInvitations(ctx, local_cb)
}

func (db *FSInvitationsDatabase) crawlInvitations(ctx context.Context, cb func(ctx context.Context, invite *invitation.Invitation) error) error {

	local_cb := func(ctx context.Context, path string) error {

		invite, err := db.readInvitation(path)

		if err != nil {
			return err
		}

		return cb(ctx, invite)
	}

	return crawlDatabase(ctx, db.root, local_cb)
}

func (db *FSInvitationsDatabase) pathForInvitation(invite *invitation.Invitation) string {
	fname := fmt.Sprintf("%s-%d", invite.Inviter, invite.Created)	// FIX ME...
	return pathForAddress(db.root, fname)
}
