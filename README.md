# Liatrio Apprenticeship Interview Exercise

A Go + Fiber API that will be containerized with Docker, validated and published through GitHub Actions, and deployed to a cloud platform.

## Overview

This repository contains my work for the Liatrio Apprenticeship Interview Exercise.

The application exposes a single HTTP endpoint that returns a JSON object containing a name and a dynamically generated timestamp.

## Stack

- Go + Fiber
- Docker
- GitHub Actions
- Docker Hub
- Cloud platform: TBD

## Status

Work in progress. See `PLAN.md` for the current approach and project stages.

## Endpoint

```text
GET /
```

Expected response:

```json
{"message":"My name is Terry","timestamp":1234567890}
```

## Running Locally

Instructions will be added as each working stage is completed.
