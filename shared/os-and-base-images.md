# Overview
This policy defines the organization-wide defaults for operating systems and OCI-compatible base images.
Apply it when creating an environment or changing its OS or base image. Do not migrate a working environment solely
to conform to this policy unless the task explicitly includes that migration.

# Decision order
Select an OS or base image in this order:

1. Correctness and compatibility with the workload.
2. Security updates and upstream support.
3. Operability, including diagnostics and recovery.
4. Measured energy, CPU, memory, storage, and network efficiency.

Image size alone does not demonstrate lower runtime energy consumption. Measure the real workload when energy or
resource efficiency is a goal.

# Standard selection
Use this table before considering an alternative distribution or image family.

| Environment | Default | Use when |
| --- | --- | --- |
| Bare metal, VM, CI runner, development environment, or other general-purpose environment | Debian Trixie | The environment needs a complete, stable OS and Debian is supported by the platform and workload. |
| OCI build, test, CI, or debugging stage | A Debian Trixie variant of the upstream tool image, otherwise `debian:trixie` | The stage needs a shell, package manager, build toolchain, or broader userspace. |
| OCI production runtime for a fully static binary | `gcr.io/distroless/static-debian13:nonroot` | The binary requires neither libc nor a shell or package manager. For Go, this normally means a build with `CGO_ENABLED=0`. |
| OCI production runtime with OS-level dependencies | `debian:trixie-slim` | The application needs glibc, shared libraries, runtime Debian packages, or shell-based diagnostics. |
| Third-party language, tool, application, or infrastructure service | Its supported Debian-based image variant | Upstream publishes and supports a Debian-based variant. Prefer Trixie when available. |

Debian Trixie is the default OS release. Use an explicit release codename or version instead of floating aliases such
as `latest`. See the [Debian Trixie release information](https://www.debian.org/releases/trixie/) and the
[Debian Official Image](https://hub.docker.com/_/debian).

# Strict rules
- MUST use Debian Trixie when selecting an OS for bare metal, virtual machines, CI runners, development environments, and other general-purpose environments unless an exception below applies.
- MUST prefer a supported Debian Trixie variant of a language, tool, application, or service image when upstream provides one.
- MUST use `debian:trixie` for custom general-purpose OCI build, test, CI, and debugging stages that need a shell, package manager, or broader userspace.
- MUST use `debian:trixie-slim` for production OCI runtimes that need glibc, shared libraries, runtime Debian packages, or shell-based diagnostics.
- MUST use `gcr.io/distroless/static-debian13:nonroot` for production OCI runtimes containing fully static binaries that need neither libc nor a shell or package manager.
- MUST separate build and runtime stages with a multi-stage build when the build environment contains tools or dependencies that the runtime does not need.
- MUST run production workloads as a non-root user unless the workload has a documented technical requirement for root.
- MUST verify required runtime assets, including CA certificates, timezone data, user and group entries, writable directories, and dynamic libraries.
- MUST NOT select an OS or base image from compressed size alone.
- MUST NOT migrate an existing working environment solely to enforce this policy unless migration is in scope for the task.

The Distroless project documents the static image's contents, supported Debian 13 tags, and lack of a shell and package
manager in its [official repository](https://github.com/GoogleContainerTools/distroless).

# Situational alternatives
- MAY use Alpine when upstream requires or explicitly recommends it, or when a measured footprint constraint justifies it and musl compatibility is verified by the relevant tests.
- MAY use Ubuntu LTS when hardware enablement, vendor support, CI infrastructure, or the hosting platform requires it.
- MUST prefer a supported Debian-based upstream image over another upstream variant unless maintainers explicitly recommend otherwise.
- MUST use the upstream official image when third-party software has no supported Debian-based image.
- MUST NOT rebuild third-party software images on Debian solely to satisfy the Debian default.
- MUST document which condition justified a situational alternative in the change or pull request.
- MUST obtain explicit user `Accept` before selecting any OS distribution or base image family not covered by this policy, and MUST document the approved exception in the relevant `Strict rules` section.
