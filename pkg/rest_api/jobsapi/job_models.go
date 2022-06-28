package jobsapi

type NewJobRequest struct {
	GitUrl           string `json:"giturl"`
	TerraformVersion string `json:"terraform_version"`
}
