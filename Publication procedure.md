# Documentation for a publication procedure of the CLI binaries

In order to generate stable releases for IMCO Cloud Orchestrator Command Line Interface, the IMCO team proceeds based on the Software Development Life Cycle (SDLC) guidelines that complement the current procedure:

- Link to Confluence space: [Software Development Life Cycle](https://cloudteam.atlassian.net/wiki/spaces/COAL/pages/357369124/Software+Development+Life+Cycle#SoftwareDevelopmentLifeCycle-Releases)

> Note: Consider current "Go Releaser Procedure" document as a complementary helpful guide for CLI in addition to SDLC document; which is the main guide used in IMCO platform.

So, along development life cycle, several stages could take place from prereleases to stable releases; which mainly differ in stability and depending on planning time and release needs of project goals.

The succesive releases will be a set of different releases no published and identified as non-production ready.

Finally, a new release will be published and identified as production ready.

This publication procedure mechanism requires to install and configure "goreleaser" tool. The 'goreleaser' tool is added in order to automate the release generation process.

>"GoReleaser is a release automation tool for Go projects, the goal is to simplify the build, release and publish steps while providing variant customization options for all steps."

- <https://goreleaser.com/>
- <https://github.com/goreleaser/goreleaser>

> NOTE: Along this document version: 0.6.0 is taken only for reference purpouses

## Procedure

A release generation process can travel along different stages: 'Alpha', 'Beta', 'Candidates'(RC) until final 'Stable' releases (RTM):

Id | Case | VERSION | Production Ready
------------------------|---------------------|---------------------|---------------------
ALPHA | Alpha Prerelease -First- | 0.6.0-alpha | No
ALPHA.n | Alpha Prerelease -Successives- | 0.6.0-alpha.n | No
BETA | Beta Prerelease -First- | 0.6.0-beta | No
BETA.n | Beta Prerelease -Successives- | 0.6.0-beta.n | No
RC | Release Candidate -First- | 0.6.0-rc | No
RCn | Release Candidates -Successives- | 0.6.0-rc[n] | No
RTM | Latest Release | 0.6.0 | Yes

> Note: for a better document comprehension and depending on release generation context, the version will be evaluated: [CURRENT_VERSION] = [ALPHA | ALPHA.n | BETA | BETA.n | RC | RCn | RTM]

### 1. Init path

- Locate in the your project development git folder

### 2. Checkout and prepare adequate branch

- The release branch has to be created as defined in SDLC procedure; taking into account to be up to date from develop branch.

  Then, and in terms of the current release stage, the adequate branch should be created and published.

  So, checkout adequate branch for current release and update with latest changes.

### 3. Set adequate "VERSION" in code file

- Set the adequate release version in code file

  Update in "utils/version.go" the adequate version: `const VERSION = [CURRENT_VERSION]`

  Samples:

  ```go
  const VERSION = "0.6.0-alpha"
  const VERSION = "0.6.0-alpha.1"
  const VERSION = "0.6.0-beta"
  const VERSION = "0.6.0-beta.1"
  const VERSION = "0.6.0-rc"
  const VERSION = "0.6.0-rc1"
  const VERSION = "0.6.0"
  ```

- Add/Commit version file

  ```bash
  git add "utils/version.go"
  git commit -m "Update version to [CURRENT_VERSION]"
  ```

- Publish the release branch, depending of current case:

  Samples:

  ```bash
  git push release/0.6.0-alpha
  git push release/0.6.0-alpha.1
  git push release/0.6.0-beta
  git push release/0.6.0-beta.1
  git push release/0.6.0-rc
  git push release/0.6.0-rc1
  git push release/0.6.0
  ```

- Set environment variable

  ```bash
  export GITHUB_TOKEN="MY_GITHUB_TOKEN"
  ```

- Check the '.goreleaser.yml' file and set:

  - "prerelease": Identify the release as production ready or not.
    - ALPHA / BETA / RC:

      ```yaml
      prerelease: true
      ```

    - RTM:
      ```yaml
      prerelease: false
      ```

    > Note: this parameter can be changed later editing the release in <https://github.com/ingrammicro/concerto/releases>

### 4. ONLY IN RTM CASE. Prepare master branch

In order to create the latest release and to public as production ready:

- Create Pull Request into MASTER and wait for approval prior to merge and continue the procedure

Once approved and merged:

- Checkout the master branch

  ```bash
  git checkout master
  ```

- Update with the latest changes

  ```bash
  git pull -r
  ```

### 5. Tag the release

- Tag the new release

  ```bash
  git tag -a v[CURRENT_VERSION] -m "Concerto v[CURRENT_VERSION]"
  ```

  Samples:

  ```bash
  git tag -a v0.6.0-alpha -m "Concerto v0.6.0-alpha"
  git tag -a v0.6.0-alpha.1 -m "Concerto v0.6.0-alpha.1"
  git tag -a v0.6.0-beta -m "Concerto v0.6.0-beta"
  git tag -a v0.6.0-beta.1 -m "Concerto v0.6.0-beta.1"
  git tag -a v0.6.0-rc -m "Concerto v0.6.0-rc"
  git tag -a v0.6.0-rc1 -m "Concerto v0.6.0-rc1"
  git tag -a v0.6.0 -m "Concerto v0.6.0"
  ```

- Publish tag

  ```bash
  git push origin v[CURRENT_VERSION]
  ```

  Samples:

  ```bash
    git push origin v0.6.0-alpha
    git push origin v0.6.0-alpha.1
    git push origin v0.6.0-beta
    git push origin v0.6.0-beta.1
    git push origin v0.6.0-rc
    git push origin v0.6.0-rc1
    git push origin v0.6.0
  ```

### 6. RUN GORELEASER

Once the scenario is ready:

- Run 'goreleaser' to generate and publish the Latest Release
  > Note: if "dist is not empty" error appears, you should remove the existing 'dist' directory and re-launch goreleaser

### 7. ONLY IN RTM CASE. Merge into develop branch

In addition, and once the final stable release has been merged into master branch as well as tagged; the master branch must be merged into develop and this should be done through a PR.

  > Note: see more details in SDLC