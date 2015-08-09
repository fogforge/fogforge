package config

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	src := ``
	config, err := FromString("foo.ff", src)
	if err != nil {
		t.Fatal(err)
	}

	if got, expected := config.SourceFilename, "foo.ff"; got != expected {
		t.Errorf("SourceFilename == %#v; want %#v", got, expected)
	}

	if got, expected := len(config.UpstreamStates), 0; got != expected {
		t.Errorf("len(UpstreamStates) == %#v; want %#v", got, expected)
	}

	if got, expected := config.Build, (*Stage)(nil); got != expected {
		t.Errorf("Build == %#v; want %#v", got, expected)
	}
	if got, expected := config.Deploy, (*Stage)(nil); got != expected {
		t.Errorf("Deploy == %#v; want %#v", got, expected)
	}
	if got, expected := config.Infrastructure, (*Stage)(nil); got != expected {
		t.Errorf("Infrastructure == %#v; want %#v", got, expected)
	}
}

func TestFull(t *testing.T) {
	src := `
upstream "build_infrastructure" {
    source = "http://release.example.com/build_infrastructure"
}
upstream "some_library" {
    source = "http://release.example.com/some_library"
}

build {
    config_dir = "release/build"
    variables {
        foo = "baz"
        bar = "bar"
    }
}

deploy {
    config_dir = "release/deploy/app"
    variables {}
}

infrastructure {
    config_dir = "release/deploy/infrastructure"
    variables {}
}
`
	config, err := FromString("foo.ff", src)
	if err != nil {
		t.Fatal(err)
	}

	if got, expected := config.SourceFilename, "foo.ff"; got != expected {
		t.Errorf("SourceFilename == %#v; want %#v", got, expected)
	}

	if got, expected := len(config.UpstreamStates), 2; got != expected {
		t.Errorf("len(UpstreamStates) == %#v; want %#v", got, expected)
	}

	if got, expected := config.UpstreamStates["build_infrastructure"].SourceURL, "http://release.example.com/build_infrastructure"; got != expected {
		t.Errorf("UpstreamStates[build_infrastructure].SourceURL == %#v; want %#v", got, expected)
	}
	if got, expected := config.UpstreamStates["some_library"].SourceURL, "http://release.example.com/some_library"; got != expected {
		t.Errorf("UpstreamStates[some_library].SourceURL == %#v; want %#v", got, expected)
	}

	if got, expected := config.Build.ConfigDir, InterpolationString("release/build"); got != expected {
		t.Errorf("Build.ConfigDir == %#v; want %#v", got, expected)
	}
	if got, expected := len(config.Build.Variables), 2; got != expected {
		t.Errorf("len(Build.Variables) == %#v; want %#v", got, expected)
	}

	if got, expected := config.Deploy.ConfigDir, InterpolationString("release/deploy/app"); got != expected {
		t.Errorf("Deploy.ConfigDir == %#v; want %#v", got, expected)
	}
	if got, expected := len(config.Deploy.Variables), 0; got != expected {
		t.Errorf("len(Deploy.Variables) == %#v; want %#v", got, expected)
	}

	if got, expected := config.Infrastructure.ConfigDir, InterpolationString("release/deploy/infrastructure"); got != expected {
		t.Errorf("Infrastructure.ConfigDir == %#v; want %#v", got, expected)
	}
	if got, expected := len(config.Infrastructure.Variables), 0; got != expected {
		t.Errorf("len(Infrastructure.Variables) == %#v; want %#v", got, expected)
	}
}
