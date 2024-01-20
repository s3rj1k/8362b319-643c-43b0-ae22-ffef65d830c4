package api

import "net"

type TimeMetadata struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GPS struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type City struct {
	ID int `json:"id"`
	GPS

	Name     string `json:"name"`
	DNSName  string `json:"dns_name"`
	HubScore int    `json:"hub_score"`
}

type Country struct {
	ID int `json:"id"`

	Name string `json:"name"`
	Code string `json:"code"`
	City City   `json:"city"`
}

type Location struct {
	ID int `json:"id"`
	TimeMetadata
	GPS

	Country Country `json:"country"`
}

type Services struct {
	ID int `json:"id"`
	TimeMetadata

	Name       string `json:"name"`
	Identifier string `json:"identifier"`
}

type Pivot struct {
	TechnologyID int    `json:"technology_id"`
	ServerID     int    `json:"server_id"`
	Status       string `json:"status"`
}

type Technologies struct {
	ID int `json:"id"`
	TimeMetadata

	Name       string `json:"name"`
	Identifier string `json:"identifier"`
	Metadata   []any  `json:"metadata"`
	Pivot      Pivot  `json:"pivot"`
}

type Type struct {
	ID int `json:"id"`
	TimeMetadata

	Title      string `json:"title"`
	Identifier string `json:"identifier"`
}

type Groups struct {
	ID int `json:"id"`
	TimeMetadata

	Title      string `json:"title"`
	Identifier string `json:"identifier"`
	Type       Type   `json:"type"`
}

type Value struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

type Specifications struct {
	ID int `json:"id"`

	Title      string  `json:"title"`
	Identifier string  `json:"identifier"`
	Values     []Value `json:"values"`
}

type IP struct {
	ID int `json:"id"`

	IP      net.IP `json:"ip"`
	Version int    `json:"version"`
}

type IPs struct {
	ID int `json:"id"`
	TimeMetadata

	ServerID int    `json:"server_id"`
	IPID     int    `json:"ip_id"`
	Type     string `json:"type"`
	IP       IP     `json:"ip"`
}

type Response struct {
	ID int `json:"id"`
	TimeMetadata

	Name           string           `json:"name"`
	Station        net.IP           `json:"station"`
	Ipv6Station    net.IP           `json:"ipv6_station"`
	Hostname       string           `json:"hostname"`
	Load           int              `json:"load"`
	Status         string           `json:"status"`
	Cpt            int              `json:"cpt"`
	Locations      []Location       `json:"locations"`
	Services       []Services       `json:"services"`
	Technologies   []Technologies   `json:"technologies"`
	Groups         []Groups         `json:"groups"`
	Specifications []Specifications `json:"specifications"`
	IPs            []IPs            `json:"ips"`
}
