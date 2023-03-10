package types

import (
	"encoding/json"
)

type ProfilesMsg struct {
	SaveProfile               *json.RawMessage `json:"save_profile"`
	DeleteProfile             *json.RawMessage `json:"delete_profile"`
	RequestDtagTransfer       *json.RawMessage `json:"request_dtag_transfer"`
	AcceptDtagTransferRequest *json.RawMessage `json:"accept_dtag_transfer_request"`
	RefuseDtagTransferRequest *json.RawMessage `json:"refuse_dtag_transfer_request"`
	CancelDtagTransferRequest *json.RawMessage `json:"cancel_dtag_transfer_request"`
	LinkChainAccount          *json.RawMessage `json:"link_chain_account"`
	UnlinkChainAccount        *json.RawMessage `json:"unlink_chain_account"`
	SetDefaultExternalAddress *json.RawMessage `json:"set_default_external_address"`
	LinkApplication           *json.RawMessage `json:"link_application"`
	UnlinkApplication         *json.RawMessage `json:"unlink_application"`
}

type ProfilesQuery struct {
	Profile                      *json.RawMessage `json:"profile"`
	IncomingDtagTransferRequests *json.RawMessage `json:"incoming_dtag_transfer_requests"`
	ChainLinks                   *json.RawMessage `json:"chain_links"`
	ChainLinkOwners              *json.RawMessage `json:"chain_link_owners"`
	DefaultExternalAddresses     *json.RawMessage `json:"default_external_addresses"`
	ApplicationLinks             *json.RawMessage `json:"application_links"`
	ApplicationLinkByClientID    *json.RawMessage `json:"application_link_by_client_id"`
	ApplicationLinkOwners        *json.RawMessage `json:"application_link_owners"`
}
