# VPSentry

Inspect, plan, and apply a verified security baseline to your Ubuntu VPS.

VPSentry is a Go CLI project for setting up basic security controls on new VPS instances and reviewing existing servers before making changes. I am building it to learn Go and explore safe, idempotent Linux automation.

## Status

**Early development. Not ready for production use.**

The current prototype queries user accounts with `getent` and checks group membership with `id`. It does not create accounts or change SSH, firewall, or package settings.

The group check still uses substring matching, and the CLI does not yet handle all query errors or distinguish every account state. Its output is not an authorization audit.

## Planned workflow

1. **Inspect:** collect the current state without changing the system. Report query failures separately from missing configuration.
2. **Plan:** compare the observed state with the requested configuration and show proposed changes, conflicts, and dependencies.
3. **Execute:** apply only authorized changes. Stop dependent operations when a prerequisite fails.
4. **Verify:** inspect again after changes and report what was confirmed, what failed, and what remains unknown.

Repeated runs should leave an already compliant configuration unchanged. If an operation fails partway through, the next run should inspect the actual state rather than assume earlier steps completed.

## Roadmap

- Create or configure an explicitly selected non-root administrative account.
- Ask for confirmation before granting `sudo` group membership to an existing account.
- Install public SSH keys without collecting client private keys.
- Verify alternative access before restricting root or password-based SSH login.
- Configure automatic security updates with an explicit reboot policy.
- Configure a firewall while preserving authorized administrative and application access.
- Consider optional Fail2ban and VPN support after the core workflow is tested.

These are planned features, not current capabilities. A baseline does not guarantee that a server or its applications are secure or production-ready.

## Development

You need the Go SDK to build VPSentry. The current inspection commands target Ubuntu Server and require `getent` and `id`.

Build from the project directory:

```bash
go build -o vpsentry .
```

On Ubuntu, run the current read-only prototype:

```bash
./vpsentry --admin-user ubuntu
```

The current queries normally do not need root privileges. Future operations that modify the system will require appropriate privileges.

### Cross-compile from macOS

For an Ubuntu server with an x86-64 CPU:

```bash
GOOS=linux GOARCH=amd64 go build -o vpsentry-linux-amd64 .
```

For an ARM64 server:

```bash
GOOS=linux GOARCH=arm64 go build -o vpsentry-linux-arm64 .
```

Copy the matching binary to the server. Go does not need to be installed there.

### Testing

Use a disposable Ubuntu VM with a snapshot and console access independent of SSH. Do not test privileged changes on a server you depend on.

Automated tests and VM integration tests are not implemented yet. Planned scenarios include existing and missing accounts, failed inspections, repeated execution, and partial failures.

## Project structure

All Go files currently belong to the same `main` package.

| File | Responsibility |
| --- | --- |
| `main.go` | Current CLI and prototype account queries |
| `inspection.go` | Placeholder for system inspection |
| `planning.go` | Placeholder for proposed changes |
| `execution.go` | Placeholder for approved operations and result checks |

## License

MIT. A license file has not been added yet.

## Author

Gabriel Rodríguez Ovares — [@garze05](https://github.com/garze05)

[gabrielro.com](https://gabrielro.com)
