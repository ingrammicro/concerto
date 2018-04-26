# Documentation for a publication procedure of the CLI binaries

In order to generate a new release for IMCO Cloud Orchestrator Command Line Interface, several stages could take place as release candidates - RCn - depending on planning time and release needs of project goals.

The succesive release candidates will be a set of different release no published and identified as non-production ready.

Finally, a new release will published and identified as production ready.

This publication procedure mechanism requires to install and configure goreleaser tool. The 'goreleaser' tool is added in order to automate the release generation process.

"GoReleaser is a release automation tool for Go projects, the goal is to simplify the build, release and publish steps while providing variant customization options for all steps."

- <https://goreleaser.com/>
- <https://github.com/goreleaser/goreleaser>

> NOTE: Along this document version: 0.6.0 is taken only for reference purpouses

## Procedure

A release can travel along different release candidates(RC) until final release (RTM):

Id | Case | VERSION | Production Ready
------------------------|---------------------|---------------------|---------------------
RC | Release Candidate -First- | 0.6.0-rc | No
RCn | Release Candidates -Successives- | 0.6.0-rc[n] | No
RTM | Latest Release | 0.6.0 | Yes

> Note: for a better document comprehension and depending on release generation context, the version will be evaluated: [CURRENT_VERSION] = [RC | RCn | RTM]

### 1. Init path

- Locate in the your project development git folder

### 2. Checkout and prepare adequate branch

- Checkout adequate branch.

    The first time for current release, the release branch has to be created.

  - RC case:
    - Checkout the develop branch
      ```bash
      git checkout origin/develop
      ```
    - Update with the latest changes
      ```bash
      git pull -r
      ```

    - Create the new release branch
      ```bash
      git checkout -b release/0.6.0
      ```

    - Publish the new release branch
      ```bash
      git push --set-upstream origin release/0.6.0
      ```

  - RCn / RTM cases:
    - Checkout the release branch
      ```bash
      git checkout origin/release/0.6.0
      ```

    - Update with the latest changes
      ```bash
      git pull -r
      ```

### 3. Set adequate "VERSION" in code file

- Set the adequate release version in code file

  Update in "utils/version.go" the adequate version: `const VERSION = [CURRENT_VERSION]`

  Samples:

  ```bash
  const VERSION = "0.6.0-rc"
  const VERSION = "0.6.0-rc1"
  const VERSION = "0.6.0"
  ```

- Add/Commit version file

  ```bash
  git add "utils/version.go"
  git commit -m "Update version to [CURRENT_VERSION]"
  ```

- Publish the release branch

  ```bash
  git push release/0.6.0
  ```

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

  ```bash
  Samples:
  git tag -a v0.6.0-rc -m "Concerto v0.6.0-rc"
  git tag -a v0.6.0-rc1 -m "Concerto v0.6.0-rc1"
  git tag -a v0.6.0 -m "Concerto v0.6.0"
  ```

- Publish tag

  ```bash
  git push origin v[CURRENT_VERSION]
  ```

  ```bash
    Samples:
    git push origin v0.6.0-rc
    git push origin v0.6.0-rc1
    git push origin v0.6.0
  ```

### 6. Configure goreleaser

- Set environment variable

  ```bash
  export GITHUB_TOKEN="MY_GITHUB_TOKEN"
  ```

- Check the '.goreleaser.yml' file and set:

  - "prerelease": Identify the release as production ready or not.
    - RC / RCn:

      ```bash
      prerelease: true
      ```

    - RTM:
      ```bash
      prerelease: false
      ```

    > Note: this parameter can be changed later editing the release in <https://github.com/ingrammicro/concerto/releases>

### 7. RUN GORELEASER

Once the scenario is ready:

- Run 'goreleaser' to generate and publish the Latest Release
  > Note: if "dist is not empty" error appears, you should remove the existing 'dist' directory and re-launch goreleaser
