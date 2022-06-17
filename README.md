# neos-site-clone

A portable utility to clone Neos CMS site dumps.

## What does it do?

It reads a Neos XML site export and maps all node identifiers of the given site dump to newly generated identifiers.
This way an existing site can be cloned.

## Installation

Download a binary from a release for your OS and architecture and move it to your local binaries directory.

E.g. for macOS on M1:

```sh
curl -LO https://github.com/networkteam/neos-site-clone/releases/download/v0.1.1/neos-site-clone_darwin_arm64 && sudo mv neos-site-clone_darwin_arm64 /usr/local/bin/neos-site-clone && sudo chmod +x /usr/local/bin/neos-site-clone
```

## How to use it:

```sh
neos-site-clone < DistributionPackages/SiteA/Resources/Private/Content/Sites.xml > DistributionPackages/SiteB/Resources/Private/Content/Sites.xml
```
