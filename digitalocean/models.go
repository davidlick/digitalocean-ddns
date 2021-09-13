package digitalocean

type DomainRecords struct {
	Data string `json:"data"`
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type RecordsResponse struct {
	DomainRecords []DomainRecords `json:"domain_records"`
}
