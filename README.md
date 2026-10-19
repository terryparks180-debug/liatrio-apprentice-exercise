# Liatrio Apprenticeship Interview Exercise

A Go + Fiber web application containerized with Docker, automatically verified and published through GitHub Actions, and deployed to Google Cloud Run using a CI/CD pipeline.

## Overview

This repository contains my work for the Liatrio Apprenticeship Interview Exercise.

The goal of this project was to build a simple web application and take it through containerization, automated verification, image publishing, and cloud deployment.

I divided the exercise into five sprints so I could focus on one part at a time, verify my progress, and improve my work based on feedback.

I also extended the project by implementing automated deployment to Google Cloud Run and adding my own minified JSON verification step.

## Tech Stack

- **Go:** Application programming language
- **Fiber:** Web framework for handling HTTP requests
- **Docker:** Application containerization
- **GitHub Actions:** CI/CD automation
- **Docker Hub:** Container image registry
- **Google Cloud Run:** Cloud hosting and deployment
- **Python:** Automated minified JSON verification

## Application

The application exposes one HTTP endpoint:

`GET /`

Example JSON response:

`{"message":"My name is Terry","timestamp":1791413589822,"version 2":"demo 2"}`

The version field is used to demonstrate application changes and may vary between deployments.

The timestamp is dynamically generated using `time.Now().UnixMilli()`, which returns the current Unix timestamp in milliseconds.

The JSON response is returned using Fiber's `c.JSON()` function.

The application listens on port `3000`.

## Docker Containerization

I chose to use a multi-stage Dockerfile to separate the build environment from the runtime environment.

**Stage 1 — Build**

The first stage uses the Go environment to:

- Download application dependencies.
- Copy the Go source code.
- Compile the application into an executable.

**Stage 2 — Runtime**

The second stage uses a lightweight Alpine Linux image to:

- Copy the compiled executable from the builder stage.
- Set the application's working directory.
- Expose port `3000`.
- Start the compiled application using `CMD`.

This approach keeps the final image smaller because the Go compiler and other development tools are not needed in the runtime image.

### Running Locally

Build the Docker image:

`docker build -t liatrio-apprentice-app .`

Run the container:

`docker run -d -p 8080:3000 liatrio-apprentice-app`

Test the application:

`curl http://localhost:8080/`

## GitHub Actions — CI/CD Pipeline

My automated workflow is located at:

`.github/workflows/ci.yml`

The workflow runs on pushes to `main` and on pull requests.

It contains two main jobs: **build** and **deploy**.

### Build and Verification Job

The build job:

1. Checks out the GitHub repository.
2. Builds the Docker image.
3. Starts the application in a Docker container.
4. Uses `curl` to verify that the application responds.
5. Runs my custom minified JSON verification.
6. Runs Liatrio's apprentice-action verification.
7. Authenticates with Docker Hub using GitHub Secrets.
8. Tags the image using the GitHub Actions run number.
9. Pushes the versioned image to Docker Hub.

Docker Hub publishing occurs on successful pushes to `main`, not during pull-request builds.

Verification must succeed before the image can be published.

### Liatrio Verification Troubleshooting

One of my biggest challenges was troubleshooting Liatrio's apprentice-action.

The newer master version contained an unfinished minified JSON test that caused verification errors.

Initially, I used `continue-on-error` to investigate whether the remaining publishing steps worked. However, I learned that this weakened the pipeline because it allowed publishing to continue even when verification failed.

After receiving feedback, I investigated the action's source code and commit history.

I decided to pin the action to an earlier working commit:

`0b41561cca6822cc8d880fe0e49e7807a41fdf91`

That version successfully runs six verification checks.

However, the earlier version does not include the newer minified JSON test.

### Additional Minified JSON Verification

Because the exercise explicitly requires minified JSON, I added my own automated verification step to GitHub Actions.

The verification uses `curl` and Python to:

