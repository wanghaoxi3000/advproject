pipeline {
  agent none
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