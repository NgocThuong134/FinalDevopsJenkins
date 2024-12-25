pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'ngocthuong/server-final'
        DOCKER_TAG = 'latest'
        TELEGRAM_BOT_TOKEN = '7649334871:AAF9YVvIXp3SlCAVS27BOUWMwisFCDZM0y4'
        TELEGRAM_CHAT_ID = '-1002452635800'
        PROD_SERVER = 'ec2-18-141-140-203.ap-southeast-1.compute.amazonaws.com'
        PROD_USER =  'ubuntu'
    }

    stages {
        stage('Clone Repository') {
            steps {
                git branch: 'main', url: 'https://github.com/NgocThuong134/FinalDevopsJenkins.git'
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
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
                script {
                    echo 'Clearing server-final-related images and containers...'
                    sh '''
                        docker container stop server-final || echo "No container named server-final to stop"
                        docker container rm server-final || echo "No container named server-final to remove"
                        docker image rm ${DOCKER_IMAGE}:${DOCKER_TAG} || echo "No image named ${DOCKER_IMAGE}:${DOCKER_TAG} to remove"
                    '''
                }
                echo 'Deploying to DEV environment...'
                sh 'docker image pull ngocthuong/server-final:latest'
                sh 'docker container stop server-final || echo "this container does not exist"'
                sh 'docker network create dev || echo "this network exists"'
                sh 'echo y | docker container prune '

                sh 'docker container run -d --rm --name server-final -p 7080:3000 --network dev ngocthuong/server-final:latest'
            }
        }

        stage ('Deploy to Production on AWS') {
            steps{
                script {
                    echo 'Deploying to Production...'
                    sshagent(['aws-ssh-key']) {
                    sh '''
                        ssh -o StrictHostKeyChecking=no ${PROD_USER}@${PROD_SERVER} << EOF
                         docker container stop server-final || echo "No container to stop"
                         docker container rm server-final || echo "No container to remove"
                         docker image rm ${DOCKER_IMAGE}:${DOCKER_TAG} || echo "No image to remove"
                         docker image pull ${DOCKER_IMAGE}:${DOCKER_TAG}
                         docker container run -d --rm --name server-final -p 7081:7080 ${DOCKER_IMAGE}:${DOCKER_TAG}
                    '''
                    }
                }
            }
        }
    }

    post {
        always {
            cleanWs()
        }
        success {
            sendTelegramMessage("✅ Build #${BUILD_NUMBER} was successful! ✅")
        }

        failure {
            sendTelegramMessage("❌ Build #${BUILD_NUMBER} failed. ❌")
        }
    }
}

def sendTelegramMessage(String message) {
    sh """
    curl -s -X POST https://api.telegram.org/bot${TELEGRAM_BOT_TOKEN}/sendMessage \
    -d chat_id=${TELEGRAM_CHAT_ID} \
    -d text="${message}"
    """
}