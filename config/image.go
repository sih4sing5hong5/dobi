package config

import (
	"fmt"
	"path/filepath"
	"strings"
)

// ImageConfig ia a data object for image resource
type ImageConfig struct {
	Image      string
	Dockerfile string
	Context    string
	Args       map[string]string
	Pull       bool
	Depends    []string
}

// Dependencies returns the list of implicit and explicit dependencies
func (c *ImageConfig) Dependencies() []string {
	return c.Depends
}

// Validate checks that all fields have acceptable values
func (c *ImageConfig) Validate(config *Config) error {
	if missing := config.missingResources(c.Depends); len(missing) != 0 {
		reason := fmt.Sprintf("missing dependencies: %s", strings.Join(missing, ", "))
		return NewResourceError(c, reason)
	}

	// TODO: validate required fields are set
	// TODO: validate no tag on image name
	return nil
}

func (c *ImageConfig) String() string {
	dir := filepath.Join(c.Context, c.Dockerfile)
	return fmt.Sprintf("Build image '%s' from '%s'", c.Image, dir)
}

// NewImageConfig creates a new ImageConfig with default values
func NewImageConfig() *ImageConfig {
	return &ImageConfig{Dockerfile: "Dockerfile", Context: "."}
}

func imageFromConfig(values map[string]interface{}) (Resource, error) {
	image := NewImageConfig()
	return image, Transform(values, image)
}

func init() {
	RegisterResource("image", imageFromConfig)
}