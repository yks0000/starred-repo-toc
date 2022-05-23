package schemas

type GitHubResponseField struct {
	Name        string
	FullName    string
	Description string
	CloneUrl    string
	OwnerName   string
	StarCount   int
	LastUpdated string
	Topics      []string
}