1. Request the JSON response from my running application.
2. Parse the response to confirm it is valid JSON.
3. Generate a compact JSON representation without unnecessary whitespace.
4. Compare the compact representation against the original response.
5. Fail the verification if the responses do not match.

This check uses Python's compact JSON serialization as the expected format.

The verification runs before the Liatrio action and before image publishing.

**Testing performed:**

- Confirmed that the automated verification passed in GitHub Actions.
- Tested the Python comparison logic locally using compact JSON.
- Tested the logic using JSON containing unnecessary whitespace.
- Confirmed that the compact example passed and the non-minified example was rejected.

The local negative test checked the verification logic independently; it was not a full failed CI/CD run.

This additional test helps maintain verification coverage without relying on the unfinished test in the newer Liatrio action.

## Docker Hub

My published Docker images are stored at:

https://hub.docker.com/r/tpgotmotion/liatrio-apprentice-app

The workflow uniquely tags each published image using its GitHub Actions run number.

Image format:

`tpgotmotion/liatrio-apprentice-app:<run_number>`

This allows each successful publishing run to produce an identifiable image version.

## Automated Google Cloud Run Deployment

The application is deployed to Google Cloud Run.

**Cloud Run configuration:**

- Service: `liatrio-apprentice-app`
- Region: `us-west1`
- Container port: `3000`
- Image registry: Docker Hub

**Public endpoint:**

https://liatrio-apprentice-app-1042979448670.us-west1.run.app/

### Continuous Deployment

I initially deployed the application to Google Cloud Run manually.

After completing that process, I extended the project by adding an automated deployment job to GitHub Actions.

The deployment job uses `needs: build`, meaning the build job must succeed before deployment can begin.

When a push to `main` successfully completes the build and verification steps, the workflow:

1. Publishes the uniquely tagged Docker image to Docker Hub.
2. Authenticates with Google Cloud.
3. Deploys the versioned Docker Hub image to Cloud Run.
4. Uses `curl` to verify that the deployed application responds successfully.

I tested the end-to-end process by updating the application, pushing the changes to GitHub, and confirming that the changes appeared at the public Cloud Run URL.

### Workload Identity Federation

To allow GitHub Actions to deploy securely to Google Cloud, I configured Workload Identity Federation.

This allows GitHub Actions to authenticate with Google Cloud using short-lived credentials instead of storing a long-lived Google Cloud service account private key.

The deployment workflow uses a Google Cloud service account with the permissions required to deploy the application.

This helped me understand how authentication and permissions are used in automated cloud deployments.

## Challenges and Lessons Learned

### Troubleshooting Verification Failures

One of the biggest lessons I learned was the importance of investigating failures rather than bypassing them.

My experience with `continue-on-error` helped me understand why verification needs to remain a required part of the CI/CD pipeline.

### Understanding Docker

Using a multi-stage Dockerfile helped me understand the difference between building an application and running a compiled executable inside a container.

I also learned how separating build and runtime environments can reduce unnecessary image contents.

### CI/CD Automation

I gained hands-on experience connecting multiple tools into a working automated pipeline.

Instead of manually repeating the build, verification, publishing, and deployment processes, GitHub Actions now handles those steps after a successful push to `main`.

### Communication and Feedback

I worked through the exercise in smaller stages while researching unfamiliar technologies, troubleshooting problems, and applying feedback.

This helped me become more comfortable investigating issues, communicating my progress, and improving my work.

## Future Improvements

If I had more time, I would consider:

- Documenting and testing a rollback procedure using Cloud Run revisions.
- Implementing monitoring and alerts for application availability.
- Further improving deployment reliability and security.

## Final Outcome

This project demonstrates an end-to-end CI/CD process.

I built a Go web application, containerized it with Docker, implemented automated verification, published versioned images to Docker Hub, and configured automatic deployment to Google Cloud Run.

I also added an independent minified JSON verification step to cover a requirement not checked by the pinned version of Liatrio's apprentice-action.

The most valuable part of this exercise was learning how to investigate failures, understand the tools I'm using, and improve a working system through testing and feedback.
