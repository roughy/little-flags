package littleflags

import (
    "os"
    "strings"
    "testing"
)

func TestGetFeatureNotInitialized(t *testing.T) {
    name        := "foo"
    actual, err := GetFeature(name)

    if actual != nil {
        t.Errorf("Expected GetFeature(\"%s\") == nil", name)
    }
    if err == nil {
        t.Error("Expected to get err == nil")
    }
    if !strings.Contains(err.Error(), "Feature toggles not initialized") {
        t.Errorf("Unexpected error: %v", err)
    }
}

func TestGetFeatureEmptyName(t *testing.T) {
    os.Setenv("TEST_FLAG_FEATURE1", "on")
    Initialize("TEST_FLAG")

    name        := ""
    actual, err := GetFeature(name)

    if actual != nil {
        t.Errorf("Expected GetFeature(\"%s\") == nil", name)
    }
    if err == nil {
        t.Error("Expected to get err == nil")
    }
    if !strings.Contains(err.Error(), "Given name is empty") {
        t.Errorf("Unexpected error: %v", err)
    }
}

func TestGetFeatureButDoNotExists(t *testing.T) {
    os.Setenv("TEST_FLAG_FEATURE1", "on")
    Initialize("TEST_FLAG")

    name        := "feature2"
    expected    := &Feature{false, name}
    actual, err := GetFeature(name)

    if actual == nil {
        t.Errorf("Expected GetFeature(\"%s\") == %s but got nil", name, expected.String())
    }

    if !expected.IsEqual(actual) {
        t.Errorf("Expected GetFeature(\"%s\") == %s but got %s", name, expected.String(), actual.String())
    }

    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
}

func TestGetFeatureCaseInsensitive(t *testing.T) {
    os.Setenv("TEST_FLAG_FEATURE1", "on")
    Initialize("TEST_FLAG")

    name        := "fEAture1"
    expected    := &Feature{true, "feature1"}
    actual, err := GetFeature(name)

    if actual == nil {
        t.Errorf("Expected GetFeature(\"%s\") == %s but got nil", name, expected.String())
    }

    if !expected.IsEqual(actual) {
        t.Errorf("Expected GetFeature(\"%s\") == %s but got %s", name, expected.String(), actual.String())
    }

    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
}

func TestGetFeature(t *testing.T) {
    os.Setenv("TEST_FLAG_FEATURE1", "on")
    Initialize("TEST_FLAG")

    name        := "feature1"
    expected    := &Feature{true, name}
    actual, err := GetFeature(name)

    if actual == nil {
        t.Errorf("Expected GetFeature(\"%s\") == %s but got nil", name, expected.String())
    }

    if !expected.IsEqual(actual) {
        t.Errorf("Expected GetFeature(\"%s\") == %s but got %s", name, expected.String(), actual.String())
    }

    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
}

func TestIsEnabledButDoNotExists(t *testing.T) {
    os.Setenv("TEST_FLAG_FEATURE1", "on")
    Initialize("TEST_FLAG")

    name     := "feature2"
    expected := false

    f, err := GetFeature(name)
    actual := f.IsEnabled()

    if expected != actual {
        t.Errorf("Expected IsEnabled() == false but got true")
    }

    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
}

func TestIsEnabled(t *testing.T) {
    os.Setenv("TEST_FLAG_FEATURE1", "on")
    Initialize("TEST_FLAG")

    name     := "feature1"
    expected := true

    f, err := GetFeature(name)
    actual := f.IsEnabled()

    if expected != actual {
        t.Errorf("Expected IsEnabled() == true but got false")
    }

    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
}

func TestAlphaEnablesAll(t *testing.T) {
    os.Setenv("TEST_FLAG_ALPHA", "on")
    Initialize("TEST_FLAG")

    name     := "feature1"
    expected := true

    f, err := GetFeature(name)
    actual := f.IsEnabled()

    if expected != actual {
        t.Errorf("Expected IsEnabled() == true but got false")
    }

    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    name = "alpha"
    f, err = GetFeature(name)
    actual = f.IsEnabled()

    if expected != actual {
        t.Errorf("Expected IsEnabled() == true but got false")
    }

    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    name = "foobar"
    f, err = GetFeature(name)
    actual = f.IsEnabled()

    if expected != actual {
        t.Errorf("Expected IsEnabled() == true but got false")
    }

    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
}

func TestIsEnabledWithDefaultInit(t *testing.T) {
    os.Setenv("LITTLE_FLAGS_FEATURE1", "on")
    InitializeLittleFlags()

    name     := "feature1"
    expected := true

    f, err := GetFeature(name)
    actual := f.IsEnabled()

    if expected != actual {
        t.Errorf("Expected IsEnabled() == true but got false")
    }

    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
}

func TestAlphaEnablesAllWithDefaultInit(t *testing.T) {
    os.Setenv("LITTLE_FLAGS_ALPHA", "on")
    InitializeLittleFlags()

    name     := "feature1"
    expected := true

    f, err := GetFeature(name)
    actual := f.IsEnabled()

    if expected != actual {
        t.Errorf("Expected IsEnabled() == true but got false")
    }

    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    name = "alpha"
    f, err = GetFeature(name)
    actual = f.IsEnabled()

    if expected != actual {
        t.Errorf("Expected IsEnabled() == true but got false")
    }

    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    name = "foobar"
    f, err = GetFeature(name)
    actual = f.IsEnabled()

    if expected != actual {
        t.Errorf("Expected IsEnabled() == true but got false")
    }

    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
}
