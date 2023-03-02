package entity

type Pharmacy struct {
	Date         string `json:"fecha"`
	StoreID      string `json:"local_id"`
	StoreName    string `json:"local_nombre"`
	CommuneName  string `json:"comuna_nombre"`
	Locality     string `json:"localidad_nombre"`
	StoreAddress string `json:"local_direccion"`
	OpenHour     string `json:"funcionamiento_hora_apertura"`
	CloseHour    string `json:"funcionamiento_hora_cierre"`
	StorePhone   string `json:"local_telefono"`
	StoreLat     string `json:"local_lat"`
	StoreLng     string `json:"local_lng"`
	FuncDay      string `json:"funcionamiento_dia"`
	FkRegion     string `json:"fk_region"`
	FkCommune    string `json:"fk_comuna"`
	FkLocality   string `json:"fk_localidad"`
}

type Request struct {
	FilterValue  string `json:"filtroValor"`
	TypeResponse string `json:"tipoRespuesta"`
}
