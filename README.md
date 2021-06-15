# packaging


# Github Actions:
GitHub Actions are based on the concept of Workflows. A workflow is nothing more than a set of jobs and steps that are executed when some condition or event is met. (Ex: a push to the repository, a pull request, a deployment, etc).

### Note: You can have multiple workflows by project, each one responding to a different set of events.

1. In our example, we will have two workflows. The "Build" or "Main" workflow which will be triggered when there is a push the master branch or when a PR is created and the "Release" workflow which will run when a new tag is pushed to GitHub, that will create a new release of the application.

Our "Build" Workflow will have 3 Jobs (Lint, Build and Test) and our "Release" workflow will have a single "release" job.

examples:Each job is made of steps. For Example, the "Unit Test" job will have steps for checkout the source code, run the tests and generating code coverage report.


## Workflow:
Workflows are defined in YAML files located in .github/workflows directory of your repository.
Each file in the directory represents a different Workflow.

1.In our case, we want it to run when there is a push to master or a pull request.
examples:The workflow contains 3 jobs, "lint", "test" and "build".

2.runs-on:ubuntu-lastest (means specify that we want this job to run on an ubuntu machine. ("runs-on" keyword)).

3.Then, we define the steps that compose our job..First thing is to install Go. GitHub already provides an action for it, so let's use it:
-   name: Set up Go
    uses: actions/setup-go@v1
    with:
        go-version: 1.12
note: with keyword allows us to specify the arguments required by the action. In this case, the "setup-go" action allows us to specify the go version to use.

4.Next step is to check-out the source code. Again we will use a built-in action:

-  name: Check out code
   uses: actions/checkout@v1
5.And finally we will install and run golint tool:

- name: Lint Go Code
  run: |
    export PATH=$PATH:$(go env GOPATH)/bin
    go get -u golang.org/x/lint/golint 
    make lint


Similary in "Test" job:
 test:
    name: Test
    runs-on: ubuntu-latest
    steps:
      - name: Set up Go
        uses: actions/setup-go@v1
        with:
          go-version: 1.12

      - name: Check out code
        uses: actions/checkout@v1

      - name: Run Unit tests.
        run: make test-coverage
      
      - name: Upload Coverage report to CodeCov
        uses: codecov/codecov-action@v1.0.0
        with:
          token: ${{secrets.CODECOV_TOKEN}}
          file: ./coverage.txt

new thing in Test job: using a third-party action, in this case, to publish the test coverage report to CodeCov.
