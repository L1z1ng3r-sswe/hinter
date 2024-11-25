# Project Management Methodologies, Quality Gates, and SDLC

1. [Project Management Methodologies](#project-management-methodologies)
    - [Waterfall](#waterfall)
    - [Agile](#agile)
    - [Scrum](#scrum)
    - [Kanban](#kanban)
    - [Blending Methodologies](#blending-methodologies)
2. [Software Development Life Cycle (SDLC)](#software-development-life-cycle-sdlc)
    - [Phases](#phases)
    - [Models](#models)
3. [Quality Gates](#quality-gates)
    - [Definition](#definition)
    - [Benefits](#benefits)
    - [Key Metrics](#key-metrics)
    - [Blockers (in Golang Backend Applications)](#blockers-in-golang-backend-applications)
    - [Implementation Tools](#implementation-tools)
    - [GitLab CI Pipelines](#gitlab-ci-pipelines)

---

## Project Management Methodologies

### Waterfall <a id="waterfall"></a>

A linear and sequential approach, ideal for projects with well-defined requirements. Phases include:
- **Requirement Gathering**: Complete upfront analysis and documentation.
- **Design**: Create system architecture and detailed plans.
- **Implementation**: Develop the solution as per design specifications.
- **Testing**: Verify that the system meets all requirements.
- **Deployment**: Deliver the final product.
- **Maintenance**: Provide updates and bug fixes post-deployment.

---

### Agile <a id="agile"></a>

An iterative methodology focused on delivering incremental value through continuous feedback and collaboration. Key features include:
- **Short Iterations**: Frequent delivery of working software.
- **Flexibility**: Adapt to changing priorities.
- **Collaboration**: Strong emphasis on stakeholder involvement.
- **Blockers Addressed Early**: Agile ceremonies like standups are used to identify and resolve blockers.

---

### Scrum <a id="scrum"></a>

A framework under Agile that organizes work into **sprints** (time-boxed iterations). Core elements:
- **Roles**: Product Owner, Scrum Master, Development Team.
- **Ceremonies**:
  - Sprint Planning
  - Daily Standups
  - Sprint Reviews
  - Retrospectives
- **Sprint Board**: Tracks tasks with columns like **To Do**, **In Progress**, **Done**.

---

### Kanban <a id="kanban"></a>

A flow-based approach focusing on visualizing work and limiting Work In Progress (WIP). Key components:
- **Kanban Board**:
  - Columns such as **Backlog**, **Ready**, **In Progress**, **Testing**, **Done**.
  - WIP limits to optimize task flow.
- **Continuous Delivery**: No fixed iterations; work is delivered as soon as it’s ready.

---

### Blending Methodologies <a id="blending-methodologies"></a>

In practice, methodologies can be combined to suit project needs:
- Start with **Waterfall** for detailed planning.
- Use **Agile/Scrum** for iterative development and deliveries.
- Use **Kanban** for maintenance and support tasks.

---

## Software Development Life Cycle (SDLC) <a id="software-development-life-cycle-sdlc"></a>

SDLC is a structured process for developing high-quality software. It defines a sequence of phases to ensure software meets user requirements and is maintainable.

---

### Phases <a id="phases"></a>

1. **Requirement Analysis**:
   - Collect and analyze business requirements.
   - Deliverable: Requirement Specification Document.
2. **Planning**:
   - Define scope, timeline, resources, and risks.
   - Deliverable: Project Plan.
3. **Design**:
   - Create the architecture, UI/UX design, and database structure.
   - Deliverable: System Design Document.
4. **Development**:
   - Write and integrate the application code.
   - Deliverable: Functional Software.
5. **Testing**:
   - Validate the software against requirements.
   - Deliverable: Test Reports.
6. **Deployment**:
   - Deploy the software to production environments.
   - Deliverable: Live Software.
7. **Maintenance**:
   - Provide updates, bug fixes, and support.
   - Deliverable: Updated Versions.

---

### Models <a id="models"></a>

SDLC can follow various models based on the project’s nature:
1. **Waterfall Model**: Sequential flow through SDLC phases.
2. **Iterative Model**: Cyclical approach, refining features incrementally.
3. **Agile Model**: Focuses on iterative development and collaboration.
4. **Spiral Model**: Combines iterative development with risk assessment.
5. **DevOps Model**: Emphasizes automation and continuous delivery.

---

## Quality Gates

### Definition <a id="definition"></a>

A **quality gate** is a set of conditions or criteria used to assess the quality of code or processes before they can advance to the next stage in the development lifecycle.

---

### Benefits <a id="benefits"></a>

1. Ensures code quality and maintainability.
2. Detects issues early in the development process.
3. Enforces compliance with standards and regulations.
4. Reduces technical debt by avoiding problematic code.

---

### Key Metrics <a id="key-metrics"></a>

1. **Code Coverage**: Ensures sufficient automated test coverage.
2. **Critical Issues**: No blocker or critical bugs, especially related to security.
3. **Code Smells**: Avoidance of redundant logic and overly complex methods.
4. **Technical Debt**: Ensures debt is kept within acceptable thresholds.

---

### Blockers (in Golang Backend Applications) <a id="blockers-in-golang-backend-applications"></a>

Blockers include: SQL injections, runtime panics, missing dependencies, broken tests, performance regressions, deadlocks, race conditions, and incorrect API contracts.

---

### Implementation Tools <a id="implementation-tools"></a>

1. **SonarQube**: Quality gate enforcement.
2. **Static Analysis Tools**:
   - **GoSec**: Detects security issues in Go code.
   - **Staticcheck**: Finds bugs and performance issues.
3. **CI/CD Platforms**: GitLab CI and Google Cloud Build.

---

### GitLab CI Pipelines <a id="gitlab-ci-pipelines"></a>

```yaml
stages: # Define the stages of the pipeline
  - build
  - test
  - lint
  - deploy

variables:
  GO_VERSION: "1.20"
  GOPATH: "$CI_PROJECT_DIR/go"

default:
  image: golang:${GO_VERSION}-alpine
  before_script:
    - apk add --no-cache git
    - go mod download

build:
  stage: build
  script:
    - go build -o app .
    - echo "Build successful!"
  artifacts:
    paths:
      - app

test:
  stage: test
  script:
    - go test ./... -v
  allow_failure: false

lint:
  stage: lint
  script:
    - go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
    - golangci-lint run
  allow_failure: false

deploy:
  stage: deploy
  script:
    - echo "Deploying to production..."
    - ./deploy_script.sh
  only:
    - main
```