pipeline{
    agent any
    tools{
        go 'go-1.12'
    }
    environment{
        GO111MODULE = 'on'
    }
    stages{
        stage('Checkout'){
            steps{
               checkout scm
            }
        }
        stage('Compile'){
            steps{
                sh 'go build'
            }
        }
    }
}
