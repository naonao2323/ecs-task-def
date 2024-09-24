# ECS Task Definition Image Tag Difference Action
This GitHub Action automates the process of detecting differences in Docker image tags within ECS task definitions and automatically creates a pull request (PR) with the updated image tag. This can be useful when you want to keep your ECS services up-to-date with the latest container versions without manual intervention.

Features
Detects changes in the Docker image tag used in your ECS task definition.
Automatically creates a pull request when an image tag change is detected.
Compatible with both amd64 and arm64 architectures.
Easily customizable to suit different ECS environments.
