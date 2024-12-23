pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'anvnt96/golang-jenkins'
        DOCKER_TAG = 'latest'
        TELEGRAM_BOT_TOKEN = credentials('telegram-bot-token')
        TELEGRAM_CHAT_ID = credentials('telegram-chat-id')
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
                    sh 'sudo chmod 666 /var/run/docker.sock'
                    docker.build("${DOCKER_IMAGE}:${DOCKER_TAG}")
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
                    docker.withRegistry('https://index.docker.io/v1/', 'docker-hub-credentials') {
                        docker.image("${DOCKER_IMAGE}:${DOCKER_TAG}").push()
                    }
                }
            }
        }

        stage('Deploy Golang to DEV') {
            steps {
                echo 'Deploying to DEV...'
                sh 'sudo docker image pull anvnt96/golang-jenkins:latest'
                
                // Check if container exists and is running
                sh '''
                    CONTAINER_ID=$(sudo docker ps -q -f name=server-golang)
                    if [ ! -z "$CONTAINER_ID" ]; then
                        echo "Container is running, stopping it..."
                        sudo docker stop $CONTAINER_ID
                    fi
                    
                    # Remove container if it exists but not running
                    sudo docker rm -f server-golang || true
                '''
                
                sh 'sudo docker network create dev || echo "this network exists"'
                sh 'sudo echo y | docker container prune '

                sh 'sudo docker container run -d --rm --name server-golang -p 4001:4001 --network dev anvnt96/golang-jenkins:latest'
            }
        }
    }

    post {
        success {
            script {
                def message = "✅ Build SUCCESS!\nJob: ${env.JOB_NAME}\nBuild Number: ${env.BUILD_NUMBER}\nBuild URL: ${env.BUILD_URL}"
                sh """
                    curl -s -X POST https://api.telegram.org/bot${TELEGRAM_BOT_TOKEN}/sendMessage \
                    -d chat_id=${TELEGRAM_CHAT_ID} \
                    -d parse_mode=HTML \
                    -d text="${message}"
                """
            }
            cleanWs()
        }
        failure {
            script {
                def message = "❌ Build FAILED!\nJob: ${env.JOB_NAME}\nBuild Number: ${env.BUILD_NUMBER}\nBuild URL: ${env.BUILD_URL}"
                sh """
                    curl -s -X POST https://api.telegram.org/bot${TELEGRAM_BOT_TOKEN}/sendMessage \
                    -d chat_id=${TELEGRAM_CHAT_ID} \
                    -d parse_mode=HTML \
                    -d text="${message}"
                """
            }
            cleanWs()
        }
        always {
        }
    }
}
