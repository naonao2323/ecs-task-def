# ECS Task Definition Image Tag Difference
This GitHub Action automates the process of detecting differences in Docker image tags within ECS task definitions and automatically creates a pull request (PR) with the updated image tag. This can be useful when you want to keep your ECS services up-to-date with the latest container versions without manual intervention.
<img width="1304" alt="スクリーンショット 2024-09-24 23 01 00" src="https://github.com/user-attachments/assets/793bd6d0-e5e2-487b-9403-42d3ebe6c69d">

Features
Detects changes in the Docker image tag used in your ECS task definition.
Automatically creates a pull request when an image tag change is detected.
Compatible with both amd64 and arm64 architectures.
Easily customizable to suit different ECS environments.
