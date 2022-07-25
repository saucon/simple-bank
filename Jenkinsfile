pipeline {
    agent any
    stages {
        stage('Build Account') {
            when {
                changeset "**/account/**"
                beforeAgent true
             }
            steps {
                sh('cd account')
                sh('docker build -f Dockerfile -t simple-bank-account:latest .')
                sh('docker rm -f simple-bank-account')
                sh('docker run --name=simple-bank-account -d -p 9980:9980 simple-bank-account:latest')
            }
        }
        stage('Build Customer') {
            when {
                changeset "**/customer/**"
                beforeAgent true
             }
            steps {
                sh('cd customer')
                sh('docker build -f Dockerfile -t simple-bank-customer:latest .')
                sh('docker rm -f simple-bank-customer')
                sh('docker run --name=simple-bank-customer -d -p 9981:9981 simple-bank-customer:latest')
            }
        }
    }
}