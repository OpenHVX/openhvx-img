package images

import "time"

type Image struct {
	ID        string  `json:"id"`
	Scope     string  `json:"scope"`    // "public" | "tenant"
	TenantID  *string `json:"tenantId"` // nil si public
	Name      string  `json:"name"`
	Path      string  `json:"path"`
	OS        string  `json:"os"`
	Arch      string  `json:"arch"`
	Gen       int     `json:"gen"`
	SizeBytes int64   `json:"sizeBytes"`
	Mtime     string  `json:"mtime"`
}

type Index struct {
	Schema      string    `json:"schema"`
	GeneratedAt time.Time `json:"generatedAt"`
	Images      []Image   `json:"images"`
}
