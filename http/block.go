package http

import (
	"github.com/aaronland/go-mailinglist/block"	
	"github.com/aaronland/go-mailinglist/database"
	gohttp "net/http"
	"net/mail"
)

func IsBlocked(db database.BlockDatabase, req *gohttp.Request, addr *mail.Address) (bool, error) {

	var is_blocked bool
	var err error

	is_blocked, err = db.IsBlocked(addr.Address, block.RULE_TYPE_ADDRESS)

	if err != nil {
		return is_blocked, err
	}

	if is_blocked {
		return true, nil
	}
	
	is_blocked, err = db.IsBlocked(req.RemoteAddr, block.RULE_TYPE_IP)

	if err != nil {
		return is_blocked, err
	}

	if is_blocked {
		return true, nil
	}

	return false, nil
}
