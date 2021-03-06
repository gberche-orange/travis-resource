package model

type Source struct {
	Repository     string   `json:"repository"`
	Branch         string   `json:"branch"`
	BranchRegex    string   `json:"branch-regex"`
	GithubToken    string   `json:"github-token"`
	TravisToken    string   `json:"travis-token"`
	CheckAllBuilds bool     `json:"check-all-builds"`
	Pro            bool     `json:"pro"`
	Url            string   `json:"travis-url"`
	CheckOnState   string   `json:"check-on-state"`
}
type OutParams struct {
	Build      interface{}  `json:"build"`
	Branch     string       `json:"branch"`
	Repository string       `json:"repository"`
	SkipWait   bool         `json:"skip-wait"`
}
type InParams struct {
	DownloadLogs bool `json:"download-logs"`
}
type Version struct {
	BuildNumber string `json:"build"`
}

type Metadata struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type CheckRequest struct {
	Source  Source  `json:"source"`
	Version Version `json:"version"`
}
type InRequest struct {
	Source   Source   `json:"source"`
	Version  Version  `json:"version"`
	InParams InParams `json:"params"`
}
type OutRequest struct {
	Source    Source    `json:"source"`
	Version   Version   `json:"version"`
	OutParams OutParams `json:"params"`
}
type CheckResponse []Version
type InResponse struct {
	Metadata []Metadata   `json:"metadata"`
	Version  Version      `json:"version"`
}

type OutResponse struct {
	Metadata []Metadata   `json:"metadata"`
	Version  Version      `json:"version"`
}