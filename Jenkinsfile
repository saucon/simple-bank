pipeline {
    agent any
    environment {
        ENV_MS_ACCOUNT = credentials('simple-bank-account')
        ENV_MS_CUSTOMER = credentials('simple-bank-customer')
    }
    stages {
        stage('Deploy Account') {
            when {
                changeset "**/account/**"
                beforeAgent true
             }
            steps {
                dir('account') {
                    sh '''
                    cp \$ENV_MS_ACCOUNT .
                    docker build -f Dockerfile -t simple-bank-account:latest .
                    docker rm -f simple-bank-account
                    docker run --name=simple-bank-account -d -p 9980:9980 simple-bank-account:latest
                    '''
                }
            }
        }
        stage('Deploy Customer') {
            when {
                changeset "**/customer/**"
                beforeAgent true
             }
            steps {
                dir('customer') {
                    sh '''
                    cp \$ENV_MS_CUSTOMER .
                    docker build -f Dockerfile -t simple-bank-customer:latest .
                    docker rm -f simple-bank-customer
                    docker run --name=simple-bank-customer -d -p 9981:9981 simple-bank-customer:latest
                    '''
                }
            }
        }

        stage('Force Deploy') {

            steps {
                script {
                    timeout(time: 30, unit: 'SECONDS') {
                        input(id: "Force Deploy", message: "Deploy All?", ok: 'Deploy')
                    }
                }
                dir('account') {
                    sh '''
                    cp \$ENV_MS_ACCOUNT .
                    docker build -f Dockerfile -t simple-bank-account:latest .
                    docker rm -f simple-bank-account
                    docker run --name=simple-bank-account -d -p 9980:9980 simple-bank-account:latest
                    '''
                }
                dir('customer') {
                    sh '''
                    cp \$ENV_MS_CUSTOMER .
                    docker build -f Dockerfile -t simple-bank-customer:latest .
                    docker rm -f simple-bank-customer
                    docker run --name=simple-bank-customer -d -p 9981:9981 simple-bank-customer:latest
                    '''
                }
            }
        }
    }
}
