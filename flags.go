package littleflags

import (
    "errors"
    "fmt"
    "os"
    "strings"
)

const (
    special_key_alpha = "alpha"
)

var (
    alpha_enabled bool = false
)

type Feature struct {
    enabled bool
    name    string
}

func (f Feature) String() string {
    status := "disabled"
    if f.enabled {
        status = "enabled"
    }
    return fmt.Sprintf("Feature{name:\"%s\", status:\"%s\"}", f.name, status)
}

func (this Feature) IsEqual(other *Feature) bool {
    if this.name != other.name {
        return false
    }

    if this.enabled != other.enabled {
        return false
    }

    return true
}

var initializedFeatures map[string]*Feature

func readOSEnvironment(prefix string) map[string]*Feature {
    features := make(map[string]*Feature)

    for _, e := range os.Environ() {
        pair := strings.Split(e, "=")

        key := pair[0]
        if strings.HasPrefix(key, prefix) {
            enabled := len(pair[1]) > 0
            name    := strings.ToLower(key[len(prefix) + 1:])

            if name == special_key_alpha {
                alpha_enabled = true
            } else {
                features[name] = &Feature{enabled, name}
            }
        }
    }

    return features
}

func InitializeLittleFlags() {
    Initialize("LITTLE_FLAGS")
}

func Initialize(prefix string) {
    initializedFeatures = readOSEnvironment(prefix)
}

func GetFeature(name string) (*Feature, error) {
    if initializedFeatures == nil {
        return nil, errors.New("Feature toggles not initialized")
    }
    if name == "" {
        return nil, errors.New("Given name is empty")
    }

    name     = strings.ToLower(name)
    feature := initializedFeatures[name]

    if feature == nil {
        feature = &Feature{false, name}
        initializedFeatures[name] = feature
    }

    return feature, nil
}

func (feature *Feature) IsEnabled() bool {
    return alpha_enabled || feature.enabled
}
