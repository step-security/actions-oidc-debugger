# actions-oidc-debugger

This action requests a JWT and prints the claims included within the JWT received from GitHub Actions.

## How to use this Action

Here's an example of how to use this action:

```yaml

name: Test Debugger Action
on: 
  pull_request:
  workflow_dispatch:

jobs:
  oidc_debug_test:
    permissions:
      contents: read
      id-token: write
    runs-on: ubuntu-latest
    name: A test of the oidc debugger
    steps:
      - name: Debug OIDC Claims
        uses: step-security/actions-oidc-debugger@v1
        with:
          audience: '${{ github.server_url }}/${{ github.repository_owner }}'
```

The resulting output in your Actions log will look something like this:

```json
{
  "actor": "amanstep",
  "actor_id": "206509459",
  "aud": "https://github.com/step-security",
  "base_ref": "main",
  "enterprise": "step-security",
  "enterprise_id": "18181",
  "event_name": "pull_request",
  "exp": 1759908794,
  "head_ref": "release",
  "iat": 1759908494,
  "iss": "https://token.actions.githubusercontent.com",
  "job_workflow_ref": "step-security/actions-oidc-debugger/.github/workflows/main.yml@refs/pull/2/merge",
  "job_workflow_sha": "0207a4bb93eb7a6af22451b92a89eb23585fd5a2",
  "jti": "8395a793-0929-40a1-bd32-e12f2e856bed",
  "nbf": 1759908194,
  "ref": "refs/pull/2/merge",
  "ref_protected": "false",
  "ref_type": "branch",
  "repository": "step-security/actions-oidc-debugger",
  "repository_id": "1071992672",
  "repository_owner": "step-security",
  "repository_owner_id": "88700172",
  "repository_visibility": "public",
  "run_attempt": "1",
  "run_id": "18337179416",
  "run_number": "1",
  "runner_environment": "github-hosted",
  "sha": "0207a4bb93eb7a6af22451b92a89eb23585fd5a2",
  "sub": "repo:step-security/actions-oidc-debugger:pull_request",
  "workflow": "Test OIDC Debugger",
  "workflow_ref": "step-security/actions-oidc-debugger/.github/workflows/main.yml@refs/pull/2/merge",
  "workflow_sha": "0207a4bb93eb7a6af22451b92a89eb23585fd5a2"
}
```

### Bootstrapping

This assumes you have `goenv` installed and the version listed in the `.go-version` file is installed as well.

```bash
go mod vendor && go mod tidy && go mod verify
```
