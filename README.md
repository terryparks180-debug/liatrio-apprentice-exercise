# Liatrio Apprenticeship Interview Exercise

A Go + Fiber API containerized with Docker, verified and published through GitHub Actions, and deployed to Google Cloud Run.

## Overview

This repository contains my work for the Liatrio Apprenticeship Interview Exercise.

The goal of the project was to build a simple web application and take it through the process of containerization, automated verification, image publishing, and cloud deployment.

The application exposes a single HTTP GET endpoint that returns my name and a dynamically generated timestamp as minified JSON.

## Tech Stack

- Go
- Fiber
- Docker
- GitHub Actions
- Docker Hub
- Google Cloud Run

## Application

The application exposes one HTTP endpoint:

`GET /`

Example response:

`{"message":"My name is Terry","timestamp":1791413589822}`

The timestamp is generated for every request using `time.Now().UnixMilli()`.

The application listens on port `3000`.

## Docker Containerization

The application uses a multi-stage Dockerfile.

The first stage uses Go to download dependencies and compile the application.

The second stage uses Alpine Linux and copies the compiled application into the final image.

This separates the build environment from the runtime environment.

### Running Locally

Build the Docker image:

`docker build -t liatrio-apprentice-app .`

Run the container:

`docker run -d -p 8080:3000 liatrio-apprentice-app`

Test the application:

`curl http://localhost:8080/`

## GitHub Actions

The GitHub Actions workflow is located at `.github/workflows/ci.yml`.

The workflow runs on pushes to `main` and on pull requests.

It performs the following steps:

1. Checks out the repository.
2. Builds the Docker image.
3. Starts the application in a Docker container.
4. Checks that the application responds using `curl`.
5. Runs Liatrio's apprentice-action verification.
6. Authenticates with Docker Hub using GitHub Secrets.
7. Tags the image using the GitHub Actions run number.
8. Pushes the versioned image to Docker Hub.

The verification step must succeed before the image can be published.

The Liatrio action is pinned to a specific commit SHA. The selected version runs six verification checks. The unfinished minified-JSON test on the newer master branch is not included in this version, so JSON formatting was also checked separately.

## Docker Hub

The published images are stored at:

`tpgotmotion/liatrio-apprentice-app`

The workflow uses the GitHub Actions run number as the image tag.

The image deployed to Google Cloud Run is:

`tpgotmotion/liatrio-apprentice-app:13`

## Cloud Deployment

The application is deployed to Google Cloud Run.

- Service: `liatrio-apprentice-app`
- Region: `us-west1`
- Container port: `3000`
- Deployed image: `tpgotmotion/liatrio-apprentice-app:13`

Public endpoint:

https://liatrio-apprentice-app-1042979448670.us-west1.run.app/

The endpoint was tested successfully in both a web browser and the terminal using `curl`.

The deployment to Cloud Run was performed manually. The build, verification, and image publishing steps are automated.

## Challenges and Lessons Learned

One of the main challenges was troubleshooting the Liatrio apprentice-action verification step.

The action's master branch contained an unfinished test that prevented the workflow from passing.

I initially used `continue-on-error` to test the publishing steps, but learned that this weakened the verification gate because publishing could proceed despite a failed check.

I removed that workaround, investigated the action's commit history, and pinned the dependency to an earlier working version using its full commit SHA.

This experience reinforced the importance of understanding failures rather than bypassing verification.

## Future Improvements

- Automate deployment to Cloud Run after successful verification and image publishing.
- Add another JSON field and verify that the updated application deploys successfully.
- Improve automated test coverage, including a minified-JSON test.
- Strengthen image versioning and dependency pinning.
- Add monitoring and deployment rollback
