package common

// type ASRequestEncrypted struct {
// 	ASClientRequest string `json:"clientRequest"`
// }

type ASClientRequest struct {
	TS   string `json:"ts"`
	ID   string `json:"id"`
	Req  string `json:"req"`
	Addr string `json:"addr"`
}

type ASResponseEncrypted struct {
	ASClientResponse string `json:"clientResponse"`
	TGT              string `json:"tgt"`
}

type ASClientResponse struct {
	SessionKey string `json:"sessionKey"`
	TS         string `json:"ts"`
	TGSAddr    string `json:"tgsAddr"`
	Exp        string `json:"exp"`
}

type TicketGrantingTicket struct {
	ID         string `json:"id"`
	SessionKey string `json:"sessionKey"`
	TS         string `json:"ts"`
	Addr       string `json:"addr"`
	Exp        string `json:"exp"`
}

type TGSRequestASEncrypted struct {
	TGSRequestAS string `json:"requestAS"`
}

type TGSRequestAS struct {
	Dest string `json:"dest"`
}

type ASResponseTGSEncrypted struct {
	ASResponseTGS string `json:"responseTGS"`
}

type ASResponseTGS struct {
	ServerKey string `json:"serverKey"`
}

type TGSRequestEncrypted struct {
	TGSClientRequest string `json:"clientRequest"`
	TGT              string `json:"tgt"`
}

type TGSClientRequest struct {
	ID  string `json:"id"`
	TS  string `json:"ts"`
	Req string `json:"req"`
}

type TGSResponseEncrypted struct {
	TGSClientResponse     string `json:"clientResponse"`
	TGSServerDataResponse string `json:"serverDataResponse"`
}

type TGSServerDataResponse struct {
	ID           string `json:"id"`
	GeneratedKey string `json:"generatedKey"`
	Addr         string `json:"addr"`
	TS           string `json:"ts"`
	Exp          string `json:"exp"`
}

type TGSClientResponse struct {
	GeneratedKey string `json:"generatedKey"`
	Dest         string `json:"dest"`
	TS           string `json:"ts"`
	Exp          string `json:"exp"`
}

type ServerRequestEncrypted struct {
	ServerClientRequest string `json:"serverClientRequest"`
	TGSServerResponse   string `json:"TGSClientResponse"`
}

type ServerClientRequest struct {
	ID string `json:"id"`
	TS string `json:"ts"`
}

type ServerResponseEncrypted struct {
	ServerResponse string `json:"serverResponse"`
}

type ServerResponse struct {
	Name string `json:"name"`
	TS   string `json:"ts"`
}
