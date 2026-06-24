# updater-terraform

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
  --certificate-identity-regexp 'https://github.com/SemRels/updater-terraform/.github/workflows/release.yml.*' \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com
```


## Configuration

```yaml
plugins:
  - name: updater-terraform
    path: ~/.semrel/plugins/semrel-plugin-updater-terraform
    env:
      SEMREL_PLUGIN_FILE: "variables.tf"
      SEMREL_PLUGIN_VARIABLE: "app_version"
```

## `SEMREL_PLUGIN_*` variables

| Name | Required | Description | Default |
| --- | --- | --- | --- |
| `SEMREL_PLUGIN_FILE` | Optional | Path to the Terraform file to update. | variables.tf |
| `SEMREL_PLUGIN_VARIABLE` | Optional | Terraform variable name to update. | app_version |

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
