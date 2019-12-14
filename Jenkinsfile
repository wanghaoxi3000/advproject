pipeline {
  agent none

  options {
      timeout(time: 3, unit: 'MINUTES')
  }

  stages {
    stage('Test') {
      agent {
        docker {
          image 'golang:1.13'
        }

      }
      steps {
        sh 'go test ./test'
      }
    }

  }
}