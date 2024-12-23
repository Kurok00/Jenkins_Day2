pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'anvnt96/golang-jenkins'
        DOCKER_TAG = 'latest'
    }

    stages {
        stage('Clone Repository') {
            steps {
                git branch: 'master', url: 'https://github.com/Kurok00/Jenkins_Day2.git'
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    // Fix Docker socket permissions
                    sh '''
                        sudo chown root:docker /var/run/docker.sock || true
                        sudo chmod 666 /var/run/docker.sock || true
                        sudo usermod -aG docker jenkins || true
                        newgrp docker || true
                    '''
                    
                    // Build with error handling
                    sh '''
                        if ! docker build -t ${DOCKER_IMAGE}:${DOCKER_TAG} .; then
                            echo "Docker build failed. Retrying with sudo..."
                            sudo docker build -t ${DOCKER_IMAGE}:${DOCKER_TAG} .
                        fi
                    '''
                }
            }
        }

        stage('Run Tests') {
            steps {
                echo 'Running tests...'
            }
        }

        stage('Push to Docker Hub') {
            steps {
                script {
                    sh 'sudo chmod 666 /var/run/docker.sock || true'
                    docker.withRegistry('https://index.docker.io/v1/', 'docker-hub-credentials') {
                        docker.image("${DOCKER_IMAGE}:${DOCKER_TAG}").push()
                    }
                }
            }
        }

        stage('Deploy Golang to DEV') {
            steps {
                script {
                    sh 'sudo chmod 666 /var/run/docker.sock || true'
                    
                    sh '''
                        sudo docker image pull ${DOCKER_IMAGE}:${DOCKER_TAG}
                        sudo docker rm -f server-golang || true
                        sudo docker network create dev || true
                        sudo docker container run -d --rm --name server-golang -p 4001:4001 --network dev ${DOCKER_IMAGE}:${DOCKER_TAG}
                    '''
                }
            }
        }
    }

    post {
        success {
            node('built-in') {
                script {
                    withCredentials([
                        string(credentialsId: 'telegram-token', variable: 'TELEGRAM_TOKEN'),
                        string(credentialsId: 'telegram-chatid', variable: 'TELEGRAM_CHAT')
                    ]) {
                        def message = "✅ Build SUCCESS!\nJob: ${env.JOB_NAME}\nBuild Number: ${env.BUILD_NUMBER}\nBuild URL: ${env.BUILD_URL}"
                        sh """#!/bin/bash
                            curl -s -X POST https://api.telegram.org/bot${TELEGRAM_TOKEN}/sendMessage \
                            -d chat_id=${TELEGRAM_CHAT} \
                            -d parse_mode=HTML \
                            -d text="${message}"
                        """
                    }
                }
                cleanWs()
            }
        }
        failure {
            node('built-in') {
                script {
                    withCredentials([
                        string(credentialsId: 'telegram-token', variable: 'TELEGRAM_TOKEN'),
                        string(credentialsId: 'telegram-chatid', variable: 'TELEGRAM_CHAT')
                    ]) {
                        def message = "❌ Build FAILED!\nJob: ${env.JOB_NAME}\nBuild Number: ${env.BUILD_NUMBER}\nBuild URL: ${env.BUILD_URL}"
                        sh """#!/bin/bash
                            curl -s -X POST https://api.telegram.org/bot${TELEGRAM_TOKEN}/sendMessage \
                            -d chat_id=${TELEGRAM_CHAT} \
                            -d parse_mode=HTML \
                            -d text="${message}"
                        """
                    }
                }
                cleanWs()
            }
        }
    }
}