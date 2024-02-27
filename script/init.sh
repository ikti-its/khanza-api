#!/bin/bash

read -p "Enter the desired module name: " module

mkdir "internal/${module}"
mkdir "internal/${module}/repository"
mkdir "internal/${module}/service"
mkdir "internal/${module}/controller"
mkdir "internal/${module}/model"
mkdir "internal/${module}/entity"
mkdir "internal/${module}/validation"

touch "internal/${module}/repository/${module}_repository.go"
touch "internal/${module}/repository/${module}_repository_impl.go"
touch "internal/${module}/service/${module}_service.go"
touch "internal/${module}/service/${module}_service_impl.go"
touch "internal/${module}/controller/${module}_controller.go"
touch "internal/${module}/controller/${module}_controller_impl.go"
touch "internal/${module}/model/${module}_model.go"
touch "internal/${module}/entity/${module}.go"
touch "internal/${module}/validation/${module}_validation.go"

echo "Module ${module} created successfully!"
