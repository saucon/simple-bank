pipeline {
    agent any
    stages {
        stage('Build Account') {
            when {
                changeset "**/account/**"
                beforeAgent true
             }
            steps {
            dir('account') {
                  sh '''
                    docker build -f Dockerfile -t simple-bank-account:latest .
                    docker rm -f simple-bank-account
                    docker run --name=simple-bank-account -d -p 9980:9980 simple-bank-account:latest
                  '''
                }
            }
        }
        stage('Build Customer') {
            when {
                changeset "**/customer/**"
                beforeAgent true
             }
            steps {
            dir('customer') {
                  sh '''
                    docker build -f Dockerfile -t simple-bank-customer:latest .
                    docker rm -f simple-bank-customer
                    docker run --name=simple-bank-customer -d -p 9981:9981 simple-bank-customer:latest
                  '''
                }
            }
        }
    }
}
