package images

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func BuildIndex(root string) (*Index, error) {
	var imgs []Image

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(d.Name()))
		if ext != ".vhdx" {
			return nil
		}

		stat, err := d.Info()
		if err != nil {
			return nil
		}

		// scope: public ou tenant
		scope := "public"
		var tenantID *string
		rel, _ := filepath.Rel(root, path)
		parts := strings.Split(rel, string(os.PathSeparator))
		if len(parts) > 1 && parts[0] == "tenants" {
			scope = "tenant"
			tenantID = &parts[1]
		}

		id := strings.TrimSuffix(d.Name(), filepath.Ext(d.Name()))
		img := Image{
			ID:        id,
			Scope:     scope,
			TenantID:  tenantID,
			Name:      id,
			Path:      path,
			OS:        GuessOS(id),
			Arch:      GuessArch(id),
			Gen:       GuessGen(id),
			SizeBytes: stat.Size(),
			Mtime:     stat.ModTime().UTC().Format(time.RFC3339),
		}
		imgs = append(imgs, img)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &Index{
		Schema:      "openhvx.images/v1",
		GeneratedAt: time.Now().UTC(),
		Images:      imgs,
	}, nil
}

func WriteIndex(idx *Index, output string, pretty bool) error {
	tmp := output + ".tmp"

	var data []byte
	var err error
	if pretty {
		data, err = json.MarshalIndent(idx, "", "  ")
	} else {
		data, err = json.Marshal(idx)
	}
	if err != nil {
		return err
	}

	if err := os.WriteFile(tmp, data, 0644); err != nil {
		return fmt.Errorf("write tmp: %w", err)
	}
	return os.Rename(tmp, output)
}
