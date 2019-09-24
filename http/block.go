package http

import (
	"github.com/aaronland/go-mailinglist/block"	
	"github.com/aaronland/go-mailinglist/database"
	gohttp "net/http"
	"net/mail"
)

func IsAddressBlocked(db database.BlockDatabase, addr *mail.Address) (bool, error) {

	return false, nil
	
	var is_blocked bool
	var err error

	is_blocked, err = db.IsBlocked(addr.Address, block.RULE_TYPE_ADDRESS)

	if err != nil {
		return is_blocked, err
	}

	if is_blocked {
		return true, nil
	}

	return false, nil
}

func IsHostBlocked(db database.BlockDatabase, req *gohttp.Request) (bool, error) {

	return false, nil

	var is_blocked bool
	var err error
	
	// something something something check CIDR blocks...
	
	is_blocked, err = db.IsBlocked(req.RemoteAddr, block.RULE_TYPE_IP)

	if err != nil {
		return is_blocked, err
	}

	if is_blocked {
		return true, nil
	}

	return false, nil
}
