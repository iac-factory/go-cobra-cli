# `go-cobra-cli`

[[_TOC_]]

## Overview

`go-cobra-cli` is a full, working implementation of a `go` cli application; the following
project is a template that can be forked, copied for private or personal usage.

The remaining `README.md`'s contents is for non-template usage. Of course, replace
section(s) where appropriate -- namely `go-cobra-cli` and VCS references.

## Getting Started

### Installing

```bash
brew tap iac-factory/homebrew-taps
# brew tap ethr/homebrew-go-cobra-cli-tap git@gitlab.com:ethrgg/templates/homebrew-go-cobra-cli-tap

# brew install iac-factory/homebrew-taps/go-cobra-cli --verbose --debug

# --> Simplified
brew install go-cobra-cli
```

### Pre-Commit

1. Install pre-commit from https://pre-commit.com/#install
2. Create a `.pre-commit-config.yaml` file at the root of your repository with the following content:

   ```
   repos:
     - repo: https://github.com/gitleaks/gitleaks
       rev: latest
       hooks:
         - id: gitleaks
   ```

3. Auto-update the config to the latest repos' versions by executing `pre-commit autoupdate`
4. Install with `pre-commit install`

---

## Local Deployment & Initial Setup Instructions

1. Install `goreleaser` if it isn't installed

    ```bash
    brew install goreleaser/tap/goreleaser
    ```

1. Initialize the repository for new repositories

    ```bash
    goreleaser init
    ```

1. Test the snapshot without VCS deployment

    ```bash
    goreleaser release --snapshot --clean
    ```

1. Configure the default system's local `gitlab_token` or `github_token` secret

    ```bash
    mkdir -p ~/.config/goreleaser && touch ~/.config/goreleaser/gitlab_token
    mkdir -p ~/.config/goreleaser && touch ~/.config/goreleaser/github_token
    ```

1. Commit and push to VCS

    ```bash
    git add . && git commit -m "CI - Example" && git push
    ```

1. List `git` tags to get a version that isn't already established

    ```bash
    git tag --list
    ```

1. Using the output from above, increment the version and push a new tag

    ```bash
    git tag -a v0.0.1 -m "Bump: Initial Release" && git push origin v0.0.1
    ```

1. Create and push a new release

    ```bash
    goreleaser release --clean
    ```

1. Update local `Formula` to include the new tag. The following step is always required for private GitLab projects

    ```bash
    sed -i -e "s/using: GitDownloadStrategy/using: GitDownloadStrategy, tag: \"$(git tag --points-at HEAD)\"/g" ./dist/homebrew/Formula/go-cobra-cli.rb
    ```

1. Copy the updated `Formula` to system clipboard

    ```bash
    cat ./dist/homebrew/Formula/go-cobra-cli.rb | pbcopy
    ```

1. Update the Homebrew Tap's *.rb file with clipboard's contents

1. Tap the repository using `git+ssh` protocol - if the repository is private, ssh access is required

    ```bash
    # brew tap iac-factory/homebrew-taps git@github.com:iac-factory/homebrew-taps
    ```

1. Update the Cask if already established

    ```bash
    brew update
    ```

1. Install the package

    ```bash
    brew install iac-factory/homebrew-taps/go-cobra-cli --verbose --debug

    # --> Simplified
    brew install go-cobra-cli
    ```

## Reference(s)

---

- [`go` Linking](https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications?utm_source=reddit&utm_medium=social&utm_campaign=do-ldflags)
- [`goreleaser`](https://goreleaser.com/install/)
  - [`goreleaser` Environment Variable(s)](https://goreleaser.com/customization/env/)
- [Cobra User Guide](https://github.com/spf13/cobra/blob/main/site/content/user_guide.md)
- [Brew Formula Cookbook](https://github.com/Homebrew/brew/blob/master/docs/Formula-Cookbook.md)
- [New Cask Formula](https://github.com/Homebrew/homebrew-cask)

### Commands

```bash
# --> Delete Remote Tag
git push origin --delete v0.1.22

# --> HEAD Tag
git tag --points-at HEAD
```
