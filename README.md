# updater-terraform

[![Latest Release](https://img.shields.io/github/v/release/SemRels/updater-terraform?label=version&color=blue)](https://github.com/SemRels/updater-terraform/releases/latest)
[![CI](https://github.com/SemRels/updater-terraform/actions/workflows/ci.yml/badge.svg)](https://github.com/SemRels/updater-terraform/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/SemRels/updater-terraform)](LICENSE)
[![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/SemRels/updater-terraform/badge)](https://scorecard.dev/viewer/?uri=github.com/SemRels/updater-terraform)
[![REUSE status](https://api.reuse.software/badge/github.com/SemRels/updater-terraform)](https://api.reuse.software/info/github.com/SemRels/updater-terraform)

Updates a Terraform variable that stores the application version.

This plugin is distributed as the standalone Go binary `semrel-plugin-updater-terraform`. Semrel executes the binary as a subprocess, provides plugin configuration through `SEMREL_PLUGIN_*` environment variables, provides release context through `SEMREL_*` environment variables, reads standard output, and treats exit code `0` as success and any non-zero exit code as failure. Install the binary in `~/.semrel/plugins/` or anywhere on your `$PATH`.

## Installation

### Binary

```bash
go install github.com/SemRels/updater-terraform/cmd/plugin@latest
```

### Docker

Pre-built, multi-platform images (linux/amd64, linux/arm64) are published to the GitHub Container Registry on every release:

```bash
docker pull ghcr.io/semrels/updater-terraform:latest
```

Images are signed with [cosign](https://github.com/sigstore/cosign) and include a full SBOM attestation. Verify the signature:

```bash
cosign verify ghcr.io/semrels/updater-terraform:latest \
  --certificate-identity-regexp 'https://github.com/SemRels/updater-terraform/.github/workflows/container.yml.*' \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com
```


## Configuration

```yaml
plugins:
  - uses: "@semrel/updater-terraform"
    phase: pre-tag
    args:
      file: "variables.tf"
      variable: "app_version"
```

## Options

Every option can be set as an entry under `args` in `.semrel.yaml`. semrel passes each entry to the plugin as the environment variable shown in the table, so you can also set that variable directly (for example in your CI environment) instead of using `args`.

| Argument | Environment variable | Required | Description | Default |
| --- | --- | --- | --- | --- |
| `file` | `SEMREL_PLUGIN_FILE` | Optional | Path to the Terraform file to update. | variables.tf |
| `variable` | `SEMREL_PLUGIN_VARIABLE` | Optional | Terraform variable name to update. | app_version |

## `SEMREL_*` release context used

| Variable | Description |
| --- | --- |
| `SEMREL_VERSION` | Resolved release version for the current run. |
| `SEMREL_NEXT_VERSION` | Next version computed by semrel for the release. |
| `SEMREL_DRY_RUN` | Whether semrel is running in dry-run mode. |

## Example behavior

The plugin updates the default value of the configured Terraform variable to the new release version.

## License

Apache-2.0
