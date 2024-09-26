# ECS Task Definition Image Tag Difference
This GitHub Action automates the process of detecting differences in Docker image tags within ECS task definitions and automatically creates a pull request (PR) with the updated image tag. This can be useful when you want to keep your ECS services up-to-date with the latest container versions without manual intervention.
<img width="1304" alt="スクリーンショット 2024-09-24 23 01 00" src="https://github.com/user-attachments/assets/793bd6d0-e5e2-487b-9403-42d3ebe6c69d">

Features
Detects changes in the Docker image tag used in your ECS task definition.
Automatically creates a pull request when an image tag change is detected.
Compatible with both amd64 and arm64 architectures.
Easily customizable to suit different ECS environments.

# Usage
pull docker image from github container registory
https://github.com/naonao2323/ecs-task-def/pkgs/container/ecs-task-def

```
start ecs-task-def

Usage:
  ecs-task-def [flags]

Flags:
      --container-name string      container name
      --container-path string      the path to the container definition
      --github-email string        git email
      --github-owner string        github owner
      --github-repository string   github repositoy (etc ecs-task-def)
      --github-token string        github token (pat)
      --github-url string          github repository https url (etc https://github.com/naonao2323/ecs-task-def.git)
      --github-username string     git username
  -h, --help                       help for ecs-task-def
      --target-tag string          target tag
      --task-path string           the path to the task definition
```
# Support
