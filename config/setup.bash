#!/bin/bash

aws ecs register-task-definition \
    --cli-input-json file://config/renderdog-server-task.json

aws ecs run-task \
    --task-definition 'renderdog:1' \
    --count 1 

aws ecs create-service \
    --service-name renderdog \
    --task-definition 'renderdog:1' \
    --desired-count 1 

