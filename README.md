# neos-site-clone

A portable utility to clone Neos CMS site dumps.

## What does it do?

It reads a Neos XML site export and maps all node identifiers of the given site dump to newly generated identifiers.
This way an existing site can be cloned.

## Installation

Download a release and unpack it into your local binaries directory.

E.g. for macOS on M1:

```sh
curl
```

## How to use it:

```sh
neos-site-clone < DistributionPackages/SiteA/Resources/Private/Content/Sites.xml > DistributionPackages/SiteB/Resources/Private/Content/Sites.xml
```
