# Contributing

We, the maintainers, love pull requests from everyone, but often find
we must say "no" despite how reasonable the proposal may seem.

For this reason, we ask that you open an issue to discuss proposed
changes prior to submitting a pull request for the implementation.
This helps us to provide direction as to implementation details, which
branch to base your changes on, and so on.

1. Open an issue to describe your proposed improvement or feature
1. Fork https://github.com/jamesdphillips/graphqlgen clone your fork to your workstation
1. Create your feature branch (`git checkout -b my-new-feature`)
1. If applicable, add a [CHANGELOG.md entry](#changelog) describing your change.
1. Push your feature branch (`git push origin my-new-feature`)
1. Create a Pull Request as appropriate based on the issue discussion

## Changelog

The [Changelog](CHANGELOG.md) is based on [keep a changelog v1.0](https://keepachangelog.com/en/1.0.0/).

All new changes go underneath the _Unreleased_ heading at the top of the Changelog.
Beyond that, here are some additional guidelines that should make it more clear where your
change goes in the Changelog.

### Added

Any _new_ functionality goes here. This may be a new field on a data type or a new data
type altogether; a new API endpoint; or possibly a whole new feature. In general, these
are sentences that start with the word "added."

Examples:

- now support 2018-10-01 version of GraphQL spec.
- added support for list type.
- updated parser to support list type.

### Changed

Changes to any existing component or functionality of the system that does not cause
breaking changes to users or developers go here. _Changed_ is distinguishable from
_Fixed_ in that it is an intentional change to existing functionality.

Examples:

- Refactored executor to support parallel execution of fields.

### Fixed

Fixed bugs go here.

Examples:

- Fixes issues with parsing list tokens with parentheses present.

### Deprecated

Deprecated should include any soon-to-be removed functionality. An entry here that
is user facing will likely yield entries in _Removed_ or _Breaking_ eventually.

Examples:

- `Parallel` configuration option is no longer supported.

### Removed

Removed is for the removal of functionality that does not directly impact users,
these entries most likely only impact implementors.  If user facing
functionality is removed, an entry should be added to the _Breaking Changes_
section instead.

Examples:

- Removed references to `encoding/json` in favor of `json-iter`.
- Removed unused `Store` interface for `BlobStore`.

### Security

Any fixes to address security exploits should be added to this section. If
available, include an associated CVE entry.

Examples:

- Upgraded build to use Go 1.9.1 to address [CVE-2017-15041](https://www.cvedetails.com/cve/CVE-2017-15041/)
- Fixed issue where users could view entities without permission

### Breaking Changes

Whenever you have to make a change that will cause implementors to make changes 
to their code.

Examples:

- Previously deprecated list types are no longer supported.
